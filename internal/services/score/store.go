package score

import (
	"database/sql"
	"fmt"
	"highscore/internal/types"
	"time"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllApprovedScoresByGame(gameId int) ([]types.Score, error) {
	var scores []types.Score
	rows, err := s.db.Query(`SELECT s.id, u1.username, g.name, u2.username, s.score, s.createdat, s.approvedat
							FROM score s
							JOIN game g ON s.gameid = g.id
							JOIN users u1 ON s.playerid = u1.id
							JOIN users u2 ON s.approverid = u2.id
							WHERE s.approvedat IS NOT NULL`)
	if err != nil {
		return nil, fmt.Errorf("error while getting the scores: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var score types.Score
		if err := rows.Scan(&score.Id, &score.Player, &score.Game, &score.Approver, &score.Score, &score.CreatedAt, &score.ApprovedAt); err != nil {
			return nil, fmt.Errorf("error while getting the scores: %v", err)
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func (s *Store) GetAllScoresPendingApproval() ([]types.Score, error) {
	var scores []types.Score
	rows, err := s.db.Query(`SELECT s.id, u.username, g.name, s.score, s.createdat
							FROM score s
							JOIN game g ON s.gameid = g.id
							JOIN users u ON s.playerid = u.id
							WHERE s.approvedat IS NULL`)
	if err != nil {
		return nil, fmt.Errorf("error while getting the scores: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var score types.Score
		if err := rows.Scan(&score.Id, &score.Player, &score.Game, &score.Score, &score.CreatedAt); err != nil {
			return nil, fmt.Errorf("error while getting the scores: %v", err)
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func (s *Store) Insert(playerId int, gameId int, score int) error {
	_, err := s.db.Exec("INSERT INTO score (playerid, gameid, score, createdat) VALUES ($1, $2, $3, $4)", playerId, gameId, score, time.Now())
	if err != nil {
		return fmt.Errorf("error while inserting the score: %v", err)
	}
	return nil
}

func (s *Store) Approve(approverId int) error {
	_, err := s.db.Exec("UPDATE score SET approverId = $1, approvedAt = $2", approverId, time.Now())
	if err != nil {
		return fmt.Errorf("error while updating the score: %v", err)
	}
	return nil
}
