package nimbostratus

import "time"

type Region struct {
	Id      string
	Name    string
	Latency time.Duration
}
