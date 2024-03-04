package main

import (
	"io"
	"os"
)

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		io.WriteString(w, v+"\n")
	}
}

func main() {
	league := League{
		Teams: []Team{
			{"Reds", []string{"r1", "r2", "r3"}},
			{"Blues", []string{"b1", "b2", "b3"}},
			{"Greens", []string{"g1", "g2", "g3"}},
			{"Purples", []string{"p1", "p2", "p3"}},
		},
		Wins: map[string]int{},
	}

	league.MatchResult("Reds", 30, "Blues", 40)
	league.MatchResult("Reds", 30, "Greens", 40)
	league.MatchResult("Reds", 30, "Purples", 40)
	league.MatchResult("Blues", 30, "Greens", 40)
	league.MatchResult("Blues", 30, "Purples", 40)
	league.MatchResult("Greens", 30, "Purples", 40)

	RankPrinter(league, os.Stdout)
}
