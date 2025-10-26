package common

import (
	"context"
	"time"
)



type DownloadWorker struct {
	ID int64
	Partiton Partiton
	IDMaster int64
	Cancel context.CancelFunc
	Ctx context.Context

}

type WorkerHandler struct {
	ID int64
	Ctx context.Context
	Cancel context.CancelFunc
	LeaderID int64
	Heartbeat []chan time.Time
	Msg []chan ACK
}


type DownloadLeader struct {
	Id int64
	// ParentCtx context.Context
	WorkerHandlerID int64 
}

const (
	HEALTHY string = "healthy"
)

type ACK struct {
	ID string
	Status string
}