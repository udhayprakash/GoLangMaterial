package main

import (
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("public/tmpl/*.html"))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./public/img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.Handle("/icn/", http.StripPrefix("/icn/", http.FileServer(http.Dir("./public/icn"))))
	http.Handle("/doc/", http.StripPrefix("/doc/", http.FileServer(http.Dir("./public/doc"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./public/download"))))
	http.Handle("/mov/", http.StripPrefix("/mov/", http.FileServer(http.Dir("./public/mov"))))
	http.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir("./.well-known/"))))
}

func main() {
	http.HandleFunc("/", endpoint)
	go http.ListenAndServeTLS(":443", "certificate.crt", "private.key", nil)
	err := http.ListenAndServe(":80", http.HandlerFunc(rehttps))

	if err != nil {
		log(err.Error())
	}
}

func rehttps(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	page := (path + ".html")

	switch path {
	case "contactme":
		contactme(w, r)
	case "": //empty
		url := geturl(path)
		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "max-age=3600")
		tpl.ExecuteTemplate(w, "home.html", url)
	default:
		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "max-age=3600")
		exist := tpl.Lookup(page)
		if exist == nil {
			redirect(w, r) // if not found -> redirect.go
		}
		url := geturl(path)
		tpl.ExecuteTemplate(w, page, url)
	}
}

func geturl(path string) string {
	//used for canonical path on each page
	switch path {
	case "":
		return ("https://gowebstatic.tk/")
	default:
		return ("https://gowebstatic.tk/" + path)
	}
}
