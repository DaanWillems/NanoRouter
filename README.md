### Usage example
---

```golang
import "NanoRouter"

func main() {
	router := NanoRouter.NewRouter()
	router.NewRoute("GET", "/get/:example", getHandler)
	router.NewRoute("POST", "/post", postHandler)
	router.SetNotFoundRoute(notfoundHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8380", nil)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unknown page")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, NanoRouter.Vars["test"])
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post")
}
```

### Not found
When an url is not found NanoRouter automatically returns the string 'not found'. This behaviour can be overridden with:
```golang
	router.SetNotFoundRoute(handler)
```
