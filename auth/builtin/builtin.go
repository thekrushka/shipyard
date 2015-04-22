package builtin

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/shipyard/shipyard/auth"
)

type (
	BuiltinAuthenticator struct {
		salt []byte
	}
)

func NewAuthenticator(salt string) auth.Authenticator {
	return &BuiltinAuthenticator{
		salt: []byte(salt),
	}
}

func (a BuiltinAuthenticator) IsUpdateSupported() bool {
	return true
}

func (a BuiltinAuthenticator) Name() string {
	return "builtin"
}

func (a BuiltinAuthenticator) Authenticate(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err == nil {
		return true
	}
	return false
}

func (a BuiltinAuthenticator) GenerateToken() (string, error) {
	return auth.GenerateToken()
}