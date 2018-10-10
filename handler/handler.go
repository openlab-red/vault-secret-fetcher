package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"fmt"
)

var log = logrus.New()

func initLog() {
	level, err := logrus.ParseLevel(viper.GetString("log-level"))
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Level = level
	}

}

func validateConfig() {
	if viper.GetString("vault-cacert") == "" {
		log.Fatalln("vault-cacert must be set and non-empty")
	}
	if viper.GetString("vault-addr") == "" {
		log.Fatalln("vault-addr must be set and non-empty")
	}
	log.Debugln("configuration is valid")
}

func Start() {
	initLog()
	log.Infoln("Starting token handler...")

	validateConfig()

	tokenHandler := tokenHandler{viper.GetString("vault-addr")}
	newWatcher(tokenHandler)
}

func newWatcher(tokenHandler tokenHandler) {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				log.Info("EVENT! %#v\n", event)
				tokenHandler.readToken()
				// watch for errors
			case err := <-watcher.Errors:
				log.Errorln(err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(viper.GetString("vault-token")); err != nil {
		log.Fatal(err)
	}

	<-done
}
