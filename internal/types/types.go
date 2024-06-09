package types

import "time"

const (
	UserRole      string = "User"
	ModeratorRole string = "Moderator"
	AdminRole     string = "Admin"
)

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
	Player     string
	Game       string
	Approver   string
	Score      int
	CreatedAt  time.Time
	ApprovedAt time.Time
}

type ScoreSubmit struct {
	GameId int
	Score  int
}

// Interfaces
type UserStore interface {
	GetById(id int) (User, error)
	GetByEmail(email string) (User, error)
	Insert(user UserRegister) error
}

type GameStore interface {
	GetAll() ([]Game, error)
	GetById(id int) (Game, error)
	Insert(game Game) error
}

type ScoreStore interface {
	GetAllApprovedScoresByGame(gameId int) ([]Score, error)
	GetAllScoresPendingApproval() ([]Score, error)
	Insert(playerId int, gameId int, score int) error
	Approve(approverId int, scoreId int) error
}
