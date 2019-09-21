package secrets

import (
	"log"
	"strings"
)

func Decrypt(secretPath string) string {
	log.Printf("processing secret: %v\n", secretPath)

	data := SecretsConfigurationInstance.EncryptionConfiguration.DecryptFile(secretPath)
	return strings.TrimSpace(string(data[:]))
}