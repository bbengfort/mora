package mora_test

import (
	. "github.com/bbengfort/mora"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mora", func() {

	It("should be at version 0.1", func() {
		Ω(Version).Should(Equal("0.1"))
	})

})
