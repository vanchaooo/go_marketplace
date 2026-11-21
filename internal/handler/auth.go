package handler

import (
	"net/http"

	"github.com/vanchaooo/go-marketplace/internal/service"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

type registerRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type tokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	user, err := a.auth.Register(r.Context(), request.Name, request.Email, request.Password)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}
	_ = writeJSON(w, http.StatusCreated, userToResponse(user))
}	

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	pair, err := a.auth.Login(r.Context(), request.Email, request.Password)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}
	_ = writeJSON(w, http.StatusOK, tokensResponse{pair.Access, pair.Refresh})
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var request refreshTokenRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	if err := a.auth.Logout(r.Context(), request.RefreshToken); err != nil {
		handleServiceErrors(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var request tokensResponse
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	pair, err := a.auth.Refresh(r.Context(), request.RefreshToken)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, tokensResponse{pair.Access, pair.Refresh})
}