package datatype

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
	Healthy = "alive"
	Dead = "dead"
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
 
type LeaderResponse struct {
	Status string
	ID int
	WorkerConditions []Partiton
}


type DatabaseRequest struct {
	// using the exited const here
	Order string
	Err error
}

type DatabaseAnswer struct {

}


type LeaderConn struct {
	LeaderConnAns chan DownloadResponse
	LeaderConnRecive chan LeaderResponse
}
