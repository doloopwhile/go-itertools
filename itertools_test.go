package itertools

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Itertools", func() {
	Context("Count", func() {
		It("should iterate integer in step by step", func() {
			ch := Count(1, 2)
			Expect(<-ch).To(Equal(1))
			Expect(<-ch).To(Equal(3))
			Expect(<-ch).To(Equal(5))
		})
	})
})
