package itertools

func Count(start int, step int) chan int {
	ch := make(chan int)
	go func() {
		n := start
		for {
			ch <- n
			n += step
		}
	}()
	return ch
}
