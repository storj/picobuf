// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"storj.io/picobuf/internal/protowire"
	"storj.io/picobuf/picojson/internal/jsontext"
)

const (
	// Range of google.protobuf.Timestamp, 0001-01-01 to 9999-12-31.
	minTimestampSeconds = -62135596800
	maxTimestampSeconds = 253402300799
	// Range of google.protobuf.Duration, about 10000 years.
	maxDurationSeconds = 315576000000
)

func isWrapper(name string) bool {
	switch name {
	case "google.protobuf.DoubleValue", "google.protobuf.FloatValue",
		"google.protobuf.Int64Value", "google.protobuf.UInt64Value",
		"google.protobuf.Int32Value", "google.protobuf.UInt32Value",
		"google.protobuf.BoolValue", "google.protobuf.StringValue",
		"google.protobuf.BytesValue":
		return true
	}
	return false
}

// wellKnown renders the well-known types that have a special JSON form.
// google.protobuf.Empty is an ordinary message.
func (m *marshaler) wellKnown(msg *message, wire []byte) (bool, error) {
	switch {
	case isWrapper(msg.Name):
		f := msg.first()
		v := wireValue{typ: wireType(f.Kind)}
		for rest := wire; len(rest) > 0; {
			num, x, n := consumeValue(rest)
			if n < 0 {
				return true, errSyntax
			}
			rest = rest[n:]
			if num == 1 && matchesWire(&f, x) {
				v = x
			}
		}
		return true, m.value(&f, v)
	case msg.Name == "google.protobuf.Timestamp":
		seconds, nanos, err := secondsNanos(wire)
		if err != nil {
			return true, err
		}
		if seconds < minTimestampSeconds || seconds > maxTimestampSeconds || nanos < 0 || nanos >= 1e9 {
			return true, fmt.Errorf("%s out of range: %ds %dns", msg.Name, seconds, nanos)
		}
		text := time.Unix(seconds, int64(nanos)).UTC().Format("2006-01-02T15:04:05.000000000")
		text = trimNanos(text) + "Z"
		m.out, err = jsontext.AppendQuote(m.out, text)
		return true, err
	case msg.Name == "google.protobuf.Duration":
		seconds, nanos, err := secondsNanos(wire)
		if err != nil {
			return true, err
		}
		if seconds < -maxDurationSeconds || seconds > maxDurationSeconds ||
			nanos <= -1e9 || nanos >= 1e9 || (seconds > 0 && nanos < 0) || (seconds < 0 && nanos > 0) {
			return true, fmt.Errorf("%s out of range: %ds %dns", msg.Name, seconds, nanos)
		}
		sign := ""
		if seconds < 0 || nanos < 0 {
			sign, seconds, nanos = "-", -seconds, -nanos
		}
		text := fmt.Sprintf("%s%d.%09d", sign, seconds, nanos)
		m.out, err = jsontext.AppendQuote(m.out, trimNanos(text)+"s")
		return true, err
	case msg.Name == "google.protobuf.FieldMask":
		var paths []string
		for rest := wire; len(rest) > 0; {
			num, v, n := consumeValue(rest)
			if n < 0 {
				return true, errSyntax
			}
			rest = rest[n:]
			if num != 1 || v.typ != protowire.BytesType {
				continue
			}
			path, err := snakeToCamel(string(v.buf))
			if err != nil {
				return true, err
			}
			paths = append(paths, path)
		}
		var err error
		m.out, err = jsontext.AppendQuote(m.out, strings.Join(paths, ","))
		return true, err
	}
	return false, nil
}

func (u *unmarshaler) wellKnown(msg *message, out []byte) (bool, []byte, error) {
	switch {
	case isWrapper(msg.Name):
		f := msg.first()
		out, err := u.value(&f, out)
		return true, out, err
	case msg.Name == "google.protobuf.Timestamp":
		tok, err := u.read('"')
		if err != nil {
			return true, nil, err
		}
		text := tok.String()
		t, err := time.Parse(time.RFC3339Nano, text)
		// Like protojson, limit the fraction to nine digits.
		dot, zone := strings.LastIndexByte(text, '.'), strings.LastIndexAny(text, "Z-+")
		if err != nil || t.Unix() < minTimestampSeconds || t.Unix() > maxTimestampSeconds ||
			(dot >= 0 && zone >= dot && zone-dot > len(".999999999")) {
			return true, nil, fmt.Errorf("invalid %s %q", msg.Name, text)
		}
		return true, appendSecondsNanos(out, t.Unix(), int32(t.Nanosecond())), nil
	case msg.Name == "google.protobuf.Duration":
		tok, err := u.read('"')
		if err != nil {
			return true, nil, err
		}
		seconds, nanos, err := parseDuration(tok.String())
		if err != nil {
			return true, nil, fmt.Errorf("invalid %s %q: %w", msg.Name, tok.String(), err)
		}
		return true, appendSecondsNanos(out, seconds, nanos), nil
	case msg.Name == "google.protobuf.FieldMask":
		tok, err := u.read('"')
		if err != nil {
			return true, nil, err
		}
		if tok.String() == "" {
			return true, out, nil
		}
		for _, path := range strings.Split(tok.String(), ",") {
			snake, err := camelToSnake(path)
			if err != nil {
				return true, nil, err
			}
			out = protowire.AppendTag(out, 1, protowire.BytesType)
			out = protowire.AppendString(out, snake)
		}
		return true, out, nil
	}
	return false, out, nil
}

// secondsNanos reads the seconds and nanos fields of a Timestamp or a
// Duration.
func secondsNanos(wire []byte) (seconds int64, nanos int32, err error) {
	for len(wire) > 0 {
		num, v, n := consumeValue(wire)
		if n < 0 {
			return 0, 0, errSyntax
		}
		wire = wire[n:]
		if v.typ != protowire.VarintType {
			continue
		}
		switch num {
		case 1:
			seconds = int64(v.num)
		case 2:
			nanos = int32(v.num)
		}
	}
	return seconds, nanos, nil
}

func appendSecondsNanos(out []byte, seconds int64, nanos int32) []byte {
	if seconds != 0 {
		out = protowire.AppendTag(out, 1, protowire.VarintType)
		out = protowire.AppendVarint(out, uint64(seconds))
	}
	if nanos != 0 {
		out = protowire.AppendTag(out, 2, protowire.VarintType)
		out = protowire.AppendVarint(out, uint64(int64(nanos)))
	}
	return out
}

// trimNanos drops a nine digit fraction down to 0, 3, or 6 digits when the
// remaining digits are zero.
func trimNanos(text string) string {
	for range 3 {
		trimmed, ok := strings.CutSuffix(text, "000")
		if !ok {
			break
		}
		text = trimmed
	}
	return strings.TrimSuffix(text, ".")
}

// parseDuration parses the JSON form of a Duration the way protojson does:
// an optionally signed decimal number of seconds, with at most nine
// fractional digits and no leading zeros, followed by "s".
func parseDuration(text string) (seconds int64, nanos int32, err error) {
	b, ok := strings.CutSuffix(text, "s")
	if !ok || b == "" {
		return 0, 0, errors.New("invalid format")
	}
	negative := false
	switch b[0] {
	case '-':
		negative, b = true, b[1:]
	case '+':
		b = b[1:]
	}
	if b == "" {
		return 0, 0, errors.New("invalid format")
	}

	var whole string
	switch {
	case b[0] == '0':
		b = b[1:]
	case '1' <= b[0] && b[0] <= '9':
		n := 1
		for n < len(b) && '0' <= b[n] && b[n] <= '9' {
			n++
		}
		whole, b = b[:n], b[n:]
	case b[0] != '.':
		return 0, 0, errors.New("invalid format")
	}
	if whole != "" {
		if seconds, err = strconv.ParseInt(whole, 10, 64); err != nil {
			return 0, 0, errors.New("invalid seconds")
		}
	}
	if b != "" {
		fraction, ok := strings.CutPrefix(b, ".")
		if !ok || len(fraction) > 9 || strings.Trim(fraction, "0123456789") != "" {
			return 0, 0, errors.New("invalid fraction")
		}
		n, _ := strconv.ParseInt(fraction+strings.Repeat("0", 9-len(fraction)), 10, 64)
		nanos = int32(n)
	}
	if negative {
		seconds, nanos = -seconds, -nanos
	}
	if seconds < -maxDurationSeconds || seconds > maxDurationSeconds {
		return 0, 0, errors.New("out of range")
	}
	return seconds, nanos, nil
}

// snakeToCamel converts a FieldMask path to its JSON form. Like protojson,
// it rejects invalid paths and paths that would not convert back.
func snakeToCamel(path string) (string, error) {
	camel := jsonCamelCase(path)
	if !validFullName(path) || jsonSnakeCase(camel) != path {
		return "", fmt.Errorf("invalid FieldMask path %q", path)
	}
	return camel, nil
}

func camelToSnake(path string) (string, error) {
	snake := jsonSnakeCase(path)
	if strings.Contains(path, "_") || !validFullName(snake) {
		return "", fmt.Errorf("invalid FieldMask path %q", path)
	}
	return snake, nil
}

// jsonCamelCase and jsonSnakeCase convert names the way protobuf does for
// JSON field names.
func jsonCamelCase(s string) string {
	var b []byte
	underscore := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != '_' {
			if underscore && 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
			b = append(b, c)
		}
		underscore = c == '_'
	}
	return string(b)
}

func jsonSnakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			b = append(b, '_')
			c += 'a' - 'A'
		}
		b = append(b, c)
	}
	return string(b)
}

// validFullName reports whether s is a dot-separated list of protobuf
// identifiers.
func validFullName(s string) bool {
	for _, part := range strings.Split(s, ".") {
		if part == "" {
			return false
		}
		for i := 0; i < len(part); i++ {
			c := part[i]
			letter := c == '_' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
			if !letter && (i == 0 || c < '0' || c > '9') {
				return false
			}
		}
	}
	return true
}
