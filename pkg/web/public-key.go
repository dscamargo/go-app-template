package web

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ReadPublicKey(filePath string) (any, error) {
	f, err := os.ReadFile(filePath)
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
