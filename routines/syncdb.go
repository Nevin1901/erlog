package routines

import (
	"fmt"
	"sync"
	"time"

	"github.com/Nevin1901/arlog/models"
)

var slice []int
var wg sync.WaitGroup
var queue chan int

// func SetupSync() {
// 	wg = sync.WaitGroup{}
// 	queue = make(chan int, 1)
// }

// todo: make append log function which can be synchronized across multiple threads calling it
func AppendLog(value int) {
	if queue == nil {
		println("queue is nil")
		queue = make(chan int, 1)
	}

	wg.Add(1)
	queue <- value
}

func SyncDB() {
	for t := range queue {
		slice = append(slice, t)
		wg.Done()
	}
}

func SetupSync() {
	for {
		if models.DB == nil {
			println("db is nil, skipping")
			continue
		}

		go SyncDB()
		wg.Wait()

		fmt.Println(slice)
		slice = nil

		time.Sleep(time.Second)
	}
}