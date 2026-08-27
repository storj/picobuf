// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picowkt_test

import (
	"bytes"
	"math"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"storj.io/picobuf"
	"storj.io/picobuf/picowkt"
)

func check(t *testing.T, pico picobuf.Message, std proto.Message) {
	t.Helper()

	got, err := picobuf.Marshal(pico)
	if err != nil {
		t.Fatal(err)
	}
	want, err := proto.Marshal(std)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%T encoded %x, want %x", pico, got, want)
	}

	fresh := reflect.New(reflect.TypeOf(pico).Elem()).Interface().(picobuf.Message)
	if err := picobuf.Unmarshal(want, fresh); err != nil {
		t.Fatal(err)
	}
	back, err := picobuf.Marshal(fresh)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(back, want) {
		t.Fatalf("%T round trip %x, want %x", pico, back, want)
	}
}

func TestTypesMatchProtobuf(t *testing.T) {
	check(t, &picowkt.DoubleValue{}, &wrapperspb.DoubleValue{})
	check(t, &picowkt.DoubleValue{Value: -1.5}, wrapperspb.Double(-1.5))
	check(t, &picowkt.FloatValue{}, &wrapperspb.FloatValue{})
	check(t, &picowkt.FloatValue{Value: 1.5}, wrapperspb.Float(1.5))
	check(t, &picowkt.DoubleValue{Value: math.Copysign(0, -1)}, wrapperspb.Double(math.Copysign(0, -1)))
	check(t, &picowkt.FloatValue{Value: float32(math.Copysign(0, -1))}, wrapperspb.Float(float32(math.Copysign(0, -1))))
	check(t, &picowkt.Int64Value{}, &wrapperspb.Int64Value{})
	check(t, &picowkt.Int64Value{Value: -1}, wrapperspb.Int64(-1))
	check(t, &picowkt.UInt64Value{}, &wrapperspb.UInt64Value{})
	check(t, &picowkt.UInt64Value{Value: 1 << 63}, wrapperspb.UInt64(1<<63))
	check(t, &picowkt.Int32Value{}, &wrapperspb.Int32Value{})
	check(t, &picowkt.Int32Value{Value: -1}, wrapperspb.Int32(-1))
	check(t, &picowkt.UInt32Value{}, &wrapperspb.UInt32Value{})
	check(t, &picowkt.UInt32Value{Value: 1 << 31}, wrapperspb.UInt32(1<<31))
	check(t, &picowkt.BoolValue{}, &wrapperspb.BoolValue{})
	check(t, &picowkt.BoolValue{Value: true}, wrapperspb.Bool(true))
	check(t, &picowkt.StringValue{}, &wrapperspb.StringValue{})
	check(t, &picowkt.StringValue{Value: "hello"}, wrapperspb.String("hello"))
	check(t, &picowkt.BytesValue{}, &wrapperspb.BytesValue{})
	check(t, &picowkt.BytesValue{Value: []byte{1, 2, 3}}, wrapperspb.Bytes([]byte{1, 2, 3}))

	check(t, &picowkt.Timestamp{}, &timestamppb.Timestamp{})
	check(t, &picowkt.Timestamp{Seconds: 1234567890, Nanos: 987654321},
		&timestamppb.Timestamp{Seconds: 1234567890, Nanos: 987654321})
	check(t, &picowkt.Duration{}, &durationpb.Duration{})
	check(t, &picowkt.Duration{Seconds: -5, Nanos: -500000000},
		&durationpb.Duration{Seconds: -5, Nanos: -500000000})
	check(t, &picowkt.Empty{}, &emptypb.Empty{})
	check(t, &picowkt.FieldMask{}, &fieldmaskpb.FieldMask{})
	check(t, &picowkt.FieldMask{Paths: []string{"a.b", "c"}},
		&fieldmaskpb.FieldMask{Paths: []string{"a.b", "c"}})
}
