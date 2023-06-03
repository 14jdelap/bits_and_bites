# Bits and Bites

This repo has programs that have helped me better understand how data is represented and manipulated by machines.

All data is ultimately represented by bits (`0`s and `1`s) organized in bytes, groupings of 8 bits. While decimal notation is easier for us 10-fingered humans, binary notation is easier for machines because two-valued signals can be more easily represented:

- The presence or abscence of a hole in a punch card
- High or low voltage on a wire
- A magnetic domain oriented clockwise or counterclockwise

It's precisely electronic circuits, millions or billions of which are packed in a single silicon chip, that represent binary data in computers.

Bits in isolation are not helpful. We need bits and context to derive meaning from the bits. For example, the bits `01000001` can represent the:

- Character `a`
- Signed or unsigned integer `65`
- Single-precision (32 bit) float `9.10E-44`
- Double-precision (64 bit) float `3.21E-322`

## Problems

- Protobuf Varint: encode an integer argument as my implementation of a Protobuf Varint
- Sneaky NaN: encode and decode a message in the mantissa of a floating point
- TCP SYN Flood: parse packets to determine the percentage of incoming SYN messages that were ACKed

Pending

- UTF-8 Truncate: manipulating data in UTF-8 encoding
- CSS color convert: hexadecimal as an output format
- Terminal beeping: ASCII as binary encoding and tty interfacee

## Key concepts

- Big and little endian ordering
- Bitwise operations
- IEEE 754 floating point spec
- Bytes and hexadecimals
- Unicode as the universal identifier of characters
- Character encodings like UTF-8, UTF-16, and UTF-32