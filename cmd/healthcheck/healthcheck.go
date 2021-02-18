package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var (
		port    string
		tlsPort string
	)

	flag.StringVar(&port, "port", "8080", "the port to listen on for insecure connections, defaults to a random value")
	flag.StringVar(&tlsPort, "tls-port", "8443", "the port to listen on for secure connections, defaults to a random value")
	flag.Parse()

	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/ping", stringEnvOverride(port, "PORT")))
	if err == nil {
		os.Exit(0)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, errtls := http.Get(fmt.Sprintf("https://127.0.0.1:%s/ping", stringEnvOverride(tlsPort, "TLS_PORT")))
	if errtls == nil {
		os.Exit(0)
	}

	os.Exit(1)
}

func stringEnvOverride(orig string, key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	return orig
}
