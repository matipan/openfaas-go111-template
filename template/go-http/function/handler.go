package function

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	return fmt.Fprintf(w, "Hello. Your UUID is: %s", uuid.New().String())
}
