// +build spec

package fpgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Immutable", func() {
	var person person

	BeforeEach(func() {
		person = createPerson("Alex")
	})

	It("allows to mutate", func() {
		changeName(&person, "Alejandro")
		Expect(person.name).To(Equal("Alejandro"))
	})

	It("doesn't allow to mutate", func() {
		cantChangeName(person, "Alejandro")
		Expect(person.name).To(Equal("Alex"))
	})
})
