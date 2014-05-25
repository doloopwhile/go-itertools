package itertools

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func values(vals ...int) chan int {
	ch := make(chan int)
	go func() {
		for _, x := range vals {
			ch <- x
		}
		close(ch)
	}()
	return ch
}

func array(ch <-chan int) []int {
	var arr []int
	for x := range ch {
		arr = append(arr, x)
	}
	return arr
}

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Itertools Suite")
}

var _ = Describe("Itertools", func() {
	Context("Accumulate", func() {
		It("should return a list of application of function", func() {
			ch := Accumulate(
				values(1, 2, 3, 4, 5),
				func(x int, y int) int { return x + y },
			)
			vals := array(ch)
			Expect(vals).To(Equal([]int{1, 3, 6, 10, 15}))
		})
	})

	Context("Chain", func() {
		It("should return concatinated channel", func() {
			ch := Chain(values(1, 2, 3), values(4, 5), values(), values(6))
			var vals []int
			for i := 0; i < 6; i++ {
				v, ok := <-ch
				vals = append(vals, v)
				Expect(ok).To(BeTrue())
			}
			Expect(vals).To(Equal([]int{1, 2, 3, 4, 5, 6}))

			_, ok := <-ch
			Expect(ok).To(BeFalse())
		})
		It("should return nil channel if no argument", func() {
			ch := Chain()
			_, ok := <-ch
			Expect(ok).To(BeFalse())
		})
	})

	Context("ChainFromIterable", func() {
		It("should return concatinated channel", func() {
			in := make(chan (chan int))
			go func() {
				defer close(in)
				in <- values(1, 2, 3)
				in <- values(4, 5)
				in <- values()
				in <- values(6)
			}()

			ch := ChainFromIterable(in)
			var vals []int
			for i := 0; i < 6; i++ {
				v, ok := <-ch
				vals = append(vals, v)
				Expect(ok).To(BeTrue())
			}
			Expect(vals).To(Equal([]int{1, 2, 3, 4, 5, 6}))

			_, ok := <-ch
			Expect(ok).To(BeFalse())
		})
	})

	Context("Consume", func() {
		It("should advance the iterator n-steps ahead", func() {
			ch := values(1, 2, 3, 4, 5)
			Consume(ch, 2)
			vals := array(ch)
			Expect(vals).To(Equal([]int{3, 4, 5}))
		})
		It("should return empty channel if n is larger than length of input channel", func() {
			ch := values(1, 2, 3, 4, 5)
			Consume(ch, 10)
			vals := array(ch)
			Expect(vals).To(BeNil())
		})
	})

	Context("Count", func() {
		It("should iterate integer in step by step", func() {
			ch := Count(1, 2)
			Expect(<-ch).To(Equal(1))
			Expect(<-ch).To(Equal(3))
			Expect(<-ch).To(Equal(5))
		})
	})

	Context("Cycle", func() {
		It("should cyclicly repeat elements", func() {
			ch := Cycle(values(1, 2, 3))
			vals := Take(ch, 12)
			Expect(vals).To(Equal([]int{
				1, 2, 3,
				1, 2, 3,
				1, 2, 3,
				1, 2, 3,
			}))
		})
	})

	Context("Nth", func() {
		It("should return nth element of the channel", func() {
			val, ok := Nth(values(1, 2, 3, 4), 2)
			Expect(ok).To(BeTrue())
			Expect(val).To(Equal(3))
		})
		It("should return error if out of range", func() {
			val, ok := Nth(values(1, 2, 3, 4), 4)
			Expect(ok).To(BeFalse())
			Expect(val).To(Equal(0))
		})
		It("should return error if less than 0", func() {
			val, ok := Nth(values(1, 2, 3, 4), -1)
			Expect(ok).To(BeFalse())
			Expect(val).To(Equal(0))
		})
		It("should return error if channel is empty", func() {
			val, ok := Nth(values(), 2)
			Expect(val).To(Equal(0))
			Expect(ok).To(BeFalse())
		})
	})

	Context("NCycle", func() {
		It("should return a channel which iterate elements n times", func() {
			ch := NCycle(values(1, 2, 3), 4)
			vals := Take(ch, 12+1)
			Expect(vals).To(Equal([]int{
				1, 2, 3,
				1, 2, 3,
				1, 2, 3,
				1, 2, 3,
			}))
		})
	})
	Context("Take", func() {
		It("should return a array of first elements", func() {
			vals := Take(values(1, 2, 3, 4, 5), 3)
			Expect(vals).To(Equal([]int{1, 2, 3}))
		})
		It("should return a short array for a shorter channel than the count", func() {
			vals := Take(values(1, 2, 3), 5)
			Expect(vals).To(Equal([]int{1, 2, 3}))
		})
		It("should return a empty array for empty channel", func() {
			vals := Take(values(), 3)
			Expect(vals).To(BeEmpty())
		})
		It("should return a empty array for non-positive count", func() {
			vals := Take(values(1, 2, 3, 4, 5), 0)
			Expect(vals).To(BeEmpty())

			vals = Take(values(1, 2, 3, 4, 5), -3)
			Expect(vals).To(BeEmpty())
		})
	})

	Context("TakeAll", func() {
		It("should return a array of all elements", func() {
			vals := TakeAll(values(1, 2, 3, 4, 5))
			Expect(vals).To(Equal([]int{1, 2, 3, 4, 5}))
		})
		It("should return a empty array for empty channel", func() {
			vals := TakeAll(values())
			Expect(vals).To(BeEmpty())
		})
	})

	Context("RepeatN", func() {
		It("should return a channel which repeat element n times", func() {
			ch := RepeatN(42, 3)
			for i := 0; i < 3; i++ {
				val, ok := <-ch
				Expect(ok).To(BeTrue())
				Expect(val).To(Equal(42))
			}
			val, ok := <-ch
			Expect(ok).To(BeFalse())
			Expect(val).To(Equal(0))
		})
		It("should return a empty channel for non-positive count", func() {
			ch := RepeatN(42, 0)
			_, ok := <-ch
			Expect(ok).To(BeFalse())
			ch = RepeatN(42, -3)
			_, ok = <-ch
			Expect(ok).To(BeFalse())
		})
	})

	Context("Repeat", func() {
		It("should return a channel which repeat value forever", func() {
			ch := Repeat(42)
			for i := 0; i < 3; i++ {
				val, ok := <-ch
				Expect(ok).To(BeTrue())
				Expect(val).To(Equal(42))
			}
		})
	})

	Context("RepeatFuncN", func() {
		It("should return a channel which repeat element n times", func() {
			ch := RepeatFuncN(func() int { return 42 }, 3)
			for i := 0; i < 3; i++ {
				val, ok := <-ch
				Expect(ok).To(BeTrue())
				Expect(val).To(Equal(42))
			}
			val, ok := <-ch
			Expect(ok).To(BeFalse())
			Expect(val).To(Equal(0))
		})
		It("should return a empty channel for non-positive count", func() {
			ch := RepeatFuncN(func() int { return 42 }, 0)
			_, ok := <-ch
			Expect(ok).To(BeFalse())
			_, ok = <-ch
			ch = RepeatFuncN(func() int { return 42 }, -3)
			Expect(ok).To(BeFalse())
		})
	})

	Context("RepeatFunc", func() {
		It("should return a channel which repeat value forever", func() {
			ch := RepeatFunc(func() int { return 42 })
			for i := 0; i < 3; i++ {
				val, ok := <-ch
				Expect(ok).To(BeTrue())
				Expect(val).To(Equal(42))
			}
		})
	})
})
