package game

import (
	"encoding/json"
	"fmt"
	"highscore/internal/auth"
	"highscore/internal/types"
	"highscore/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.GameStore
	userStore types.UserStore
}

func NewHandler(store types.GameStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/games", auth.WithJWTAuth(utils.Handler(h.GetAll), h.userStore, types.UserRole)).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}", auth.WithJWTAuth(utils.Handler(h.GetById), h.userStore, types.UserRole)).Methods("GET")
	router.HandleFunc("/games", auth.WithJWTAuth(utils.Handler(h.Insert), h.userStore, types.AdminRole))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	games, err := h.store.GetAll()
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Successfully got the list of games",
		Data:    games,
	}

	utils.WriteJson(w, http.StatusOK, res)
	return nil
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	game, err := h.store.GetById(id)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Successfully got the game",
		Data:    game,
	}

	utils.WriteJson(w, http.StatusOK, res)
	return nil
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) error {
	var game types.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		return fmt.Errorf("failed to parse the data")
	}

	err = h.store.Insert(game)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Game created successfully",
	}

	utils.WriteJson(w, http.StatusCreated, res)
	return nil
}
