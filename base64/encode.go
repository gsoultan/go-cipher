package base64

import "encoding/base64"

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
