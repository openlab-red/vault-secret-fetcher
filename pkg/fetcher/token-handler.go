package fetcher

import (
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/micro/go-config/encoder/yaml"
	"github.com/micro/go-config/encoder/json"
	"strconv"
	"path/filepath"
)

type tokenHandler struct {
	vaultAddr string
}

func (h tokenHandler) createAPIClient() (*api.Client, error) {
	//creates the vault config
	log.Debugln("Creating vault config")
	insecure, _ := strconv.ParseBool(viper.GetString("vault-insecure"))
	vConfig := api.Config{
		Address: viper.GetString("vault-addr"),
	}
	tlsConfig := api.TLSConfig{
		CAPath:   viper.GetString("vault-capath"),
		Insecure: insecure,
	}
	err := vConfig.ConfigureTLS(&tlsConfig)
	if err != nil {
		log.Warnln(err)
		return nil, err
	}
	log.Debugln("Created vault config")

	//creates the vault client
	log.Debugln("Creating vault client")
	client, err := api.NewClient(&vConfig)
	if err != nil {
		log.Warnln(err)
		return client, err
	}
	client.SetAddress(h.vaultAddr)
	log.Debugln("Created vault client")
	return client, err
}

func (h tokenHandler) readToken() {
	propertiesFile := viper.GetString("properties-file")
	vaultToken := viper.GetString("vault-token")
	retrieveSecret := viper.GetString("vault-secret")

	if err := os.Remove(propertiesFile); err != nil {
		log.WithFields(logrus.Fields{
			"propertiesFile": propertiesFile,
		}).Warn(err)
	}

	data, err := ioutil.ReadFile(vaultToken)
	check(err)

	clientToken := string(data)
	client, err := h.createAPIClient()
	if err != nil {
		log.Warnln(err)
		return
	}

	if retrieveSecret != "" {
		log.Debugln("Using token: ", clientToken)
		log.Debugln("Retrieving secret: ", retrieveSecret)
		client.SetToken(clientToken)
		secret, err := client.Logical().Read(retrieveSecret)
		if err != nil {
			log.Warnln(err)
			return
		}
		f, err := os.Create(propertiesFile)
		if err != nil {
			log.Warnln(err)
			return
		}
		defer f.Close()

		var content [] byte

		switch    ext := filepath.Ext(propertiesFile); ext {
		case "yaml":
			content, err = yaml.NewEncoder().Encode(&secret.Data)
		case "json":
			content, err = json.NewEncoder().Encode(&secret.Data)
		default:
			log.Fatalln("Type %s not supported", ext)
		}

		check(err)
		f.Write(content)
		log.Infoln("Wrote secret: ", propertiesFile)
	}

}

func check(e error) {
	if e != nil {
		log.Error(e)
		panic(e)
	}
}
