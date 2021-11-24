// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// Message an interface for the generate protobuf message types.
type Message interface {
	// Picobuf returns true when the message is set.
	Picobuf(*Codec) bool
}

// Marshal encodes msg as bytes.
func Marshal(msg Message) ([]byte, error) {
	enc := NewEncoder()
	msg.Picobuf(enc.Codec())
	return enc.Buffer(), nil
}

// Unmarshal decodes msg as bytes.
func Unmarshal(data []byte, msg Message) error {
	dec := NewDecoder(data)
	msg.Picobuf(dec.Codec())
	return dec.Err()
}
