// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"errors"
	"math"
	"slices"
	"sync"

	"storj.io/picobuf/internal/protowire"
)

// Message is an interface that all generated messages implement.
type Message interface {
	Encode(*Encoder) bool
	Decode(*Decoder)
}

// RequiredFields is the interface that generated messages implement to
// validate legacy required fields. depth is the nesting depth of the message,
// zero for the root. A validator passes it on to ValidateRequiredDepth for its
// children.
type RequiredFields interface {
	PicoValidateRequiredDepth(depth int) error
}

// ValidateRequired validates all legacy required fields known to msg.
func ValidateRequired(msg any) error {
	return ValidateRequiredDepth(msg, 0)
}

// ValidateRequiredDepth validates a child at the given nesting depth.
// Generated validators call it so that one depth budget covers the whole
// message graph. The root message has depth zero.
func ValidateRequiredDepth(msg any, depth int) error {
	if depth > protowire.DefaultRecursionLimit {
		return errors.New("exceeded maximum recursion depth while validating required fields")
	}
	if required, ok := msg.(RequiredFields); ok {
		return required.PicoValidateRequiredDepth(depth)
	}
	return nil
}

// MissingRequiredField returns an error for an unset legacy required field.
func MissingRequiredField(field FieldNumber) error {
	return errors.New("required field " + field.String() + " is not set")
}

// CustomType is the interface that a custom type implements to encode and
// decode itself.
type CustomType interface {
	PicoEncode(*Encoder, FieldNumber)
	PicoDecode(*Decoder, FieldNumber)
}

// FieldNumber is a protobuf field number.
type FieldNumber protowire.Number

// IsValid reports whether the field number is in the valid range.
func (field FieldNumber) IsValid() bool { return protowire.Number(field).IsValid() }

// String formats the field number in decimal.
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

// Marshal encodes msg.
//
// It sizes the message first and then encodes into a buffer of that size, so
// msg.Encode runs twice and must do the same thing both times. A custom
// encoder can tell the passes apart with Encoder.IsSizing.
func Marshal(msg Message) ([]byte, error) {
	// The encoder is returned to the pool without defer, which costs several
	// kilobytes in TinyGo. If msg.Encode panics, the encoder is dropped.
	enc := getEncoder()
	data, err := enc.marshal(msg)
	putEncoder(enc)
	return data, err
}

func (enc *Encoder) marshal(msg Message) ([]byte, error) {
	if err := enc.sizeForMarshal(msg); err != nil || enc.size == 0 {
		return nil, err
	}
	enc.buffer = make([]byte, 0, enc.size)
	msg.Encode(enc)
	if enc.err != nil {
		return nil, enc.err
	}
	return enc.Buffer(), nil
}

// MarshalAppend appends the encoding of msg to buffer and returns the extended
// buffer. On error it returns buffer with its original length.
//
// Like Marshal, it sizes the message first to reserve the output space, so
// msg.Encode runs twice and must do the same thing both times.
func MarshalAppend(buffer []byte, msg Message) ([]byte, error) {
	enc := getEncoder()
	data, err := enc.marshalAppend(buffer, msg)
	putEncoder(enc)
	return data, err
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
	// The sizing pass encoded the same strings, so the writing pass does not
	// need to validate them again.
	enc.validatedUTF8 = true
	return nil
}

// Size returns the encoded size of msg without constructing its wire data.
func Size(msg Message) (int, error) {
	if err := ValidateRequired(msg); err != nil {
		return 0, err
	}
	enc := getEncoder()
	enc.sizing = true
	msg.Encode(enc)
	size, err := enc.size, error(nil)
	if enc.err != nil {
		size, err = 0, enc.err
	}
	putEncoder(enc)
	return size, err
}

// UnmarshalOptions configures decoding. The decoder caps positive recursion
// and repeated-element limits at math.MaxInt32. That is enough, because a
// serialized protobuf message is smaller than 2 GiB.
//
// The decoder allocates the elements of a repeated message field together in
// blocks of up to 4 KiB, so keeping one decoded element keeps the rest of its
// block alive. Copy the elements that outlive the rest of the message.
type UnmarshalOptions struct {
	// AliasInput lets decoded bytes fields point into data instead of holding a
	// copy. That saves an allocation per bytes field, but it ties the decoded
	// message to data. Keep data alive and unchanged for as long as the message
	// is in use, otherwise the message changes with it. The decoder always copies
	// string fields.
	AliasInput bool
	// AllowInvalidUTF8 permits string fields that contain invalid UTF-8. By
	// default such a string fails unmarshaling.
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

// Unmarshal decodes data into msg. It merges into msg and does not reset it
// first.
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

// Unmarshal decodes data into msg. It merges into msg and does not reset it
// first.
func Unmarshal(data []byte, msg Message) error {
	return UnmarshalOptions{}.Unmarshal(data, msg)
}

// encoderPool keeps encoders, and the storage for their length prefix hints,
// between calls to Marshal, MarshalAppend and Size. Otherwise every call
// allocates the encoder, which escapes through msg.Encode, and regrows the
// hints of a message with many nested fields.
var encoderPool = sync.Pool{New: func() any { return new(Encoder) }}

// maxPooledLengthHints limits the hint storage that goes back into the pool,
// so that a single huge message does not keep its storage alive.
const maxPooledLengthHints = 64 << 10

func getEncoder() *Encoder { return encoderPool.Get().(*Encoder) }

func putEncoder(enc *Encoder) {
	enc.reset()
	if hints := enc.lengthHints; hints != nil && cap(hints.widths) > maxPooledLengthHints {
		enc.lengthHints = nil
	}
	encoderPool.Put(enc)
}
