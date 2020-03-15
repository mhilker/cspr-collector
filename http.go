package csprcollector

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func NewHttpClient(certFile, keyFile, caFile string) *http.Client {
	var certs []tls.Certificate
	if certFile != "" && keyFile != "" {
		certs = append(certs, newClientCert(certFile, keyFile))
	}

	var caCertPool *x509.CertPool
	if caFile != "" {
		caCertPool = newCaCertPool(caFile)
	}

	tlsConfig := &tls.Config{
		Certificates: certs,
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return &http.Client{Transport: transport}
}

func newClientCert(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func newCaCertPool(caFile string) *x509.CertPool {
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool
}
