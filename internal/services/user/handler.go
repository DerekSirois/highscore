package user

import (
	"encoding/json"
	"fmt"
	"highscore/internal/auth"
	"highscore/internal/types"
	"highscore/internal/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {

	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", utils.Handler(h.HandleLogin)).Methods("POST")
	router.HandleFunc("/register", utils.Handler(h.HandleRegister)).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	var user types.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to parse the data")
	}

	userDb, err := h.store.GetByEmail(user.Email)
	if err != nil {
		return err
	}

	log.Println(user.Password, userDb.Password)

	if !auth.CheckPasswordHash(user.Password, userDb.Password) {
		return fmt.Errorf("wrong username/password")
	}

	// START TEMP
	token := "asdf"
	// END TEMP

	res := utils.JsonResponse{
		Message: "You logged in successfully",
		Token:   token,
	}

	utils.WriteJson(w, http.StatusOK, res)
	return nil
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	var user types.UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to parse the data")
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash the password")
	}

	err = h.store.Insert(user)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "You registered successfully",
	}

	utils.WriteJson(w, http.StatusCreated, res)
	return nil
}
