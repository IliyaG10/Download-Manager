package common

import(
	"module github.com/IliyaG10/Download-Manager/common/dataType"
)

type Connections struct {

	mu sync.RWMutex
	LeaderConns map[int]*datatype.LeaderConn 
	DatabaseConnAns chan datatype.DatabaseRequest
	DatabaseConnRecive chan datatype.DatabaseAnswer

}


var  (
	instance *Connections
	once sync.Once
)