# Bloom Filters

[![Checks](https://github.com/DaanV2/go-bloom-filters/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-bloom-filters/actions/workflows/pipeline.yaml)
[![🧬 Fuzz Testing](https://github.com/DaanV2/go-bloom-filters/actions/workflows/fuzz-testing.yaml/badge.svg)](https://github.com/DaanV2/go-bloom-filters/actions/workflows/fuzz-testing.yaml)

A fast, generic, and concurrent [Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) implementation in Go.

## Installation

```bash
go get github.com/daanv2/go-bloom-filters
```

## Features

- **Standard Bloom filter** — basic, high-performance implementation
- **Concurrent Bloom filter** — safe for concurrent reads and writes using a spinlock
- **Generic wrapper** — use any type with a custom serializer via `GenericBloomFilter[T]`
- **Configurable hash functions** — ships with FNV, CRC-64, SHA, and MD5; bring your own with `WithHashFunctions`
- **Functional options** — clean builder pattern with `WithSize`, `WithDefaultHashFunctions`, etc.

## Quick Start

```go
package main

import (
	"fmt"

	bloomfilters "github.com/daanv2/go-bloom-filters"
)

func main() {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithDefaultHashFunctions(),
	)
	if err != nil {
		panic(err)
	}

	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))

	fmt.Println(bf.Test([]byte("hello"))) // true
	fmt.Println(bf.Test([]byte("world"))) // true
	fmt.Println(bf.Test([]byte("foo")))   // false (probably)
}
```

### Concurrent Bloom Filter

```go
bf, err := bloomfilters.NewConcurrentBloomFilter(
	bloomfilters.WithSize(1024),
	bloomfilters.WithDefaultHashFunctions(),
)
if err != nil {
	panic(err)
}

// Safe to call Add / Test from multiple goroutines
bf.Add([]byte("hello"))
fmt.Println(bf.Test([]byte("hello"))) // true
```

### Generic Bloom Filter

Wrap any `IBloomFilter` to work with a custom type:

```go
bf, _ := bloomfilters.NewBloomFilter(
	bloomfilters.WithSize(1024),
	bloomfilters.WithDefaultHashFunctions(),
)

sbf := bloomfilters.NewGenericBloomFilter(bf, func(s string) []byte {
	return []byte(s)
})

sbf.Add("hello")
fmt.Println(sbf.Test("hello")) // true
```
