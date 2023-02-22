package main

import (
	"fmt"
	"net/http"
)

// ハンドラ
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}
