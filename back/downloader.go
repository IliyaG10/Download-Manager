package main

// import (
// 	"container/list"
// 	"context"
// 	"crypto/rand"
// 	"fmt"
// 	"io"
// 	"log/slog"
// 	"math/big"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/IliyaG10/Download-Manager/common"
// )



// type DownloadWorker struct {
// 	ID int64
// 	Partiton common.Partiton
// 	IDMaster int64
// 	Cancel context.CancelFunc
// 	Ctx context.Context

// }


// type WorkerHandler struct {
// 	ID int64
// 	Ctx context.Context
// 	Cancel context.CancelFunc
// 	LeaderID int64
// 	Heartbeat []chan time.Time
// 	Msg []chan ACK
// }


// type DownloadLeader struct {
// 	Id int64
// 	// ParentCtx context.Context
// 	WorkerHandlerID int64 
// }

// const (
// 	HEALTHY string = "healthy"
// )

// type ACK struct {
// 	ID string
// 	Status string
// }

// func generateRandomID() int64 {
// 	n, _ := rand.Int(rand.Reader, big.NewInt(1_000_000_000_000))
// 	return n.Int64()
// }

// var workersList list.List

// func Partitiong(link string, workerNumbers int) {
// 	resp , err := http.Get(link)
	
// 	if err != nil {
// 		slog.Info(err.Error())
// 	}

// 	contentLngth := resp.Header.Get("Content-length")

// 	fileSize, err := strconv.Atoi(contentLngth)

// 	if err != nil {
// 		slog.Info(err.Error())
// 	}

// 	step := fileSize/workerNumbers

// 	var leader DownloadLeader
// 	leader.Id = generateRandomID()
// 	// leader.ParentCtx , _ = context.WithCancel(context.Background())
// 	parentCtx, cancelCtx := context.WithCancel(context.Background())

// 	wh := &WorkerHandler {
// 		ID: generateRandomID(),
// 		Ctx: parentCtx,
// 		Cancel: cancelCtx,
// 		LeaderID: leader.Id,
// 	}


// 	for i := 0; i < workerNumbers; i++ {
// 		p := &common.Partiton{
// 			StartBytes: int64(i*step),
// 			EndBytes: int64(i*step + step),
// 			Status: common.NotCompelete,
// 			CurrentBytePosition: int64(i*step),
// 		}

// 		wCtx, wCancel := context.WithCancel(wh.Ctx)

// 		w := &DownloadWorker {
// 			ID: generateRandomID(),
// 			Partiton: *p,
// 			IDMaster: leader.Id,
// 			Cancel: wCancel,
// 			Ctx: wCtx,
			
// 		}

// 		workersList.PushFront(w)


// 	}

// }

// func MakingFile(path string,size int64) *os.File {
// 	// making the file from the scratch
// 	file, err := os.OpenFile(path,os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		slog.Info(err.Error())
// 	}

// 	file.Truncate(size)

// 	return file
// }

// func WorkerDownloader(link string,worker *DownloadWorker ,file os.File) error {
// 	req, err := http.NewRequest("GET",link,nil)
// 	if err != nil {
// 		slog.Info(err.Error())
// 	}

//     req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", worker.Partiton.StartBytes, worker.Partiton.EndBytes))
// 	resp, err := http.DefaultClient.Do(req)

// 	if err != nil {
// 		slog.Info(err.Error())
// 	}

// 	defer resp.Body.Close()

// 	buff := make([]byte, 128*1024) // 128 kilo byte per second 
// 	offset := worker.Partiton.StartBytes
// 	for {
//         n, err := resp.Body.Read(buff)
//         if n > 0 {
//             _, werr := file.WriteAt(buff[:n], offset)
//             if werr != nil {
//                 return werr
//             }
//             offset += int64(n)
//         }
//         if err == io.EOF {
//             break
//         }
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }

// func healthCondiniting() {
// 	// make a reliable heartbeat
// }