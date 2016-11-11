package main

import (
	"bufio"
	"net"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	name    = "relayovic"
	version = "0.0.1"
	desc    = "statsd compatible forwarder, with consistent hashing"

	port = 3889
)

var log = logrus.WithFields(logrus.Fields{"app": "relayovic"})

func process(conn net.Conn) {
	buffy, _ := bufio.NewReader(conn).ReadString('\n')
	//	time.Sleep(1 * time.Second)
	log.Info("Received: " + buffy)
	time.Sleep(1 * time.Second)
	conn.Close()
}

func main() {
	log.Info("Aaaaand begin")

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Cannae listen on " + strconv.Itoa(port))
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
			log.Error("Accept errored, giving up")
		}
		go process(conn)

	}
}
