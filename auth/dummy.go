package auth

import (
	"fmt"

	"github.com/stormentt/dumbstored/random"
)

// we have a dummy so we have something to check against incase someone tries to log into a user that doesn't exist.
// this prevents a possible username enumeration timing attack
type Dummy struct {
	Password string
	Salt     string
	Hash     string
}

func (d Dummy) String() string {
	return fmt.Sprintf("Dummy<%s,%s,%s>", d.Salt, d.Password, d.Hash)
}

var dummy Dummy

func init() {
	dummy.Password = random.AlphaNum(16)

	dummy.Salt, dummy.Hash = HashPassword(dummy.Password)

	fmt.Printf("Initialized %s\n", dummy)
}

func DummyCheckPassword() {
	CheckPassword(dummy.Password, dummy.Hash, dummy.Salt)
}
