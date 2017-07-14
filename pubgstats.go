package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Player struct {
	Name    string
	Kdr     float64
	Skill   float64
	Winrate float64
}

type Target struct {
}

func NewPlayer(name, apikey string) *Player {
	// make api call to gather player information https://pubgtracker.com/site-api
	// curl -X GET https://pubgtracker.com/api/profile/pc/phosphoric -H TRN-API-KEY:1234
	// returns json..
	var kdr, winrate, skill float64
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://pubgtracker.com/api/profile/pc/"+name, nil)
	request.Header.Add("TRN-API-KEY", apikey)
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var data map[string]interface{}
	body, _ := ioutil.ReadAll(response.Body)

	// decode json into data interface
	json.Unmarshal(body, &data)

	// assert array of arbitrary type items
	stats := data["Stats"].([]interface{})
	for i := 0; i < len(stats); i++ {
		// make sure game mode is solo and region is agg
		if stats[i].(map[string]interface{})["Match"] == "solo" && stats[i].(map[string]interface{})["Region"] == "agg" {

			ratings := stats[i].(map[string]interface{})["Stats"].([]interface{})
			for j := 0; j < len(ratings); j++ {
				// print all available stats and their values
				//fmt.Println(ratings[j].(map[string]interface{})["label"], ratings[j].(map[string]interface{})["value"])

				// convert final value string to appropriate values and save them
				if ratings[j].(map[string]interface{})["label"] == "K/D Ratio" {
					kdr, _ = strconv.ParseFloat(ratings[j].(map[string]interface{})["value"].(string), 64)
				}
				if ratings[j].(map[string]interface{})["label"] == "Win %" {
					winrate, _ = strconv.ParseFloat(ratings[j].(map[string]interface{})["value"].(string), 64)
				}
				if ratings[j].(map[string]interface{})["label"] == "Rating" {
					skill, _ = strconv.ParseFloat(ratings[j].(map[string]interface{})["value"].(string), 64)
				}
			}
		}
	}

	// finally pack these stats into a player struct
	player := &Player{
		Name:    name,
		Kdr:     kdr,
		Skill:   skill,
		Winrate: winrate,
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
	fmt.Printf("%14v %10v %10v\n\n", "Win rate %", p.Winrate, p2.Winrate)

	// player2 has x% better kdr.
	// player2 has x% better skill rating.
	// player1 has x% better win rate.

	// for effeciency, should break the following out into a re-usable function
	if p.Kdr == p2.Kdr {
		fmt.Println("Both players have equal K/D ratios.")
	} else if p.Kdr > p2.Kdr {
		fmt.Printf("%v has a %.2f percent better K/D ratio.\n", p.Name, p.Kdr/p2.Kdr*100)
	} else if p.Kdr < p2.Kdr {
		fmt.Printf("%v has a %.2f percent better K/D ratio.\n", p2.Name, p2.Kdr/p.Kdr*100)
	}

	if p.Skill == p2.Skill {
		fmt.Println("Both players have equal skill ratings.")
	} else if p.Skill > p2.Skill {
		fmt.Printf("%v has a %.2f percent better skill rating.\n", p.Name, p.Skill/p2.Skill*100)
	} else if p.Skill < p2.Skill {
		fmt.Printf("%v has a %.2f percent better skill rating.\n", p2.Name, p2.Skill/p.Skill*100)
	}

	if p.Winrate == p2.Winrate {
		fmt.Println("Both players have equal win rates.")
	} else if p.Winrate > p2.Winrate {
		fmt.Printf("%v has a %.2f percent better win rate.\n", p.Name, p.Winrate/p2.Winrate*100)
	} else if p.Winrate < p2.Winrate {
		fmt.Printf("%v has a %.2f percent better win rate.\n", p2.Name, p2.Winrate/p.Winrate*100)
	}
}

func main() {
	fmt.Println("PUGB Player Stat Comparison!\n")

	var apikey string

	// process args
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Please use the '-k' flag to specify and API key from pubgtracker.com.")
		os.Exit(1)
	}
	for x := 0; x < len(args); x++ {
		switch args[x] {
		case "-k", "--key":
			apikey = args[x+1]
			x++
		case "-h", "--help":
			printHelp()
		}
	}

	// passing the api key here is a result of not splitting up the api call form the NewPlayer factory function
	player1 := NewPlayer("Phosphoric", apikey)
	player2 := NewPlayer("Subnova", apikey)
	player1.CompareStats(player2)
}

func printHelp() {
	fmt.Println("Usage: pubgstats [options]")
	fmt.Println("       pubgstats -k 1234-abcd")
}
