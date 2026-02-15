# Hash Cost

using the benchmarks in `tests/benchmarks/bloomhashes` we can index the costs of hashes in terms of time and allocations.

```log
goos: windows
goarch: amd64
pkg: github.com/daanv2/go-bloom-filters/tests/benchmarks/bloomhashes
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
Benchmark_Hashes_Cost/MD5-8               117894             97986 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha1-8               51582            220881 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha224-8             59685            205605 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha3_384-8           28364            403428 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha256-8             64435            187986 ns/op               8 B/op          0 allocs/op
Benchmark_Hashes_Cost/Sha512-8             49522            246200 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Crc64_ISO-8         195788             60333 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Crc64_ECMA-8        196144             60534 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_64-8           490876             23990 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_64a-8          461137             23895 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_128-8          195676             61446 ns/op               0 B/op          0 allocs/op
Benchmark_Hashes_Cost/Fnv1_128a-8         192081             62571 ns/op               0 B/op          0 allocs/op
```

conclusions sorted by time, higher is faster:

| Hash         |       N |          Time | Bytes  | Allocs      |
| ------------ | ------: | ------------: | ------ | ----------- |
| Fnv1_64a-8   | 461.137 |  23.895 ns/op | 0 B/op | 0 allocs/op |
| Fnv1_64-8    | 490.876 |  23.990 ns/op | 0 B/op | 0 allocs/op |
| Crc64_ISO-8  | 195.788 |  60.333 ns/op | 0 B/op | 0 allocs/op |
| Crc64_ECMA-8 | 196.144 |  60.534 ns/op | 0 B/op | 0 allocs/op |
| Fnv1_128-8   | 195.676 |  61.446 ns/op | 0 B/op | 0 allocs/op |
| Fnv1_128a-8  | 192.081 |  62.571 ns/op | 0 B/op | 0 allocs/op |
| MD5-8        | 117.894 |  97.986 ns/op | 0 B/op | 0 allocs/op |
| Sha256-8     |  64.435 | 187.986 ns/op | 8 B/op | 0 allocs/op |
| Sha224-8     |  59.685 | 205.605 ns/op | 0 B/op | 0 allocs/op |
| Sha1-8       |  51.582 | 220.881 ns/op | 0 B/op | 0 allocs/op |
| Sha512-8     |  49.522 | 246.200 ns/op | 0 B/op | 0 allocs/op |
| Sha3_384-8   |  28.364 | 403.428 ns/op | 0 B/op | 0 allocs/op |