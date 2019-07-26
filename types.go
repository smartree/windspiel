package main

import (
	"os"
	"time"
)

type event struct {
	Time   time.Time `json:"time"`
	Proto  string    `json:"protocol"`
	Src    string    `json:"source address"`
	Target string    `json:"target system"`
	Port   string    `json:"target port"`
	Data   []byte    `json:"data sent"`
}

func newEvent(occured time.Time, proto string, src string, port string, data []byte) event {
	if len(data) > 1024 { // Just in case...
		newData := data[:1024]
		data = newData
	}
	hostname, err := os.Hostname()
	checkError(err, false)
	e := event{
		Time:   occured,
		Proto:  proto,
		Src:    src,
		Target: hostname,
		Port:   port,
		Data:   data,
	}
	return e
}
