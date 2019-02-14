package cluster

import (
	"do/restapi"
	"io/ioutil"
	"log"
)

// CreateCluster is to Send request to Master for starting cluster
func CreateCluster(fileName string) (err error) {

	clusteryaml, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		return err
	}
	restapi.Post(MasterAddr+":"+MasterPort+ClusterDirectory, clusteryaml)
	return nil

}

// DeleteCluster is to Send request to Master for Deleting cluster
func DeleteCluster() {


	return

}

// GetCluster is to Send request to Master for Getting cluster information
func GetCluster() {


	return

}
