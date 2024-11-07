package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandruchiku/go-ecom/service/auth"
	"github.com/chandruchiku/go-ecom/types"
	"github.com/chandruchiku/go-ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("login")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}
	// Check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.RespondError(w, http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	// Create user
	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}
