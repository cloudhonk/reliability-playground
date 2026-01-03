package system

import (
	"runtime"
	"time"
)

var StartTime = time.Now()

func Uptime() time.Duration {
	return time.Since(StartTime)
}

func MemoryMB() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / 1024 / 1024
}
