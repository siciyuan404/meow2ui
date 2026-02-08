package compliance

import "testing"

func TestCheckKeyRotationPolicy(t *testing.T) {
	ok := CheckKeyRotationPolicy(30)
	if !ok.Passed {
		t.Fatal("expected pass for 30 days")
	}
	boundary := CheckKeyRotationPolicy(90)
	if !boundary.Passed {
		t.Fatal("expected pass for 90 days")
	}
	fail := CheckKeyRotationPolicy(91)
	if fail.Passed {
		t.Fatal("expected fail for 91 days")
	}
}

func TestCheckAuditRetentionPolicy(t *testing.T) {
	ok := CheckAuditRetentionPolicy(365)
	if !ok.Passed {
		t.Fatal("expected pass for 365 days")
	}
	fail := CheckAuditRetentionPolicy(366)
	if fail.Passed {
		t.Fatal("expected fail for 366 days")
	}
}
