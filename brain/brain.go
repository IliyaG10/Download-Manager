package brain

import (
	"context"
	"fmt"
	"sync"

	"github.com/IliyaG10/Download-Manager/dataType"
)


type BrainConn struct {

	mu sync.RWMutex
	LeaderConns map[int]*datatype.LeaderConn 
	DatabaseConnAns chan datatype.DatabaseRequest
	DatabaseConnRecive chan datatype.DatabaseAnswer

}


var  (
	instance *BrainConn
	once sync.Once
)
func CreateDownloadRequest(r *datatype.DownloadRequest) {

	numberOfPartition := TestDownloadConnection(r.FileLink,r.Partitionable)
	capacity := int64(r.Size / numberOfPartition)
	partitions := make([]datatype.Partiton,numberOfPartition)

	for i := 0; i < int(numberOfPartition); i++ {
		p := &datatype.Partiton{
			StartBytes: int64(i) * capacity,
			EndBytes: int64(i + 1) * capacity,
			Status: datatype.NotCompelete,
			CurrentBytePosition: int64(i) * capacity,
		}
		partitions[i] = *p
	}

	leaderCtx, err := context.WithCancel(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	response := &datatype.DownloadResponse{
		Link: r.FileLink,
		FilePath: r.FilePath,
		ParentCtx: leaderCtx,
		Size: r.Size,
		Partitionable: r.Partitionable,
		Partitions: partitions,
	}

	fmt.Println(response)
}


func TestDownloadConnection (link string, partitionAble bool) (int64) {
	return 1
}

// singlton intance 
func getBrain() *BrainConn {
	once.Do(func() {
        instance = &BrainConn{
			LeaderConns: make(map[int]*datatype.LeaderConn),
            DatabaseConnAns:    make(chan datatype.DatabaseRequest),
            DatabaseConnRecive: make(chan datatype.DatabaseAnswer),
        }
    })
	return instance
}


func IdGenerator() int {
	return 1
}


func AddLeader(leader *datatype.LeaderConn) {
	getBrain().mu.Lock()
	defer getBrain().mu.Unlock()
	getBrain().LeaderConns[IdGenerator()] = leader
}


func GetLeaderChannel(Id int64) *datatype.LeaderConn {
	return getBrain().LeaderConns[int(Id)]
}

func GetDatabaseConnAns() chan datatype.DatabaseRequest {
	
	return getBrain().DatabaseConnAns
}


func GetDatabaseConnRecive() chan datatype.DatabaseAnswer {
	return getBrain().DatabaseConnRecive
}