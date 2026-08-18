// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestGenerateUTF8ValidationFeatures(t *testing.T) {
	for _, test := range []struct {
		name            string
		syntax          string
		edition         descriptorpb.Edition
		fileFeature     descriptorpb.FeatureSet_Utf8Validation
		fieldFeature    descriptorpb.FeatureSet_Utf8Validation
		wantUnvalidated bool
	}{
		{name: "proto3", syntax: "proto3"},
		{name: "proto2", syntax: "proto2", wantUnvalidated: true},
		{name: "edition default", syntax: "editions", edition: descriptorpb.Edition_EDITION_2023},
		{name: "edition file none", syntax: "editions", edition: descriptorpb.Edition_EDITION_2023, fileFeature: descriptorpb.FeatureSet_NONE, wantUnvalidated: true},
		{name: "edition field none", syntax: "editions", edition: descriptorpb.Edition_EDITION_2023, fieldFeature: descriptorpb.FeatureSet_NONE, wantUnvalidated: true},
		{name: "field verify overrides file", syntax: "editions", edition: descriptorpb.Edition_EDITION_2023, fileFeature: descriptorpb.FeatureSet_NONE, fieldFeature: descriptorpb.FeatureSet_VERIFY},
	} {
		t.Run(test.name, func(t *testing.T) {
			fileOptions := &descriptorpb.FileOptions{GoPackage: proto.String("example.com/test;test")}
			if test.fileFeature != descriptorpb.FeatureSet_UTF8_VALIDATION_UNKNOWN {
				fileOptions.Features = &descriptorpb.FeatureSet{Utf8Validation: test.fileFeature.Enum()}
			}
			fieldOptions := &descriptorpb.FieldOptions{}
			if test.fieldFeature != descriptorpb.FeatureSet_UTF8_VALIDATION_UNKNOWN {
				fieldOptions.Features = &descriptorpb.FeatureSet{Utf8Validation: test.fieldFeature.Enum()}
			}
			file := &descriptorpb.FileDescriptorProto{
				Name:    proto.String("test.proto"),
				Package: proto.String("test"),
				Syntax:  proto.String(test.syntax),
				Options: fileOptions,
				MessageType: []*descriptorpb.DescriptorProto{{
					Name: proto.String("Message"),
					Field: []*descriptorpb.FieldDescriptorProto{{
						Name: proto.String("value"), Number: proto.Int32(1),
						Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), Options: fieldOptions,
					}},
				}},
			}
			if test.edition != descriptorpb.Edition_EDITION_UNKNOWN {
				file.Edition = test.edition.Enum()
			}
			plugin, err := (protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{
				ProtoFile: []*descriptorpb.FileDescriptorProto{file}, FileToGenerate: []string{"test.proto"},
			})
			if err != nil {
				t.Fatal(err)
			}
			genFile(plugin, plugin.Files[0], config{})
			content := plugin.Response().File[0].GetContent()
			gotUnvalidated := strings.Contains(content, "WithoutUTF8Validation")
			if gotUnvalidated != test.wantUnvalidated {
				t.Fatalf("generated unvalidated=%v, want %v:\n%s", gotUnvalidated, test.wantUnvalidated, content)
			}
		})
	}
}

func TestSupportedFeatures(t *testing.T) {
	plugin, err := (protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{})
	if err != nil {
		t.Fatal(err)
	}
	setSupportedFeatures(plugin)

	wantFeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) |
		uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
	if plugin.SupportedFeatures != wantFeatures {
		t.Fatalf("supported features = %d, want %d", plugin.SupportedFeatures, wantFeatures)
	}
	if plugin.SupportedEditionsMinimum != descriptorpb.Edition_EDITION_2023 ||
		plugin.SupportedEditionsMaximum != descriptorpb.Edition_EDITION_2023 {
		t.Fatalf("supported editions = %v..%v, want 2023..2023", plugin.SupportedEditionsMinimum, plugin.SupportedEditionsMaximum)
	}
}

func TestGenerateExplicitScalarPresence(t *testing.T) {
	file := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{GoPackage: proto.String("example.com/test;test")},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:      proto.String("Message"),
			OneofDecl: []*descriptorpb.OneofDescriptorProto{{Name: proto.String("_explicit")}},
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name: proto.String("implicit"), Number: proto.Int32(1),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
				},
				{
					Name: proto.String("explicit"), Number: proto.Int32(2),
					Label:          descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:           descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					Proto3Optional: proto.Bool(true), OneofIndex: proto.Int32(0),
				},
			},
		}},
	}
	plugin, err := (protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{file}, FileToGenerate: []string{"test.proto"},
	})
	if err != nil {
		t.Fatal(err)
	}
	genFile(plugin, plugin.Files[0], config{})
	content := plugin.Response().File[0].GetContent()

	if !strings.Contains(content, "c.Int32(1, &m.Implicit)") {
		t.Fatalf("implicit scalar does not use zero-eliding encoder:\n%s", content)
	}
	if !strings.Contains(content, "c.AlwaysInt32(2, m.Explicit)") {
		t.Fatalf("explicit scalar does not use presence-preserving encoder:\n%s", content)
	}
}

func TestGenerateEnumPresence(t *testing.T) {
	file := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("proto3"),
		Options: &descriptorpb.FileOptions{GoPackage: proto.String("example.com/test;test")},
		EnumType: []*descriptorpb.EnumDescriptorProto{{
			Name: proto.String("State"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: proto.String("STATE_UNSPECIFIED"), Number: proto.Int32(0)},
				{Name: proto.String("STATE_READY"), Number: proto.Int32(1)},
			},
		}},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Message"),
			OneofDecl: []*descriptorpb.OneofDescriptorProto{
				{Name: proto.String("choice")},
				{Name: proto.String("_optional_state")},
			},
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name: proto.String("optional_state"), Number: proto.Int32(1),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(), TypeName: proto.String(".test.State"),
					Proto3Optional: proto.Bool(true), OneofIndex: proto.Int32(1),
				},
				{
					Name: proto.String("selected_state"), Number: proto.Int32(2),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(), TypeName: proto.String(".test.State"),
					OneofIndex: proto.Int32(0),
				},
			},
		}},
	}
	plugin, err := (protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{file}, FileToGenerate: []string{"test.proto"},
	})
	if err != nil {
		t.Fatal(err)
	}
	genFile(plugin, plugin.Files[0], config{})
	content := plugin.Response().File[0].GetContent()

	for _, want := range []string{
		"OptionalState *State",
		"c.AlwaysInt32(1, (*int32)(m.OptionalState))",
		"m.OptionalState = new(State)",
		"c.Int32(1, (*int32)(m.OptionalState))",
		"c.AlwaysInt32(2, (*int32)(&m.SelectedState))",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("generated enum presence code does not contain %q:\n%s", want, content)
		}
	}
}

func TestGenerateEditionsStringBytesPresence(t *testing.T) {
	file := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("test"),
		Syntax:  proto.String("editions"),
		Edition: descriptorpb.Edition_EDITION_2023.Enum(),
		Options: &descriptorpb.FileOptions{GoPackage: proto.String("example.com/test;test")},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Message"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name: proto.String("explicit_string"), Number: proto.Int32(1),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				},
				{
					Name: proto.String("explicit_bytes"), Number: proto.Int32(2),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum(),
				},
				{
					Name: proto.String("implicit_string"), Number: proto.Int32(3),
					Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					Options: &descriptorpb.FieldOptions{Features: &descriptorpb.FeatureSet{
						FieldPresence: descriptorpb.FeatureSet_IMPLICIT.Enum(),
					}},
				},
			},
		}},
	}
	plugin, err := (protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{file}, FileToGenerate: []string{"test.proto"},
	})
	if err != nil {
		t.Fatal(err)
	}
	genFile(plugin, plugin.Files[0], config{})
	content := plugin.Response().File[0].GetContent()

	for _, want := range []string{
		"ExplicitString *string",
		"ExplicitBytes  *[]byte",
		"ImplicitString string",
		"c.AlwaysString(1, m.ExplicitString)",
		"c.AlwaysBytes(2, m.ExplicitBytes)",
		"c.String(3, &m.ImplicitString)",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("generated editions presence code does not contain %q:\n%s", want, content)
		}
	}
}
