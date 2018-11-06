package fetcher

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvPathToMapOdd(t *testing.T) {
	input := "database/creds/pg-readwrite"

	expected := make(map[string]interface{})
	creds := make(map[string]interface{})
	pgReadwrite := make(map[string]interface{})

	pgReadwrite["pg-readwrite"] = make(map[string]interface{})
	creds["creds"] = pgReadwrite
	expected["database"] = creds

	actual := convPathToMap(input, make(map[string]interface{}))

	assert.Equal(t, expected, actual)
}

func TestConvPathToMapEven(t *testing.T) {
	input := "secret/example"

	expected := make(map[string]interface{})
	example := make(map[string]interface{})

	example["example"] = make(map[string]interface{})
	expected["secret"] = example

	actual := convPathToMap(input, make(map[string]interface{}))

	assert.Equal(t, expected, actual)
}
