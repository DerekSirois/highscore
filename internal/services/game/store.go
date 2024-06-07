package game

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

func (s *Store) GetAll() ([]types.Game, error) {
	var games []types.Game
	rows, err := s.db.Query("SELECT * FROM game")
	if err != nil {
		return nil, fmt.Errorf("error while getting the games: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var game types.Game
		if err := rows.Scan(&game.Id, &game.Name, &game.CreatedAt); err != nil {
			return nil, fmt.Errorf("error while getting the games: %v", err)
		}
		games = append(games, game)
	}

	return games, nil
}

func (s *Store) GetById(id int) (types.Game, error) {
	var game types.Game
	row := s.db.QueryRow("SELECT * FROM game WHERE id = $1", id)

	if err := row.Scan(&game.Id, &game.Name, &game.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return game, fmt.Errorf("game not found")
		}
		return game, fmt.Errorf("error while getting the game: %v", err)
	}
	return game, nil
}

func (s *Store) Insert(game types.Game) error {
	_, err := s.db.Exec("INSERT INTO game (name, createdat) VALUES ($1, $2)", game.Name, time.Now())
	if err != nil {
		return fmt.Errorf("error while inserting the game: %v", err)
	}
	return nil
}
