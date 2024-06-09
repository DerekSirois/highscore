package score

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
	store     types.ScoreStore
	userStore types.UserStore
}

func NewHandler(store types.ScoreStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/score/{gameId:[0-9]+}", auth.WithJWTAuth(utils.Handler(h.GetLeaderboard), h.userStore, types.UserRole)).Methods("GET")
	router.HandleFunc("/score/pending", auth.WithJWTAuth(utils.Handler(h.GetScoresPendingApproval), h.userStore, types.ModeratorRole)).Methods("GET")
	router.HandleFunc("/score", auth.WithJWTAuth(utils.Handler(h.SubmitScore), h.userStore, types.UserRole)).Methods("POST")
	router.HandleFunc("/score/approve/{scoreId:[0-9]+}", auth.WithJWTAuth(utils.Handler(h.ApproveScore), h.userStore, types.ModeratorRole)).Methods("PUT")
}

func (h *Handler) GetLeaderboard(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		return fmt.Errorf("invalid game id")
	}

	scores, err := h.store.GetAllApprovedScoresByGame(gameId)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Successfully got the leaderboard",
		Data:    scores,
	}

	utils.WriteJson(w, http.StatusOK, res)
	return nil
}

func (h *Handler) GetScoresPendingApproval(w http.ResponseWriter, r *http.Request) error {
	scores, err := h.store.GetAllScoresPendingApproval()
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Successfully got the scores",
		Data:    scores,
	}

	utils.WriteJson(w, http.StatusOK, res)
	return nil
}

func (h *Handler) SubmitScore(w http.ResponseWriter, r *http.Request) error {
	var score types.ScoreSubmit
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		return fmt.Errorf("failed to parse the data")
	}

	playerId := auth.GetUserIDFromContext(r.Context())

	err = h.store.Insert(playerId, score.GameId, score.Score)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Score submitted successfully",
	}

	utils.WriteJson(w, http.StatusCreated, res)
	return nil
}

func (h *Handler) ApproveScore(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	scoreId, err := strconv.Atoi(vars["scoreId"])
	if err != nil {
		return fmt.Errorf("invalid score id")
	}

	approverId := auth.GetUserIDFromContext(r.Context())

	err = h.store.Approve(approverId, scoreId)
	if err != nil {
		return err
	}

	res := utils.JsonResponse{
		Message: "Score approved successfully",
	}

	utils.WriteJson(w, http.StatusCreated, res)
	return nil
}
