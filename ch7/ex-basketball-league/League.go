package main

import (
	"sort"
)

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	var team1InLeague, team2InLeague bool
	for _, v := range l.Teams {
		if v.name == team1 {
			team1InLeague = true
		} else if v.name == team2 {
			team2InLeague = true
		}
	}
	if !(team1InLeague && team2InLeague) {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		names = append(names, team.name)
	}
	sort.Slice(names, func(i int, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}
