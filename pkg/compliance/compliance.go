package compliance

type CheckResult struct {
	Name   string
	Passed bool
	Detail string
}

func CheckKeyRotationPolicy(lastRotatedDays int) CheckResult {
	if lastRotatedDays <= 90 {
		return CheckResult{Name: "key_rotation", Passed: true, Detail: "ok"}
	}
	return CheckResult{Name: "key_rotation", Passed: false, Detail: "rotation overdue"}
}

func CheckAuditRetentionPolicy(days int) CheckResult {
	if days <= 365 {
		return CheckResult{Name: "audit_retention", Passed: true, Detail: "ok"}
	}
	return CheckResult{Name: "audit_retention", Passed: false, Detail: "retention too long"}
}
