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


func (p *Properties) save() (err error) {

	if p.File == nil {
		err = p.create()
	}

	log.Debugln("Saving it on ", p.File.Name())

	log.Debugf("Map %v", p.Content)

	switch p.Format {
	case ".yaml", ".yml":
		err = yaml.NewEncoder(p.File).Encode(&p.Content)
	case ".json":
		err = json.NewEncoder(p.File).Encode(&p.Content)
	default:
		log.Fatalf("Type %s not supported", p.Format)
	}

	return
}

func (p *Properties) close() {
	defer p.File.Close()
}
