package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//type Game interface {
//	checkWin(): bool
//}

type Game struct {
	Cells         []string
	Winner        string
	GameOn        bool
	CurrentPlayer *Player
}

func NewGame() *Game {
	return &Game{
		Cells:  []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"},
		Winner: "",
		GameOn: true,
		CurrentPlayer: &Player{
			Symbol:   "X",
			Nickname: "",
		},
	}
}

func (g *Game) Display() {
	fmt.Printf("%s | %s | %s      1|2|3\n", g.Cells[0], g.Cells[1], g.Cells[2])
	fmt.Printf("%s | %s | %s      4|5|6\n", g.Cells[3], g.Cells[4], g.Cells[5])
	fmt.Printf("%s | %s | %s      7|8|9\n", g.Cells[6], g.Cells[7], g.Cells[8])
}

func (g *Game) Players() {
	fmt.Println("Select Player - X or O")
	g.CurrentPlayer.Symbol = strings.ToUpper(g.Input("Player1: "))
	g.CurrentPlayer.SetNickname(g.Input("Enter Player1's nickname: "))

	if g.CurrentPlayer.Symbol == "X" {
		fmt.Printf("Player2 (%s): O")
		g.Input("Enter Player2's nickname: ")
	} else if g.CurrentPlayer.Symbol == "O" {
		fmt.Printf("Player2 (%s): X\n", g.Input("Enter Player2's nickname: "))
	} else {
		fmt.Println("Sorry, invalid input. Type X or O")
		g.Players()
	}
}

func (g *Game) Input(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}

func (g *Game) PlayerPosition() {
	fmt.Printf("%s's turn (%s)\n", g.CurrentPlayer.Nickname, g.CurrentPlayer.Symbol)
	position := g.Input("Choose position from 1 - 9: ")

	valid := false
	for !valid {
		for !contains([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, position) {
			position = g.Input("Invalid. Choose position from 1 - 9: ")
		}
		pos := g.parsePosition(position)

		if g.Cells[pos] == "-" {
			valid = true
		} else {
			fmt.Println("Position already selected, choose another position!")
		}
		g.Cells[pos] = g.CurrentPlayer.Symbol
		g.Display()
	}
}

func (g *Game) CheckWinner() {
	var win bool
	if g.Winner, win = utils; win {
		g.GameOn = false
		fmt.Printf("Congratulations %s (%s), you WON!\n", g.CurrentPlayer.Symbol, g.CurrentPlayer.Nickname)
		g.ExitGame()
	} else if !contains(g.Cells, "-") {
		g.GameOn = false
		g.Winner = "-"
		fmt.Println("It's a Tie")
		g.ExitGame()
	}
}

func contains(cells []string, s string) bool {
	for _, cell := range cells {
		if cell == s {
			return true
		}
	}
	return false
}

func (g *Game) FlipPlayer() {
	if g.CurrentPlayer.Symbol == "X" {
		g.CurrentPlayer.Symbol = "O"
	} else {
		g.CurrentPlayer.Symbol = "X"
	}
}

func (g *Game) ExitGame() {
	parse := map[string]int{g.CurrentPlayer.Symbol: 0, "O": 1, "-": 2}
	for i, cell := range g.Cells {
		g.Cells[i] = fmt.Sprint(parse[cell])
	}
	g.Winner = fmt.Sprint(parse[g.Winner])
	game := rand.Intn(1000)

	file, err := os.Create("../zk/Prover.toml")
	if err != nil {
		fmt.Println("Error creating Prover.toml file:", err)
		return
	}
	defer file.Close()

	data := fmt.Sprintf("board = %v\n"+
		"game = \"%d\"\n"+
		"winner = \"%s\"", g.Cells, game, g.Winner)

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to Prover.toml file:", err)
		return
	}

	fmt.Println("\nGenerating Prover and a proof...")
	fmt.Println(os.ExpandEnv("cd ../zk && nargo prove"))

	fmt.Println("\nGenerating Verifier.sol...")
	fmt.Println(os.ExpandEnv("cd ../zk && nargo codegen-verifier"))
	os.Rename("../zk/contract/noirTicTacToe/plonk_vk.sol", "../zk/contracts/Verifier.sol")
	os.RemoveAll("../zk/contract")

	proofBytes, err := os.ReadFile("../zk/proofs/noirTicTacToe.proof")
	if err != nil {
		fmt.Println("Error reading proof file:", err)
		return
	}
	proof := string(proofBytes)

	fmt.Printf("\n---------------------\n"+
		"Your public inputs are:\n"+
		" - game: %d\n"+
		" - winner: %s\n"+
		"Your proof is: %s\n\n"+
		"This data was saved to ../zk/Prover.toml and ../zk/proofs/noirTicTacToe.proof respectively.\n\n"+
		"Now run `cd ../zk && truffle test`\n", game, g.Winner, proof)
}

func (g *Game) parsePosition(input string) int {
	position, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return position - 1
}

func (g *Game) PlayGame() {
	fmt.Println("My Tic Tac Toe Game")
	g.Display()
	g.Players()

	for g.GameOn {
		g.PlayerPosition()
		g.CheckWinner()
		g.FlipPlayer()
	}
}

func main() {
	game := NewGame()
	game.PlayGame()
}
