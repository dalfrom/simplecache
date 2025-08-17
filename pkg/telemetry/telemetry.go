package telemetry

import (
	"fmt"
	"time"
)

type Telemetry struct {
	startTime time.Time
}

func NewTelemetry() *Telemetry {
	return &Telemetry{
		startTime: time.Now(),
	}
}

func (t *Telemetry) Report() {
	elapsed := time.Since(t.startTime)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
