package fetcher

import (
	"os"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"strings"
)

func (p *Properties) create() (err error) {
	p.File, err = os.Create(p.Path)
	return
}

func (p *Properties) save(secret Secret) (err error) {

	log.Debugln("Saving it on ", p.File.Name())

	data := convPathToMap(secret.Name, secret.VaultSecret.Data)
	log.Debugf("Map %v", data)

	switch p.Format {
	case ".yaml":
		err = yaml.NewEncoder(p.File).Encode(&data)
	case ".json":
		err = json.NewEncoder(p.File).Encode(&data)
	default:
		log.Fatalf("Type %s not supported", p.Format)
	}

	return
}

func (p *Properties) close() {
	defer p.File.Close()
}

func convPathToMap(secret string, content map[string]interface{}) (map[string]interface{}) {
	data := make(map[string]interface{})

	if secret != "" {
		paths := strings.Split(secret, "/")
		var prev string
		var parent map[string]interface{}
		for _, key := range paths {
			tmp := createMap(key, content)
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

func createMap(key string, content map[string]interface{}) map[string]interface{} {
	tmp := make(map[string]interface{})
	tmp[key] = content
	return tmp
}
