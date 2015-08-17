package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PageData struct{}

func main() {
	http.HandleFunc("/", indexResponse)
	http.Handle("/blog", BlogHandler{
		loadedArticle: NewArticle("Making web services in Go.", "Guy Moore", "Much fun."),
	})
	http.Handle("/video", VideoHandler{})
	/*
		TODO: Get DB authentication done.
		http.Handle("/uptime", AuthHandler{
			next: NewUptimeHandler(),
			auth: "fish",
		})
	*/
	http.HandleFunc("/assets/", assetResponse)
	http.HandleFunc("/info/", infoPageResponse)
	http.ListenAndServe(":8080", nil)
}

func indexResponse(w http.ResponseWriter, r *http.Request) {
	// Testing out adding header data.
	w.Header().Add("Server", "Go Server")
	w.Header().Add("Auth-token", "1")
	fmt.Fprintf(w, "Hello Grace!")
}

func assetResponse(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/assets/"):]
	asset, err := ioutil.ReadFile("/Users/jeannichols/Git/gocode/src/website/Assets/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(asset)
	if err != nil {
		log.Fatal(err)
	}
}

func infoPageResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: Make an interface to hold data for the info pages.
	data := PageData{}
	pageTemplateName := r.URL.Path[len("/info/"):]
	templates := LoadTemplate(pageTemplateName)
	err := templates.ExecuteTemplate(w, "Header", data)
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, pageTemplateName, data)
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, "Footer", data)
	if err != nil {
		log.Fatal(err)
	}
}
