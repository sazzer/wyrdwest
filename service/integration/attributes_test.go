package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sazzer/wyrdwest/service/integration/http"
)

var _ = Describe("Attributes", func() {
	Context("Get By ID", func() {
		It("Returns the correct error when the ID is invalid", func() {
			expected := `{
				"status": 400,
				"type": "tag:wyrdwest,2019:problems/attributes/invalid-id",
				"title": "The Attribute ID was invalid"
			}`

			err := http.StartTest().Get("/attributes/invalid").
				Expect(GinkgoT()).
				Status(400).
				Type("json").
				JSON(http.ParseJSONToMap(expected)).
				Done()
			Expect(err).To(BeNil())
		})

		It("Returns the correct error when the ID is unknown", func() {
			expected := `{
				"status": 404,
				"type": "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
				"title": "The Attribute was not found"
			}`

			err := http.StartTest().Get("/attributes/00000000-0000-0000-0000-000000000000").
				Expect(GinkgoT()).
				Status(404).
				Type("json").
				JSON(http.ParseJSONToMap(expected)).
				Done()
			Expect(err).To(BeNil())
		})
	})
})
