package handler

import (
	"encoding/json"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/micro/go-config/encoder/yaml"
	"strconv"
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
	retrieveSecret := viper.GetString("retrieve-secret")
	var swi api.SecretWrapInfo

	if err := os.Remove(propertiesFile); err != nil {
		log.WithFields(logrus.Fields{
			"propertiesFile": propertiesFile,
		}).Warn(err)
	}

	data, err := ioutil.ReadFile(vaultToken)
	check(err)

	err = json.Unmarshal(data, &swi)
	if err != nil {
		log.Warnln(err)
		return
	}

	client, err := h.createAPIClient()
	if err != nil {
		log.Warnln(err)
		return
	}

	// Now unwrap it
	ahWrapping := true
	switch {
	case swi.TTL != 10:
		log.Errorln("bad wrap info: ", swi.TTL)
	case !ahWrapping && swi.CreationPath != "sys/wrapping/wrap":
		log.Errorln("bad wrap path:", swi.CreationPath)
	case ahWrapping && swi.CreationPath != "auth/kubernetes/login":
		log.Errorln("bad wrap path:", swi.CreationPath)
	case swi.Token == "":
		log.Errorln("wrap token is empty")
	}
	client.SetToken(swi.Token)
	secret, err := client.Logical().Unwrap("")
	if err != nil {
		log.Errorln(err)
	}
	var clientToken string

	if ahWrapping {
		switch {
		case secret.Auth == nil:
			log.Errorln("unwrap secret auth is nil")
		case secret.Auth.ClientToken == "":
			log.Errorln("unwrap token is nil")
		}
		clientToken = secret.Auth.ClientToken
	} else {
		switch {
		case secret.Data == nil:
			log.Errorln("unwrap secret data is nil")
		case secret.Data["token"] == nil:
			log.Errorln("unwrap token is nil")
		}
		clientToken = secret.Data["token"].(string)
	}

	if retrieveSecret != "" {
		log.Debugln("Using token: ", clientToken)
		log.Debugln("Retrieving secret: ", retrieveSecret)
		client.SetToken(swi.Token)
		//secret, err := client.Logical().Read(retrieveSecret)
		secret, err := client.Auth().Token().LookupSelf()
		if err != nil {
			log.Warnln(err)
			return
		}
		//log.Debugln("Executed secret request", client.Address(), retrieveSecret)
		f, err := os.Create(propertiesFile)
		if err != nil {
			log.Warnln(err)
			return
		}
		defer f.Close()
		//err = json.NewEncoder(f).Encode(&secret.Data)
		content, err := yaml.NewEncoder().Encode(&secret.Data)
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
