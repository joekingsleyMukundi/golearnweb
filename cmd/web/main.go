package main

import (
	"fmt"
	"net/http"

	"github.com/joekingsleyMukundi/golearnweb/pkg/handlers"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// n is the number of bytes written
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Println((err))
		}

		fmt.Println(fmt.Sprintf("number of bytes written : %d", n))
	})

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)

	fmt.Println(fmt.Sprintf("Start application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
