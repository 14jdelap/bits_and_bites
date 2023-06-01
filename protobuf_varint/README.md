# Protobuf Base 128 varint

Encoder takes a `uint64` as an argument and returns a string representing a hexadecimal value.

Decoder takes a string hexadecimal and returns a `uint64`.

```go
func encode(i uint64) string {
  ...
}
func decode(s string) uint64 {
  ...
}
```

The hexadecimal can be between 1 and 10 bytes long in little endian ordering.

Each byte is composed of 2 parts:

- Sign: first bit, which marks if there's a successive byte
  - If final byte: sign bit is off
- Payload: next 7 bits, which when concatenated form the encoded value in little endian format

## `encode`

- In a loop, take the lowest order 7 bits (bitwise XOR operator against `0000000`)
  - `while n > 0`
  - What does it mean to use a bitmask?
- Add the correct MSB
- Push to a `[]byte`
- Reduce `n` by 7 bits through bitwise shifting
- Return a byte sequence


Questions:

- Should I convert the `int` to `uint64` or another type?
- How can I programmatically move the bits to set the sign bit and move the payload bit to a new byte?
