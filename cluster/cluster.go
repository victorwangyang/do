package cluster

//InitPort is starting number of Nodes port
//ClusterAddr is address and port of visiting mater
const (


	MasterAddr       = "http://localhost"
	MasterPort       = "8085"
	ClusterDirectory = "/cluster"

	MasterExePosition = "../master/master"
	MasterIniPosition = "../master/masterpid.ini"

	NodeExePosition = "../node/node"

	NodeAddr       = "http://localhost"
	InitNodePort   = 500
	NodeDirectory = "/node"
)

// Yamlcluster is struct of cluster
type Yamlcluster struct {
	ClusterNumber   string   `yaml:"clusterNumber"`
	MachineNamelist []string `yaml:"machineNamelist"`
}

// NodesInfo to record nodes name,port,PID information of Cluster
type NodesInfo struct {
	NodeName   string
	PortNum    string
	NodePID    string
	NodeStatus bool
}

//Nodes is slice instance of NodesInfo
var Nodes []NodesInfo
