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
