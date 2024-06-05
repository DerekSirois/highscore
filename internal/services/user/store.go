package user

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
	return &Store{
		db: db,
	}
}

func (s *Store) GetById(id int) (types.User, error) {
	var user types.User

	row := s.db.QueryRow(`SELECT u.id, u.username, u.email, u.password, u.createdat, r.name as role 
						  FROM users u JOIN role r on u.roleid = r.id 
					      WHERE id = $1`, id)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error while getting the user: %v", err)
	}
	return user, nil
}

func (s *Store) GetByEmail(email string) (types.User, error) {
	var user types.User

	row := s.db.QueryRow(`SELECT u.id, u.username, u.email, u.password, u.createdat, r.name as role 
						  FROM users u JOIN role r on u.roleid = r.id 
						  WHERE email = $1`, email)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error while getting the user: %v", err)
	}
	return user, nil
}

func (s *Store) Insert(user types.UserRegister) error {
	_, err := s.db.Exec(`INSERT INTO users (username, email, password, roleid, createdat)
							  VALUES ($1, $2, $3, 1, $4)`, user.Username, user.Email, user.Password, time.Now())
	if err != nil {
		return fmt.Errorf("error while inserting the user: %v", err)
	}
	return nil
}
