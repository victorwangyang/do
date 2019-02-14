package restapi

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//InitRestSvr is to prepare resource of path/func and start to listen at the addr
func InitRestSvr(path string, port string, f func(http.ResponseWriter,
	*http.Request)) {

	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc(path, f)

	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + port, r))
}
