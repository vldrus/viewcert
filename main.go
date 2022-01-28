package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "cert.crt")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var certs [][]byte

	for {
		var block *pem.Block

		block, b = pem.Decode(b)
		if block == nil {
			break
		}

		if block.Type == "CERTIFICATE" {
			certs = append(certs, block.Bytes)
		}
	}

	if len(certs) == 0 {
		fmt.Println("Error: cannot parse certificate file", os.Args[1])
	}

	for _, c := range certs {
		cert, err := x509.ParseCertificate(c)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Type:                  ", cert.PublicKeyAlgorithm)
		fmt.Println("Serial:                ", cert.SerialNumber)
		fmt.Println("Version:               ", cert.Version)
		fmt.Println("Subject:               ", cert.Subject)
		fmt.Println("Issuer:                ", cert.Issuer)
		fmt.Println("Not Before:            ", cert.NotBefore)
		fmt.Println("Not After:             ", cert.NotAfter)
		fmt.Println("IsCA:                  ", cert.IsCA)
		fmt.Println("BasicConstraintsValid: ", cert.BasicConstraintsValid)
		fmt.Println("KeyUsage:              ", cert.KeyUsage)
		fmt.Println("ExtKeyUsage:           ", cert.ExtKeyUsage)
		fmt.Println("DNS Names:             ", cert.DNSNames)
		fmt.Println("IP Addresses:          ", cert.IPAddresses)

		fmt.Println()
	}
}
