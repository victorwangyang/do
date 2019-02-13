package cluster

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"net/http"
)

//InitNodes to init Nodesinfo struct
func InitNodes(len int) {

	Nodes = make([]NodesInfo, len, len)
}

//StartNode is to start node process by name and port
func StartNode(Index int, NodeName string, Port string) (PID string) {

	var Args = []string{NodeName, Port}

	cmd := exec.Command(NodeExePosition, Args...) 
	cmd.Stdin = os.Stdin                 
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start() 

	Nodes[Index].NodeName = NodeName
	Nodes[Index].PortNum = Port
	Nodes[Index].NodePID = strconv.Itoa(cmd.Process.Pid)
	Nodes[Index].NodeStatus = true
	fmt.Println("start....", NodeName, Port, Nodes[Index].NodePID)

	return
}

// KillNode is to kill a master process
func KillNode(Index int) {

	Pid, _ := strconv.Atoi(Nodes[Index].NodePID)

	if Pid != 0 {
		syscall.Kill(int(Pid), syscall.SIGKILL)
		Nodes[Index].NodeName = ""
		Nodes[Index].PortNum = ""
		Nodes[Index].NodeStatus = false
	}
	return
}

//NodeProcessRequestFromMaster is to Process request from Client
func NodeProcessRequestFromMaster(r *http.Request) {

	log.Println("NodeProcessRequestFromMaster .......")
}
