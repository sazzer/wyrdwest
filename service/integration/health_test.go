package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sazzer/wyrdwest/service/integration/http"
)

var _ = Describe("Health", func() {
	It("is healthy", func() {
		expected := `{
			"database": {
				"status": "ok"
			}
		}`

		err := http.StartTest().Get("/health").
			Expect(GinkgoT()).
			Status(200).
			Type("json").
			JSON(http.ParseJsonToMap(expected)).
			Done()
		Expect(err).To(BeNil())
	})
})
