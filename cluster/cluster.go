package cluster

// constant of cluster
const (
	MasterAddr       = "http://localhost"
	MasterPort       = "8085"
	ClusterDirectory = "/cluster"

	MasterExePosition = "../master/master"
	StatusPosition = "./status.yaml"

	NodeExePosition = "../node/node"

	NodeAddr       = "http://localhost"
	DefaultNodePort   = 500
	NodeDirectory = "/node"
)

// NodesInfo to record nodes name,port,PID information of Cluster
type NodesInfo struct {
	NodeName   string
	PortNum    string
	NodePID    string
	NodeStatus bool
}

//Nodes is slice instance of NodesInfo
var Nodes []NodesInfo

func initNodes(len int) {

	Nodes = make([]NodesInfo, len, len)
}
