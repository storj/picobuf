// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package presencecompat

import (
	"bytes"
	"slices"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/presencecompat/editionpico"
	"storj.io/picobuf/internal/presencecompat/editionprot"
	"storj.io/picobuf/internal/presencecompat/pico"
	"storj.io/picobuf/internal/presencecompat/pico2"
	"storj.io/picobuf/internal/presencecompat/prot"
	"storj.io/picobuf/internal/presencecompat/prot2"
)

func pointerTo[T any](value T) *T { return &value }

func TestPresentDefaultsMatchProtobuf(t *testing.T) {
	emptyBytes := []byte{}
	picoMessage := &pico.Message{
		Number:        new(int32),
		Text:          new(string),
		Data:          &emptyBytes,
		State:         new(pico.State),
		Boolean:       new(bool),
		Int64Value:    new(int64),
		Uint32Value:   new(uint32),
		Uint64Value:   new(uint64),
		Sint32Value:   new(int32),
		Sint64Value:   new(int64),
		Fixed32Value:  new(uint32),
		Fixed64Value:  new(uint64),
		Sfixed32Value: new(int32),
		Sfixed64Value: new(int64),
		FloatValue:    new(float32),
		DoubleValue:   new(float64),
		Selection: &pico.Message_SelectedState{
			SelectedState: pico.State_STATE_UNSPECIFIED,
		},
	}
	protoMessage := &prot.Message{
		Number:        proto.Int32(0),
		Text:          proto.String(""),
		Data:          []byte{},
		State:         prot.State_STATE_UNSPECIFIED.Enum(),
		Boolean:       pointerTo(false),
		Int64Value:    pointerTo(int64(0)),
		Uint32Value:   pointerTo(uint32(0)),
		Uint64Value:   pointerTo(uint64(0)),
		Sint32Value:   pointerTo(int32(0)),
		Sint64Value:   pointerTo(int64(0)),
		Fixed32Value:  pointerTo(uint32(0)),
		Fixed64Value:  pointerTo(uint64(0)),
		Sfixed32Value: pointerTo(int32(0)),
		Sfixed64Value: pointerTo(int64(0)),
		FloatValue:    pointerTo(float32(0)),
		DoubleValue:   pointerTo(float64(0)),
		Selection: &prot.Message_SelectedState{
			SelectedState: prot.State_STATE_UNSPECIFIED,
		},
	}

	picoData, err := picobuf.Marshal(picoMessage)
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(protoMessage)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded present defaults differ: pico %x, protobuf %x", picoData, protoData)
	}

	var decodedProto prot.Message
	if err := proto.Unmarshal(picoData, &decodedProto); err != nil {
		t.Fatal(err)
	}
	sent, received := protoMessage.ProtoReflect(), decodedProto.ProtoReflect()
	fields := received.Descriptor().Fields()
	for i := range fields.Len() {
		fd := fields.Get(i)
		if sent.Has(fd) != received.Has(fd) {
			t.Errorf("protobuf field %s presence changed: sent %v, decoded %v", fd.Name(), sent.Has(fd), received.Has(fd))
		}
	}

	var decodedPico pico.Message
	if err := picobuf.Unmarshal(protoData, &decodedPico); err != nil {
		t.Fatal(err)
	}
	if decodedPico.Number == nil || decodedPico.Text == nil || decodedPico.Data == nil || decodedPico.State == nil ||
		decodedPico.Boolean == nil || decodedPico.Int64Value == nil || decodedPico.Uint32Value == nil ||
		decodedPico.Uint64Value == nil || decodedPico.Sint32Value == nil || decodedPico.Sint64Value == nil ||
		decodedPico.Fixed32Value == nil || decodedPico.Fixed64Value == nil || decodedPico.Sfixed32Value == nil ||
		decodedPico.Sfixed64Value == nil || decodedPico.FloatValue == nil || decodedPico.DoubleValue == nil {
		t.Fatalf("picobuf optional field lost presence: %+v", decodedPico)
	}
	if _, ok := decodedPico.Selection.(*pico.Message_SelectedState); !ok {
		t.Fatalf("picobuf oneof lost selected zero enum: %T", decodedPico.Selection)
	}
}

func TestProto2PresentDefaultsMatchProtobuf(t *testing.T) {
	emptyBytes := []byte{}
	picoData, err := picobuf.Marshal(&pico2.Message{
		Number: new(int32), Text: new(string), Data: &emptyBytes, State: new(pico2.State),
	})
	if err != nil {
		t.Fatal(err)
	}
	protoMessage := &prot2.Message{
		Number: proto.Int32(0), Text: proto.String(""), Data: []byte{}, State: prot2.State_STATE_UNSPECIFIED.Enum(),
	}
	protoData, err := proto.Marshal(protoMessage)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded proto2 defaults differ: pico %x, protobuf %x", picoData, protoData)
	}

	var decodedProto prot2.Message
	if err := proto.Unmarshal(picoData, &decodedProto); err != nil {
		t.Fatal(err)
	}
	sent, received := protoMessage.ProtoReflect(), decodedProto.ProtoReflect()
	fields := received.Descriptor().Fields()
	for i := range fields.Len() {
		fd := fields.Get(i)
		if sent.Has(fd) != received.Has(fd) {
			t.Errorf("protobuf proto2 field %s presence changed: sent %v, decoded %v", fd.Name(), sent.Has(fd), received.Has(fd))
		}
	}
	var decodedPico pico2.Message
	if err := picobuf.Unmarshal(protoData, &decodedPico); err != nil {
		t.Fatal(err)
	}
	if decodedPico.Number == nil || decodedPico.Text == nil || decodedPico.Data == nil || decodedPico.State == nil {
		t.Fatalf("picobuf proto2 field lost presence: %+v", decodedPico)
	}
}

func TestProto2InvalidUTF8MatchesProtobuf(t *testing.T) {
	invalid := string([]byte{0xff})

	picoData, err := picobuf.Marshal(&pico2.Message{Text: &invalid})
	if err != nil {
		t.Fatalf("picobuf rejected invalid UTF-8 in a proto2 string: %v", err)
	}
	protoData, err := proto.Marshal(&prot2.Message{Text: proto.String(invalid)})
	if err != nil {
		t.Fatalf("protobuf rejected invalid UTF-8 in a proto2 string: %v", err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded invalid UTF-8 differs: pico %x, protobuf %x", picoData, protoData)
	}

	var decoded pico2.Message
	if err := picobuf.Unmarshal(protoData, &decoded); err != nil {
		t.Fatalf("picobuf rejected decoding invalid UTF-8 in a proto2 string: %v", err)
	}
	if decoded.Text == nil || *decoded.Text != invalid {
		t.Fatalf("decoded %+v", decoded)
	}
}

func TestEditionsPresenceMatchesProtobuf(t *testing.T) {
	emptyBytes := []byte{}
	picoData, err := picobuf.Marshal(&editionpico.Message{
		ExplicitNumber: new(int32), ExplicitText: new(string), ExplicitData: &emptyBytes,
		ImplicitNumber: 0, ImplicitText: "", ImplicitData: []byte{},
	})
	if err != nil {
		t.Fatal(err)
	}
	protoMessage := &editionprot.Message{
		ExplicitNumber: proto.Int32(0), ExplicitText: proto.String(""), ExplicitData: []byte{},
		ImplicitNumber: 0, ImplicitText: "", ImplicitData: []byte{},
	}
	protoData, err := proto.Marshal(protoMessage)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded Editions presence differs: pico %x, protobuf %x", picoData, protoData)
	}

	var decodedProto editionprot.Message
	if err := proto.Unmarshal(picoData, &decodedProto); err != nil {
		t.Fatal(err)
	}
	sent, received := protoMessage.ProtoReflect(), decodedProto.ProtoReflect()
	fields := received.Descriptor().Fields()
	for i := range fields.Len() {
		fd := fields.Get(i)
		if sent.Has(fd) != received.Has(fd) {
			t.Errorf("protobuf Editions field %s presence changed: sent %v, decoded %v", fd.Name(), sent.Has(fd), received.Has(fd))
		}
		if !fd.HasPresence() && received.Has(fd) {
			t.Errorf("protobuf Editions implicit field %s unexpectedly has presence", fd.Name())
		}
	}
	var decodedPico editionpico.Message
	if err := picobuf.Unmarshal(protoData, &decodedPico); err != nil {
		t.Fatal(err)
	}
	if decodedPico.ExplicitNumber == nil || decodedPico.ExplicitText == nil || decodedPico.ExplicitData == nil {
		t.Fatalf("picobuf Editions explicit field lost presence: %+v", decodedPico)
	}
}

func TestEditionsRepeatedEncodingMatchesProtobuf(t *testing.T) {
	picoMessage := &editionpico.Message{
		PackedNumbers:   []int32{0, 1, 150},
		ExpandedNumbers: []int32{0, 1, 150},
	}
	protoMessage := &editionprot.Message{
		PackedNumbers:   []int32{0, 1, 150},
		ExpandedNumbers: []int32{0, 1, 150},
	}

	picoData, err := picobuf.Marshal(picoMessage)
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(protoMessage)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded Editions repeated fields differ: pico %x, protobuf %x", picoData, protoData)
	}

	var decodedProto editionprot.Message
	if err := proto.Unmarshal(picoData, &decodedProto); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(&decodedProto, protoMessage) {
		t.Fatalf("protobuf decoded repeated fields as %v", &decodedProto)
	}
	var decodedPico editionpico.Message
	if err := picobuf.Unmarshal(protoData, &decodedPico); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(decodedPico.PackedNumbers, picoMessage.PackedNumbers) ||
		!slices.Equal(decodedPico.ExpandedNumbers, picoMessage.ExpandedNumbers) {
		t.Fatalf("picobuf decoded repeated fields as %+v", decodedPico)
	}
}

func TestEditionsDelimitedMessageMatchesProtobuf(t *testing.T) {
	for _, value := range []*int32{new(int32), pointerTo(int32(123))} {
		picoMessage := &editionpico.Message{Nested: &editionpico.Nested{Value: value}}
		protoMessage := &editionprot.Message{Nested: &editionprot.Nested{Value: value}}

		picoData, err := picobuf.Marshal(picoMessage)
		if err != nil {
			t.Fatal(err)
		}
		protoData, err := proto.Marshal(protoMessage)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(picoData, protoData) {
			t.Fatalf("encoded Editions delimited message differs: pico %x, protobuf %x", picoData, protoData)
		}

		var decodedProto editionprot.Message
		if err := proto.Unmarshal(picoData, &decodedProto); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(&decodedProto, protoMessage) {
			t.Fatalf("protobuf decoded delimited message as %v", &decodedProto)
		}
		var decodedPico editionpico.Message
		if err := picobuf.Unmarshal(protoData, &decodedPico); err != nil {
			t.Fatal(err)
		}
		if decodedPico.Nested == nil || decodedPico.Nested.Value == nil || *decodedPico.Nested.Value != *value {
			t.Fatalf("picobuf decoded delimited message as %+v", decodedPico.Nested)
		}
	}
}

func TestEditionsClosedEnumsMatchProtobuf(t *testing.T) {
	// Unknown values occur in a singular field, a packed repeated field, and
	// a oneof. The repeated field also contains the known values 0 and 1.
	data := []byte{0x50, 0x7b, 0x5a, 0x03, 0x00, 0x7b, 0x01, 0x60, 0x7b}

	var protoMessage editionprot.Message
	if err := proto.Unmarshal(data, &protoMessage); err != nil {
		t.Fatal(err)
	}
	fields := protoMessage.ProtoReflect().Descriptor().Fields()
	if !protoMessage.ProtoReflect().Has(fields.ByNumber(10)) || protoMessage.GetClosedState() != 123 {
		t.Fatalf("protobuf decoded singular closed enum as %v", protoMessage.GetClosedState())
	}
	if !protoMessage.ProtoReflect().Has(fields.ByNumber(12)) || protoMessage.GetSelectedClosedState() != 123 {
		t.Fatalf("protobuf decoded closed enum oneof as %v", protoMessage.GetClosedSelection())
	}
	if got := protoMessage.GetClosedStates(); !slices.Equal(got, []editionprot.ClosedState{0, 123, 1}) {
		t.Fatalf("protobuf decoded closed enum values as %v", got)
	}
	protoData, err := proto.Marshal(&protoMessage)
	if err != nil {
		t.Fatal(err)
	}

	var picoMessage editionpico.Message
	if err := picobuf.Unmarshal(data, &picoMessage); err != nil {
		t.Fatal(err)
	}
	if picoMessage.ClosedState == nil || *picoMessage.ClosedState != 123 {
		t.Fatalf("picobuf decoded singular closed enum as %v", picoMessage.ClosedState)
	}
	selection, ok := picoMessage.ClosedSelection.(*editionpico.Message_SelectedClosedState)
	if !ok || selection.SelectedClosedState != 123 {
		t.Fatalf("picobuf decoded closed enum oneof as %T", picoMessage.ClosedSelection)
	}
	if !slices.Equal(picoMessage.ClosedStates, []editionpico.ClosedState{0, 123, 1}) {
		t.Fatalf("picobuf decoded closed enum values as %v", picoMessage.ClosedStates)
	}
	picoData, err := picobuf.Marshal(&picoMessage)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("reserialized closed enums differ: pico %x, protobuf %x", picoData, protoData)
	}
}

func TestEditionsLegacyRequiredMatchesProtobuf(t *testing.T) {
	if _, err := proto.Marshal(&editionprot.RequiredMessage{}); err == nil {
		t.Fatal("protobuf marshaled a message with a missing required field")
	}
	if _, err := picobuf.Marshal(&editionpico.RequiredMessage{}); err == nil {
		t.Fatal("picobuf marshaled a message with a missing required field")
	}

	var protoMissing editionprot.RequiredMessage
	if err := proto.Unmarshal(nil, &protoMissing); err == nil {
		t.Fatal("protobuf unmarshaled a message with a missing required field")
	}
	var picoMissing editionpico.RequiredMessage
	if err := picobuf.Unmarshal(nil, &picoMissing); err == nil {
		t.Fatal("picobuf unmarshaled a message with a missing required field")
	}

	picoData, err := picobuf.Marshal(&editionpico.RequiredMessage{RequiredNumber: new(int32)})
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(&editionprot.RequiredMessage{RequiredNumber: new(int32)})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded required defaults differ: pico %x, protobuf %x", picoData, protoData)
	}
}

func TestEditionsNestedLegacyRequiredMatchesProtobuf(t *testing.T) {
	if _, err := proto.Marshal(&editionprot.RequiredParent{Child: &editionprot.RequiredMessage{}}); err == nil {
		t.Fatal("protobuf marshaled a nested message with a missing required field")
	}
	if _, err := picobuf.Marshal(&editionpico.RequiredParent{Child: &editionpico.RequiredMessage{}}); err == nil {
		t.Fatal("picobuf marshaled a nested message with a missing required field")
	}

	data := []byte{0x0a, 0x00}
	var protoMessage editionprot.RequiredParent
	if err := proto.Unmarshal(data, &protoMessage); err == nil {
		t.Fatal("protobuf unmarshaled a nested message with a missing required field")
	}
	var picoMessage editionpico.RequiredParent
	if err := picobuf.Unmarshal(data, &picoMessage); err == nil {
		t.Fatal("picobuf unmarshaled a nested message with a missing required field")
	}
}

func TestRepeatedAndMapDoNotTrackPresence(t *testing.T) {
	picoNil, err := picobuf.Marshal(&pico.Message{})
	if err != nil {
		t.Fatal(err)
	}
	picoEmpty, err := picobuf.Marshal(&pico.Message{Numbers: []int32{}, Names: map[string]string{}})
	if err != nil {
		t.Fatal(err)
	}
	protoNil, err := proto.Marshal(&prot.Message{})
	if err != nil {
		t.Fatal(err)
	}
	protoEmpty, err := proto.Marshal(&prot.Message{Numbers: []int32{}, Names: map[string]string{}})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoNil, picoEmpty) || !bytes.Equal(picoNil, protoNil) || !bytes.Equal(picoNil, protoEmpty) {
		t.Fatalf("empty repeated/map fields encoded differently: pico nil %x, pico empty %x, protobuf nil %x, protobuf empty %x", picoNil, picoEmpty, protoNil, protoEmpty)
	}
}

func TestAlwaysPresentIntentionallyLosesPresence(t *testing.T) {
	picoData, err := picobuf.Marshal(&pico.Message{AlwaysText: ""})
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(&prot.Message{AlwaysText: proto.String("")})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(picoData, protoData) {
		t.Fatalf("always_present unexpectedly retained presence: %x", picoData)
	}
	var decoded prot.Message
	if err := proto.Unmarshal(picoData, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.AlwaysText != nil {
		t.Fatal("always_present empty value unexpectedly decoded as present")
	}
}

func TestPresentEmptyOneofStringMatchesProtobuf(t *testing.T) {
	picoData, err := picobuf.Marshal(&pico.Message{
		Selection: &pico.Message_SelectedText{},
	})
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(&prot.Message{
		Selection: &prot.Message_SelectedText{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded present empty strings differ: pico %x, protobuf %x", picoData, protoData)
	}
}

func TestUnsetFieldsMatchProtobuf(t *testing.T) {
	picoData, err := picobuf.Marshal(&pico.Message{})
	if err != nil {
		t.Fatal(err)
	}
	protoData, err := proto.Marshal(&prot.Message{})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(picoData, protoData) {
		t.Fatalf("encoded unset fields differ: pico %x, protobuf %x", picoData, protoData)
	}
}
