package redisClient

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"io/ioutil"
	"log"
	"os"
)

func CreateTLSConfig(caCertFile, clientCertFile, clientKeyFile string) *tls.Config {
	// Load CA certificate
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Failed to read CA certificate file: %v", err)
	}

	// Load client certificate and key
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatalf("Failed to load client certificate/key: %v", err)
	}

	// Create a CertPool containing the CA certificate
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{clientCert},
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
	}

	return tlsConfig
}

func GetConnection() *redis.Client {
	tlsConfig := CreateTLSConfig(os.Getenv("caCertFile"), os.Getenv("clientCertFile"), os.Getenv("clientKeyFile"))

	return redis.NewClient(&redis.Options{
		Addr:      "redis:6379",
		Password:  os.Getenv("REDIS_PASSWORD"),
		DB:        0,
		TLSConfig: tlsConfig,
	})
}
