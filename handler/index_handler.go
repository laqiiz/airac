package handler

import (
	"html/template"
	"net/http"
)

type IndexHandler struct {
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("view/index.tpl"))

	if err := tpl.Execute(w, map[string]string{}); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
}
