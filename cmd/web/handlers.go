package main

import (
	"fmt"
	"net/http"
)

// home handles GET / requests.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello go")
}
