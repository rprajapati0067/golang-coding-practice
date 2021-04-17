package waitgroups

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var cpuUsed = 1
var maxRandomNums = 20

func init() {
	maxCpu := runtime.NumCPU()

	cpuUsed = 4
	runtime.GOMAXPROCS(cpuUsed)

	fmt.Printf("Number of CPUs (Total=%d - Used=%d) \n", maxCpu, cpuUsed)

}

func ConCurrencyAndParallelism() {
	start := time.Now()
	ids := []string{"#", "!", "^", "*"}

	waitGroup.Add(4)
	for i := range ids {
		go numbers(ids[i])
	}
	waitGroup.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\n program took %s \n", elapsed)
}

func numbers(id string) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxRandomNums; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s%d ", id, rand.Intn(20)+20)
	}
	waitGroup.Done()
}

//func WaitGroups() {
//	waitGroup.Add(2)
//
//	go numbers()
//	go alphabets()
//
//	waitGroup.Wait()
//
//}
//
//func numbers() {
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i <= 15; i++ {
//		time.Sleep(200 * time.Millisecond)
//		fmt.Printf("%d ", rand.Intn(20)+20)
//	}
//	waitGroup.Done()
//}
//
//func alphabets() {
//	for i := 'C'; i <= 'G'; i++ {
//		time.Sleep(400 * time.Millisecond)
//		fmt.Printf("%c ", i)
//	}
//	waitGroup.Done()
//}
