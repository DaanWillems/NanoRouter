package NanoRouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test handler
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "succes")
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

	r.NewRoute("GET", "/test/:id", test)

	req, err := http.NewRequest("GET", "/test/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "1"

	if Vars["id"] != expected {
		t.Errorf("handler returned unexpected variable: got %v want %v", Vars["id"], expected)
	}
}

func TestSingleUrlVar(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/:id", test)

	req, err := http.NewRequest("GET", "/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expected := "1"

	if Vars["id"] != expected {
		t.Errorf("handler returned unexpected variable: got %v want %v", Vars["id"], expected)
	}
}

func TestTwoUrlVar(t *testing.T) {
	rr := httptest.NewRecorder()

	r := NewRouter()

	r.NewRoute("GET", "/varone/:id/vartwo/:name", test)

	req, err := http.NewRequest("GET", "/varone/1/vartwo/name", nil)

	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	expectedVarOne := "1"
	expectedVarTwo := "name"

	if Vars["id"] != expectedVarOne {
		t.Errorf("handler returned unexpected variable: got %v want %v", Vars["id"], expectedVarOne)
	}

	if Vars["name"] != expectedVarTwo {
		t.Errorf("handler returned unexpected variable: got %v want %v", Vars["id"], expectedVarTwo)
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
