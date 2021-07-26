package checksum

import (
	"crypto/sha1"
	"io"
	"strings"
)

const HEX_DIGITS = "0123456789abcdef"

func Encode(appSecret, nonce, time string) (string, error) {
	contentString := appSecret + nonce + time
	messageDigest := sha1.New()
	_, err := io.WriteString(messageDigest, contentString)
	if err != nil {
		return "", nil
	}

	return getFormattedText(messageDigest.Sum(nil)), nil
}

func getFormattedText(bytes []byte) string {
	length := len(bytes)
	buf := strings.Builder{}
	for i := 0; i < length; i++ {
		buf.WriteByte(HEX_DIGITS[(bytes[i]>>4)&0x0f])
		buf.WriteByte(HEX_DIGITS[bytes[i]&0x0f])
	}
	return buf.String()
}
