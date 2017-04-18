package nano

import (
	"fmt"
	"net/http"
)

//Stores the path and the handler
type Route struct {
	path    string
	handler http.Handler
}

//Matches the url against its own path, returns true when there is a match
func (r *Route) match(req *http.Request) bool {
	fmt.Println(r.path + "" + req.URL.Path)
	if r.path == req.URL.Path {
		return true
	}
	return false
}

//Executes the handler function attached to this route
func (r *Route) Handle(w http.ResponseWriter, req *http.Request) {
	r.handler.ServeHTTP(w, req)
}
