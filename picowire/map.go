// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import "storj.io/picobuf"

// MapBoolBool implements map<bool,bool>.
type MapBoolBool map[bool]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolBool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolBool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolInt32 implements map<bool,int32>.
type MapBoolInt32 map[bool]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolInt32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolInt32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolInt64 implements map<bool,int64>.
type MapBoolInt64 map[bool]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolInt64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolInt64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolUint32 implements map<bool,uint32>.
type MapBoolUint32 map[bool]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolUint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolUint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolUint64 implements map<bool,uint64>.
type MapBoolUint64 map[bool]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolUint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolUint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolSint32 implements map<bool,sint32>.
type MapBoolSint32 map[bool]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolSint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolSint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolSint64 implements map<bool,sint64>.
type MapBoolSint64 map[bool]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolSint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolSint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolFixed32 implements map<bool,fixed32>.
type MapBoolFixed32 map[bool]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolFixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolFixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolFixed64 implements map<bool,fixed64>.
type MapBoolFixed64 map[bool]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolFixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolFixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolSfixed32 implements map<bool,sfixed32>.
type MapBoolSfixed32 map[bool]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolSfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolSfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolSfixed64 implements map<bool,sfixed64>.
type MapBoolSfixed64 map[bool]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolSfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolSfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolFloat implements map<bool,float>.
type MapBoolFloat map[bool]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolFloat) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolFloat) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolDouble implements map<bool,double>.
type MapBoolDouble map[bool]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolDouble) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolDouble) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolString implements map<bool,string>.
type MapBoolString map[bool]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolString) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolString) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapBoolBytes implements map<bool,bytes>.
type MapBoolBytes map[bool][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapBoolBytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Bool(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapBoolBytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key bool
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[bool][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Bool(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Bool implements map<int32,bool>.
type MapInt32Bool map[int32]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Int32 implements map<int32,int32>.
type MapInt32Int32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Int64 implements map<int32,int64>.
type MapInt32Int64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Uint32 implements map<int32,uint32>.
type MapInt32Uint32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Uint64 implements map<int32,uint64>.
type MapInt32Uint64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Sint32 implements map<int32,sint32>.
type MapInt32Sint32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Sint64 implements map<int32,sint64>.
type MapInt32Sint64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Fixed32 implements map<int32,fixed32>.
type MapInt32Fixed32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Fixed64 implements map<int32,fixed64>.
type MapInt32Fixed64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Sfixed32 implements map<int32,sfixed32>.
type MapInt32Sfixed32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Sfixed64 implements map<int32,sfixed64>.
type MapInt32Sfixed64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Float implements map<int32,float>.
type MapInt32Float map[int32]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Double implements map<int32,double>.
type MapInt32Double map[int32]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32String implements map<int32,string>.
type MapInt32String map[int32]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt32Bytes implements map<int32,bytes>.
type MapInt32Bytes map[int32][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt32Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int32(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt32Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int32(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Bool implements map<int64,bool>.
type MapInt64Bool map[int64]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Int32 implements map<int64,int32>.
type MapInt64Int32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Int64 implements map<int64,int64>.
type MapInt64Int64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Uint32 implements map<int64,uint32>.
type MapInt64Uint32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Uint64 implements map<int64,uint64>.
type MapInt64Uint64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Sint32 implements map<int64,sint32>.
type MapInt64Sint32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Sint64 implements map<int64,sint64>.
type MapInt64Sint64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Fixed32 implements map<int64,fixed32>.
type MapInt64Fixed32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Fixed64 implements map<int64,fixed64>.
type MapInt64Fixed64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Sfixed32 implements map<int64,sfixed32>.
type MapInt64Sfixed32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Sfixed64 implements map<int64,sfixed64>.
type MapInt64Sfixed64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Float implements map<int64,float>.
type MapInt64Float map[int64]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Double implements map<int64,double>.
type MapInt64Double map[int64]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64String implements map<int64,string>.
type MapInt64String map[int64]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapInt64Bytes implements map<int64,bytes>.
type MapInt64Bytes map[int64][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapInt64Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Int64(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapInt64Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Int64(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Bool implements map<uint32,bool>.
type MapUint32Bool map[uint32]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Int32 implements map<uint32,int32>.
type MapUint32Int32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Int64 implements map<uint32,int64>.
type MapUint32Int64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Uint32 implements map<uint32,uint32>.
type MapUint32Uint32 map[uint32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Uint64 implements map<uint32,uint64>.
type MapUint32Uint64 map[uint32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Sint32 implements map<uint32,sint32>.
type MapUint32Sint32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Sint64 implements map<uint32,sint64>.
type MapUint32Sint64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Fixed32 implements map<uint32,fixed32>.
type MapUint32Fixed32 map[uint32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Fixed64 implements map<uint32,fixed64>.
type MapUint32Fixed64 map[uint32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Sfixed32 implements map<uint32,sfixed32>.
type MapUint32Sfixed32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Sfixed64 implements map<uint32,sfixed64>.
type MapUint32Sfixed64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Float implements map<uint32,float>.
type MapUint32Float map[uint32]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Double implements map<uint32,double>.
type MapUint32Double map[uint32]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32String implements map<uint32,string>.
type MapUint32String map[uint32]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint32Bytes implements map<uint32,bytes>.
type MapUint32Bytes map[uint32][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint32Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint32(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint32Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint32(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Bool implements map<uint64,bool>.
type MapUint64Bool map[uint64]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Int32 implements map<uint64,int32>.
type MapUint64Int32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Int64 implements map<uint64,int64>.
type MapUint64Int64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Uint32 implements map<uint64,uint32>.
type MapUint64Uint32 map[uint64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Uint64 implements map<uint64,uint64>.
type MapUint64Uint64 map[uint64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Sint32 implements map<uint64,sint32>.
type MapUint64Sint32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Sint64 implements map<uint64,sint64>.
type MapUint64Sint64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Fixed32 implements map<uint64,fixed32>.
type MapUint64Fixed32 map[uint64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Fixed64 implements map<uint64,fixed64>.
type MapUint64Fixed64 map[uint64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Sfixed32 implements map<uint64,sfixed32>.
type MapUint64Sfixed32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Sfixed64 implements map<uint64,sfixed64>.
type MapUint64Sfixed64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Float implements map<uint64,float>.
type MapUint64Float map[uint64]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Double implements map<uint64,double>.
type MapUint64Double map[uint64]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64String implements map<uint64,string>.
type MapUint64String map[uint64]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapUint64Bytes implements map<uint64,bytes>.
type MapUint64Bytes map[uint64][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapUint64Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Uint64(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapUint64Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Uint64(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Bool implements map<sint32,bool>.
type MapSint32Bool map[int32]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Int32 implements map<sint32,int32>.
type MapSint32Int32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Int64 implements map<sint32,int64>.
type MapSint32Int64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Uint32 implements map<sint32,uint32>.
type MapSint32Uint32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Uint64 implements map<sint32,uint64>.
type MapSint32Uint64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Sint32 implements map<sint32,sint32>.
type MapSint32Sint32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Sint64 implements map<sint32,sint64>.
type MapSint32Sint64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Fixed32 implements map<sint32,fixed32>.
type MapSint32Fixed32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Fixed64 implements map<sint32,fixed64>.
type MapSint32Fixed64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Sfixed32 implements map<sint32,sfixed32>.
type MapSint32Sfixed32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Sfixed64 implements map<sint32,sfixed64>.
type MapSint32Sfixed64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Float implements map<sint32,float>.
type MapSint32Float map[int32]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Double implements map<sint32,double>.
type MapSint32Double map[int32]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32String implements map<sint32,string>.
type MapSint32String map[int32]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint32Bytes implements map<sint32,bytes>.
type MapSint32Bytes map[int32][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint32Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint32(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint32Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint32(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Bool implements map<sint64,bool>.
type MapSint64Bool map[int64]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Int32 implements map<sint64,int32>.
type MapSint64Int32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Int64 implements map<sint64,int64>.
type MapSint64Int64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Uint32 implements map<sint64,uint32>.
type MapSint64Uint32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Uint64 implements map<sint64,uint64>.
type MapSint64Uint64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Sint32 implements map<sint64,sint32>.
type MapSint64Sint32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Sint64 implements map<sint64,sint64>.
type MapSint64Sint64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Fixed32 implements map<sint64,fixed32>.
type MapSint64Fixed32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Fixed64 implements map<sint64,fixed64>.
type MapSint64Fixed64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Sfixed32 implements map<sint64,sfixed32>.
type MapSint64Sfixed32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Sfixed64 implements map<sint64,sfixed64>.
type MapSint64Sfixed64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Float implements map<sint64,float>.
type MapSint64Float map[int64]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Double implements map<sint64,double>.
type MapSint64Double map[int64]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64String implements map<sint64,string>.
type MapSint64String map[int64]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSint64Bytes implements map<sint64,bytes>.
type MapSint64Bytes map[int64][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSint64Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sint64(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSint64Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sint64(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Bool implements map<fixed32,bool>.
type MapFixed32Bool map[uint32]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Int32 implements map<fixed32,int32>.
type MapFixed32Int32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Int64 implements map<fixed32,int64>.
type MapFixed32Int64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Uint32 implements map<fixed32,uint32>.
type MapFixed32Uint32 map[uint32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Uint64 implements map<fixed32,uint64>.
type MapFixed32Uint64 map[uint32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Sint32 implements map<fixed32,sint32>.
type MapFixed32Sint32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Sint64 implements map<fixed32,sint64>.
type MapFixed32Sint64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Fixed32 implements map<fixed32,fixed32>.
type MapFixed32Fixed32 map[uint32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Fixed64 implements map<fixed32,fixed64>.
type MapFixed32Fixed64 map[uint32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Sfixed32 implements map<fixed32,sfixed32>.
type MapFixed32Sfixed32 map[uint32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Sfixed64 implements map<fixed32,sfixed64>.
type MapFixed32Sfixed64 map[uint32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Float implements map<fixed32,float>.
type MapFixed32Float map[uint32]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Double implements map<fixed32,double>.
type MapFixed32Double map[uint32]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32String implements map<fixed32,string>.
type MapFixed32String map[uint32]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed32Bytes implements map<fixed32,bytes>.
type MapFixed32Bytes map[uint32][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed32Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed32(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed32Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint32
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint32][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed32(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Bool implements map<fixed64,bool>.
type MapFixed64Bool map[uint64]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Int32 implements map<fixed64,int32>.
type MapFixed64Int32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Int64 implements map<fixed64,int64>.
type MapFixed64Int64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Uint32 implements map<fixed64,uint32>.
type MapFixed64Uint32 map[uint64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Uint64 implements map<fixed64,uint64>.
type MapFixed64Uint64 map[uint64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Sint32 implements map<fixed64,sint32>.
type MapFixed64Sint32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Sint64 implements map<fixed64,sint64>.
type MapFixed64Sint64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Fixed32 implements map<fixed64,fixed32>.
type MapFixed64Fixed32 map[uint64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Fixed64 implements map<fixed64,fixed64>.
type MapFixed64Fixed64 map[uint64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Sfixed32 implements map<fixed64,sfixed32>.
type MapFixed64Sfixed32 map[uint64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Sfixed64 implements map<fixed64,sfixed64>.
type MapFixed64Sfixed64 map[uint64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Float implements map<fixed64,float>.
type MapFixed64Float map[uint64]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Double implements map<fixed64,double>.
type MapFixed64Double map[uint64]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64String implements map<fixed64,string>.
type MapFixed64String map[uint64]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapFixed64Bytes implements map<fixed64,bytes>.
type MapFixed64Bytes map[uint64][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapFixed64Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Fixed64(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapFixed64Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key uint64
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[uint64][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Fixed64(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Bool implements map<sfixed32,bool>.
type MapSfixed32Bool map[int32]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Int32 implements map<sfixed32,int32>.
type MapSfixed32Int32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Int64 implements map<sfixed32,int64>.
type MapSfixed32Int64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Uint32 implements map<sfixed32,uint32>.
type MapSfixed32Uint32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Uint64 implements map<sfixed32,uint64>.
type MapSfixed32Uint64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Sint32 implements map<sfixed32,sint32>.
type MapSfixed32Sint32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Sint64 implements map<sfixed32,sint64>.
type MapSfixed32Sint64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Fixed32 implements map<sfixed32,fixed32>.
type MapSfixed32Fixed32 map[int32]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Fixed64 implements map<sfixed32,fixed64>.
type MapSfixed32Fixed64 map[int32]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Sfixed32 implements map<sfixed32,sfixed32>.
type MapSfixed32Sfixed32 map[int32]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Sfixed64 implements map<sfixed32,sfixed64>.
type MapSfixed32Sfixed64 map[int32]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Float implements map<sfixed32,float>.
type MapSfixed32Float map[int32]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Double implements map<sfixed32,double>.
type MapSfixed32Double map[int32]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32String implements map<sfixed32,string>.
type MapSfixed32String map[int32]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed32Bytes implements map<sfixed32,bytes>.
type MapSfixed32Bytes map[int32][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed32(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed32Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int32
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int32][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed32(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Bool implements map<sfixed64,bool>.
type MapSfixed64Bool map[int64]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Bool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Bool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Int32 implements map<sfixed64,int32>.
type MapSfixed64Int32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Int32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Int32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Int64 implements map<sfixed64,int64>.
type MapSfixed64Int64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Int64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Int64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Uint32 implements map<sfixed64,uint32>.
type MapSfixed64Uint32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Uint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Uint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Uint64 implements map<sfixed64,uint64>.
type MapSfixed64Uint64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Uint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Uint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Sint32 implements map<sfixed64,sint32>.
type MapSfixed64Sint32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Sint64 implements map<sfixed64,sint64>.
type MapSfixed64Sint64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Fixed32 implements map<sfixed64,fixed32>.
type MapSfixed64Fixed32 map[int64]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Fixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Fixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Fixed64 implements map<sfixed64,fixed64>.
type MapSfixed64Fixed64 map[int64]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Fixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Fixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Sfixed32 implements map<sfixed64,sfixed32>.
type MapSfixed64Sfixed32 map[int64]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Sfixed64 implements map<sfixed64,sfixed64>.
type MapSfixed64Sfixed64 map[int64]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Sfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Float implements map<sfixed64,float>.
type MapSfixed64Float map[int64]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Float) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Float) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Double implements map<sfixed64,double>.
type MapSfixed64Double map[int64]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Double) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Double) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64String implements map<sfixed64,string>.
type MapSfixed64String map[int64]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64String) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64String) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapSfixed64Bytes implements map<sfixed64,bytes>.
type MapSfixed64Bytes map[int64][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Bytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.Sfixed64(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapSfixed64Bytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key int64
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[int64][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.Sfixed64(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringBool implements map<string,bool>.
type MapStringBool map[string]bool

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringBool) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Bool(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringBool) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val bool
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]bool{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Bool(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringInt32 implements map<string,int32>.
type MapStringInt32 map[string]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringInt32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Int32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringInt32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Int32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringInt64 implements map<string,int64>.
type MapStringInt64 map[string]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringInt64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Int64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringInt64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Int64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringUint32 implements map<string,uint32>.
type MapStringUint32 map[string]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringUint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Uint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringUint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Uint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringUint64 implements map<string,uint64>.
type MapStringUint64 map[string]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringUint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Uint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringUint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Uint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringSint32 implements map<string,sint32>.
type MapStringSint32 map[string]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringSint32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Sint32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringSint32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Sint32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringSint64 implements map<string,sint64>.
type MapStringSint64 map[string]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringSint64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Sint64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringSint64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Sint64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringFixed32 implements map<string,fixed32>.
type MapStringFixed32 map[string]uint32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringFixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Fixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringFixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val uint32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]uint32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Fixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringFixed64 implements map<string,fixed64>.
type MapStringFixed64 map[string]uint64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringFixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Fixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringFixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val uint64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]uint64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Fixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringSfixed32 implements map<string,sfixed32>.
type MapStringSfixed32 map[string]int32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringSfixed32) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Sfixed32(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringSfixed32) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Sfixed32(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringSfixed64 implements map<string,sfixed64>.
type MapStringSfixed64 map[string]int64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringSfixed64) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Sfixed64(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringSfixed64) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val int64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]int64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Sfixed64(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringFloat implements map<string,float>.
type MapStringFloat map[string]float32

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringFloat) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Float(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringFloat) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val float32
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]float32{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Float(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringDouble implements map<string,double>.
type MapStringDouble map[string]float64

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringDouble) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Double(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringDouble) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val float64
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]float64{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Double(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringString implements map<string,string>.
type MapStringString map[string]string

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringString) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.String(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringString) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val string
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string]string{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.String(2, &val)
		})
		(*m)[key] = val
	})
}

// MapStringBytes implements map<string,bytes>.
type MapStringBytes map[string][]byte

// PicoEncode encodes as protobuf.
//
//go:noinline
func (m *MapStringBytes) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {
	for key, val := range *m {
		enc.AlwaysAnyBytes(field, func() {
			enc.String(1, &key)
			enc.Bytes(2, &val)
		})
	}
}

// PicoDecode decodes as protobuf.
//
//go:noinline
func (m *MapStringBytes) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {
	var key string
	var val []byte
	dec.RepeatedMessage(field, func(c *picobuf.Decoder) {
		if *m == nil {
			*m = map[string][]byte{}
		}
		c.Loop(func(c *picobuf.Decoder) {
			c.String(1, &key)
			c.Bytes(2, &val)
		})
		(*m)[key] = val
	})
}
