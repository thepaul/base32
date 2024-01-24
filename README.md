# base32

Base32 encoding and decoding, using the RFC 4648 alphabet ("abcdefghijklmnopqrstuvwxyz234567").

## Usage

By default, encodes all input into base32 and outputs the result.

With `-d`/`--decode`, decodes all input _from_ base32 and outputs the result.

With `-x`/`--hex`, expects hexadecimal input (or, if decoding, produce hexadecimal output).

## Examples

```
$ echo hello | base32
nbswy3dpbi

$ echo af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e7190800000000 | base32 -x
v4weeab67sbgvnbwd5z7tweqsqqun7qox2agpbxy44mqqaaaaaaa

$ echo v4weeab67sbgvnbwd5z7tweqsqqun7qox2agpbxy44mqqaaaaaaa | base32 -dx
af2c42003efc826ab4361f73f9d890942146fe0ebe806786f8e7190800000000

$ echo nbswy3dpbi | base32 -d
hello
```
