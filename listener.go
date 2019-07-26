package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

func createListener(proto string, port string, timeout time.Duration) {
	l, err := net.Listen(proto, fmt.Sprintf(":%s", port))
	checkError(err, true)
	for {
		con, err := l.Accept()
		checkError(err, false)
		err = con.SetDeadline(time.Now().Add(timeout))
		checkError(err, false)

		go handleCon(con)
	}
}

func handleCon(con net.Conn) {
	defer con.Close()
	now := time.Now()
	rhost := con.RemoteAddr().String()
	proto := con.LocalAddr().Network()
	log.Printf("Received a connection from %s:%s", rhost, proto)
	localAddrStrSlice := strings.Split(con.LocalAddr().String(), ":")
	port := localAddrStrSlice[len(localAddrStrSlice)-1] // handles the IPv6 [::1]:XXX case
	data, err := ioutil.ReadAll(con)
	if err != nil {
		log.Printf("Terminating connection to %s due to timeout", con.RemoteAddr().String())
	}

	e := newEvent(now, proto, rhost, port, data)
	processEvent(e)
}
