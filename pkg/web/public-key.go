package pkg

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ReadPublicKey() (any, error) {
	f, err := os.ReadFile("ssl/public.key")
	if err != nil {
		return "", fmt.Errorf("error reading public key: %w", err)
	}
	pubPem, _ := pem.Decode(f)
	parsedKey, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return "", fmt.Errorf("error parsing public key: %v", err)
	}
	return parsedKey, nil
}
