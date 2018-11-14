package util

import (
	"strings"
	"github.com/imdario/mergo"
)

func CreateMap(key string, content map[string]interface{}) map[string]interface{} {
	tmp := make(map[string]interface{})
	tmp[key] = content
	return tmp
}

func MergeMap(s map[string]interface{}, t map[string]interface{}) (error) {
	return mergo.Merge(&t, s)
}

func PathToMap(secret string, content map[string]interface{}) (map[string]interface{}) {
	data := make(map[string]interface{})

	if secret != "" {
		paths := strings.Split(secret, "/")
		var prev string
		var parent map[string]interface{}
		for _, key := range paths {
			tmp := CreateMap(key, content)
			if len(prev) == 0 {
				data = tmp
				parent = data
			} else {
				parent[prev] = tmp
				parent = tmp
			}
			prev = key
		}
	}
	return data

}
