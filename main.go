package main

import (
	"log"
	"os"
	"secret-share/app"
)

const (
	defaultCertificate = "./ssl/cert.pem"
	defaultPricateKey  = "./ssl/key.pem"
	defaultHTTPPort    = "8080"
	defaultHTTPSPort   = "8443"
)

var (
	tlsCert    = os.Getenv("TLS_CERTIFICATE")
	tlsKey     = os.Getenv("TLS_PRIVATE_KEY")
	serveTLS   = os.Getenv("SERVE_TLS")
	listenPort = os.Getenv("PORT")
	listenAddr = getenv("ADDR", "127.0.0.1")
)

func main() {
	var err error
	app := app.New()
	cert, key := tlsFiles()

	if cert != "" && key != "" {
		app.Addr = addr(defaultHTTPSPort, "https")
		err = app.ListenAndServeTLS(cert, key)
	} else {
		app.Addr = addr(defaultHTTPPort, "http")
		err = app.ListenAndServe()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func addr(port, protocol string) string {
	address := listenAddr + ":"
	if listenPort != "" {
		address += listenPort
	} else {
		address += port
	}
	log.Printf("Listening on: %s://%s", protocol, address)
	return address
}

func tlsFiles() (string, string) {
	if serveTLS == "false" {
		return "", ""
	}

	cert := defaultCertificate
	key := defaultPricateKey

	if _, err := os.Stat(cert); err != nil {
		cert = ""
	}
	if _, err := os.Stat(key); err != nil {
		key = ""
	}

	if tlsCert != "" {
		if _, err := os.Stat(tlsCert); err != nil {
			log.Fatalf("unable to access TLS certificate: %s", err.Error())
		}
		cert = tlsCert
	}

	if tlsKey != "" {
		if _, err := os.Stat(tlsKey); err != nil {
			log.Fatalf("unable to access TLS private key: %s", err.Error())
		}
		key = tlsKey
	}

	return cert, key
}

func getenv(envVar, defVar string) string {
	v := os.Getenv(envVar)
	if v == "" {
		return defVar
	}
	return v
}
