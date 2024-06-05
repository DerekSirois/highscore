package types

import "time"

// Users
type User struct {
	Id        int
	Username  string
	Email     string
	Password  string `json:"-"`
	Role      string
	CreatedAt time.Time
}

type UserLogin struct {
	Email    string
	Password string
}

type UserRegister struct {
	Username string
	Email    string
	Password string
}

// Games
type Game struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

// Scores
type Score struct {
	Id         int
	PlayerId   int
	GameId     int
	ApproverId int
	Score      int
	CreatedAt  time.Time
	ApprovedAt time.Time
}

// Interfaces
type UserStore interface {
	GetById(id int) (User, error)
	GetByEmail(email string) (User, error)
	Insert(user UserRegister) error
}
