package function

import "net/http"

func F(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.HandleFunc("/function-1", one)
	router.ServeHTTP(w, r)
}

func one(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from one"))
}
