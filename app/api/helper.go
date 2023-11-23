package handler

func factorialCount(n int, ch chan int) {
	result := 1

	if n == 0 || n == 1 {
		ch <- 1
		return
	}

	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
}
