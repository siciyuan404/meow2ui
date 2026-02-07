package tracing

import "time"

type Span struct {
	Name      string
	StartedAt time.Time
	EndedAt   time.Time
	Error     string
}

func StartSpan(name string) Span {
	return Span{Name: name, StartedAt: time.Now()}
}

func EndSpan(span *Span, err error) {
	span.EndedAt = time.Now()
	if err != nil {
		span.Error = err.Error()
	}
}
