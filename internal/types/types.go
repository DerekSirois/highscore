package types

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string `json:"-"`
	Role      string
	CreatedAt time.Time
}

type Game struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

type Score struct {
	Id         int
	PlayerId   int
	GameId     int
	ApproverId int
	Score      int
	CreatedAt  time.Time
	ApprovedAt time.Time
}

type UserStore interface {
	GetById(id int) (User, error)
	GetByEmail(email string) (User, error)
	Insert(user User) (int64, error)
}
