package main

import (
	"github.com/pkeorley/alloc"
	"time"
)

func main() {
	var allocator alloc.Supplier = alloc.NewAllocator()
	allocator.Alloc(alloc.GB)

	time.Sleep(time.Second * 1000)
}
