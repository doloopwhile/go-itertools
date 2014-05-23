package itertools

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func array(vals ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range vals {
			out <- v
		}
	}()
	return out
}

func mul(x int, y int) int {
	return x * y
}

var _ = Describe("Book", func() {
	// itertools functions
	Context("Accumulate", func() {
		It("should return accumulated sums", func() {
			out := Accumulate(array(1, 2, 3, 4, 5), mul)
			Expect(out).To(Equal([]int{1, 2, 6, 24, 120}))
		})
		It("should be nil channel if no elements", func() {
			Expect(Accumulate(array(), mul)).To(BeNil())
		})
	})
	Context("Chain", func() {
		It("should iterate all elements in the channels", func() {
		})
		It("should be nil if no channels given", func() {
			Expect(Chain()).To(BeNil())
		})
	})
	Context("ChainFromIterable", func() {
		It("should iterate all elements in channels", func() {
		})
		It("should be nil if no channels in the input channel", func() {
			empty := make(chan (chan int))
			Expect(ChainFromIterable(empty)).To(BeNil())
		})
	})
	Context("Combinations", func() {
		It("should r length subsequences of elements from the input channel", func() {
		})
		It("should be nil if less elements than r", func() {
			Expect(Combinations(array(1, 2), 3)).To(BeNil())
		})
	})
	Context("CombinationsWithReplacement", func() {
		It("should r length subsequences of elements from the input iterable allowing duplication.", func() {
		})
		It("should be nil if no elements in the input channel", func() {
			Expect(CombinationsWithReplacement(array(), 3)).To(BeNil())
		})
	})
	Context("Compress", func() {
		It("should iterate elements that correspond to true", func() {
		})
		It("should be nil if no elements in the input channel", func() {
		})
	})
	Context("Count", func() {
		It("should repeat integers start with given n", func() {
		})
	})
	Context("Cycle", func() {
		It("should cyclicly repeat elements", func() {
		})
		It("should be nil channel if no elements", func() {
		})
	})
	Context("DropWhile", func() {
		It("should drop elements as long as the predicate is true", func() {
		})
		It("should be nil channel if no elements", func() {
		})
	})
	Context("Filter", func() {
		It("should iterate elements with which the predicate is true", func() {
		})
	})
	Context("FilterFalse", func() {
		It("should iterate elements with which the predicate is false", func() {
		})
	})
	Context("GroupBy", func() {
		It("should iterate keys and groups", func() {
		})
	})
	// Slice
	Context("Permutations", func() {
		It("should iterate successive r length permutations of elements in the iterable.", func() {
		})
	})
	Context("Product", func() {
		It("should be a cartesian product of the input channels.", func() {
		})
	})
	Context("Repeat", func() {
		It("should repeat a given element infinitely", func() {
		})
	})
	Context("RepeatN", func() {
		It("should repeat a given element for times", func() {
		})
		It("should be nil channel if times is non-positive", func() {
		})
	})
	Context("TakeWhile", func() {
		It("should returns elements from the channel as long as the predicate is true.", func() {
		})
	})
	Context("Tee", func() {
		It("should returns n copies of the channel", func() {
		})
	})
})
