package auth

import (
  "crypto/sha512"
  "crypto/hmac"
  "github.com/stormentt/dumbstored/random"
  "github.com/stormentt/dumbstored/config"
  "encoding/base64"
  "golang.org/x/crypto/bcrypt"
)

func prehashPassword(pw, salt string) (string) {
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

  return string(pwhash), salt
}

func CheckPassword(pw, hash, salt string) bool {
  prehash := prehashPassword(pw, salt)

  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(prehash))
  if err != nil {
    return false
  }

  return true
}
