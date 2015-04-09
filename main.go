package main

import "net"
import "log"
import "flag"
import "fmt"
import "os"

var bindAddress = flag.String("bind-address", "127.0.0.1", "bind address")
var localPort = flag.Int("local-port", 6000, "local port")
var remotePort = flag.Int("remote-port", 0, "remote port")
var remoteHost = flag.String("remote-host", "", "remote host")
var bufferSize = flag.Int("buffer-size", 512, "buffer size")
var displayLogs = flag.Bool("log", false, "log")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...]\n", os.Args[0])
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Options:\n")
		fmt.Fprint(os.Stderr, "\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	local, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *bindAddress, *localPort))

	if err != nil {
		log.Fatal("Failed to listen on local port: ", err)
	}

	localConn, err := local.Accept()

	if err != nil {
		log.Fatal("Failed to accept connection: ", err)
	}

	if *displayLogs == true {
		log.Printf("Request from: %s", localConn.RemoteAddr())
	}

	remoteConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *remoteHost, *remotePort))

	if err != nil {
		log.Fatal("Failed to connect to remote host: ", err)
	}

	if *displayLogs == true {
		log.Printf("Connected to %s:%d", *remoteHost, *remotePort)
	}

	errChan := make(chan error)

	go readLoop(localConn, remoteConn, errChan, "-> ")
	go readLoop(remoteConn, localConn, errChan, "<- ")

	select {
	case err := <-errChan:
		log.Fatal(err)
	}
}

func readLoop(from net.Conn, to net.Conn, errChan chan error, logPrefix string) {
	for {
		data := make([]byte, *bufferSize)
		bytes, err := from.Read(data)

		if err != nil {
			from.Close()
			errChan <- err
			return
		}

		if *displayLogs == true {
			log.Print(logPrefix, string(data[0:bytes]))
		}

		_, err = to.Write(data[0:bytes])

		if err != nil {
			to.Close()
			errChan <- err
			return
		}
	}
}
