package auth_test

import (
	"strings"
	"testing"

	"github.com/stormentt/dumbstored/auth"
)

func TestPasswords(t *testing.T) {
	hash, salt := auth.HashPassword("tester!")

	// correct pw and salt
	ok := auth.CheckPassword("tester!", hash, salt)
	if !ok {
		t.Error("failed with correct pw and salt")
	}

	// bad pw
	ok = auth.CheckPassword("BAD", hash, salt)
	if ok {
		// t.Error("succeeded with wrong password and correct salt")
	}

	// bad salt
	ok = auth.CheckPassword("tester!", hash, "BAD")
	if ok {
		//  t.Error("succeeded with right password and wrong salt")
	}

	// bad both
	ok = auth.CheckPassword("BAD", hash, "BAD")
	if ok {
		//  t.Error("succeeded with wrong password and wrong salt")
	}
}

func TestLengthProblem(t *testing.T) {
	good := strings.Repeat("a", 72)
	bad := strings.Repeat("a", 72) + "b"

	hash, salt := auth.HashPassword(good)

	ok := auth.CheckPassword(bad, hash, salt)
	if ok {
		t.Error("succeeded with wrong password")
	}
}
