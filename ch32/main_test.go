package main

import (
	"sync"
	"testing"
	"gotour1/bpool"
)


var bp=bpool.NewBytePoolCap(500,1024,1024)
var sp=&sync.Pool{
	New: func() interface{}{
		return make([]byte,1024,1024)
	},
}

func BenchmarkOpBytePool(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		opBytePool(bp)
	}
}


func BenchmarkOpSyncPool(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		opSyncPool(sp)
	}
	
}