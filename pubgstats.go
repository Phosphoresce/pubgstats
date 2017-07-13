package main

import "fmt"

type Player struct {
	Name    string
	Kdr     float32
	Skill   int
	Winrate float32
}

func NewPlayer() *Player {
	// make api call to gather player information https://pubgtracker.com/site-api
	// player := &Player{
	//	Name: thatguy1
	//	Kdr: 0.4
	//	Skill: 1356
	//	Winrate: 0.1
	// }
	return player
}

func (p *Player) CompareStats(p2 *Player) {
	// Pretty print data!

	// 		player1 player2
	// kdr		0.4	0.5
	// skill rating	1234	1345
	// win rate	20%	10%

	// player2 has x% better kdr.
	// player2 has x% better skill rating.
	// player1 has x% better win rate.
}

func main() {
	fmt.Println("PUGB Player Stat Comparison!")
	player1 := NewPlayer()
	player2 := NewPlayer()
	player1.CompareStats(player2)
}
