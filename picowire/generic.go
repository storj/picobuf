// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import (
	"cmp"
	"iter"
	"slices"

	"storj.io/picobuf"
)

// sortedMap iterates over map entries in key order, so that encoding is
// deterministic. Single-entry maps are iterated directly to avoid the key
// slice allocation.
func sortedMap[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if len(m) < 2 {
			for key, val := range m {
				if !yield(key, val) {
					return
				}
			}
			return
		}
		keys := make([]K, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		slices.Sort(keys)
		for _, key := range keys {
			if !yield(key, m[key]) {
				return
			}
		}
	}
}

// boolMap iterates over a map with bool keys in key order, false first.
// There are only two keys, so there is nothing to sort.
func boolMap[V any](m map[bool]V) iter.Seq2[bool, V] {
	return func(yield func(bool, V) bool) {
		for _, key := range [2]bool{false, true} {
			if val, ok := m[key]; ok && !yield(key, val) {
				return
			}
		}
	}
}

// MapCodec encodes and decodes a map key or value.
type MapCodec[T any] interface {
	PicoEncode(*picobuf.Encoder, picobuf.FieldNumber, T)
	PicoDecode(*picobuf.Decoder, picobuf.FieldNumber, *T)
	PicoDefault(*T)
}

// MessagePointer is a pointer to a protobuf message implementation.
type MessagePointer[M any] interface {
	*M
	picobuf.Message
}

// MessageCodec encodes and decodes message-valued map entries.
type MessageCodec[M any, P MessagePointer[M]] struct{}

// PicoEncode encodes a message-valued map entry, which is always emitted.
func (MessageCodec[M, P]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber, value P) {
	v := value
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

// PicoDecode decodes a message-valued map entry.
func (MessageCodec[M, P]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber, value *P) {
	dec.Message(field, func(dec *picobuf.Decoder) {
		dec.Loop((*value).Decode)
	})
}

// PicoDefault resets a message-valued map entry to an empty message, which is
// what an omitted value means.
func (MessageCodec[M, P]) PicoDefault(value *P) {
	*value = P(new(M))
}

// EnumCodec encodes and decodes enum-valued map entries.
type EnumCodec[E ~int32] struct{}

// PicoEncode encodes an enum-valued map entry, which is always emitted.
func (EnumCodec[E]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber, value E) {
	v := int32(value)
	enc.AlwaysInt32(field, &v)
}

// PicoDecode decodes an enum-valued map entry.
func (EnumCodec[E]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber, value *E) {
	v := int32(*value)
	dec.Int32(field, &v)
	*value = E(v)
}

// PicoDefault resets an enum-valued map entry to the default.
func (EnumCodec[E]) PicoDefault(value *E) {
	*value = 0
}

// MapOf implements a protobuf map using custom key and value wire codecs.
type MapOf[K cmp.Ordered, V any, KC MapCodec[K], VC MapCodec[V]] map[K]V

// PicoEncode encodes as protobuf, with entries sorted by key.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	// key and val are passed by value, so that they do not escape through the
	// dictionary dispatched codec calls.
	for key, val := range sortedMap(*m) {
		enc.AlwaysAnyBytes(field, func() {
			keyCodec.PicoEncode(enc, 1, key)
			valueCodec.PicoEncode(enc, 2, val)
		})
	}
}

// maxMapSizeHint caps the presized map capacity. The hint counts wire
// entries, which a peer controls and which may repeat the same key, so an
// uncapped hint would let the input dictate the table size.
const maxMapSizeHint = 1024

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	decodeMap[K, V, KC, VC]((*map[K]V)(m), dec, field)
}

// MapOfBool implements a protobuf map with bool keys using a custom value wire
// codec. It is separate from MapOf, so that MapOf can order its keys with
// cmp.Ordered, which excludes bool.
type MapOfBool[V any, VC MapCodec[V]] map[bool]V

// PicoEncode encodes as protobuf, with entries sorted by key.
//
//go:noinline
func (m *MapOfBool[V, VC]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	var valueCodec VC
	for key, val := range boolMap(*m) {
		enc.AlwaysAnyBytes(field, func() {
			enc.AlwaysBool(1, &key)
			valueCodec.PicoEncode(enc, 2, val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapOfBool[V, VC]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	decodeMap[bool, V, BoolCodec, VC]((*map[bool]V)(m), dec, field)
}

// decodeMap decodes map entries into m with the key and value codecs.
func decodeMap[K comparable, V any, KC MapCodec[K], VC MapCodec[V]](m *map[K]V, dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	// key and val are hoisted out of the entry, because the codec calls are
	// dictionary dispatched, so &key and &val always escape; this costs one
	// allocation per map field instead of one per entry.
	var key K
	var val V
	if *m == nil {
		if n := dec.RepeatedMessageCount(field); n > 0 {
			*m = make(map[K]V, min(n, maxMapSizeHint))
		}
	}
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[K]V{}
		}
		// Each entry starts from the default, because the key or the value may
		// be omitted from the entry.
		keyCodec.PicoDefault(&key)
		valueCodec.PicoDefault(&val)
		c.Loop(func(c *picobuf.Decoder) {
			keyCodec.PicoDecode(c, 1, &key)
			valueCodec.PicoDecode(c, 2, &val)
		})
		(*m)[key] = val
	})
}
