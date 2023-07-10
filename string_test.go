package nstring_test

import (
	"github.com/smartwalle/nstring"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nstring.Sub("testhahahaha", 2, 5)
	}
}

func BenchmarkChunk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nstring.Chunk("1234567", 2)
	}
}
