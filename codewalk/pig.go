package main

import (
	"fmt"
	"math/rand"
)

const (
	win = 100
	gamesPerSeries = 10
)

type score struct {
	player, opponent, thisTurn int
}

type action func(current score) (result score, turnIsOver bool)

func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}


