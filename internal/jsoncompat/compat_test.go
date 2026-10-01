// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package jsoncompat

import (
	"encoding/json"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/jsoncompat/pico"
	"storj.io/picobuf/internal/jsoncompat/prot"
	"storj.io/picobuf/picowkt"
)

func seeds() []*prot.All {
	scalars := &prot.Scalars{
		Int32: -5, Int64: math.MinInt64, Uint32: math.MaxUint32, Uint64: math.MaxUint64,
		Sint32: -7, Sint64: math.MaxInt64, Fixed32: 1, Fixed64: 1 << 60,
		Sfixed32: math.MinInt32, Sfixed64: -1, Float: 1.5e-7, Double: 1e21,
		Bool: true, String_: "hé\"\\\n\x01<>& ", Bytes: []byte{0, 0xff, 0xfe},
		Color: prot.Color_COLOR_BLUE, SnakeCaseName: "snake", Renamed: "renamed",
	}
	return []*prot.All{
		{},
		{Scalars: scalars},
		{Scalars: &prot.Scalars{Float: float32(math.Inf(1)), Double: math.NaN(), Color: prot.Color(42)}},
		{Scalars: &prot.Scalars{Double: math.Copysign(0, -1), Float: float32(math.Inf(-1))}},
		{Optionals: &prot.Optionals{
			Int32: proto.Int32(0), Int64: proto.Int64(0), String_: proto.String(""),
			Bytes: []byte{}, Double: proto.Float64(0), Color: prot.Color_COLOR_UNSPECIFIED.Enum(),
			Bool: proto.Bool(false), Float: proto.Float32(0), Uint64: proto.Uint64(0),
		}},
		{Lists: &prot.Lists{
			Int32: []int32{-1, 0, 1}, Int64: []int64{math.MinInt64, 0}, Uint64: []uint64{math.MaxUint64},
			Sint32: []int32{-2, 2}, Fixed32: []uint32{7}, Sfixed64: []int64{-8},
			Float: []float32{0.1, -2}, Double: []float64{0.1, 1e-7}, Bool: []bool{true, false},
			String_: []string{"", "a"}, Bytes: [][]byte{{}, {1}}, Color: []prot.Color{0, 1, -2, 9},
			Messages: []*prot.Scalars{{}, scalars},
		}},
		{Maps: &prot.Maps{
			StringInt64:  map[string]int64{"": 0, "a": -1},
			Int32String:  map[int32]string{-5: "x", 0: ""},
			BoolMessage:  map[bool]*prot.Scalars{true: {}, false: scalars},
			Uint64Color:  map[uint64]prot.Color{math.MaxUint64: prot.Color_COLOR_RED, 1: 7},
			Sint32Bytes:  map[int32][]byte{-1: {1, 2}},
			StringDouble: map[string]float64{"nan": math.NaN(), "big": 1e300},
			Fixed64Bool:  map[uint64]bool{1: true, 2: false},
		}},
		{Oneofs: &prot.Oneofs{Kind: &prot.Oneofs_Name{Name: ""}}},
		{Oneofs: &prot.Oneofs{Kind: &prot.Oneofs_Scalars{Scalars: &prot.Scalars{}}}},
		{Oneofs: &prot.Oneofs{Kind: &prot.Oneofs_Count{Count: 0}}},
		{Oneofs: &prot.Oneofs{Kind: &prot.Oneofs_Color{Color: prot.Color_COLOR_UNSPECIFIED}}},
		{Oneofs: &prot.Oneofs{Kind: &prot.Oneofs_Time{Time: timestamppb.New(time.Unix(0, 0))}}},
		{WellKnown: &prot.WellKnown{
			Timestamp: timestamppb.New(time.Date(1972, 1, 1, 10, 0, 20, 21e6, time.UTC)),
			Duration:  durationpb.New(-1500 * time.Millisecond),
			Double:    wrapperspb.Double(0), Float: wrapperspb.Float(-1.25),
			Int64: wrapperspb.Int64(-3), Uint64: wrapperspb.UInt64(math.MaxUint64),
			Int32: wrapperspb.Int32(0), Uint32: wrapperspb.UInt32(3),
			Bool: wrapperspb.Bool(false), String_: wrapperspb.String("s"), Bytes: wrapperspb.Bytes(nil),
			Mask:  &fieldmaskpb.FieldMask{Paths: []string{"foo_bar", "baz.qux_quux"}},
			Empty: &emptypb.Empty{},
			Timestamps: []*timestamppb.Timestamp{
				{Seconds: -62135596800}, {Seconds: 253402300799, Nanos: 999999999}, {Nanos: 1000},
			},
			Durations: map[string]*durationpb.Duration{
				"zero": {}, "nanos": {Nanos: -1}, "max": {Seconds: 315576000000, Nanos: 999999999},
			},
		}},
		{Custom: &prot.Custom{
			Id:           []byte("0123456789ab"),
			Time:         timestamppb.New(time.Date(2026, 9, 30, 12, 0, 0, 5, time.UTC)),
			OptionalTime: timestamppb.New(time.Time{}),
		}},
		{Recursive: &prot.All{Recursive: &prot.All{Scalars: scalars}}},

		// Values that protojson refuses to marshal.
		{WellKnown: &prot.WellKnown{Timestamp: &timestamppb.Timestamp{Seconds: 253402300800}}},
		{WellKnown: &prot.WellKnown{Timestamp: &timestamppb.Timestamp{Nanos: -1}}},
		{WellKnown: &prot.WellKnown{Duration: &durationpb.Duration{Seconds: 315576000001}}},
		{WellKnown: &prot.WellKnown{Duration: &durationpb.Duration{Seconds: 1, Nanos: -1}}},
		{WellKnown: &prot.WellKnown{Mask: &fieldmaskpb.FieldMask{Paths: []string{"fooBar"}}}},
		{WellKnown: &prot.WellKnown{Mask: &fieldmaskpb.FieldMask{Paths: []string{"foo__bar"}}}},
	}
}

// FuzzJSON checks that picojson and protojson agree on messages decoded from
// the same wire data, in both directions.
func FuzzJSON(f *testing.F) {
	for _, seed := range seeds() {
		data, err := proto.Marshal(seed)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(data)
	}
	// Custom.raw set to "" on the wire.
	f.Add([]byte{0x3a, 0x02, 0x22, 0x00})
	f.Fuzz(checkWire)
}

func checkWire(t *testing.T, wire []byte) {
	var got pico.All
	if err := picobuf.Unmarshal(wire, &got); err != nil {
		return
	}
	// Compare against what picobuf holds, since custom types can normalize
	// values.
	held, ok := toProt(t, &got)
	if !ok {
		return
	}
	want, wantErr := protojson.Marshal(held)
	data, err := got.MarshalJSON()
	if wantErr != nil {
		if err == nil {
			t.Fatalf("picojson marshaled what protojson rejects (%v):\n%s", wantErr, data)
		}
		return
	}
	if err != nil {
		t.Fatalf("picojson failed (%v), protojson:\n%s", err, want)
	}
	if !sameJSON(t, data, want) {
		t.Fatalf("picojson:\n%s\nprotojson:\n%s", data, want)
	}

	var back pico.All
	if err := back.UnmarshalJSON(want); err != nil {
		t.Fatalf("picojson failed to parse protojson output (%v):\n%s", err, want)
	}
	if backProt, ok := toProt(t, &back); !ok || !proto.Equal(backProt, held) {
		t.Fatalf("round trip through picojson changed the message:\n%s", want)
	}
}

// TestJSONOutOfOrderWire checks wire data that picobuf does not encode by
// itself, but that unrecognized fields can add, such as when a referenced
// package is generated from an older version of a message.
func TestJSONOutOfOrderWire(t *testing.T) {
	varint := func(num protowire.Number, v uint64) []byte {
		return protowire.AppendVarint(protowire.AppendTag(nil, num, protowire.VarintType), v)
	}
	entry := func(key string, value uint64) []byte {
		e := protowire.AppendString(protowire.AppendTag(nil, 1, protowire.BytesType), key)
		e = append(e, varint(2, value)...)
		return protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), e)
	}
	for _, m := range []*pico.All{
		// Before the last field.
		{Scalars: &pico.Scalars{Int64: 5, XXX_unrecognized: varint(1, 7)}},
		// A singular field twice.
		{Scalars: &pico.Scalars{Int32: 3, XXX_unrecognized: varint(1, 9)}},
		// A singular message twice, which merges.
		{
			Scalars:          &pico.Scalars{Int32: 3},
			XXX_unrecognized: protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), varint(2, 4)),
		},
		// A repeated field in two places.
		{Lists: &pico.Lists{Int32: []int32{1}, Bool: []bool{true}, XXX_unrecognized: varint(1, 2)}},
		// A repeated map key.
		{Maps: &pico.Maps{StringInt64: map[string]int64{"a": 1}, XXX_unrecognized: entry("a", 2)}},
		// Map keys out of order.
		{Maps: &pico.Maps{StringInt64: map[string]int64{"b": 1}, XXX_unrecognized: entry("a", 2)}},
		// An invalid value that a later occurrence replaces.
		{WellKnown: &pico.WellKnown{
			Timestamp:        &picowkt.Timestamp{Seconds: 1 << 40},
			XXX_unrecognized: protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), varint(1, 1)),
		}},
	} {
		held, ok := toProt(t, m)
		if !ok {
			t.Fatal("invalid message")
		}
		want, err := protojson.Marshal(held)
		if err != nil {
			t.Fatal(err)
		}
		got, err := m.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if !sameJSON(t, got, want) {
			t.Errorf("picojson:\n%s\nprotojson:\n%s", got, want)
		}
	}
}

// TestJSONValue checks that encoding/json uses MarshalJSON for messages that
// are not addressable.
func TestJSONValue(t *testing.T) {
	got, err := json.Marshal(map[string]pico.Scalars{"a": {Int64: 5}})
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"a":{"int64":"5"}}`; string(got) != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

// TestJSONNestedError checks that an error deep in a message does not
// render the message again at every level, which took time exponential in
// the depth.
func TestJSONNestedError(t *testing.T) {
	m := &pico.All{WellKnown: &pico.WellKnown{Timestamp: &picowkt.Timestamp{Seconds: 1 << 40}}}
	for range 60 {
		m = &pico.All{Recursive: m}
	}
	if _, err := m.MarshalJSON(); err == nil {
		t.Fatal("expected an error")
	}
}

// TestJSONNestedIrregular checks that wire data that needs grouping at every
// level does not render each subtree again at every level.
func TestJSONNestedIrregular(t *testing.T) {
	// Scalars after Recursive is out of declaration order.
	scalars := protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType),
		protowire.AppendVarint(protowire.AppendTag(nil, 1, protowire.VarintType), 7))
	m := &pico.All{}
	for range 40 {
		m = &pico.All{Recursive: m, XXX_unrecognized: scalars}
	}
	held, ok := toProt(t, m)
	if !ok {
		t.Fatal("invalid message")
	}
	want, err := protojson.Marshal(held)
	if err != nil {
		t.Fatal(err)
	}
	got, err := m.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !sameJSON(t, got, want) {
		t.Fatalf("picojson:\n%s\nprotojson:\n%s", got, want)
	}
}

// TestJSONNull checks that null leaves a message unchanged, as encoding/json
// expects, both for message values and pointers.
func TestJSONNull(t *testing.T) {
	var v struct {
		Value   pico.Scalars
		Pointer *pico.Scalars
	}
	v.Value.Int32 = 5
	if err := json.Unmarshal([]byte(`{"Value":null,"Pointer":null}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.Value.Int32 != 5 || v.Pointer != nil {
		t.Fatalf("got %+v", v)
	}
}

// TestJSONErrorPrefix checks that nested errors name the package once.
func TestJSONErrorPrefix(t *testing.T) {
	m := &pico.All{WellKnown: &pico.WellKnown{Timestamp: &picowkt.Timestamp{Seconds: 1 << 40}}}
	_, err := m.MarshalJSON()
	if err == nil || strings.Count(err.Error(), "picojson:") != 1 {
		t.Fatalf("got %v", err)
	}
	err = new(pico.All).UnmarshalJSON([]byte(`{"wellKnown":{"timestamp":"x"}}`))
	if err == nil || strings.Count(err.Error(), "picojson:") != 1 {
		t.Fatalf("got %v", err)
	}
}

// toProt converts a picobuf message to its protobuf-go equivalent through
// the wire format, without unknown fields.
func toProt(t *testing.T, m *pico.All) (*prot.All, bool) {
	t.Helper()
	wire, err := picobuf.Marshal(m)
	if err != nil {
		return nil, false
	}
	// ProtoJSON drops unknown fields, which picobuf captures and re-emits.
	var p prot.All
	if err := (proto.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(wire, &p); err != nil {
		return nil, false
	}
	return &p, true
}

func sameJSON(t *testing.T, a, b []byte) bool {
	t.Helper()
	var x, y any
	if err := json.Unmarshal(a, &x); err != nil {
		t.Fatalf("invalid JSON (%v): %s", err, a)
	}
	if err := json.Unmarshal(b, &y); err != nil {
		t.Fatalf("invalid JSON (%v): %s", err, b)
	}
	return reflect.DeepEqual(x, y)
}

// jsonInputs are ProtoJSON documents that exercise the parser, including
// ones that protojson rejects.
var jsonInputs = []string{
	`{}`,
	`{"scalars":{"snake_case_name":"a","renamed":"b"}}`,
	`{"scalars":{"snakeCaseName":"a","customName":"b"}}`,
	`{"scalars":{"int32":"5","int64":7,"uint32":1e2,"sint64":"-3","fixed64":"18446744073709551615"}}`,
	`{"scalars":{"int32":1.0,"sfixed32":-2147483648,"uint64":"1e3"}}`,
	// Integer conversion must stay exact beyond float64's integer precision.
	`{"scalars":{"int64":9007199254740993e0,"uint64":"9007199254740993.0"}}`,
	`{"scalars":{"int64":"9223372036854775807e0","uint64":18446744073709551615.0}}`,
	`{"scalars":{"int64":-9223372036854775808.0,"sint64":"90071992547409930e-1"}}`,
	`{"scalars":{"int32":2147483647.0,"uint32":"4294967295e0","sfixed64":"-9007199254740993.0"}}`,
	`{"scalars":{"int32":0e999999999999999999999,"uint64":-0.0,"fixed64":"1.000e3"}}`,
	`{"scalars":{"int64":1.00000000000000001}}`,
	`{"scalars":{"uint64":"1.00000000000000001"}}`,
	`{"scalars":{"int64":-9223372036854775809.0}}`,
	`{"scalars":{"uint64":18446744073709551616e0}}`,
	`{"scalars":{"int32":1e-999999999999999999999}}`,
	`{"scalars":{"uint64":1e999999999999999999999}}`,
	`{"scalars":{"float":"NaN","double":"-Infinity"}}`,
	`{"scalars":{"float":"1.5","double":1e-7}}`,
	`{"scalars":{"bytes":"_-8"}}`,
	`{"scalars":{"bytes":"AQI"}}`,
	`{"scalars":{"bytes":"AQI="}}`,
	`{"scalars":{"color":1}}`,
	`{"scalars":{"color":5}}`,
	`{"scalars":{"color":"COLOR_BLUE"}}`,
	`{"scalars":null,"optionals":{"int32":null,"float":0,"uint64":"0"}}`,
	`{"optionals":{"int32":0,"string":"","bytes":"","bool":false,"color":"COLOR_UNSPECIFIED"}}`,
	`{"lists":{"int32":[1,"2",3e0],"color":["COLOR_RED",0],"messages":[{},{"int32":1}]}}`,
	`{"lists":{"int32":null}}`,
	`{"maps":{"int32String":{"-5":"a"},"boolMessage":{"true":{}},"uint64Color":{"18446744073709551615":"COLOR_RED"}}}`,
	`{"maps":{"fixed64Bool":{"1":true},"stringDouble":{"x":"Infinity"},"sint32Bytes":{"-1":"AQ=="}}}`,
	// Distinct JSON names can decode to the same numeric map key.
	`{"maps":{"int32String":{"1":"a","01":"b"}}}`,
	`{"maps":{"int32String":{"1":"a","+1":"b"}}}`,
	`{"maps":{"int32String":{"0":"a","-0":"b"}}}`,
	`{"maps":{"uint64Color":{"1":"COLOR_RED","01":"COLOR_BLUE"}}}`,
	`{"maps":{"fixed64Bool":{"1":true,"01":false}}}`,
	`{"maps":{"sint32Bytes":{"-1":"AQ==","-01":"Ag=="}}}`,
	`{"maps":{"int32String":{"01":"a","+2":"b","-0":"c"}}}`,
	`{"maps":{"int32String":{"1e0":"a"}}}`,
	`{"maps":{"int32String":{"1.0":"a"}}}`,
	`{"maps":{"int32String":{"2147483648":"a"}}}`,
	`{"maps":{"uint64Color":{"-0":"COLOR_RED"}}}`,
	`{"maps":{"uint64Color":{"+1":"COLOR_RED"}}}`,
	`{"maps":{"fixed64Bool":{"18446744073709551616":true}}}`,
	`{"oneofs":{"name":"x"}}`,
	`{"oneofs":{"time":"1970-01-01T00:00:00Z"}}`,
	`{"wellKnown":{"timestamp":"1972-01-01T10:00:20.021+01:00","duration":"-1.5s"}}`,
	`{"wellKnown":{"timestamp":"0001-01-01T00:00:00Z","duration":"0.000000001s"}}`,
	`{"wellKnown":{"double":1,"float":"NaN","int64":"-1","uint64":1,"int32":2,"uint32":3,"bool":true,"string":"s","bytes":"AQ=="}}`,
	`{"wellKnown":{"int32":null,"mask":"fooBar,baz.quxQuux","empty":{}}}`,
	`{"wellKnown":{"mask":""}}`,
	`{"wellKnown":{"timestamps":["2000-01-01T00:00:00.5Z"],"durations":{"a":"3s"}}}`,
	`{"custom":{"id":"MDEyMzQ1Njc4OWFi","time":"2026-09-30T12:00:00Z","optionalTime":"0001-01-01T00:00:00Z"}}`,
	`{"recursive":{"recursive":{"scalars":{"bool":true}}}}`,

	// Inputs that both should reject.
	`{"unknown":1}`,
	`{"scalars":{"int32":1,"int32":2}}`,
	`{"scalars":{"snakeCaseName":"a","snake_case_name":"b"}}`,
	`{"oneofs":{"name":"x","count":1}}`,
	`{"scalars":{"int32":2147483648}}`,
	`{"scalars":{"uint32":-1}}`,
	`{"scalars":{"int32":1.5}}`,
	`{"scalars":{"float":1e39}}`,
	`{"scalars":{"bool":"true"}}`,
	`{"scalars":{"string":1}}`,
	`{"scalars":{"color":"COLOR_GREEN"}}`,
	`{"scalars":{"bytes":"!"}}`,
	`{"wellKnown":{"timestamp":"10000-01-01T00:00:00Z"}}`,
	`{"wellKnown":{"duration":"1.5"}}`,
	`{"wellKnown":{"duration":"315576000001s"}}`,
	`{"wellKnown":{"mask":"foo_bar"}}`,
	`{"wellKnown":{"empty":{"x":1}}}`,
	`{"lists":{"int32":[null]}}`,
	`{"maps":{"int32String":{"x":"a"}}}`,
	`{} {}`,
	`[]`,
}

// TestJSONInputMatchesProtojson checks that picojson accepts and rejects the
// same inputs as protojson, and decodes accepted ones to the same message.
func TestJSONInputMatchesProtojson(t *testing.T) {
	for _, input := range jsonInputs {
		t.Run(input, func(t *testing.T) { checkInput(t, input, false) })
	}
}

// FuzzJSONInput is TestJSONInputMatchesProtojson on arbitrary input. It lets
// picojson be stricter than protojson, which accepts some malformed input.
func FuzzJSONInput(f *testing.F) {
	for _, input := range jsonInputs {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, input string) { checkInput(t, input, true) })
}

// checkInput compares picojson with protojson on input. With stricter,
// picojson may reject input that protojson accepts, as long as it accepts
// the protojson encoding of the same message.
func checkInput(t *testing.T, input string, stricter bool) {
	// Generated UnmarshalJSON follows encoding/json: a literal null is a no-op.
	if input == "null" {
		got := pico.All{Scalars: &pico.Scalars{Int32: 7}}
		if err := got.UnmarshalJSON([]byte(input)); err != nil || got.Scalars == nil || got.Scalars.Int32 != 7 {
			t.Fatal("null must leave the receiver unchanged")
		}
		return
	}
	var want prot.All
	wantErr := protojson.Unmarshal([]byte(input), &want)
	var got pico.All
	err := got.UnmarshalJSON([]byte(input))
	switch {
	case wantErr != nil && err == nil:
		t.Fatalf("picojson accepted what protojson rejects: %v", wantErr)
	case wantErr == nil && err != nil:
		// protojson accepts some malformed input, such as the number 0e or
		// the string "10!" for an integer.
		if !stricter {
			t.Fatalf("picojson rejected what protojson accepts: %v", err)
		}
		canonical, merr := protojson.Marshal(&want)
		if merr != nil {
			return
		}
		// Custom types, such as pic.ID, may reject values regardless of JSON.
		if wire, err := proto.Marshal(&want); err != nil || picobuf.Unmarshal(wire, new(pico.All)) != nil {
			return
		}
		var back pico.All
		if err := back.UnmarshalJSON(canonical); err != nil {
			t.Fatalf("picojson rejected protojson output (%v):\n%s", err, canonical)
		}
		return
	case wantErr != nil:
		return
	}
	if gotProt, ok := toProt(t, &got); !ok || !proto.Equal(gotProt, &want) {
		t.Fatalf("decoded %v, protojson decoded %v", gotProt, &want)
	}
}
