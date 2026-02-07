package backup

import (
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Job struct {
	ID        string
	Type      string
	Status    string
	Artifact  string
	SizeBytes int64
	Checksum  string
	StartedAt time.Time
	EndedAt   *time.Time
	Error     string
}

func RunFullBackup() Job {
	now := time.Now()
	end := now.Add(2 * time.Second)
	return Job{ID: util.NewID("bkp"), Type: "full", Status: "completed", Artifact: "file://backup.dump", SizeBytes: 1024, Checksum: "mock-checksum", StartedAt: now, EndedAt: &end}
}
