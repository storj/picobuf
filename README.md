# Picobuf

Picobuf is a minimal replacement for [google.golang.org/protobuf](https://github.com/protocolbuffers/protobuf-go)
that focuses on small binary size.

It does not support the whole Protocol Buffers feature set. Features get added when someone needs them.

Picobuf is in beta and its API can change.

## How?

You generate Go code from proto files with a `protoc` plugin, as you do with the official implementation.

First, install the `protoc` protobuf compiler.

Second, install the picobuf plugin with:
```
go install storj.io/picobuf/protoc-gen-pico@latest
```

Finally, add a `go:generate` line for your proto file:

```
//go:generate protoc  -I. --pico_out=paths=source_relative:. example.proto
```

`example.proto` is an ordinary proto file:
```
syntax = "proto3";
option go_package = "example.test/pb";

package example;

message Person {
  string  name = 1;
  Address address = 2;
}

message Address {
  string street = 1;
}
```

Picobuf also has options that the official implementation lacks.
[pico.proto](pico.proto) defines them, so it has to be on the `protoc` include
path. One way to add it is `-I$(go list -m -f '{{.Dir}}' storj.io/picobuf)`.

```
syntax = "proto3";
option go_package = "example.test/pb";

package example;

import "pico.proto";

message Timestamp {
  int64 seconds = 1;
  int32 nanos = 2;
}

message Person {
  string name = 1;

  Timestamp created_at = 5 [
    // "always_present" generates the field by value instead of by pointer.
    // That saves an allocation, but the field can no longer tell absent from empty.
    (pico.field).always_present = true,
    // "custom_type" replaces the generated Go type of the field with your own type.
    (pico.field).custom_type = "time.Time",
    // "custom_serialize" names the type that encodes and decodes the field.
    // `time.Time` needs one, because we cannot add methods to it.
    (pico.field).custom_serialize = "storj.io/picobuf/picoconv.Timestamp"
  ];
}
```

[picoconv.Timestamp](picoconv/timestamp.go) shows how to implement custom serialization.

### JSON

By default picobuf implements only the protobuf binary wire format. Go's
`encoding/json` is not a substitute for ProtoJSON. It can produce different
JSON for a generated struct, and it fails on fields generated with
`value_presence`.

Generating with `--pico_opt=protojson=true` adds `MarshalJSON` and
`UnmarshalJSON` methods that implement the
[ProtoJSON mapping](https://protobuf.dev/programming-guides/json/), using
[picojson](picojson). Before Go 1.27, picojson uses
[go-json-experiment/json](https://github.com/go-json-experiment/json) in
place of encoding/json/jsontext. The conversion goes through the wire format,
so a custom type maps to JSON as the protobuf type that its field declares.
Messages from other packages that yours refer to do not need the option.
Tests compare the mapping against protojson. It differs in two ways:

* picojson does not support `google.protobuf.Any`, `Struct`, `Value`,
  `ListValue`, or group fields. Generation fails if a message refers to them.
* The parser follows the JSON specification, so it rejects malformed input
  that protojson accepts, such as the number `0e`.

The Editions `features.json_format` setting does not change the generated
code. `protoc` still checks the schema requirements of `ALLOW` before it
invokes the generator.

## Why?

We wanted smaller binaries. At the time of writing, the official library
calls methods through reflection, which switches the compiler into
conservative mode. In that mode much of dead code elimination stops working
and the binaries bloat.

Picobuf also supports custom types, so a field can use a different Go
representation than the official implementation gives it.

## Should I use it?

Probably not. The official implementation supports many features that picobuf
lacks, and for most programs it is fine.

Picobuf makes sense when:

* you need a small binary, such as a WASM module, a C library, or a TinyGo
  build for an embedded system
* you want custom serialization, which avoids some allocations
* you want custom types, which avoid annoying conversions across the codebase

You only get the smaller binary when nothing else pulls in the official
protobuf library, or any other package that switches the compiler into
conservative mode, such as `text/template` or `html/template`.

Do not use picobuf when you need:

* `Any` messages
* reflection
* descriptors
* JSON for `google.protobuf.Any` or `Struct`, see [JSON](#json)

## Benchmark

[sizebench](./internal/sizebench) measures the binary size.

```
go1.19.5
goos: linux
goarch: amd64
// A basic "fmt.Println" executable size.
ExeBase               1     1814.766 KB
// A basic "fmt.Println" executable size, but contains code
// that switches the compiler into conservative mode.
ExeBaseConserv        1     2549.444 KB
// Picobuf generated binary size (non-conservative).
ExePico/Sml           1     2000.987 KB      186.221 KB-Δ
// Picobuf generated binary size (conservative).
ExePicoConserv/Sml    1     2700.553 KB      151.109 KB-Δ
// Protobuf generated binary size (conservative).
ExeProto/Sml          1     4978.815 KB     2429.371 KB-Δ
```

The message definitions are in [msg-sml.proto](./internal/sizebench/msg-sml.proto).

The protobuf binary is about 3 MB bigger than the picobuf binary in
non-conservative mode, and about 2.3 MB bigger than the one in conservative
mode.
