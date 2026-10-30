package handler

import "net/http"

func Health(w http.ResponseWriter, _ *http.Request) {
	_ = writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}