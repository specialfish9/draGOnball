package main

import (
	"dragonball/card"
	"dragonball/engine"
	"dragonball/player"
	"math/rand"
	"time"
)

func main() {
	var deck1 []*card.Card
	var deck2 []*card.Card

	for i := 0; i < 10; i++ {
		deck1 = append(deck1, generateDragonBallCard())
		deck2 = append(deck2, generateDragonBallCard())
	}

	p1 := player.NewHuman("Mario", deck1)
	p2 := player.NewHuman("Luigi", deck2)

	engine.StartNewGame(p1, p2)
}

func generateDragonBallCard() *card.Card {
	dragonBallCharacters := []string{
		"Goku",
		"Vegeta",
		"Piccolo",
		"Gohan",
		"Frieza",
		"Cell",
		"Bulma",
		"Krillin",
		"Trunks",
		"Android 18",
		"Yamcha",
		"Tien",
		"Master Roshi",
		"Chi-Chi",
		"Beerus",
		"Whis",
		"Bardock",
		"Goten",
		"Majin Buu",
		"Shenron",
	}
	attack := rand.Intn(50) + 1
	life := rand.Intn(100) + 50

	effect := card.Effect{
		AttackIncr:  rand.Intn(30) + 1, // Incremento casuale dell'attacco
		DefenceIncr: rand.Intn(30) + 1, // Incremento casuale della difesa
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(dragonBallCharacters))
	characterName := dragonBallCharacters[index]

	card := card.New(characterName, attack, life, effect)

	return card
}
