package unique

func Unique(nums <-chan int, result chan<- int) {
	m := make(map[int]struct{})
	for v := range nums {
		m[v] = struct{}{}
	}

	for k := range m {
		result <- k
	}
	close(result)
}
