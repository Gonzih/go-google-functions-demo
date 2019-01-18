package function

import (
	"net/http"
)

var mux = newMux()

//F represents cloud function entry point
func F(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/one", one)
	mux.HandleFunc("/two", two)
	mux.HandleFunc("/subroute/three", three)

	return mux
}

func one(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from one"))
}

func two(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from two"))
}

func three(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from three"))
}
