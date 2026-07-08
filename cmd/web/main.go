package main

import (
	// log is used for startup and fatal error messages from the web server.
	"log"

	// net/http provides Go's standard HTTP server, request, and response types.
	"net/http"
)

// port is the TCP port that the web server listens on.
// It is kept as a string because http.ListenAndServe expects addresses in the
// form ":4000", which can be built directly from this value.
const port string = "4000"

func main() {
	// Print a startup message so it is clear which port the process is using.
	log.Printf("server listening on port %s", port)

	// Start the HTTP server on the configured port.
	// The leading colon means "listen on all available network interfaces" for
	// this port. routes() returns the chi router that handles incoming requests.
	if err := http.ListenAndServe(":"+port, routes()); err != nil {
		// ListenAndServe returns an error if the server cannot start or stops
		// unexpectedly. Fatal logs the error and exits the program.
		log.Fatal(err)
	}
}
