package cluster

import (
	// "log"
	// "io/ioutil"
	 "net/http"
	// "strconv"	
	// yaml "gopkg.in/yaml.v2"
)

//MasterProcessRequestFromClient is to Process request from Client
func MasterProcessRequestFromClient(r *http.Request) {
/*
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
	*/
}
