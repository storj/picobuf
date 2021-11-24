// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"testing"

	"github.com/zeebo/assert"
	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
)

func TestEncoding(t *testing.T) {
	tests := []TypesPico{
		{},
		{
			Int32:    1,
			Int64:    1,
			Uint32:   1,
			Uint64:   1,
			Sint32:   1,
			Sint64:   1,
			Fixed32:  1,
			Fixed64:  1,
			Sfixed32: 1,
			Sfixed64: 1,
			Float:    1,
			Double:   1,
			Bool:     true,
			String_:  "1",
			Bytes:    []byte{1},
			Message: &MessagePico{
				Int32: 1,
			},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p Types
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		canonical, err := proto.Marshal(&p)
		assert.NoError(t, err)
		assert.Equal(t, canonical, data)
	}
}
