package algorithm

import (
	"encoding/base64"
)

func EncodeBase64(plainText string) string {
	// Read the plain text from the request body
	text := plainText

	// Convert the plain text to Base64
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))

	// Return the Base64 encoded text as the response
	return encodedText
}
