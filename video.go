package main

import (
	"fmt"
	"net/http"
)

type VideoHandler struct {
}

func (h VideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><video><source src=\"localhost:8099/test1.mpeg\" type=\"video/webm\" />Wah Wah</video></body></html>")

}
