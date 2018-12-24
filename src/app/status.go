//
// status
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2017-07-27 11:37:13
//

package app

import (
	"encoding/json"
	"os"
	"runtime"
	"time"
)

var InstanceStartTime = time.Now().Unix()

// Status - the standard status struct
type Status struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	Service   string `json:"service"`
	PID       int    `json:"pid"`
	CPUs      int    `json:"cpus"`
	GoVers    string `json:"go"`
	TimeStamp int64  `json:"ts"`
	UpTime    int64  `json:"uptime-seconds"`
	LogLevel  int    `json:"loglevel"`
	Hostname  string `json:"hostname"`
}

// GetStatus return the current status struct
func GetStatus(cfg *Config) Status {
	now := time.Now().Unix()

	s := Status{
		PID:     os.Getpid(),
		Service: "list-service",
	}

	s.Status = "ok"
	s.Version = Version()
	s.CPUs = runtime.NumCPU()
	s.GoVers = runtime.Version()
	s.TimeStamp = now
	s.UpTime = now - InstanceStartTime
	s.LogLevel = log.GetLevel()

	if host, err := os.Hostname(); err == nil {
		s.Hostname = host
	}

	return s
}

// GetStatusAsJSON return the current status as a json string
func GetStatusAsJSON(cfg *Config) string {
	status := GetStatus(cfg)
	blob, _ := json.Marshal(status)

	return string(blob)
}
