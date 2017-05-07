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

func (r *Router) SetStaticPath(dir string) {
	r.StaticHandler = r.NewRoute("GET", "", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/") {
			http.NotFound(w, req)
			return
		}
		http.ServeFile(w, req, dir+req.URL.Path[1:])
	})
}

func (r *Router) find(req *http.Request) *Route {
	for _, route := range r.routes {
		if route.match(req) {
			return route
		}
	}

	if r.StaticHandler == nil {
		return r.NotFound
	}

	return r.StaticHandler
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}
