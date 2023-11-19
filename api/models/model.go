package models

import "time"

type Log struct {
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	ResourceId string    `json:"resourceId"`
	Timestamp  time.Time `json:"timestamp"`
	TraceId    string    `json:"traceId"`
	SpanId     string    `json:"spanId"`
	Commit     string    `json:"commit"`
	Metadata   Metadata  `json:"metadata"`
}

type Metadata struct {
	ParentResouceId string `json:"parentResourceID"`
}
