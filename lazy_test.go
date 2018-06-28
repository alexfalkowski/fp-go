// +build spec

package fpgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lazy", func() {
	lazyInt := createLazyInt(func() int { return 23 })

	It("gets lazy value", func() {
		Expect(lazyInt()).To(Equal(23))
	})
})
