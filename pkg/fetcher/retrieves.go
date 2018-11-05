package fetcher

import (
	"errors"
)

func (secret *Secret) retrieve() (error) {

	if secret.Name == "" {
		return errors.New("secret name is empty")
	}

	log.Debugln("Using token: ", secret.Token)
	log.Debugln("Retrieving secret: ", secret.Name)

	client := secret.Client
	client.SetToken(secret.Token)
	vaultSecret, err := client.Logical().Read(secret.Name)
	if err != nil {
		log.Errorln(err)
		return err
	}
	secret.VaultSecret = vaultSecret

	return nil

}
