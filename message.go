// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"fmt"
	"math"

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

// CustomType defines methods that are used for custom encode or decode behaviors.
type CustomType interface {
	PicoEncode(*Encoder, FieldNumber)
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
func Marshal(msg Message) ([]byte, error) {
	if err := ValidateRequired(msg); err != nil {
		return nil, err
	}
	enc := &Encoder{}
	msg.Encode(enc)
	if enc.err != nil {
		return nil, enc.err
	}
	return enc.Buffer(), nil
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

// UnmarshalOptions configures decoding.
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
	dec := newDecoder(data)
	dec.aliasInput = opts.AliasInput
	dec.allowInvalidUTF8 = opts.AllowInvalidUTF8
	if opts.MaxRecursionDepth > 0 {
		dec.maxRecursionDepth = opts.MaxRecursionDepth
	}
	dec.maxRepeatedElements = opts.MaxRepeatedElements
	dec.Loop(msg.Decode)
	if dec.err == nil {
		dec.err = ValidateRequired(msg)
	}
	return dec.err
}

// Unmarshal decodes msg as bytes.
func Unmarshal(data []byte, msg Message) error {
	return UnmarshalOptions{}.Unmarshal(data, msg)
}
