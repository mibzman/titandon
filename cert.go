package main

import (
	"crypto/tls"
	"time"

	gemCert "github.com/a-h/gemini/cert"
)

func CreateCert() (tls.Certificate, error) {
	parsedTime, err := time.ParseDuration("1000000h")
	if err != nil {
		return tls.Certificate{}, err
	}

	cert, key, err := gemCert.Generate("localhost", "localhost", "localhost", parsedTime)
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.X509KeyPair(cert, key)
}
