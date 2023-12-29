package handler

import (
	"net/http"

	"github.com/tomazcx/go-chat-app/internal/application/templates"
)

type ChatHandler struct {}

func (h *ChatHandler) Index(w http.ResponseWriter, r *http.Request){
	templates.Index().Render(r.Context(), w)
}
