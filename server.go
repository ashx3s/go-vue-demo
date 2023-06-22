package main

import "fmt"

func main() {
	http.HandleFunc("/", homePageHandler)	
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}