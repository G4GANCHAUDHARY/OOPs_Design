package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	PlayerId = 0
)

type Grid struct {
	Rows    int
	Columns int
	Target  int
	Board   [][]int
}

func (g *Grid) isPlayerWon(p *Player) bool {

	direction := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			I := i
			J := j
			count := 0
			for _, x := range direction {
				for t := 0; t < g.Target; t++ {
					if I < 0 || I >= g.Rows || J < 0 || J >= g.Columns {
						break
					}

					if g.Board[I][J] != p.Id {
						break
					}

					count++
					I += x[0]
					J += x[1]
				}
				if count == g.Target {
					return true
				}
				I = i
				J = j
				count = 0
			}
		}
	}

	return false
}

func (g *Grid) initialiseGrid(reader *bufio.Reader) {
	fmt.Print("Enter grid rows : ")
	rows, _ := reader.ReadString('\n')
	g.Rows, _ = strconv.Atoi(strings.TrimSpace(rows))

	fmt.Print("Enter grid columns : ")
	columns, _ := reader.ReadString('\n')
	g.Columns, _ = strconv.Atoi(strings.TrimSpace(columns))

	fmt.Print("Enter target : ")
	target, _ := reader.ReadString('\n')
	g.Target, _ = strconv.Atoi(strings.TrimSpace(target))

	var r []int
	for i := 0; i < g.Rows; i++ {
		r = make([]int, 0)
		for j := 0; j < g.Columns; j++ {
			r = append(r, 0)
		}
		g.Board = append(g.Board, r)
	}
}

type Player struct {
	Id   int
	Name string
}

func (p *Player) initialisePlayer(reader *bufio.Reader) {
	fmt.Print("Enter player name : ")
	p.Name, _ = reader.ReadString('\n')
	p.Name = strings.TrimSpace(p.Name)
	p.Id = PlayerId + 1
	PlayerId = p.Id
}

func (p *Player) takeTurn(g *Grid) bool {

	count := 0
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if g.Board[i][j] == 0 {
				count += 1
			}
		}
	}

	if count == 0 {
		return false
	}

	temp := 0
	random := rand.Intn(count)

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if g.Board[i][j] == 0 {
				temp += 1
			}
			if temp == random {
				g.Board[i][j] = p.Id
				return true
			}
		}
	}

	return false
}

func Game(g *Grid, p1 *Player, p2 *Player) {
	for {
		// Player 1 takes turn
		if !p1.takeTurn(g) {
			fmt.Print("No space left to take any turn")
			break
		}

		// check if player won
		if g.isPlayerWon(p1) {
			fmt.Print(fmt.Sprintf("Player %s has won the game", p1.Name))
			break
		}

		// Player 2 takes turn
		if !p2.takeTurn(g) {
			fmt.Print("No space left to take any turn")
			break
		}

		// check if player won
		if g.isPlayerWon(p2) {
			fmt.Print(fmt.Sprintf("Player %s has won the game", p2.Name))
			break
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	grid := &Grid{}
	grid.initialiseGrid(reader)

	player1 := &Player{}
	player1.initialisePlayer(reader)

	player2 := &Player{}
	player2.initialisePlayer(reader)

	Game(grid, player1, player2)
}
