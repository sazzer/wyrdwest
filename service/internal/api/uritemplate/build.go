package uritemplate

import (
	"github.com/jtacoma/uritemplates"
	"github.com/sirupsen/logrus"
)

// BuildURI will construct a URI out of the provided template and parameters
func BuildURI(template string, params map[string]interface{}) string {
	parsed, err := uritemplates.Parse(template)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse URI Template")
		panic(err)
	}

	result, err := parsed.Expand(params)
	if err != nil {
		logrus.WithError(err).Error("Failed to expand URI Template")
		panic(err)
	}

	return result
}
