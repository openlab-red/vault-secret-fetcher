package fetcher

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"github.com/openlab-red/vault-secret-fetcher/pkg/util"
)

func TestSaveJSON(t *testing.T) {

	input := mockSecret("secret/example")

	expected := "{\"secret\":{\"example\":{\"password\":\"pwd\"}}}\n"

	var p = Properties{
		Format: ".json",
		Path:   "/tmp/application.json",
		Content:  util.PathToMap(input.Name, input.VaultSecret.Data),
	}

	p.save()

	actual, _ := ioutil.ReadFile("/tmp/application.json")
	assert.Equal(t, expected, string(actual))

}

func TestSaveYaml(t *testing.T) {

	input := mockSecret("secret/example")

	expected := "secret:\n  example:\n    password: pwd\n"

	var p = Properties{
		Format: ".yaml",
		Path:   "/tmp/application.yaml",
		Content:  util.PathToMap(input.Name, input.VaultSecret.Data),
	}

	p.save()

	actual, _ := ioutil.ReadFile("/tmp/application.yaml")
	assert.Equal(t, expected, string(actual))

}

func mockSecret(name string) Secret {
	data := make(map[string]interface{})
	data["password"] = "pwd"
	var input = Secret{
		Name: name,
		VaultSecret: &api.Secret{
			Data: data,
		},
	}
	return input
}
