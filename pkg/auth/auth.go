package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func IssueToken() string {
	return util.NewID("tok")
}

func HashToken(raw string) string {
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

func ExpireAt(hours int) time.Time {
	if hours <= 0 {
		hours = 24
	}
	return time.Now().Add(time.Duration(hours) * time.Hour)
}
