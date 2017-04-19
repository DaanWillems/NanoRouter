package NanoRouter

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//Stores the path and the handler
type Route struct {
	method  string
	Path    string
	handler http.Handler
}

//Matches the url against its own path, returns true when there is a match
func (r *Route) match(req *http.Request) bool {
	if !r.matchURL(req.URL.String()) {
		return false
	}
	//See if http method matches the route
	if req.Method != r.method {
		return false
	}

	return true
}

func (r *Route) matchURL(rawURL string) bool {
	url := strings.Split(rawURL, "/")
	path := strings.Split(r.Path, "/")
	reg, _ := regexp.Compile(":[a-zA-Z0-9]")

	for i, c := range url {
		if len(url) != len(path) {
			return false
		}

		if c != path[i] && !(reg.MatchString(path[i])) {
			return false
		}
	}
	return true
}

//Handle executes the handler function attached to this route
func (r *Route) Handle(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v", r.handler)
	r.handler.ServeHTTP(w, req)
}

func (r *Route) parseVars(req *http.Request) {
	Vars = make(map[string]string)
	for k, v := range req.URL.Query() {
		Vars[k] = v[0]
	}
	if req.URL.String() != "/" {
		return
	}
	reg, _ := regexp.Compile(":[a-zA-Z0-9]")
	url := strings.Split(req.URL.String(), "/")
	path := strings.Split(r.Path, "/")

	for i, p := range path {
		if reg.MatchString(p) {
			Vars[p[1:len(p)]] = url[i]
		}
	}
}
