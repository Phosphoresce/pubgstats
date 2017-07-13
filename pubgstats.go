package main

import "fmt"

type Player struct {
	Name    string
	Kdr     float32
	Skill   int
	Winrate float32
}

func NewPlayer(name string) *Player {
	// make api call to gather player information https://pubgtracker.com/site-api
	// curl -X GET https://pubgtracker.com/api/profile/pc/phosphoric -H TRN-API-KEY:1234
	// returns json

	player := &Player{
		Name:    name,
		Kdr:     0.4,
		Skill:   1356,
		Winrate: 0.1,
	}
	return player
}

func (p *Player) CompareStats(p2 *Player) {
	// Pretty print data!

	// 		player1 player2
	// kdr		0.4	0.5
	// skill rating	1234	1345
	// win rate	20%	10%
	fmt.Printf("%14v %10v %10v\n", "", p.Name, p2.Name)
	fmt.Printf("%14v %10v %10v\n", "K/D ratio", p.Kdr, p2.Kdr)
	fmt.Printf("%14v %10v %10v\n", "Skill rating", p.Skill, p2.Skill)
	fmt.Printf("%14v %10v %10v\n\n", "Win rate", p.Winrate, p2.Winrate)

	// player2 has x% better kdr.
	// player2 has x% better skill rating.
	// player1 has x% better win rate.

	// for effeciency, should break the following out into a re-usable function
	if p.Kdr == p2.Kdr {
		fmt.Println("Both players have equal K/D ratios.")
	} else if p.Kdr > p2.Kdr {
		fmt.Printf("%v has a %v percent better K/D ratio.\n", p.Name, p.Kdr/p2.Kdr*100)
	} else if p.Kdr < p2.Kdr {
		fmt.Printf("%v has a %v percent better K/D ratio.\n", p2.Name, p2.Kdr/p.Kdr*100)
	}

	if p.Skill == p2.Skill {
		fmt.Println("Both players have equal skill ratings.")
	} else if p.Skill > p2.Skill {
		fmt.Printf("%v has a %v percent better skill rating.\n", p.Name, p.Skill/p2.Skill*100)
	} else if p.Skill < p2.Skill {
		fmt.Printf("%v has a %v percent better skill rating.\n", p2.Name, p2.Skill/p.Skill*100)
	}

	if p.Winrate == p2.Winrate {
		fmt.Println("Both players have equal win rates.")
	} else if p.Winrate > p2.Winrate {
		fmt.Printf("%v has a %v percent better win rate.\n", p.Name, p.Winrate/p2.Winrate*100)
	} else if p.Winrate < p2.Winrate {
		fmt.Printf("%v has a %v percent better win rate.\n", p2.Name, p2.Winrate/p.Winrate*100)
	}
}

func main() {
	fmt.Println("PUGB Player Stat Comparison!\n")
	player1 := NewPlayer("thatguy1")
	player2 := NewPlayer("thatgal2")
	player1.CompareStats(player2)
}
