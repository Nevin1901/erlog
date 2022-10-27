package models

import (
	"log"
	"testing"
	"time"
)

func TestSaveModelSingle(t *testing.T) {
	ConnectDB("benchmark.db")
	start := time.Now()
	var test_log ErLog
	for i := 0; i < 1000; i++ {
		test_log = ErLog{LogType: "info", Message: "test message", Title: "test title"}
		DB.Create(&test_log)
	}

	// fmt.Printf("%+v", test_log)
	_ = test_log
	elapsed := time.Since(start)
	log.Printf("Single: Took %s", elapsed)
}


func TestSaveModelBatch(t *testing.T) {
	ConnectDB("benchmark.db")
	start := time.Now()
	var test_log []ErLog
	for i := 0; i < 1000; i++ {
		test_log = append(test_log, ErLog{LogType: "info", Message: "test message", Title: "test title"})
	}

	DB.CreateInBatches(&test_log, 512)

	_ = test_log
	elapsed := time.Since(start)
	log.Printf("Batches: Took %s", elapsed)
}