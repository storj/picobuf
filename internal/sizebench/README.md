This program tries to measure the overhead of the generated code.

There are several things it tries to measure:

* library size
* message size
* unused message size
* library size with single field message

We compare these across conditions:

* picobuf
* picobuf conservative
* protobuf conservative

Conservative mode happens when program uses MethodByName to dynamically
invoke a method and deadcode elimination is able to remove less code.
Protobuf triggers conservative mode, so there's no alternative comparison.

Library size measures how large the library is when all the
features are being used.

Message size measures size of a message that has all different
fields.

Unused message size measures how much does a unused message
with all fields cost.

Small size measures how large is the overhead of library and message
when there is a message with a single field. It's quite common to
not need the encoding of all field types.
