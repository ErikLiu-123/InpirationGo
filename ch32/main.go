package main

import(
	//"fmt"
	"sync"
	"os"
	"io"
	"gotour1/bpool"
)


func main() {
	bp := bpool.NewBytePoolCap(500,1024,1024)
	buf:=bp.Get()
	defer bp.Put(buf)

	//使用buf，不再举例
}

func opBytePool(bp *bpool.BytePoolCap){
	var wg sync.WaitGroup
	wg.Add(500)
	for i:=0;i<500;i++{
		go func(bp *bpool.BytePoolCap){
			buffer:=bp.Get()
			defer bp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(bp)
	}
	wg.Wait()
}

func opSyncPool(sp *sync.Pool){
	var wg sync.WaitGroup
	wg.Add(500)
	for i:=0;i<500;i++{
		go func(sp *sync.Pool){
			buffer:=sp.Get().([]byte)
			defer sp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(sp)
	}
	wg.Wait()
}

func mockReadFile(b []byte){
	f,_:=os.Open("Water")
	for {
		n,err:=io.ReadFull(f,b)
		if n==0 || err==io.EOF{
			break
		}
	}
}