package NanoRouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test handler
func test(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	fmt.Fprintf(w, "succes")
}

func testvar(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	s := ""
	for k, v := range vars {
		s += k + ":" + v
	}
	fmt.Fprintf(w, s)
}

func TestEmptyUrl(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "", test)

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "succes"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestBasicUrl(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/", test)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "succes"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestOnePartBasicUrl(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/test", test)

	req, err := http.NewRequest("GET", "/test", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "succes"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTwoPartBasicUrl(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/test/test2", test)

	req, err := http.NewRequest("GET", "/test/test2", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "succes"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/dummy", test)

	req, err := http.NewRequest("GET", "/notfound", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "not found"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUrlVar(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/test/:id", testvar)

	req, err := http.NewRequest("GET", "/test/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "id:1"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected variable: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSingleUrlVar(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/:id", testvar)

	req, err := http.NewRequest("GET", "/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "id:1"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected variable: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTwoUrlVar(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/varone/:id/vartwo/:name", testvar)

	req, err := http.NewRequest("GET", "/varone/1/vartwo/name", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "id:1name:name"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected variable: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCustomNotfound(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.SetNotFoundRoute(test)

	req, err := http.NewRequest("GET", "/unknown", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "succes"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", Vars["id"], expected)
	}
}
