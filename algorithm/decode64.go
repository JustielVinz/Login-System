package algorithm

import (
	"encoding/base64"
	"fmt"
)

func DecodeBase64(plainText string) string {
	// Read the plain text from the request body
	text := plainText

	// Decode the base64 string
	decodedData, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		// Handle decoding error
		fmt.Println("decode error: ", err.Error())
		return err.Error()
	}

	// Convert the decoded bytes to a string (if it's text)
	originalString := string(decodedData)

	return originalString
}
