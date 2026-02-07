package auditexport

import (
	"fmt"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Job struct {
	ID          string
	Format      string
	Status      string
	Artifact    string
	RequestedBy string
	CreatedAt   time.Time
	FinishedAt  *time.Time
}

func Create(format, requestedBy string) (Job, error) {
	if format != "json" && format != "csv" {
		return Job{}, fmt.Errorf("invalid format")
	}
	return Job{ID: util.NewID("aexp"), Format: format, Status: "queued", RequestedBy: requestedBy, CreatedAt: time.Now()}, nil
}
