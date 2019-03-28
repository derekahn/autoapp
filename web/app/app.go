package app

import (
	"html/template"
	"net/http"
)

type (
	// Server encapsulates all required configurations
	Server struct {
		Router *http.ServeMux
		*Welcome
	}
	// Welcome holds information to be displayed in our HTML file
	Welcome struct {
		Name, Time string
	}
)

const (
	indexRoute   = "/"
	staticRoute  = "/static/"
	faviconRoute = "/favicon.ico"
)

// Routes acts as the root mounting point of all endpoints
func (s *Server) Routes(basePath string) {
	s.Router.HandleFunc(indexRoute, s.handleIndex(basePath))
	s.Router.HandleFunc(staticRoute, s.handleStatic(basePath))
	s.Router.HandleFunc(faviconRoute, s.handleFavicon(basePath))
}

func (s *Server) handleStatic(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, basePath+r.RequestURI)
	}
}

func (s *Server) handleIndex(basePath string) http.HandlerFunc {
	templates := template.Must(template.ParseFiles(basePath + "/template/welcome.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			s.Welcome.Name = name
		}
		templates.ExecuteTemplate(w, "welcome.html", s.Welcome)
	}
}

func (s *Server) handleFavicon(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, basePath+"/static/asset/favicon.ico")
	}
}
