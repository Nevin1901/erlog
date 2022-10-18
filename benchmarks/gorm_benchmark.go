package benchmarks

func Gorm_Benchmark() {
	var test_log models.ErLog
	for i := 0; i < 1000; i++ {
		test_log := models.ErLog{LogType: "info", Message: "test message", Title: "test title"}
	}
	// make these have the ones from readable
}
