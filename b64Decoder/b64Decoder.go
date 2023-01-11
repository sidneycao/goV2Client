package b64Decoder

import (
	"encoding/base64"
	"log"
)

func Decoder(body []byte) []byte {
	b64 := make([]byte, base64.RawStdEncoding.DecodedLen(len(body)))
	d, err := base64.StdEncoding.Decode(b64, body)
	if err != nil {
		log.Panic(err)
	}
	return b64[:d]
}
