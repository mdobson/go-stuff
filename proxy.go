package main

import "fmt"
import "net/http"
import "net/http/httputil"



// func main(){
//     proxy := httputil.NewSingleHostReverseProxy( &url.URL{Scheme:"http",Host:"www.google.com",Path:"/"})

//     http.ListenAndServe(":8080", proxy)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	//proxy_url := r.URL.Path[1:]
	//&url.URL{Scheme:"http",Host:proxy_url,Path:"/"}
	println(fmt.Sprintf("REQUEST:%s", r.URL.Path))
	proxy := httputil.NewSingleHostReverseProxy(r.URL)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", nil)
}