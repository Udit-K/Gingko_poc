package adder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example/hello/adder"
)

var _ = Describe("Adder", func() {

	Context("When the inputs are valid", func() {

		Describe("Add", func() {

			It("should add two numbers", func() {
				sum, err := adder.Add(2, 3)
				Expect(err).ToNot(HaveOccurred())
				Expect(sum).To(Equal(5))
			})
		})
	})

	Context("When the inputs are invalid", func() {
		It("Should return an error", func() {
			_, err := adder.Add(-1, 3)
			Expect(err).To(HaveOccurred())
		})
	})
})
