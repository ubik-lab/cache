package cahce

import (
	"fmt"
	"testing"

	"github.com/ubik-lab/cache/lru"

	"github.com/hashicorp/golang-lru/simplelru"
)

func BenchmarkLRUSet(b *testing.B) {
	const items = 1 << 16
	c, _ := lru.New[int, int](256, nil)

	b.ReportAllocs()
	b.SetBytes(items)

	for i := 0; i < b.N; i++ {
		for j := 0; j < items; j++ {
			c.Add(j, j+i)
		}
	}
}

func BenchmarkHoshiCorpkLRUSet(b *testing.B) {
	const items = 1 << 16
	c, _ := simplelru.NewLRU(256, nil)

	b.ReportAllocs()
	b.SetBytes(items)

	for i := 0; i < b.N; i++ {
		for j := 0; j < items; j++ {
			c.Add(j, j+i)
		}
	}
}

func BenchmarkLRUGet(b *testing.B) {
	const items = 1 << 16
	c, _ := lru.New[int, int](256, nil)

	// Set
	for j := 0; j < items; j++ {
		c.Add(j, j)
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < items; j++ {
			if v, ok := c.Get(j); ok && v != j {
				panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", v, j))
			}
		}
	}
}

func BenchmarkHoshiCorpkLRUGet(b *testing.B) {
	const items = 1 << 16
	c, _ := simplelru.NewLRU(256, nil)

	// Set
	for j := 0; j < items; j++ {
		c.Add(j, j)
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < items; j++ {
			if v, ok := c.Get(j); ok && v != j {
				panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", v, j))
			}
		}
	}
}
