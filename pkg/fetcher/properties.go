package fetcher

import (
	"os"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

func (p *Properties) create() (err error) {
	p.File, err = os.Create(p.Path)
	return
}

func (p *Properties) save(secret Secret) (err error) {

	log.Debugln("Saving it on ", p.File.Name())

	data := make(map[string]interface{})
	data[secret.Name] = secret.VaultSecret.Data

	switch p.Format {
	case ".yaml":
		err = yaml.NewEncoder(p.File).Encode(&data)
	case ".json":
		err = json.NewEncoder(p.File).Encode(&data)
	default:
		log.Fatalln("Type %s not supported", p.Format)
	}

	return
}

func (p *Properties) close() {
	defer p.File.Close()
}
