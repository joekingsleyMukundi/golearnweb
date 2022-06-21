package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/joekingsleyMukundi/golearnweb/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// render function we created now we can render the home page
	render.RenderTemplate(w, "home.page.tmpl")
}

// About  is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	var sum int
	sum = AddValues(6, 5)

	fmt.Fprintf(w, fmt.Sprintf("this is the about page and the value is %d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, error := divideValues(10.0, 2.0)
	if error != nil {
		fmt.Fprintf(w, "cannot divide with 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("result fafterdividing %f and %f  is %f", 10.0, 2.0, result))
}

func divideValues(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot  divide by 0")
	}
	return x / y, nil
}
