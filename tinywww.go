package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	myname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hi BCIT from %s!", myname)
}

func main() {
	myport := ":8000"
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", myport)
	http.ListenAndServe(myport, nil)
}
