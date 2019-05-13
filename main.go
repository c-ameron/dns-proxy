package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	cloudflare := "1.1.1.1:853"

	fmt.Println("Running proxy")

	listener, err := net.Listen("tcp", ":8853")
	checkError(err)

	for {
		// accept local connections
		client, err := listener.Accept()
		checkError(err)

		// create a connection to cloudflare
		upstream, err := tls.Dial("tcp", cloudflare, nil)
		checkError(err)

		// really basic piping between the connections
		go io.Copy(client, upstream)
		go io.Copy(upstream, client)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
