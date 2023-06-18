# UTF-8 Truncate

## Assumptions

- Restriction: 280 **byte** limit for text in a social network
- The first byte per line in "cases" is an unsigned integer with the number of bytes to which to truncate
- Avoid truncating the string in the middle of a single Unicode codepoint: in that case err on the side of more rather than less bytes
  - Why: some characters (e.g. emoticons) require multiple bytes to encode, so truncating the string to 280 bytes may leave junk bytes trailing

## Process

1. Program should read the "cases" file
2. Program should parse and write one correctly truncated string per line
3. Program should compare the result against the "expected" file

## Context

ASCII is limited to 1 byte and is what we most use, but to encode more complex characters we need multiple bytes.

## Questions

- **Assume 1 character (even if > 1 byte) is 1 byte or does 1 byte strictly mean 1 byte?**

## Steps

1. Open the `cases` file
2. Understand how `binary.Read` works to read per one line
3. Interpret the first byte as an unsigned integer to get the byte length the line has to be truncated to
4. Truncation
  - If the byte length < 280, then return the entire line untruncated
  - If the byte length is > 280, then according to how "byte" is defined
    - Iterate over every byte and count the bytes (if 1 byte is a unicode point)
    - Go up to 280 - `MAX_CHARACTER_LENGTH` in the `[]byte` and then iterate over the remaining bytes to ensure no junk byte is left if the truncation is in the middle of a Unicode point
5. Write the results, line by line, into a new file
6. Do a diff between the generated and expected files (in CLI or programmatically? Latter is better)
