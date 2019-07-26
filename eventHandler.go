package main

import (
	"github.com/spf13/viper"
)

var availableHandlers = []string{"file", "mail"}
var registeredHandlers = []eventHandler{}

func getEventHandler(EventHandlerName string) eventHandler {
	switch EventHandlerName {
	case "file":
		return &fileEventHandler{}
	case "mail":
		return &mailEventHandler{}
	default:
		return nil
	}
}

type eventHandler interface {
	configure(config *viper.Viper) error
	processEvent(e event)
}

// Is called every time a new event is generated.
// Calls event processors of registered event handlers.
func processEvent(e event) {
	for _, handler := range registeredHandlers {
		handler.processEvent(e)
	}
}
