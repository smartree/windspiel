package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type fileEventHandler struct {
	filedescriptor *os.File
}

func (handler *fileEventHandler) configure(config *viper.Viper) error {
	if !config.IsSet("filename") {
		log.Fatalf("Filename missing, aborting")
	}
	log.Printf("Enabling logfile %s", config.GetString("filename"))
	registeredHandlers = append(registeredHandlers, handler)
	var err error

	if handler.filedescriptor != nil {
		return fmt.Errorf("This handler is already configured")
	}

	mode := os.O_RDWR | os.O_CREATE | os.O_EXCL
	if config.GetBool("overwriteExisting") {
		mode = os.O_RDWR | os.O_CREATE
	}
	handler.filedescriptor, err = os.OpenFile(config.GetString("filename"), mode, 0644)

	return err

}

func (handler *fileEventHandler) processEvent(e event) {
	marshalled, err := json.Marshal(e)
	checkError(err, false)
	handler.filedescriptor.Write(append(marshalled, []byte("\n")...))
}
