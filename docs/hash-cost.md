# Hash Cost

using the benchmarks in `tests/benchmarks/bloomhashes` we can index the costs of hashes in terms of time and allocations.

```log
goos: windows
goarch: amd64
pkg: github.com/daanv2/go-bloom-filters/tests/benchmarks/bloomhashes
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
Benchmark_Hashes_Cost/Sha1-8                5696            220097 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha224-8              5649            200529 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha3_384-8            2704            378257 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha256-8              6446            178860 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha512-8              4866            234646 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Crc64_ISO-8          20148             59614 ns/op               1 B/op          0 allocs/op
Benchmark_Hashes_Cost/Crc64_ECMA-8         20125             59497 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_64-8            54276             22083 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_64a-8           53697             22135 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_128-8           19348             63846 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_128a-8          18854             62965 ns/op               0 B/op          0 allocs/op
```

conclusions sorted by time, higher is faster:

| Hash         |      N |          Time | Bytes  | Allocs      |
| ------------ | -----: | ------------: | ------ | ----------- |
| Fnv1_64-8    | 54.276 |  22.083 ns/op | 0 B/op | 0 allocs/op |
| Fnv1_64a-8   | 53.697 |  22.135 ns/op | 0 B/op | 0 allocs/op |
| Crc64_ECMA-8 | 20.125 |  59.497 ns/op | 0 B/op | 0 allocs/op |
| Crc64_ISO-8  | 20.148 |  59.614 ns/op | 1 B/op | 0 allocs/op |
| Fnv1_128a-8  | 18.854 |  62.965 ns/op | 0 B/op | 0 allocs/op |
| Fnv1_128-8   | 19.348 |  63.846 ns/op | 0 B/op | 0 allocs/op |
| Sha256-8     |  6.446 | 178.860 ns/op | 0 B/op | 0 allocs/op |
| Sha224-8     |  5.649 | 200.529 ns/op | 0 B/op | 0 allocs/op |
| Sha1-8       |  5.696 | 220.097 ns/op | 0 B/op | 0 allocs/op |
| Sha512-8     |  4.866 | 234.646 ns/op | 0 B/op | 0 allocs/op |
| Sha3_384-8   |  2.704 | 378.257 ns/op | 0 B/op | 0 allocs/op |