package fetcher

import (
	"github.com/hashicorp/vault/api"
	"os"
)

type TokenHandler struct {
	VaultAddr  string
	Token      Token
	Properties Properties
	Client     *api.Client
}

type Token struct {
	Path  string
	Value string
}

type Secret struct {
	Name        string
	VaultSecret *api.Secret
}

type Properties struct {
	Path    string
	Format  string
	File    *os.File
	Content map[string]interface{}
}
