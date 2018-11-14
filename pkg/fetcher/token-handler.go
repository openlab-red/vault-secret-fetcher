package fetcher

import (
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"github.com/openlab-red/vault-secret-fetcher/pkg/util"
	"errors"
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

	data, err := ioutil.ReadFile(h.Token.Path)
	check(err)

	h.Token.Value = string(data)
	h.Client, err = h.createAPIClient()
	check(err)

	err = h.Properties.create()
	check(err)

	secrets := strings.Split(viper.GetString("vault-secret"), ",")
	log.Debugln("List secrets", secrets)

	h.Properties.Content = make(map[string]interface{})
	for _, secret := range secrets {
		secret := Secret{
			Name: secret,
		}

		err = h.retrieve(&secret)
		check(err)

		path := util.PathToMap(secret.Name, secret.VaultSecret.Data)
		log.Infof("Path: %v", path)
		err = util.MergeMap(path, h.Properties.Content)
		check(err)

	}

	err = h.Properties.save()
	check(err)
	log.Infoln("Wrote secret: ", h.Properties.Path)
	h.Properties.close()
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func (h TokenHandler) retrieve(secret *Secret) (error) {

	if secret.Name == "" {
		return errors.New("secret name is empty")
	}

	log.Debugln("Using token: ", h.Token.Value)
	log.Debugln("Retrieving secret: ", secret.Name)

	client := h.Client
	client.SetToken(h.Token.Value)
	vaultSecret, err := client.Logical().Read(secret.Name)
	if err != nil {
		log.Errorln(err)
		return err
	}
	secret.VaultSecret = vaultSecret

	return nil

}
