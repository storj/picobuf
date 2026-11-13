This program measures the size overhead of the generated code.

It measures:

* library size
* message size
* unused message size
* library size with single field message

It compares them for:

* picobuf
* picobuf conservative
* protobuf conservative

A program enters conservative mode when it calls MethodByName to invoke a
method dynamically. Dead code elimination then removes less code. Protobuf
always triggers conservative mode, so it has no non-conservative comparison.

Library size is the size of the library when a program uses all of its
features.

Message size is the size of a message that has a field of every type.

Unused message size is the cost of such a message when the program never
uses it.

Small size is the overhead of the library and a message with a single field.
Programs often do not need the encoding of every field type.
