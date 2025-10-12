package main


import (
	"github.com/IliyaG10/Download-Manager/common"
)

type DownloadWorker struct {
	ID int64
	Partiton common.Partiton
	IDMaster int64
}


type DownloadMaster struct {
	Id int64
	Workers []DownloadWorker
}


func Partiting(size int, speed int) {
	// make partitioning
}

func downloading(link string) {
	// donwloading process 
}

func healthCondiniting() {
	// make a reliable heartbeat
}
