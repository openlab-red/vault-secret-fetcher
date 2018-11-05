package fetcher

import (
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func (h TokenHandler) createAPIClient() (*api.Client, error) {
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
	client.SetAddress(h.VaultAddr)
	log.Debugln("Created vault client")
	return client, err
}

func (h TokenHandler) readToken() {

	if err := os.Remove(h.Properties.Path); err != nil {
		log.WithFields(logrus.Fields{
			"Properties Path": h.Properties.Path,
		}).Warn(err)
	}

	data, err := ioutil.ReadFile(h.Token)
	check(err)

	clientToken := string(data)
	client, err := h.createAPIClient()
	check(err)

	err = h.Properties.create()
	check(err)

	secrets := strings.Split(viper.GetString("vault-secret"), ",")
	log.Debugln("List secrets", secrets)

	for _, secret := range secrets {
		secret := Secret{
			Name:   secret,
			Token:  clientToken,
			Client: client,
		}

		err = secret.retrieve()
		check(err)

		log.Debugln(secret.VaultSecret.Data)

		err = h.Properties.save(secret)
		check(err)

	}
	log.Infoln("Wrote secret: ", h.Properties.Path)
	h.Properties.close()
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
