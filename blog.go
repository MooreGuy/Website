package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type BlogHandler struct {
	loadedArticle Article
}

type Article struct {
	Title         string
	PublishedTime time.Time
	Author        string
	Body          string
	Draft         bool
}

func NewArticle(title string, author string, body string) Article {
	article := Article{
		Title:         title,
		PublishedTime: time.Now(),
		Author:        author,
		Body:          body,
	}

	return article
}

func (h BlogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates := LoadTemplate("BlogArticle")
	err := templates.ExecuteTemplate(w, "Header", h.loadedArticle)
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, "BlogArticle", h.loadedArticle)
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, "Footer", h.loadedArticle)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadTemplate(name string) *template.Template {
	templates := template.Must(template.ParseFiles("templates/header.html", "templates/"+name+".html", "templates/footer.html", "templates/css.html", "templates/js.html"))
	return templates
}
