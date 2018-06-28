// +build spec

package fpgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Currying", func() {
	It("adds two values via named functions", func() {
		addFive := createAdd(5)
		ten := addFive(5)
		Expect(ten).To(Equal(10))

		ten = createAdd(5)(5)
		Expect(ten).To(Equal(10))
	})
})
