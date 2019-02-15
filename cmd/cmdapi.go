package cmd

import (
	"os"
	"log"
	"io/ioutil"
	"syscall"
	"strconv"	
	"os/exec"
	"do/cluster"
	yaml "gopkg.in/yaml.v2"
)

// YamlClusterConfig is struct of cluster config
type YamlClusterConfig struct {
	ClusterNumber   string   `yaml:"clusterNumber"`
	MachineNamelist []string `yaml:"machineNamelist"`
}

// YamlClusterStatus is struct of cluster status
type YamlClusterStatus struct{
	MasterPID    string   `yaml:"masterpid"`
	NodePID       []string `yaml:"nodepid"`
	NodeNumber    string   `yaml:"nodenumber"`
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
	status.NodePID = make([]string,clusterNumber,clusterNumber)

	// start Master process and start Nodes process ,finally save status to file
	var Args = []string{" of Do Project"}
	cmd := exec.Command(cluster.MasterExePosition, Args...)
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
		log.Printf("Node Deamon(PID:%s)  Stopping......", status.NodePID[i])
		Pid,_:= strconv.Atoi(status.NodePID[i])
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

	err = ioutil.WriteFile(cluster.StatusPosition, buf, 0x777)
	if err != nil {
		panic(err.Error())
	}
	return
}

func getClusterStatus(status *YamlClusterStatus) {

	buf, err := ioutil.ReadFile(cluster.StatusPosition)
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

		port := strconv.Itoa(cluster.DefaultNodePort + i)
		var Args = []string{config.MachineNamelist[i], port}

		cmd := exec.Command(cluster.NodeExePosition, Args...) 
		cmd.Stdin = os.Stdin                 
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() 

		status.NodePID[i] = strconv.Itoa(cmd.Process.Pid)
		log.Printf("Node Deamon(PID:%s)  Starting......", status.NodePID[i])
	}
	return
}



