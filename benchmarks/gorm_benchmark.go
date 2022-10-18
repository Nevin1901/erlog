package benchmarks

import "github.com/Nevin1901/arlog/models"

func TestGorm_Benchmark() {
	var test_log models.ErLog
	for i := 0; i < 1000; i++ {
		test_log = models.ErLog{LogType: "info", Message: "test message", Title: "test title"}
	}

	println(test_log)
	// make these have the ones from readable
}
