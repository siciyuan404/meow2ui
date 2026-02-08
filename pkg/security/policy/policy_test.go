package policy

import "testing"

func TestValidateMediaRef(t *testing.T) {
	e := NewEngine()

	blocked := e.ValidateMediaRef("http://127.0.0.1/a.png", []string{"example.com"})
	if blocked.Allowed {
		t.Fatalf("expected local address blocked")
	}

	notAllowedHost := e.ValidateMediaRef("https://evil.com/a.png", []string{"example.com"})
	if notAllowedHost.Allowed {
		t.Fatalf("expected host not allowed")
	}

	allowed := e.ValidateMediaRef("https://cdn.example.com/a.png", []string{"example.com"})
	if !allowed.Allowed {
		t.Fatalf("expected allowed host")
	}

	bucket := e.ValidateMediaRef("s3://bucket/path/a.png", nil)
	if !bucket.Allowed {
		t.Fatalf("expected s3 allowed")
	}
}
