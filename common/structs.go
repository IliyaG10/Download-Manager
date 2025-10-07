package common

type Downloadable struct {
	FilePath      string
	FileLink      string
	Size          int64
	Partitionable bool
	Partitions    []Partiton
	TotalDownloadingBytes int64
}

const (
	Compelete = "compelete"
	Paused    = "paused"
	Failed    = "failed"
)

type Partiton struct {
	StartBytes int64
	EndBytes   int64
	Status     string
	// current bytes position
	CurrentBytePosition int64
	
}
