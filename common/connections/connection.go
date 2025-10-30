package connections

import(
	"github.com/IliyaG10/Download-Manager/common/dataType"
	"sync"
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


// singlton intance 
func GetConnections() *Connections{
	once.Do(func() {
        instance = &Connections{
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
	GetConnections().mu.Lock()
	defer GetConnections().mu.Unlock()
	GetConnections().LeaderConns[IdGenerator()] = leader
}


func GetLeaderChannel(Id int64) *datatype.LeaderConn {
	return GetConnections().LeaderConns[int(Id)]
}

func GetDatabaseConnAns() chan datatype.DatabaseRequest {
	
	return GetConnections().DatabaseConnAns
}


func GetDatabaseConnRecive() chan datatype.DatabaseAnswer {
	return GetConnections().DatabaseConnRecive
}