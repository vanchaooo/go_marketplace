package handler

import (
	"log"
	"errors"
	"strconv"
	"net/http"
	"github.com/vanchaooo/go-marketplace/internal/service"
)

func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

func handleServiceErrors(w http.ResponseWriter, err error) {
	switch {
		case
			errors.Is(err, service.ErrInvalidID),
			errors.Is(err, service.ErrInvalidName),
			errors.Is(err, service.ErrInvalidPrice),
			errors.Is(err, service.ErrInvalidStock):
			writeError(w, http.StatusBadRequest, err.Error())
	
		case
			errors.Is(err, service.ErrUserNotFound),
			errors.Is(err, service.ErrProductNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case 
			errors.Is(err, service.ErrInvalidCredentials),
			errors.Is(err, service.ErrInvalidToken):
			writeError(w, http.StatusUnauthorized, err.Error())

		case 
			errors.Is(err, service.ErrEmailTaken):
			writeError(w, http.StatusConflict, err.Error())

		case 
			errors.Is(err, service.ErrInvalidEmail),
			errors.Is(err, service.ErrInvalidPassword):
			writeError(w, http.StatusBadRequest, err.Error())
		
		default:
			log.Println("internal error:", err)
			writeError(w, http.StatusInternalServerError, "внутренняя ошибка")
	}
}