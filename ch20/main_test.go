package main

import (
	"testing"
)

func BenchmarkStringPlus1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt1000(b *testing.B) {
	p := initStringi(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder10000(b *testing.B) {
	p := initStrings(10000)
	cap := 10 * len(BLOG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder(p, cap)
	}
}
