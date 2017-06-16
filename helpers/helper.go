package helpers

import (
	"bytes"
	"encoding/json"
	"os"
)

// GetENVorDefault returns the env value or the default
func GetENVorDefault(env string, dft string) (strcon string) {
	strcon = dft

	if envStr := os.Getenv(env); envStr != "" {
		strcon = envStr
	}

	return
}

// JSONMarshal is a abstraction that permit correct the json scapes
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}
