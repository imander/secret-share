package cache

import (
	"testing"
	"time"
)

func BenchmarkCache(b *testing.B) {
	c := New(time.Second)
	s := randString()
	for n := 0; n < b.N; n++ {
		c.AddWithTTL(s, time.Second)
	}
}
