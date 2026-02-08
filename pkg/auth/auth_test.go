package auth

import (
	"strings"
	"testing"
	"time"
)

func TestHashPassword_And_Check(t *testing.T) {
	hash, err := HashPassword("secret123")
	if err != nil {
		t.Fatalf("HashPassword: %v", err)
	}
	if !CheckPassword(hash, "secret123") {
		t.Fatal("expected password to match")
	}
	if CheckPassword(hash, "wrong") {
		t.Fatal("expected password not to match")
	}
}

func TestIssueToken(t *testing.T) {
	tok := IssueToken()
	if !strings.HasPrefix(tok, "tok-") {
		t.Fatalf("expected tok- prefix, got %s", tok)
	}
}

func TestHashToken(t *testing.T) {
	h := HashToken("mytoken")
	if len(h) != 64 {
		t.Fatalf("expected 64 char hex, got %d", len(h))
	}
	// Same input should produce same hash
	if HashToken("mytoken") != h {
		t.Fatal("expected deterministic hash")
	}
}

func TestExpireAt(t *testing.T) {
	exp := ExpireAt(1)
	if exp.Before(time.Now()) {
		t.Fatal("expected future time")
	}
	// Default to 24h when 0
	exp0 := ExpireAt(0)
	if exp0.Before(time.Now().Add(23 * time.Hour)) {
		t.Fatal("expected ~24h from now")
	}
}
