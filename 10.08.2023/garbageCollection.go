package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func printMemStat(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)
	fmt.Println("-------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total bytes of heap objects: ", ms.TotalAlloc)
	fmt.Println("bytes obtained from operating system: ", ms.Sys)
	fmt.Println("count of heap objects :", ms.Mallocs)
	fmt.Println("count heap objects deallocated", ms.Frees)
}

func main() {
	var ms runtime.MemStats
	printMemStat(ms)

	intArr := make([]int, 90000)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}

	time.Sleep(5 * time.Second)

	printMemStat(ms)
}
