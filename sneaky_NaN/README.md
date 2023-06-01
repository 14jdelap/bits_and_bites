# Sneaky NaN

## Goals

Encode a message in a `NaN` floating point.

1. Encode a message in a floating-point value with an `conceal` function.
2. Decode said message with an `extract` function

## Floats under the hood

Floating point numbers, as defined in the IEEE 754 spec, are composed of 3 parts:

1. Signbit
2. Fraction (`a`)
3. Exponent (`b`)

These 3 parts are used by the computer to calculate the float's result. This is because floats use scientific notation to represent numeric values: `a * 10 ^ b`.

### Why floats exist

With integers you can only write 100 values with 2 integers: 0 to 99.

With floats, since you use `a * 10 ^ b` you can also only express 100 values but of a much wider range: from 0 to 9 billion.

|Number|`b`|`a`|
|-|-|-|
|0|0|0|
|1|0|1|
|2|0|2|
|3|0|3|
|4|0|4|
|5|0|5|
|6|0|6|
|7|0|7|
|8|0|8|
|9|0|9|
|10|1|1|
|20|1|2|
|30|1|3|
|...|...|...|
|100|2|1|
|200|2|2|
|300|2|3|
|...|...|...|
|1.000|3|1|
|...|...|...|
|10.000|4|1|
|...|...|...|
|100.000|5|1|
|...|...|...|
|1.000.000.000|9|1|
|...|...|...|
|9.000.000.000|9|9|

The caveat is that you lose precision. 0 to 9 can be exactly represented, but 11 can't be encoded exactly - the next highest number than can be is 20.

Thus, the strides of representable numbers increases with `b`. When:

- `b = 0` all 10 integers are perfectly representable
- `b = 1` 10/100 are perfectly representable
- `b = 2` 10/1.000 are perfectly representable
- `b = 3` 10/10.000 are perfectly representable
- `b = 9` 10/10.000.000.000.000 are perfectly representable

**Use cases**

Floats are perfect for scientific work.

As a scientist you work to a certain degree of precision within a certain order of magnitude. Thus, floats can be ideal as you can set both of these.

**Floats shouldn't be used for currency because of its inherent imprecision**.

### `NaN`s as floats

`NaN`s represent floating point numbers that are not a number (e.g. `0/0`, square root of `-1`).

In the IEEE 754 spec, a float is a NaN that satisfies 2 conditions:

- The exponent is composed only of `1` digits
- The fraction is non-zero

## Implementation

Assumption: use `float64` type

- Signbit: 1 bit
- Exponent: 11 bits
- Fraction: 53 bits
  - If using 8 bit characters: 6 bytes/characters
  - Implication: 5 spare bits (can represent 0 to 31)
    - Break into:
      - First bit always on (1)
      - Next 4 bits represent length? Though not necessary here (I think) -> 0 to 15
  - Alternative: use 7 bit characters (using ASCII format)
    - First bit always on
    - Next 4 bits represent length of concealed message (0 to 15)
    - Next 49 bits store 0 to 7 binary representations of 7 bit ASCII characters

**Implementation #1**: using 8 bit characters

- Check if all the characters are ASCII (i.e. up to 7 bits in binary)
  - If not, raise an error
- Convert the characters into 8 bit binary
  - Create a new output variable of type string
  - Iterate over a string: for all its length or until the length is 6 (because there can only be 6 characters)
    - Append to the new string the 8 bit binary representation of the current iteration's character
  - If string is less than 48, pad it with 0s until it's of length 48
- Create the IEEE-754 string
  - Signbit: `1`
  - Exponent: `11111111111`
  - Fraction: `10` + length (in 3 bits) + binary of message
    - Examples
      - `a`: `10` + `001` + `01100001` + 5 bytes of 0s
      - `apple`: `10` + `101` + `01100001 01110000 01110000 01101100 01100101` + `00000000`
- Create a `NaN` with the IEEE-754 string
- Decode the string
  - Access the fraction
  - Check the 0-index binary 2-4 to get the length
  - For length, get each successive byte from but 5 to 53 and convert them to ASCII characters
  - Return the decoded string

**Implementation #2**: using 7 bit characters

Why: using ASCII characters we use 7 bits, so we're able to squeeze in an additional character.

## First try: reflextion

- Working with bits in strings was a headache
  - Having a string that I would increase piece by piece
  - Converting ints to bits with 3-bit formatting
  - Parsing the bit string to convert it to a `uint64` to then convert it to the `float64`
  - Then doing the same in reverse order for the decoding: `float64` to `uint64` to a string of bits to having to work out the maths for the right slicing of the array segments to convert them to bytes
- Ideally, I would like to work directly with `uint64` or a similar format — `byte`s are probably a good options
  - Challenge: that involves knowing what bits to insert at each byte
    - Byte 1: `1111 1111`
    - Byte 2: `1111 1` + msg length
    - Byte 3 to 8: each character

Key data structures:

- String
  - Store the IEEE 754 in binary
- Byte (aslias for uint8, which ranges from 0 to 255)
