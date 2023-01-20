// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import "storj.io/picobuf"

// MapKey is a protobuf map key.
type MapKey interface {
	~bool | ~int32 | ~int64 | ~uint32 | ~uint64 | ~string
}

// MapCodec encodes and decodes a map key or value.
// A value codec may also implement PicoInit(*T) to initialize the protobuf
// default for each entry, including entries that omit the value field.
type MapCodec[T any] interface {
	PicoEncode(*picobuf.Encoder, picobuf.FieldNumber, *T)
	PicoDecode(*picobuf.Decoder, picobuf.FieldNumber, *T)
}

// MessagePointer is a pointer to a protobuf message implementation.
type MessagePointer[M any] interface {
	*M
	picobuf.Message
}

// MessageCodec encodes and decodes message-valued map entries.
type MessageCodec[M any, P MessagePointer[M]] struct{}

// PicoEncode encodes a message-valued map entry.
func (MessageCodec[M, P]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber, value *P) {
	v := *value
	if v == nil {
		v = P(new(M))
		// A nil map value represents an empty message, whose required fields
		// must be checked even though the original nil pointer passed validation.
		if err := picobuf.ValidateRequired(v); err != nil {
			enc.Fail(field, err.Error())
			return
		}
	}
	enc.AlwaysMessage(field, v.Encode)
}

// PicoInit initializes an absent map value to an empty message.
func (MessageCodec[M, P]) PicoInit(value *P) {
	if *value == nil {
		*value = P(new(M))
	}
}

// PicoDecode decodes a message-valued map entry.
func (codec MessageCodec[M, P]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber, value *P) {
	dec.Message(field, func(dec *picobuf.Decoder) {
		codec.PicoInit(value)
		dec.Loop((*value).Decode)
	})
}

// EnumCodec encodes and decodes enum-valued map entries.
type EnumCodec[E ~int32] struct{}

// PicoEncode encodes an enum-valued map entry.
func (EnumCodec[E]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber, value *E) {
	// *E cannot be converted to *int32, when E is a type parameter.
	v := int32(*value)
	enc.Int32(field, &v)
}

// PicoDecode decodes an enum-valued map entry.
func (EnumCodec[E]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber, value *E) {
	v := int32(*value)
	dec.Int32(field, &v)
	*value = E(v)
}

// MapOf implements a protobuf map using custom key and value wire codecs.
type MapOf[K MapKey, V any, KC MapCodec[K], VC MapCodec[V]] map[K]V

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	// key and val are hoisted out of the loop, because the codec calls are
	// dictionary dispatched, so &key and &val always escape; this costs one
	// allocation per map field instead of one per entry.
	var key K
	var val V
	for k, v := range *m {
		key, val = k, v
		enc.AlwaysAnyBytes(field, func() {
			keyCodec.PicoEncode(enc, 1, &key)
			valueCodec.PicoEncode(enc, 2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	initializer, hasInitializer := any(valueCodec).(interface{ PicoInit(*V) })
	// key and val are hoisted, see PicoEncode.
	var key K
	var val V
	var zeroK K
	var zeroV V
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[K]V{}
		}
		key, val = zeroK, zeroV
		if hasInitializer {
			initializer.PicoInit(&val)
		}
		c.Loop(func(c *picobuf.Decoder) {
			keyCodec.PicoDecode(c, 1, &key)
			valueCodec.PicoDecode(c, 2, &val)
		})
		(*m)[key] = val
	})
}
