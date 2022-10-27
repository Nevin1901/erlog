package routines

import (
	"fmt"
	"sync"
	"time"

	"github.com/Nevin1901/arlog/models"
)

var slice []models.ErLog
var wg sync.WaitGroup
var queue chan models.ErLog

func AppendLog(value models.ErLog) {
	if queue == nil {
		queue = make(chan models.ErLog, 1)
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

		models.DB.CreateInBatches(&slice, 512)
		fmt.Println(len(slice))
		slice = nil

		time.Sleep(time.Second)
	}
}