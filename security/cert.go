package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

// LoadServerTLSCredentials provides loading of server tls credentials.
func LoadServerTLSCredentials(cert_path string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key.
	serverCert, err := tls.LoadX509KeyPair(cert_path+"/server-cert.pem", cert_path+"/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it.
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

// LoadCATLSCredentials provides loading of cat tls credentials.
func LoadCATLSCredentials(cert_path string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(cert_path + "/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
