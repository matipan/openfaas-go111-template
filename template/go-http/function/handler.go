package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello. You said: %s", string(b))
}
