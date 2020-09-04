package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		funcaoMaluca()
		fmt.Fprintf(w, "Hello!!!")
	})
	http.ListenAndServe(":8080", nil)
}

func funcaoMaluca() {
	x := 0.0001
	for i := 0; i < 100000; i++ {
		x += math.Sqrt(x)
	}
}
