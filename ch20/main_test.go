package main

import (
	"testing"
)

func BenchmarkStringPlus10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt10000(b *testing.B) {
	p := initStringi(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func BenchmarkStringJoin10000(b *testing.B) {
	p := initStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin(p)
	}
}

func BenchmarkStringBuffer10000(b *testing.B) {
	p := initStrings(10000)
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
