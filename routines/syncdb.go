package routines

import (
	"time"

	"github.com/Nevin1901/arlog/models"
)

var logs []models.ErLog

// todo: make append log function which can be synchronized across multiple threads calling it
func AppendLog(log models.ErLog) {
}

func SyncDB(quit chan bool) {
	for {
		if models.DB == nil {
			println("db is nil, skipping")
		}

		time.Sleep(8 * time.Second)
	}
}