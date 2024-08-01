package url

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func hashEmail(email string) string {
	email = strings.ToLower(email)

	inputData := []byte(email)
	outputData := sha256.Sum256(inputData)
	hash := hex.EncodeToString(outputData[:])
	// log.Printf("%s\n", hash)
	return hash

}
