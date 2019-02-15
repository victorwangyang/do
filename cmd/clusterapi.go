package cmd

import (
	"net/http"
	"io/ioutil"
	"log"
)

// constant of cluster
const (
	MasterAddr       = "http://localhost"
	MasterPort       = "8085"
	ClusterGet       = "/api/v1/cluster/get"

	NodeAddr       = "http://localhost"
	NodeDirectory = "/node"
)
// GetCluster is to Send request to Master for Getting cluster information
func GetCluster() {

	ret, err := http.Get(MasterAddr + ":" + MasterPort + ClusterGet)
	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))

	return

}
