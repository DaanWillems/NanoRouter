package NanoRouter

import (
	"fmt"
	"net/http"
	"strings"
)

var Vars map[string]string

type Router struct {
	static        string
	routes        []*Route
	NotFound      *Route
	StaticHandler *Route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ROUTE: " + req.URL.String())
	route := r.find(req)
	route.parseVars(req)
	route.Handle(w, req)
}

func NewRouter() *Router {
	r := &Router{}
	r.SetNotFoundRoute(notFound)
	return r
}

func (r *Router) NewRoute(httpMethod string, route string, f func(http.ResponseWriter, *http.Request)) *Route {
	nr := &Route{method: httpMethod, Path: route, handler: http.HandlerFunc(f)}
	r.routes = append(r.routes, nr)
	return nr
}

func (r *Router) SetNotFoundRoute(f func(http.ResponseWriter, *http.Request)) {
	r.NotFound = &Route{Path: "", handler: http.HandlerFunc(f)}
}

func (r *Router) SetFaviconRoute(f func(http.ResponseWriter, *http.Request)) {
	r.NotFound = &Route{Path: "/favicon.ico", handler: http.HandlerFunc(f)}
}

func (r *Router) SetStaticPath(path string, dir string) {
	r.StaticHandler = r.NewRoute("GET", path, func(w http.ResponseWriter, req *http.Request) {
		http.StripPrefix(path, http.FileServer(http.Dir(dir)))
		http.ServeFile(w, req, req.URL.Path[1:])
	})
}

func (r *Router) find(req *http.Request) *Route {
	url := strings.Split(req.URL.String(), "/")
	if r.StaticHandler != nil {
		path := strings.Split(r.StaticHandler.Path, "/")
		if url[1] == path[1] {
			return r.StaticHandler
		}
	}

	for _, route := range r.routes {
		if route.match(req) {
			return route
		}
	}
	return r.NotFound
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}
