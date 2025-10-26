package common

import "context"

type DownloadRequest struct {
	FilePath      string
	FileLink      string
	Size          int64
	Partitionable bool
}

const (
	Compelete    = "compelete"
	Paused       = "paused"
	Failed       = "failed"
	NotCompelete = "notcompelete"
)

type Partiton struct {
	StartBytes int64
	EndBytes   int64
	Status     string
	// current bytes position
	CurrentBytePosition int64
}

type DownloadResponse struct {
	Link          string
	FilePath      string
	ParentCtx     context.Context
	Size          int64
	Partitionable bool
	Partitions    []Partiton
}
 

type Result struct {
	// using the exited const here
	Status string
	Err error
}
