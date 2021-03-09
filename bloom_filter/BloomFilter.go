package bloom_filter

import (
	"BloomFilter/bitmap"
)

type BloomFilter struct {
	DefultSize  uint32
	localBitmap *bitmap.Bitmap
}

func New() *BloomFilter {
	bloomFilter := &BloomFilter{}
	bloomFilter.DefultSize = 2 << 28
	bloomFilter.localBitmap = bitmap.New(bloomFilter.DefultSize)
	return bloomFilter
}

func (bloomFilter *BloomFilter) IsContain(str string) bool {

	if !bloomFilter.localBitmap.Has(hash1(str) & (bloomFilter.DefultSize - 1)) {
		return false
	}

	if !bloomFilter.localBitmap.Has(hash2(str) & (bloomFilter.DefultSize - 1)) {
		return false
	}

	if !bloomFilter.localBitmap.Has(hash3(str) & (bloomFilter.DefultSize - 1)) {
		return false
	}
	return true
}

func (bloomFilter *BloomFilter) Add(str string) {
	bloomFilter.localBitmap.Add(hash1(str) & (bloomFilter.DefultSize - 1))
	bloomFilter.localBitmap.Add(hash2(str) & (bloomFilter.DefultSize - 1))
	bloomFilter.localBitmap.Add(hash3(str) & (bloomFilter.DefultSize - 1))
}

func hash1(str string) uint32 {

	if len(str) == 0 {
		return 0
	}

	chars := []rune(str)
	lastIndex := len(chars) - 1
	hashcode := 0
	for i := range chars {
		if i == lastIndex {
			hashcode += int(chars[i])
		}
		//31 * i ==  (i << 5) - i
		//更好的分散hash
		hashcode += (hashcode + int(chars[i])) * 31
	}
	return uint32(hashcode)
}

func hash2(str string) uint32 {
	return hash1(str + "2")
}

func hash3(str string) uint32 {
	return hash1(str + "3")
}
