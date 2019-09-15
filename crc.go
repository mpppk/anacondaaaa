package anacondaaaa

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type CRCRequest struct {
	CRCToken string `query:"crc_token" mapstructure:"crc_token"`
}

type CRCResponse struct {
	ResponseToken string `json:"response_token"`
}

func CreateCRCToken(crcToken, consumerSecret string) string {
	mac := hmac.New(sha256.New, []byte(consumerSecret))
	mac.Write([]byte(crcToken))
	return "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
