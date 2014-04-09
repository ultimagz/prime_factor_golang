package PrimeFactor_test

import (
	. "../../../Kata/Golang/PrimeFactor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrimeFactor", func() {
	Context("PrimeFactor 1", func() {
		It("Should return []", func() {
			Expect(PrimeFactor(1)).To(Equal( []int{} ))
		})
	})

	Context("PrimeFactor 2", func() {
		It("Should return [2]", func() {
			Expect(PrimeFactor(2)).To(Equal( []int{2} ))
		})
	})

	Context("PrimeFactor 4", func() {
		It("Should return [2 2]", func() {
			Expect(PrimeFactor(4)).To(Equal( []int{2, 2} ))
		})
	})

	Context("PrimeFactor 8", func() {
		It("Should return [2 2 2]", func() {
			Expect(PrimeFactor(8)).To(Equal( []int{2, 2, 2} ))
		})
	})

	Context("PrimeFactor 3", func() {
		It("Should return [3]", func() {
			Expect(PrimeFactor(3)).To(Equal( []int{3} ))
		})
	})

	Context("PrimeFactor 9", func() {
		It("Should return [3 3]", func() {
			Expect(PrimeFactor(9)).To(Equal( []int{3, 3} ))
		})
	})

	Context("PrimeFactor 6", func() {
		It("Should return [2 3]", func() {
			Expect(PrimeFactor(6)).To(Equal( []int{2, 3} ))
		})
	})

	Context("PrimeFactor 5", func() {
		It("Should return [5]", func() {
			Expect(PrimeFactor(5)).To(Equal( []int{5} ))
		})
	})

	Context("PrimeFactor 10", func() {
		It("Should return [2 5]", func() {
			Expect(PrimeFactor(10)).To(Equal( []int{2, 5} ))
		})
	})

	Context("PrimeFactor 30", func() {
		It("Should return [2 3 5]", func() {
			Expect(PrimeFactor(30)).To(Equal( []int{2, 3, 5} ))
		})
	})
})
