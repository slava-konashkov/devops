package main

import (
	"fmt"
	"net/http"
)

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

func main() {
	// http.HandleFunc("/", helloWorld)
	const port = "80"
	fmt.Println("Server starting on port " + port + "...")
	srv := http.FileServer(http.Dir("./html"))
	if err := http.ListenAndServe(":"+port, srv); err != nil {
		panic(err)
	}
}
