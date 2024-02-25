package player

import (
	"dragonball/board"
	"dragonball/card"
)

type Player interface {
	GetDeck() []*card.Card
	NextMove(board.Board, board.ActiveCard) Move
	DrawCard(*card.Card) DrawMove
	UpdateHelpers([]*card.Card) int
	Win()
	ShowError(error)
}

type Move int

const (
	MoveDraw Move = iota
	MoveAttack
)

type DrawMove int

const (
	SetActive DrawMove = 0
	AddHelper DrawMove = 1
	Discard   DrawMove = 2
)
