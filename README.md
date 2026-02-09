# Go Bloom Filters

[![Checks](https://github.com/DaanV2/go-bloom-filters/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-bloom-filters/actions/workflows/pipeline.yaml)
[![🧬 Fuzz Testing](https://github.com/DaanV2/go-bloom-filters/actions/workflows/fuzz-testing.yaml/badge.svg)](https://github.com/DaanV2/go-bloom-filters/actions/workflows/fuzz-testing.yaml)

A fast and efficient Go implementation of Bloom filters - probabilistic data structures for set membership testing.

## What is a Bloom Filter?

A Bloom filter is a space-efficient probabilistic data structure that is used to test whether an element is a member of a set. It can have false positives (saying an element is in the set when it's not), but never false negatives (it will never miss an element that was added).

## Features

- Multiple hash function support (SHA, FNV, CRC64, etc.)
- Configurable filter size
- Binary marshaling/unmarshaling for persistence
- Zero dependencies (except for testing)
- Fuzz tested for reliability
- Clean, idiomatic Go API

## Installation

```bash
go get github.com/daanv2/go-bloom-filters
```

## Quick Start

```go
package main

import (
    "fmt"
    bloomfilters "github.com/daanv2/go-bloom-filters"
    "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

func main() {
    // Create a new bloom filter with 10,000 bits and default hash functions
    filter, err := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(10000),
        bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
    )
    if err != nil {
        panic(err)
    }

    // Add some items
    filter.Add([]byte("hello"))
    filter.Add([]byte("world"))
    filter.Add([]byte("bloom"))

    // Test membership
    fmt.Println(filter.Test([]byte("hello")))  // true
    fmt.Println(filter.Test([]byte("world")))  // true
    fmt.Println(filter.Test([]byte("bloom")))  // true
    fmt.Println(filter.Test([]byte("filter"))) // false (probably)
}
```

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    bloomfilters "github.com/daanv2/go-bloom-filters"
    "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

func main() {
    // Create a bloom filter with 1024 bits and default hash functions
    filter, _ := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(1024),
        bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
    )

    // Add elements
    emails := []string{"user@example.com", "admin@example.com", "test@example.com"}
    for _, email := range emails {
        filter.Add([]byte(email))
    }

    // Check for existence
    if filter.Test([]byte("user@example.com")) {
        fmt.Println("Email might be in the set")
    }

    if !filter.Test([]byte("unknown@example.com")) {
        fmt.Println("Email is definitely not in the set")
    }
}
```

### Using Different Hash Functions

```go
package main

import (
    bloomfilters "github.com/daanv2/go-bloom-filters"
    "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

func main() {
    // Use all available hash functions for better accuracy
    filter, _ := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(10000),
        bloomfilters.WithHashFunctions(bloomhashes.AllHashFunctions()),
    )

    // Or use specific hash functions
    customFilter, _ := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(10000),
        bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
            bloomhashes.Sha256,
            bloomhashes.Fnv1_64,
            bloomhashes.Crc64_ISO,
        }),
    )

    // Add and test
    filter.Add([]byte("data"))
    customFilter.Add([]byte("data"))
}
```

### Appending Hash Functions

```go
package main

import (
    bloomfilters "github.com/daanv2/go-bloom-filters"
    "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

func main() {
    // Start with default hash functions
    filter, _ := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(10000),
        bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
        // Append additional hash functions
        bloomfilters.WithAppendHashFunctions([]bloomhashes.HashFunction{
            bloomhashes.Fnv1_64,
            bloomhashes.Crc64_ISO,
        }),
    )

    filter.Add([]byte("example"))
}
```

### URL Deduplication Example

```go
package main

import (
    "fmt"
    bloomfilters "github.com/daanv2/go-bloom-filters"
    "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

func main() {
    // Create a bloom filter for tracking visited URLs
    // Using a larger size to reduce false positive rate
    urlFilter, _ := bloomfilters.NewBloomFilter(
        bloomfilters.WithSize(1000000),  // 1M bits for ~125KB memory
        bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
    )

    urls := []string{
        "https://example.com",
        "https://example.com/page1",
        "https://example.com",  // duplicate
        "https://example.com/page2",
        "https://example.com/page1",  // duplicate
    }

    uniqueUrls := 0
    for _, url := range urls {
        urlBytes := []byte(url)
        if !urlFilter.Test(urlBytes) {
            // URL is new
            urlFilter.Add(urlBytes)
            uniqueUrls++
            fmt.Printf("New URL: %s\n", url)
        } else {
            fmt.Printf("Duplicate URL (probably): %s\n", url)
        }
    }

    fmt.Printf("\nTotal unique URLs: %d\n", uniqueUrls)
}
```

## API Reference

### Creating a Bloom Filter

```go
func NewBloomFilter(opts ...BloomFilterOptions) (*BloomFilter, error)
```

Creates a new bloom filter with the specified options. Returns an error if the size is 0 or no hash functions are provided.

### Options

- `WithSize(size uint64)` - Sets the size of the bloom filter in bits
- `WithHashFunctions(hashFunctions []bloomhashes.HashFunction)` - Sets the hash functions to use
- `WithAppendHashFunctions(hashFunctions []bloomhashes.HashFunction)` - Appends additional hash functions

### Methods

- `Add(data []byte)` - Adds an element to the bloom filter
- `Test(data []byte) bool` - Tests if an element might be in the set (returns true for possible membership, false for definite non-membership)

### Available Hash Functions

Default hash functions:
- `Sha1`
- `Sha224`
- `Sha256`
- `Sha512`
- `Sha3_384`

Additional hash functions:
- `Crc64_ISO`
- `Crc64_ECMA`
- `Fnv1_64`
- `Fnv1_64a`
- `Fnv1_128`
- `Fnv1_128a`

Access them via:
- `bloomhashes.DefaultHashFunctions()` - Returns default set of 5 hash functions
- `bloomhashes.AllHashFunctions()` - Returns all 11 available hash functions

## Choosing Parameters

The effectiveness of a Bloom filter depends on:
1. **Size (m)**: Number of bits in the filter
2. **Number of hash functions (k)**: More hash functions = fewer false positives (up to a point)
3. **Number of elements (n)**: Expected number of items to insert

A good rule of thumb:
- Use `m = -n * ln(p) / (ln(2)^2)` bits where `p` is desired false positive rate
- Use `k = (m/n) * ln(2)` hash functions
- For example, for 10,000 items with 1% false positive rate: m ≈ 95,851 bits (~12 KB), k ≈ 7 hash functions

## Performance Considerations

- Adding an element: O(k) where k is the number of hash functions
- Testing an element: O(k)
- Memory usage: Approximately m/8 bytes where m is the size in bits
- More hash functions = more CPU time but fewer false positives

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.