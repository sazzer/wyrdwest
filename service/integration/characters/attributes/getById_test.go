package attributes_test

func (suite *Suite) TestGetInvalidID() {
	expected := `{
		"status": 400,
		"type": "tag:wyrdwest,2019:problems/attributes/invalid-id",
		"title": "The Attribute ID was invalid"
	}`

	suite.StartTest().Get("/attributes/invalid").
		Expect(suite.T()).
		Status(400).
		Type("application\\/problem\\+json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func (suite *Suite) TestGetUnknownID() {
	expected := `{
		"status": 404,
		"type": "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
		"title": "The Attribute was not found"
	}`

	suite.StartTest().Get("/attributes/00000000-0000-0000-0000-000000000000").
		Expect(suite.T()).
		Status(404).
		Type("application\\/problem\\+json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}

func (suite *Suite) TestGetKnownID() {
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
		"self": "/attributes/00000000-0000-0000-0000-000000000000",
		"name": "Strength",
		"description": "How Strong I am"
	}`

	suite.StartTest().Get("/attributes/00000000-0000-0000-0000-000000000000").
		Expect(suite.T()).
		Status(200).
		Type("application\\/json").
		JSON(suite.ParseJSONToMap(expected)).
		Done()
}
