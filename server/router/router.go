package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const indexFile = "index.html"

type Router struct {
	Mux *http.ServeMux
}

/*
* Fonction de Setup du router pour le site web
 */
func Setup(mux *http.ServeMux) *Router {
	router := newRouter(mux)

	router.Mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./client-app/dist/assets"))),
	)
	router.Mux.HandleFunc("/", router.routeHome)

	return router
}

func newRouter(mux *http.ServeMux) *Router {
	return &Router{
		Mux: mux,
	}
}

func (Router) RenderTemplate(fileName string) (*template.Template, error) {
	t, err := template.ParseFiles(fmt.Sprintf("./client-app/dist/%s", fileName))

	if t == nil {
		return nil, errHTMLNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "Unknown error")
	}

	return t, nil
}

func (rt Router) routeHome(w http.ResponseWriter, r *http.Request) {
	t, err := rt.RenderTemplate(indexFile)
	if err != nil {
		if err == errHTMLNotFound {
			log.Printf("%s HTML FILE NOT FOUND\n", indexFile)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("An error occurred while sending HTML file: %s \n", err)
	}
}
