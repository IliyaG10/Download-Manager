package brain

import (
	"context"
	"fmt"

	"github.com/IliyaG10/Download-Manager/common/dataType"
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

