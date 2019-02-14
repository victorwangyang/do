package cluster

import (
	"os"
	"log"
	"io/ioutil"
	"syscall"
	"net/http"
	"strconv"	
	"os/exec"
	yaml "gopkg.in/yaml.v2"
)

// StartMasterDeamon is to exec a master process
func StartMasterDeamon() {

	//if parent proccess exits,system No.1(PID=1) will take over child process
	if os.Getppid() != 1 { // is child process
		cmd := exec.Command(MasterExePosition, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		saveMasterPID(cmd.Process.Pid)
		log.Println("Master Deamon is starting......")
		return
	}
}

// KillMasterDeamon is to kill a master process
func KillMasterDeamon() {

	Pid := getMasterPID()
	if Pid != 0 {
		syscall.Kill(int(Pid), syscall.SIGKILL)
		log.Println("Master Deamon is stopping......")
	}
	return
}

//MasterProcessRequestFromClient is to Process request from Client
func MasterProcessRequestFromClient(r *http.Request) {

	conf := new(Yamlcluster)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(body, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	machinNumber, _ := strconv.Atoi(conf.ClusterNumber)

	InitNodes(machinNumber)

	for i := 0; i < machinNumber; i++ {
		port := strconv.Itoa(InitNodePort + i)
		StartNode(i, conf.MachineNamelist[i], port)
	}
}

func saveMasterPID(PID int) {

	Pid := strconv.Itoa(PID)
	buf := []byte(Pid)

	err := ioutil.WriteFile(MasterIniPosition, buf, 0x777)
	if err != nil {
		panic(err.Error())
	}
	return
}

func getMasterPID() (PID int) {

	buf, err := ioutil.ReadFile(MasterIniPosition)
	if err != nil {
		panic(err.Error())
	}

	Pid := string(buf)
	MasterPid, err := strconv.Atoi(Pid)
	if err != nil {
		panic(err.Error())
	}
	return MasterPid
}
