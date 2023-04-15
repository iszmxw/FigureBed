package api

import (
	"fmt"
	"net/http"
)

func Tool(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "tools")
	if err != nil {
		return
	}
}
