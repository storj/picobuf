// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// MapStringString codes a map<string, string>.
func (codec *Codec) MapStringString(field FieldNumber, m *map[string]string) {
	if codec.decode.codec == nil {
		codec.encode.MapStringString(field, m)
	} else {
		codec.decode.MapStringString(field, m)
	}
}

// MapStringString encodes a map<string, string>.
func (enc *Encoder) MapStringString(field FieldNumber, m *map[string]string) {
	for key, val := range *m {
		enc.encodeAnyBytes(field, func() bool {
			enc.String(1, &key)
			enc.String(2, &val)
			return true
		})
	}
}

// MapStringString decodes a map<string, string>.
func (dec *Decoder) MapStringString(field FieldNumber, m *map[string]string) {
	var key, val string
	var filled bool

	dec.RepeatedMessage(field, func(c *Codec, index int) bool {
		if *m == nil {
			// only create a map when there exists at least one entry
			*m = map[string]string{}
		}

		if index == -1 && filled {
			// flush the previous value,
			// we cannot immediately do this, because the entry
			// might have the key, value in a different order
			(*m)[key] = val
			filled = false
		}
		c.String(1, &key)
		c.String(2, &val)
		filled = true
		return true
	})

	if filled {
		(*m)[key] = val
	}
}
