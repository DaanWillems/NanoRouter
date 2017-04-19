### Usage
---

```golang
func main() {
	router := nano.NewRouter()
	router.NewRoute("GET", "/get/:example", get)
	router.NewRoute("POST", "/post", post)
	router.SetNotFoundRoute(notfound)
	http.Handle("/", router)
	http.ListenAndServe(":8380", nil)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unknown page")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, nano.Vars["test"]+"get")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post")
}
```
