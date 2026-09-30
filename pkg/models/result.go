package models

import "time"

type ScanResult struct {
	Host      string        `json:"host"`
	Port      int           `json:"port"`
	State     string        `json:"state"` // "open", "closed", "filtered"
	Banner    string        `json:"banner"`
	Service   string        `json:"service"`
	Latency   time.Duration `json:"latency_ms"`
	Timestamp time.Time     `json:"timestamp"`
}

type ScanSummary struct {
	Target        string        `json:"target"`
	TotalPorts    int           `json:"total_ports"`
	OpenPorts     int           `json:"open_ports"`
	ClosedPorts   int           `json:"closed_ports"`
	FilteredPorts int           `json:"filtered_ports"`
	Duration      time.Duration `json:"duration"`
	Results       []ScanResult  `json:"results"`
}
