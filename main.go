package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/pbnjay/memory"
	"math"
)

func occupyMemory(slice *[]int, bytesToOccupy int) {
	for i := 0; i < bytesToOccupy/8; i++ {
		*slice = append(*slice, 1)
	}
}

func main() {
	var MB = int(math.Pow(2, 20))
	var GB = int(math.Pow(2, 30))

	toOccupy := 0 // value to occupy bytes
	slice := make([]int, 0, GB)

	for {
		var availableMemory = humanize.Bytes(memory.FreeMemory())
		fmt.Print("[" + humanize.Bytes(uint64(len(slice)*8)) + " of " + availableMemory + "] How many megabytes do you want to use? ")
		_, err := fmt.Scanln(&toOccupy)
		if err != nil {
			return
		}

		occupyMemory(&slice, toOccupy*MB)
	}

}
