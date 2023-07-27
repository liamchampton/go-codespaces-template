package main

import (
	"log"
	"net/http"
)

func main() {
	
	// read in the static file directory
	fs := http.FileServer(http.Dir("static"))

	// main home route handler
	http.Handle("/", fs)

	// start http server
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on :3000...")
}
