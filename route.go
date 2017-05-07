package NanoRouter

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	static        string
	routes        []*Route
	NotFound      *Route
	StaticHandler *Route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ROUTE: " + req.URL.String())
	route := r.find(req)
	route.Handle(w, req)
}

func NewRouter() *Router {
	r := &Router{}
	r.SetNotFoundRoute(notFound)
	return r
}

func (r *Router) NewRoute(httpMethod string, route string, f func(http.ResponseWriter, *http.Request, map[string]string)) *Route {
	nr := &Route{method: httpMethod, Path: route, handler: f}
	r.routes = append(r.routes, nr)
	return nr
}

func (r *Router) SetNotFoundRoute(f func(http.ResponseWriter, *http.Request, map[string]string)) {
	r.NotFound = &Route{Path: "", handler: f}
}

func (r *Router) SetFaviconRoute(f func(http.ResponseWriter, *http.Request, map[string]string)) {
	r.NotFound = &Route{Path: "/favicon.ico", handler: f}
}

func (r *Router) SetStaticPath(dir string) {
	r.StaticHandler = r.NewRoute("GET", "", func(w http.ResponseWriter, req *http.Request, vars map[string]string) {
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

func notFound(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	fmt.Fprintf(w, "not found")
}
