package base64encryptionservice

import (
	"encoding/base64"
)

func EncodeBase64(plainText string) string {
	return base64.StdEncoding.EncodeToString([]byte(plainText))[:8]
}