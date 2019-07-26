package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var configFile *string

func init() {
	configFile = flag.String("f", "", "configuration file")
}

func checkError(err error, abort bool) {
	if err == nil {
		return
	}
	log.Printf("Error occured: %s", err.Error())
	if abort {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")

	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	} else {
		checkError(err, true)
	}
}

func configureEventHandlers() {
	var err error
	for _, handler := range availableHandlers {
		if !viper.IsSet(handler) {
			continue
		}
		eventHandler := getEventHandler(handler)
		err = eventHandler.configure(viper.Sub(handler))
		checkError(err, true)
	}
}

func setupListeners() {
	portsToListen := viper.GetStringSlice("listenPort")
	timeout := viper.GetDuration("connectionTimeout")
	for _, elem := range portsToListen {
		log.Printf("Listening on %s", elem)
		splitted := strings.Split(elem, "/")
		go createListener(splitted[0], splitted[1], timeout)
	}
}

func main() {
	flag.Parse()
	initConfig()
	configureEventHandlers()
	setupListeners()
	for {
	}
}
