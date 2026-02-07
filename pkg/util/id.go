package util

import (
	"fmt"
	"sync/atomic"
	"time"
)

var seq uint64

func NewID(prefix string) string {
	n := atomic.AddUint64(&seq, 1)
	return fmt.Sprintf("%s-%d-%d", prefix, time.Now().UnixNano(), n)
}
