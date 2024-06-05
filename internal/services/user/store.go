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

func (s *Store) GetById(id int) (types.User, error) {
	var user types.User

	row := s.db.QueryRow(`SELECT u.id, u.username, u.email, u.password, u.createdat, r.name as role 
						  FROM users u JOIN role r on u.roleid = r.id 
					      WHERE id = ?`, id)
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
						  WHERE email = ?`, email)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error while getting the user: %v", err)
	}
	return user, nil
}

func (s *Store) Insert(user types.User) (int64, error) {
	result, err := s.db.Exec(`INSERT INTO users (username, email, password, roleid, createdat)
							  (?, ?, ?, 1, ?)`, user.Username, user.Email, user.Password, time.Now())
	if err != nil {
		return 0, fmt.Errorf("error while inserting the user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erroe while getting the id of the inserted user: %v", err)
	}

	return id, nil
}
