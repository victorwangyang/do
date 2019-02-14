package main

import (
	"os"
	"log"
	"net/http"
	"do/cluster"
	"do/restapi"
)


func nodeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("GET.....from Master")
	case "POST":
		{
			cluster.NodeProcessRequestFromMaster(r)
		}
	default:
		log.Println("DEFAULT.......")

	}

}


func main() {

	var NodeName = os.Args[1]
	var PortNum = os.Args[2]

	log.Println(" Listening.....", NodeName, PortNum)
	
	restapi.InitRestSvr(cluster.NodeDirectory, PortNum, nodeHandler)


}
