// +build !tor

package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"time"
)

func connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		return nil, err
	}
	log.Println("connect")
	caPool := x509.NewCertPool()
	caPool.AddCert(s.cert)

	config := tls.Config{RootCAs: caPool, ServerName: s.domain}
	tlsconn := tls.Client(conn, &config)
	if err != nil {
		return nil, err
	}
	return tlsconn, nil
}

func NetClient() {
	log.Println("NetClient")
	initServer()
	for {
		conn, err := connect()
		if err != nil {
			log.Println("Could not connect:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		RPC(conn)
	}
}
