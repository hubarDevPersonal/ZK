package game

import (
	"tic-tac-toe/client/players"
)

type Game struct {
	Cells         []string
	Winner        string
	GameOn        bool
	CurrentPlayer *players.Player
}

func NewGame() *Game {
	return &Game{
		Cells:  []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"},
		Winner: "",
		GameOn: true,
		CurrentPlayer: &players.Player{
			Symbol:   "X",
			Nickname: "",
		},
	}
}
