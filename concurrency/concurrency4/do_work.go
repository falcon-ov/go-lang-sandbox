package concurrency4

func DoWork(n int) int {
	result := 0
	for i := 0; i < 500; i++ {
		result += (n + i) * (n - i) % 17
	}
	return result
}
