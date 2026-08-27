// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import "errors"

// Optional is an explicit-presence field stored by value, avoiding the
// pointer allocation of a *T field. The zero value is absent.
//
// It is generated for fields declaring (pico.field).value_presence = true.
type Optional[T any] struct {
	value T
	set   bool
}

// NewOptional returns an Optional holding value.
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{value: value, set: true}
}

// Get returns the value and whether it is present.
func (o Optional[T]) Get() (T, bool) { return o.value, o.set }

// GetOr returns the value, or def when absent.
func (o Optional[T]) GetOr(def T) T {
	if o.set {
		return o.value
	}
	return def
}

// Present returns whether a value is present.
func (o Optional[T]) Present() bool { return o.set }

// IsZero returns whether the value is absent, making `json:",omitzero"` omit
// absent fields.
func (o Optional[T]) IsZero() bool { return !o.set }

// Set replaces the value.
func (o *Optional[T]) Set(value T) {
	o.value = value
	o.set = true
}

// Clear removes the value.
func (o *Optional[T]) Clear() {
	*o = Optional[T]{}
}

// Ptr returns a pointer to the value, or nil when absent. Writes through the
// pointer modify the held value.
func (o *Optional[T]) Ptr() *T {
	if !o.set {
		return nil
	}
	return &o.value
}

// PicoSetPtr marks the value present and returns a pointer to it. It is
// intended for generated decoding code.
func (o *Optional[T]) PicoSetPtr() *T {
	o.set = true
	return &o.value
}

// errJSON reports that Optional does not implement encoding/json. Picobuf
// does not implement the ProtoJSON mapping for encoding/json, see README.
var errJSON = errors.New("picobuf: Optional does not support encoding/json")

// MarshalJSON reports an error, so that encoding/json fails instead of
// silently encoding the field as {}.
func (o Optional[T]) MarshalJSON() ([]byte, error) { return nil, errJSON }

// UnmarshalJSON reports an error, see MarshalJSON.
func (o *Optional[T]) UnmarshalJSON(data []byte) error { return errJSON }
