package attributes_test

import "fmt"

func (suite *Suite) TestGetNoAttributesNoParams() {
	expected := `{
        "self": "/attributes",
        "offset": 0,
        "total": 0,
        "data": []
    }`

	suite.StartTest().Get("/attributes").
		Expect(suite.T()).
		Status(200).
		Type("application\\/json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func (suite *Suite) TestGetNoAttributesParams() {
	expected := `{
        "self": "/attributes",
        "offset": 10,
        "total": 0,
        "data": []
    }`

	suite.StartTest().Get("/attributes").
		AddQuery("offset", "10").
		AddQuery("count", "5").
		AddQuery("sort", "-name").
		AddQuery("name", "Strength").
		Expect(suite.T()).
		Status(200).
		Type("application\\/json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func (suite *Suite) TestListAttributesInvalidParams() {
	type test struct {
		name   string
		params map[string]string
		errors string
	}

	tests := []test{
		{
			name:   "Negative Offset",
			params: map[string]string{"offset": "-1"},
			errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`,
		},
		{
			name:   "Negative Count",
			params: map[string]string{"count": "-1"},
			errors: `{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`,
		},
		{
			name:   "Negative Offset And Count",
			params: map[string]string{"count": "-1", "offset": "-2"},
			errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/negative-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`,
		},

		{
			name:   "Non-Numeric Offset",
			params: map[string]string{"offset": "a"},
			errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`,
		},
		{
			name:   "Non-Numeric Count",
			params: map[string]string{"count": "b"},
			errors: `{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`,
		},
		{
			name:   "Non-Numeric Offset And Count",
			params: map[string]string{"count": "a", "offset": "b"},
			errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`,
		},

		{
			name:   "Non-Numeric Offset And Negative Count",
			params: map[string]string{"count": "-1", "offset": "b"},
			errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			expectedResponse := fmt.Sprintf(`{
                "type": "tag:wyrdwest,2019:problems/validation-problem",
                "title": "Validation Falure",
                "status": 400,
                "validationErrors": [%s]
            }`, test.errors)

			suite.StartTest().Get("/attributes").
				SetQueryParams(test.params).
				Expect(suite.T()).
				Status(400).
				Type("application\\/problem\\+json").
				JSON(suite.ParseJSONToMap(expectedResponse)).
				Done()
		})
	}
}

func (suite *Suite) TestGetOneAttributesNoParams() {
	suite.Seed(`
    attributes:
        - attribute_id: 00000000-0000-0000-0000-000000000000
          version: 11111111-1111-1111-1111-111111111111
          created: 2019-03-29T07:09:00Z
          updated: 2019-03-29T07:09:00Z
          name: Strength
          description: How Strong I am
    `)

	expected := `{
        "self": "/attributes",
        "offset": 0,
        "total": 1,
        "data": [
            {
                "self": "/attributes/00000000-0000-0000-0000-000000000000",
                "name": "Strength",
                "description": "How Strong I am"
            }
        ]
    }`

	suite.StartTest().Get("/attributes").
		Expect(suite.T()).
		Status(200).
		Type("application\\/json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func (suite *Suite) TestGetAttributes() {
	suite.Seed(`
    attributes:
        - attribute_id: 00000000-0000-0000-0000-000000000000
          version: 11111111-1111-1111-1111-111111111111
          created: 2019-03-29T07:09:00Z
          updated: 2019-03-29T07:09:00Z
          name: Strength
          description: How Strong I am
        - attribute_id: 00000000-0000-0000-0000-000000000001
          version: 11111111-1111-1111-1111-111111111111
          created: 2019-03-29T07:10:00Z
          updated: 2019-03-29T07:10:00Z
          name: Intelligence
          description: How Smart I am
    `)

	type test struct {
		name     string
		params   map[string]string
		expected string
	}

	tests := []test{
		{
			name:   "No Params",
			params: map[string]string{},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000001",
                        "name": "Intelligence",
                        "description": "How Smart I am"
                    },
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Offset=1",
			params: map[string]string{"offset": "1"},
			expected: `{
                "self": "/attributes",
                "offset": 1,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Count=1",
			params: map[string]string{"count": "1"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000001",
                        "name": "Intelligence",
                        "description": "How Smart I am"
                    }
                ]
            }`,
		},
		{
			name:   "Sort Name Descending",
			params: map[string]string{"sort": "-name"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    },
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000001",
                        "name": "Intelligence",
                        "description": "How Smart I am"
                    }
                ]
            }`,
		},
		{
			name:   "Sort Created Ascending",
			params: map[string]string{"sort": "+created"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    },
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000001",
                        "name": "Intelligence",
                        "description": "How Smart I am"
                    }
                ]
            }`,
		},
		{
			name:   "Sort Unknown",
			params: map[string]string{"sort": "unknown"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 2,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000001",
                        "name": "Intelligence",
                        "description": "How Smart I am"
                    },
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Filtered name=strength",
			params: map[string]string{"name": "strength"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 1,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Filtered name=STRENGTH",
			params: map[string]string{"name": "STRENGTH"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 1,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Filtered name=Strength",
			params: map[string]string{"name": "Strength"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 1,
                "data": [
                    {
                        "self": "/attributes/00000000-0000-0000-0000-000000000000",
                        "name": "Strength",
                        "description": "How Strong I am"
                    }
                ]
            }`,
		},
		{
			name:   "Filtered name=Other",
			params: map[string]string{"name": "Other"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 0,
                "data": []
            }`,
		},
		{
			name:   "Filtered name=☃",
			params: map[string]string{"name": "☃"},
			expected: `{
                "self": "/attributes",
                "offset": 0,
                "total": 0,
                "data": []
            }`,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.StartTest().Get("/attributes").
				SetQueryParams(test.params).
				Expect(suite.T()).
				Status(200).
				Type("application\\/json").
				JSON(suite.ParseJSONToMap(test.expected)).
				Done()
		})
	}
}
