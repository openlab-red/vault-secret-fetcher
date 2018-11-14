package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPathToMapOdd(t *testing.T) {
	input := "database/creds/pg-readwrite"

	expected := make(map[string]interface{})
	creds := make(map[string]interface{})
	pgReadwrite := make(map[string]interface{})

	pgReadwrite["pg-readwrite"] = make(map[string]interface{})
	creds["creds"] = pgReadwrite
	expected["database"] = creds

	actual := PathToMap(input, make(map[string]interface{}))

	assert.Equal(t, expected, actual)
}

func TestPathToMapEven(t *testing.T) {
	input := "secret/example"

	expected := make(map[string]interface{})
	example := make(map[string]interface{})

	example["example"] = make(map[string]interface{})
	expected["secret"] = example

	actual := PathToMap(input, make(map[string]interface{}))

	assert.Equal(t, expected, actual)
}

func TestMergeMap(t *testing.T) {
	map1 := make(map[string]interface{})
	map1["1"] = "a"
	map1["2"] = "b"
	map1["3"] = "c"

	map2 := make(map[string]interface{})
	map2["4"] = "d"
	map2["5"] = "e"
	map2["6"] = "f"
	map2["2"] = "B"

	expected := make(map[string]interface{})
	expected["1"] = "a"
	expected["2"] = "B"
	expected["3"] = "c"
	expected["4"] = "d"
	expected["5"] = "e"
	expected["6"] = "f"

	MergeMap(map1, map2)
	actual := map2

	assert.Equal(t, expected, actual)

}
