package helpers

import (
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
