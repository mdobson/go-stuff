package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"success": true, "message": fmt.Sprintf("data: %s", r.URL.Path[1:])})
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", nil)
}