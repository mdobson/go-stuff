package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
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
	if r.Method == "POST" {
		buf, err := ioutil.ReadAll(r.Body)
		if err == nil {
			jsonStr := string(buf)
			res := &Response{}
			json.Unmarshal([]byte(jsonStr), &res)
			fmt.Fprint(w, Response{"success": true, "data": res})
		} else {
			fmt.Fprint(w, Response{"success": false, "data": "Error parsing body"})
		}
	} else {
		fmt.Fprint(w, Response{"success": true, "message": fmt.Sprintf("%s", r.URL.Path[1:])})
	}
	
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080", nil)
}