# Bitwise operations

These files are explorations of how to work with bits, bytes, and bitwise operations in Go.

The functions I've written are:

- `isAnyBitOn`: determines if an `int32` has a byte with one or more bits that are on
- `getEndian`: determines the byte ordering used by the machine running the function
- `showBytes`: prints the hex representation of all the bytes in a variable
- `showInt`, `showString`, and `showFloat`: convenience functions to work with `showBytes`

## Further work

- How can isAnyBitOn take a value of any type and still work?
  - Requirements
    - Should be able to take an argument of any type
    - Should be able to determine the byte length of the argument
