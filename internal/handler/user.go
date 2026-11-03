package handler

import (
	"net/http"
	"time"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/service"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(users *service.UserService) *UserHandler {
	return &UserHandler{users: users}
}

type createUserRequest struct {
	Name string `json:"name"`
}

type renameUserRequest struct {
	Name string `json:"name"`
}

type userResponse struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func userToResponse(user model.User) userResponse {
	return userResponse{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	user, err := u.users.AddUser(r.Context(), request.Name)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusCreated, userToResponse(user))
}

func (u *UserHandler) RenameUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	var request renameUserRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	user, err := u.users.RenameUser(r.Context(), id, request.Name)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, userToResponse(user))
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	if err := u.users.DeleteUser(r.Context(), id); err != nil {
		handleServiceErrors(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	user, err := u.users.GetUser(r.Context(), id)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, userToResponse(user))
}