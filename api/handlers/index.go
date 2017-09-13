package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index router")
}
