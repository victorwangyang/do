package cmd


import (
	"os"
	"log"
	"io/ioutil"
	"syscall"
	"strconv"	
	"os/exec"
	yaml "gopkg.in/yaml.v2"
)

//Exe and status position file
const (
	MasterExePosition = "../master/master"
	StatusPosition = "./status.yaml"
	NodeExePosition = "../node/node"
	DefaultNodePort   = 500
)

// YamlClusterConfig is struct of cluster config file
type YamlClusterConfig struct {
	ClusterNumber   string   `yaml:"clusterNumber"`
	MachineNamelist []string `yaml:"machineNamelist"`
}

// YamlClusterStatus is struct of cluster status file
type YamlClusterStatus struct{
	MasterPID     string   `yaml:"masterpid"`
	NodeNumber    string   `yaml:"nodenumber"`
	Nodes        []YamlNodeStatus `yaml:"nodestatus"`
}

// YamlNodeStatus is struct of cluster status file
type YamlNodeStatus struct {
	NodePID       string `yaml:"nodepid"`
	NodePort      string `yaml:"nodeport"`
	NodeName      string `yaml:"nodename"`
}

// StartCluster is to start a cluster
func StartCluster(yamlfile string) {

	// read config file
	config := new(YamlClusterConfig)
	body, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(body, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// new status struct and set staus values
	status := new(YamlClusterStatus)
	status.NodeNumber = config.ClusterNumber
	clusterNumber,_ := strconv.Atoi(config.ClusterNumber)
	status.Nodes = make([]YamlNodeStatus,clusterNumber,clusterNumber)


	// start Master process and start Nodes process ,finally save status to file
	var Args = []string{" of Do Project"}
	cmd := exec.Command(MasterExePosition, Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	status.MasterPID = strconv.Itoa(cmd.Process.Pid)
	log.Printf("Master Deamon(PID:%s) is starting......",status.MasterPID)

	startNodesDeamon(status,config)
	saveClusterStatus(status)
	return
}

// KillCluster is to kill the Cluster
func KillCluster() {

	// get status from file
	status := new(YamlClusterStatus)
	getClusterStatus(status)
	nodeNumber,_ := strconv.Atoi(status.NodeNumber)

	// kill nodes first
	for i := 0; i < nodeNumber; i++ {
		log.Printf("Node Deamon(PID:%s)  Stopping......", status.Nodes[i].NodePID)
		Pid,_:= strconv.Atoi(status.Nodes[i].NodePID)
		syscall.Kill(Pid, syscall.SIGKILL)
	}

	// then kill master last
	Pid,_ := strconv.Atoi(status.MasterPID)
	syscall.Kill(int(Pid), syscall.SIGKILL)
	log.Printf("Master Deamon(PID:%s) is stopping......",status.MasterPID)
	return
}


func saveClusterStatus(status *YamlClusterStatus) {

	buf, err := yaml.Marshal(&status)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(StatusPosition, buf, 0x777)
	if err != nil {
		panic(err.Error())
	}
	return
}

func getClusterStatus(status *YamlClusterStatus) {

	buf, err := ioutil.ReadFile(StatusPosition)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(buf, status)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return
}

func startNodesDeamon(status *YamlClusterStatus,config *YamlClusterConfig){
	
	nodeNumber,_ := strconv.Atoi(status.NodeNumber)

	for i := 0; i < nodeNumber; i++ {

	
		status.Nodes[i].NodeName = config.MachineNamelist[i]
		status.Nodes[i].NodePort = strconv.Itoa(DefaultNodePort + i)
	
		var Args = []string{config.MachineNamelist[i], status.Nodes[i].NodePort}

		cmd := exec.Command(NodeExePosition, Args...) 
		cmd.Stdin = os.Stdin                 
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() 

		status.Nodes[i].NodePID = strconv.Itoa(cmd.Process.Pid)
		log.Printf("Node Deamon(PID:%s)  Starting......", status.Nodes[i].NodePID)
	}
	return
}



