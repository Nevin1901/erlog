package models

import (
	"log"
	"testing"
	"time"
)

func TestSaveModel(t *testing.T) {
	start := time.Now()
	var test_log ErLog
	for i := 0; i < 1001; i++ {
		test_log = ErLog{LogType: "info", Message: "test message", Title: "test title"}
	}

	// fmt.Printf("%+v", test_log)
	_ = test_log
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

	// println(elapsed)
	// make these have the ones from readable
}
