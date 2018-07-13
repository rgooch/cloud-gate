package httpd

import (
	"io"
	"net/http"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/static/common.css":
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, commonCSS)
		return

	case "/static/customization.css":
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, customizationCSS)
		return
	}
	http.Error(w, "error not found", http.StatusNotFound)
}

func (s *Server) consoleAccessHandler(w http.ResponseWriter, r *http.Request) {
	displayData := consolePageTemplateData{
		Title: "Cloud-Gate console access",
	}
	err := s.htmlTemplate.ExecuteTemplate(w, "consoleAccessPage", displayData)
	if err != nil {
		s.logger.Printf("Failed to execute %v", err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	return
}