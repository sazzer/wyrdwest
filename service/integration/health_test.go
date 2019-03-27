package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health", func() {
	It("Fails", func() {
		Expect(1).To(Equal(2))
	})
})
