package main

import "fmt"
import "net/http"
import "net/url"
import "net/http/httputil"


func handler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("REQUEST:%s", r.URL.Path))
	url, err := url.Parse("http://127.0.0.1:1337/")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", nil)
}