package auth

import (
	"encoding/base64"
	"strings"
)

func DecodeHeader(authHeader string) (string, string, bool) {
	if len(authHeader) == 0 {
		return "", "", false
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 {
		return "", "", false
	}

	authMethod := authSplit[0]
	if authMethod != "Basic" {
		return "", "", false
	}

	authChunk, err := base64.StdEncoding.DecodeString(authSplit[1])
	if err != nil {
		return "", "", false
	}

	decodedSplit := strings.SplitN(string(authChunk), ":", 2)
	if len(decodedSplit) != 2 {
		return "", "", false
	}

	username := decodedSplit[0]
	password := decodedSplit[1]

	return username, password, true
}
