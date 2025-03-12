package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

func ValidateSign(unsignedStr, sign, token string) bool {
	mac := hmac.New(sha256.New, []byte(token))
	mac.Write([]byte(unsignedStr))
	expectedMAC := mac.Sum(nil)

	return hmac.Equal([]byte(sign), expectedMAC)
}

func Sign(str, token string) string {
	mac := hmac.New(sha256.New, []byte(token))
	mac.Write([]byte(str))
	result := mac.Sum(nil)

	return string(result)
}
