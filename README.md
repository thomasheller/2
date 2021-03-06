# 2

[![Go Report Card](https://goreportcard.com/badge/github.com/thomasheller/2)](https://goreportcard.com/report/github.com/thomasheller/2)

Handy power of 2 table/calculator.

## Usage

`2` Prints 2^1 through 2^40 (default)

`2 n` Prints 2^n (single value)

`2 n,m` Prints 2^n and 2^m (multiple values)

`2 n-m` Prints 2^n through 2^m (range)

`2 n-m,x-y,z` Prints 2^n through 2^m, 2^x through 2^y and 2^z

## Example

```text
$ 2

2^x    Value                Approx.       Mnemonic    Byte size
===============================================================
  1                    2                                    2 B
  2                    4                                    4 B
  3                    8                                    8 B
  4                   16                                   16 B
  5                   32                                   32 B
  6                   64                                   64 B
  7                  128     1 hundred                    128 B
  8                  256                                  256 B
  9                  512                                  512 B
 10                1,024    1 thousand        1 KB         1 KB
 11                2,048                                   2 KB
 12                4,096                                   4 KB
 13                8,192                                   8 KB
 14               16,384                                  16 KB
 15               32,768                                  32 KB
 16               65,536                     64 KB        64 KB
 17              131,072                                 128 KB
 18              262,144                                 256 KB
 19              524,288                                 512 KB
 20            1,048,576     1 million        1 MB         1 MB
 21            2,097,152                                   2 MB
 22            4,194,304                                   4 MB
 23            8,388,608                                   8 MB
 24           16,777,216                                  16 MB
 25           33,554,432                                  32 MB
 26           67,108,864                                  64 MB
 27          134,217,728                                 128 MB
 28          268,435,456                                 256 MB
 29          536,870,912                                 512 MB
 30        1,073,741,824     1 billion        1 GB         1 GB
 31        2,147,483,648                                   2 GB
 32        4,294,967,296                      4 GB         4 GB
 33        8,589,934,592                                   8 GB
 34       17,179,869,184                                  16 GB
 35       34,359,738,368                                  32 GB
 36       68,719,476,736                                  64 GB
 37      137,438,953,472                                 128 GB
 38      274,877,906,944                                 256 GB
 39      549,755,813,888                                 512 GB
 40    1,099,511,627,776    1 trillion        1 TB         1 TB
```
