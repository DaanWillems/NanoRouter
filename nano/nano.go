package nano

import (
	"net/http"
)

var Vars map[string]string

type Router struct {
	routes   []*Route
	NotFound *Route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := r.find(req)
	route.parseVars(req)
	route.Handle(w, req)
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) NewRoute(httpMethod string, route string, f func(http.ResponseWriter, *http.Request)) {
	r.routes = append(r.routes, &Route{method: httpMethod, path: route, handler: http.HandlerFunc(f)})
}

func (r *Router) SetNotFoundRoute(f func(http.ResponseWriter, *http.Request)) {
	r.NotFound = &Route{path: "/pagenotfound", handler: http.HandlerFunc(f)}
}

func (r *Router) find(req *http.Request) *Route {
	for _, route := range r.routes {
		if route.match(req) {
			return route
		}
	}
	return r.NotFound
}
