package auth

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"strings"

	"github.com/stormentt/dumbstored/config"
	"github.com/stormentt/dumbstored/random"
	"golang.org/x/crypto/bcrypt"
)

type Dummy struct {
	Password string
	Salt     string
	Hash     string
}

var dummy Dummy

func GenerateDummy() {
	dummy.Password = random.AlphaNum(16)

	dummy.Salt, dummy.Hash = HashPassword(dummy.Password)
}

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
func prehashPassword(pw, salt string) string {
	hasher := hmac.New(sha512.New384, []byte(salt))
	hasher.Write([]byte(pw))

	prehash := hasher.Sum(nil)
	prehashB64 := base64.StdEncoding.EncodeToString(prehash)

	return prehashB64
}

func HashPassword(pw string) (string, string) {
	salt := random.Salt()
	prehash := prehashPassword(pw, salt)

	pwhash, _ := bcrypt.GenerateFromPassword([]byte(prehash), config.C.BcryptFactor)

	return salt, string(pwhash)
}

func CheckPassword(pw, hash, salt string) bool {
	prehash := prehashPassword(pw, salt)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(prehash))
	if err != nil {
		return false
	}

	return true
}

func DummyCheckPassword() {
	CheckPassword(dummy.Password, dummy.Hash, dummy.Salt)
}
