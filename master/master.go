package main

import (
	"do/cmd"
	"fmt"
	"net/http"
	"do/cluster"
	"do/restapi"
	"encoding/json"
)

func apiV1ClusterGet(w http.ResponseWriter, r *http.Request) {

	buf,_ := json.MarshalIndent(cluster.GCluster,"","")
	fmt.Fprintf(w, string(buf))
}

func main() {

	cluster.InitClusterInfo(cmd.StatusPosition)


	restapi.InitRestSvr(cmd.ClusterGet, 
						cmd.MasterPort,
						apiV1ClusterGet)

}
