package main

import "fmt"
import "net/http"
import "net/http/httputil"


func handler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("REQUEST:%s", r.URL.Path))
	proxy := httputil.NewSingleHostReverseProxy(r.URL)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", nil)
}