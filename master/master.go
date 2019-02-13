package main

import (
	"log"
	"net/http"
	"goexample/cluster"
	"goexample/restapi"
)

func masterHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("GET.....from Cli")
	case "POST":
		{
			cluster.MasterProcessRequestFromClient(r)
		}
	default:
		log.Println("DEFAULT.......")

	} 

}

func main() {

	restapi.InitRestSvr(cluster.ClusterDirectory, 
						cluster.MasterPort,
						 masterHandler)

}
