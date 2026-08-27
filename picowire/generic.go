// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import (
	"cmp"
	"iter"
	"reflect"
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

// sortedMapFunc is sortedMap for key types that cmp.Ordered cannot express,
// such as bool and named types.
func sortedMapFunc[K MapKey, V any](m map[K]V, compare func(a, b K) int) iter.Seq2[K, V] {
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
		slices.SortFunc(keys, compare)
		for _, key := range keys {
			if !yield(key, m[key]) {
				return
			}
		}
	}
}

// compareBool orders false before true, matching numeric key order.
func compareBool(a, b bool) int {
	switch {
	case a == b:
		return 0
	case b:
		return -1
	default:
		return 1
	}
}

// compareMapKeys orders map keys by value, so that map encoding is
// deterministic. Named key types prevent a cmp.Ordered constraint, hence the
// reflection.
func compareMapKeys[K MapKey](a, b K) int {
	av, bv := reflect.ValueOf(a), reflect.ValueOf(b)
	switch av.Kind() {
	case reflect.Bool:
		return compareBool(av.Bool(), bv.Bool())
	case reflect.Int32, reflect.Int64:
		return cmp.Compare(av.Int(), bv.Int())
	case reflect.Uint32, reflect.Uint64:
		return cmp.Compare(av.Uint(), bv.Uint())
	default:
		return cmp.Compare(av.String(), bv.String())
	}
}

// MapKey is a protobuf map key.
type MapKey interface {
	~bool | ~int32 | ~int64 | ~uint32 | ~uint64 | ~string
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
type MapOf[K MapKey, V any, KC MapCodec[K], VC MapCodec[V]] map[K]V

// PicoEncode encodes as protobuf, with entries sorted by key.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	// key and val are passed by value, so that they do not escape through the
	// dictionary dispatched codec calls.
	for key, val := range sortedMapFunc(*m, compareMapKeys[K]) {
		enc.AlwaysAnyBytes(field, func() {
			keyCodec.PicoEncode(enc, 1, key)
			valueCodec.PicoEncode(enc, 2, val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapOf[K, V, KC, VC]) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var keyCodec KC
	var valueCodec VC
	// key and val are hoisted out of the entry, because the codec calls are
	// dictionary dispatched, so &key and &val always escape; this costs one
	// allocation per map field instead of one per entry.
	var key K
	var val V
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
