package main

import ("fmt"
"net/http")

func main() {

	// root directory
	http.HandleFunc("/", handler)
	// root/earth directory
	http.HandleFunc("/earth", handler2)

	// serve at port 8080
	http.ListenAndServe(":8080", nil)
}


// handle this function
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Earth\n")
}