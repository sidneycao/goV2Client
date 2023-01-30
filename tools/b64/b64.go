package b64

import (
	"encoding/base64"
	"log"
)

func B64Decoder(body string) string {
	b := make([]byte, base64.RawStdEncoding.DecodedLen(len(body)))
	d, err := base64.StdEncoding.Decode(b, []byte(body))
	if err != nil {
		log.Panic(err)
	}
	return string(b[:d])
}
