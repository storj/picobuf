// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import "storj.io/picobuf"

// MapStringString implements map<string, string>.
type MapStringString map[string]string

// PicoEncode encodes a map<string, string>.
//go:noinline
func (m *MapStringString) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes a map<string, string>.
//go:noinline
func (m *MapStringString) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key, val string
	var filled bool

	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			// only create a map when there exists at least one entry
			*m = map[string]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			if filled {
				(*m)[key] = val
				filled = false
			}
			c.String(1, &key)
			c.String(2, &val)
			filled = true
		})
	})

	if filled {
		(*m)[key] = val
	}
}
