package telemetry

import "testing"

func TestService_RecordAndSnapshot(t *testing.T) {
	svc := NewService()
	svc.RecordRun(true, 100)
	svc.RecordRun(false, 200)
	svc.RecordRun(true, 50)

	snap := svc.Snapshot()
	if snap.TotalRuns != 3 {
		t.Fatalf("expected 3 total, got %d", snap.TotalRuns)
	}
	if snap.FailedRuns != 1 {
		t.Fatalf("expected 1 failed, got %d", snap.FailedRuns)
	}
	if snap.TotalLatencyMS != 350 {
		t.Fatalf("expected 350ms, got %d", snap.TotalLatencyMS)
	}
	if snap.LastUpdated.IsZero() {
		t.Fatal("expected LastUpdated to be set")
	}
}

func TestService_EmptySnapshot(t *testing.T) {
	svc := NewService()
	snap := svc.Snapshot()
	if snap.TotalRuns != 0 || snap.FailedRuns != 0 || snap.TotalLatencyMS != 0 {
		t.Fatal("expected all zeros")
	}
}
