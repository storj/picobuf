// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"fmt"
	"math"
	"slices"

	"storj.io/picobuf/internal/protowire"
)

// Message is an interface that all generated messages implement.
type Message interface {
	Encode(*Encoder) bool
	Decode(*Decoder)
}

// RequiredFields is implemented by generated messages that can validate
// legacy required fields.
type RequiredFields interface {
	PicoValidateRequired() error
}

// ValidateRequired validates all legacy required fields known to msg.
func ValidateRequired(msg any) error {
	if required, ok := msg.(RequiredFields); ok {
		return required.PicoValidateRequired()
	}
	return nil
}

// MissingRequiredField returns an error for an unset legacy required field.
func MissingRequiredField(field FieldNumber) error {
	return fmt.Errorf("required field %s is not set", field)
}

// CustomType defines methods that are used for custom encode or decode
// behaviors. PicoEncode reports whether it encoded the field.
type CustomType interface {
	PicoEncode(*Encoder, FieldNumber) bool
	PicoDecode(*Decoder, FieldNumber)
}

// FieldNumber corresponds to a protobuf field number.
type FieldNumber protowire.Number

// IsValid returns whether the field number is in correct range.
func (field FieldNumber) IsValid() bool { return protowire.Number(field).IsValid() }

// String converts field number to a number.
func (field FieldNumber) String() string {
	const maxDigitsUint32 = len("-2147483648")
	var z [maxDigitsUint32]byte // for 32bits
	if field == 0 {
		return "0"
	}
	if field == math.MinInt32 {
		return "-2147483648"
	}
	negative := field < 0
	if negative {
		field = -field
	}
	i := len(z) - 1
	for ; i >= 0 && field > 0; i-- {
		z[i] = byte(field%10) + '0'
		field /= 10
	}
	if negative {
		z[i] = '-'
	} else {
		i++
	}
	return string(z[i:])
}

// Marshal encodes msg as bytes.
//
// It sizes the message first and encodes into an exactly sized buffer, so
// msg.Encode runs twice and must behave deterministically. Custom encoders
// can use Encoder.IsSizing to distinguish the passes.
func Marshal(msg Message) ([]byte, error) {
	enc, err := sizeForMarshal(msg)
	if err != nil || enc.size == 0 {
		return nil, err
	}
	enc.buffer = make([]byte, 0, enc.size)
	msg.Encode(enc)
	if enc.err != nil {
		return nil, enc.err
	}
	return enc.Buffer(), nil
}

// MarshalAppend appends the encoding of msg to buffer and returns the
// extended buffer. On error it returns buffer with its original length.
//
// Like Marshal, it sizes the message first to reserve output space, so
// msg.Encode runs twice and must behave deterministically.
func MarshalAppend(buffer []byte, msg Message) ([]byte, error) {
	return new(Encoder).marshalAppend(buffer, msg)
}

func (enc *Encoder) marshalAppend(buffer []byte, msg Message) ([]byte, error) {
	if err := enc.sizeForMarshal(msg); err != nil {
		return buffer, err
	}
	enc.buffer = slices.Grow(buffer, enc.size)
	msg.Encode(enc)
	if enc.err != nil {
		return buffer, enc.err
	}
	return enc.Buffer(), nil
}

type lengthHints struct {
	widths []byte
	inline [64]byte
	read   int
}

// sizeForMarshal retains prefix-width hints and reuses the sizing encoder for
// the writing pass. Small messages pack their hints into the encoder itself;
// larger messages allocate overflow storage lazily. Size does not collect hints.
func sizeForMarshal(msg Message) (*Encoder, error) {
	enc := new(Encoder)
	if err := enc.sizeForMarshal(msg); err != nil {
		return nil, err
	}
	return enc, nil
}

func (enc *Encoder) sizeForMarshal(msg Message) error {
	if err := ValidateRequired(msg); err != nil {
		return err
	}
	hints := enc.lengthHints
	if hints != nil {
		hints.widths = hints.widths[:0]
		hints.read = 0
	}
	*enc = Encoder{sizing: true, hintCount: 1, lengthHints: hints}
	msg.Encode(enc)
	if enc.err != nil {
		return enc.err
	}
	enc.sizing = false
	return nil
}

// MarshalBuffer encodes msg as bytes with buffer.
func MarshalBuffer(msg Message, buffer []byte) ([]byte, error) {
	if err := ValidateRequired(msg); err != nil {
		return nil, err
	}
	enc := &Encoder{buffer: buffer[:0]}
	msg.Encode(enc)
	if enc.err != nil {
		return nil, enc.err
	}
	return enc.Buffer(), nil
}

// Size returns the encoded size of msg without constructing its wire data.
func Size(msg Message) (int, error) {
	if err := ValidateRequired(msg); err != nil {
		return 0, err
	}
	enc := &Encoder{sizing: true}
	msg.Encode(enc)
	if enc.err != nil {
		return 0, enc.err
	}
	return enc.size, nil
}

// UnmarshalOptions configures decoding. Positive recursion and repeated-element
// limits are capped at math.MaxInt32, sufficient for serialized protobuf messages
// smaller than 2 GiB.
//
// Elements of a repeated message field are allocated in one block per
// field, so retaining a single decoded element keeps the whole block
// alive. Copy elements that outlive the rest of the message.
type UnmarshalOptions struct {
	// AliasInput allows decoded bytes fields to point into data rather than
	// being copied out of it. It avoids an allocation per bytes field, at the
	// cost of tying the decoded message to data: data must outlive the message
	// and must not be modified or reused, otherwise the message changes with
	// it. String fields are always copied.
	AliasInput bool
	// AllowInvalidUTF8 permits string fields containing invalid UTF-8. By
	// default, invalid protobuf strings cause unmarshalling to fail.
	AllowInvalidUTF8 bool

	// MaxInputSize rejects inputs larger than this many bytes. Zero disables
	// the limit.
	MaxInputSize int
	// MaxRecursionDepth limits nested messages. Zero uses the default limit.
	MaxRecursionDepth int
	// MaxRepeatedElements limits the total number of repeated values decoded
	// across the message. Zero disables the limit.
	MaxRepeatedElements int
}

// Unmarshal decodes data into msg.
func (opts UnmarshalOptions) Unmarshal(data []byte, msg Message) error {
	if opts.MaxInputSize > 0 && len(data) > opts.MaxInputSize {
		return parseError{message: "input exceeds maximum size"}
	}
	return opts.unmarshal(data, msg, new(Decoder))
}

// unmarshal initializes caller-provided storage. Ordinary Unmarshal leaves the
// resulting state intact; Decoder.Unmarshal clears it before returning.
func (opts UnmarshalOptions) unmarshal(data []byte, msg Message, dec *Decoder) error {
	*dec = Decoder{
		messageDecodeState: messageDecodeState{buffer: data},
		depthBudget:        protowire.DefaultRecursionLimit,
	}
	dec.aliasInput = opts.AliasInput
	dec.allowInvalidUTF8 = opts.AllowInvalidUTF8
	if opts.MaxRecursionDepth > 0 {
		dec.depthBudget = int32(min(opts.MaxRecursionDepth, math.MaxInt32))
	}
	if opts.MaxRepeatedElements > 0 {
		dec.maxRepeatedElements = int32(min(opts.MaxRepeatedElements, math.MaxInt32))
	}
	dec.Loop(msg.Decode)
	err := dec.Err()
	if err == nil {
		err = ValidateRequired(msg)
	}
	return err
}

// Unmarshal decodes msg as bytes.
func Unmarshal(data []byte, msg Message) error {
	return UnmarshalOptions{}.Unmarshal(data, msg)
}
