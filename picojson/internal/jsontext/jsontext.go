// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

//go:build go1.27

// Package jsontext selects encoding/json/jsontext from Go 1.27, and
// github.com/go-json-experiment/json/jsontext, where it originates, before.
package jsontext

import "encoding/json/jsontext"

type (
	// Decoder is jsontext.Decoder.
	Decoder = jsontext.Decoder
	// Kind is jsontext.Kind.
	Kind = jsontext.Kind
	// Token is jsontext.Token.
	Token = jsontext.Token
	// Value is jsontext.Value.
	Value = jsontext.Value
)

// NewDecoder is jsontext.NewDecoder.
var NewDecoder = jsontext.NewDecoder

// AppendQuote is jsontext.AppendQuote.
func AppendQuote[Bytes ~[]byte | ~string](dst []byte, src Bytes) ([]byte, error) {
	return jsontext.AppendQuote(dst, src)
}
