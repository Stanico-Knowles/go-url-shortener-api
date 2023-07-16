package base64encryptionservice

import (
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"
)

func EncodeBase64(plainText string) string {
	data := strconv.FormatInt(time.Now().Unix(), 10) + plainText
	hash := sha1.Sum([]byte(data))
	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}
