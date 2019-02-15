package cluster

import (

	"log"
	"io/ioutil"
	"do/cmd"
	"strconv"	
	yaml "gopkg.in/yaml.v2"
)
// constant of cluster
const (
	NodeAddr       = "http://localhost"
	NodeDirectory = "/node"
)

// Cluster to record cluster info
type Cluster struct {
	NodesNumber string
	Nodes [] NodesInfo
}

// NodesInfo to record nodes name,port,PID information of Cluster
type NodesInfo struct {
	NodeName   string
	PortNum    string
	NodePID    string
	NodeStatus bool
}

var GCluster = Cluster{}

// InitClusterInfo is to Init sturct from status file 
func InitClusterInfo(file string) {

	// read status file
	status := new(cmd.YamlClusterStatus)
	body, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(body, status)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	GCluster.NodesNumber = status.NodeNumber
	nodeNumber,_ := strconv.Atoi(status.NodeNumber)
	GCluster.Nodes = make([]NodesInfo,nodeNumber, nodeNumber)

	for i := 0; i < nodeNumber; i++ {
		GCluster.Nodes[i].NodeName = status.Nodes[i].NodeName
		GCluster.Nodes[i].PortNum = status.Nodes[i].NodePort
		GCluster.Nodes[i].NodePID = status.Nodes[i].NodePID
		GCluster.Nodes[i].NodeStatus = true
	}

}
