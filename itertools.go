package itertools

func Accumulate(vals chan int, f func(int, int) int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		left, ok := <-vals
		if !ok {
			return
		}
		ch <- left
		for {
			var right int
			right, ok = <-vals
			if !ok {
				return
			}
			v := f(left, right)
			ch <- v
			left = v
		}
	}()
	return ch
}

func Chain(chans ...chan int) chan int {
	chanOfChans := make(chan (chan int))
	go func() {
		defer close(chanOfChans)
		for _, ch := range chans {
			chanOfChans <- ch
		}
	}()
	return ChainFromIterable(chanOfChans)
}

func ChainFromIterable(chans chan (chan int)) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for in := range chans {
			for x := range in {
				out <- x
			}
		}
	}()
	return out
}

func Count(start int, step int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		n := start
		for {
			ch <- n
			n += step
		}
	}()
	return ch
}

func Consume(ch chan int, n int) {
	for i := 0; i < n; i++ {
		_, ok := <-ch
		if !ok {
			break
		}
	}
}

func Cycle(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var vals []int
		for v := range in {
			out <- v
			vals = append(vals, v)
		}
		for {
			for _, v := range vals {
				out <- v
			}
		}
	}()
	return out
}

func NCycle(in chan int, n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		if n <= 0 {
			return
		}
		var vals []int
		for v := range in {
			out <- v
			vals = append(vals, v)
		}
		for i := 0; i < n-1; i++ {
			for _, v := range vals {
				out <- v
			}
		}
	}()
	return out
}

func Nth(ch chan int, n int) (int, bool) {
	if n < 0 {
		return 0, false
	}
	Consume(ch, n)
	v, ok := <-ch
	return v, ok
}

func Take(in chan int, n int) []int {
	var vals []int
	for i := 0; i < n; i++ {
		v, ok := <-in
		if !ok {
			break
		}
		vals = append(vals, v)
	}
	return vals
}

func TakeAll(in chan int) []int {
	var vals []int
	for v := range in {
		vals = append(vals, v)
	}
	return vals
}

func RepeatN(v int, n int) chan int {
	return NCycle(single(v), n)
}

func Repeat(v int) chan int {
	return Cycle(single(v))
}

func RepeatFuncN(f func() int, n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- f()
		}
	}()
	return out
}

func RepeatFunc(f func() int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			out <- f()
		}
	}()
	return out
}

func single(v int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		out <- v
	}()
	return out
}
