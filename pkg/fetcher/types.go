package fetcher

import (
	"github.com/hashicorp/vault/api"
	"os"
)

type TokenHandler struct {
	VaultAddr  string
	Token      string
	Properties Properties
}

type Secret struct {
	Name        string
	Token       string
	Client      *api.Client
	VaultSecret *api.Secret
}

type Properties struct {
	Path    string
	Format  string
	File    *os.File
	Content []byte
}
