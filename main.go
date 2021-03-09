package main

import (
	"BloomFilter/bloom_filter"
)

func main() {
	bloomFilter := bloom_filter.New()
	bloomFilter.Add("hello")
	bloomFilter.Add("fuck")
	print(bloomFilter.IsContain("hello"))
	print(bloomFilter.IsContain("fuck"))
}
