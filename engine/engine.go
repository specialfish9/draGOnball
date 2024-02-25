package engine

import (
	"dragonball/board"
	"dragonball/card"
	"dragonball/player"
	"fmt"
)

type Game struct {
	pb1 *board.Board
	pb2 *board.Board

	turn turn
}

func NewGame(p1Cards, p2Cards []*card.Card) *Game {
	return &Game{
		pb1:  board.NewBoard(p1Cards),
		pb2:  board.NewBoard(p2Cards),
		turn: turnOne,
	}
}

func StartNewGame(p1, p2 player.Player) {
	game := NewGame(p1.GetDeck(), p2.GetDeck())

	for {
		var activeBoard, opponentBoard *board.Board
		var activePlayer, opponentPlayer player.Player

		if game.turn == turnOne {
			activeBoard = game.pb1
			opponentBoard = game.pb2
			activePlayer = p1
			opponentPlayer = p2
		} else if game.turn == turnTwo {
			activeBoard = game.pb2
			opponentBoard = game.pb1
			activePlayer = p2
			opponentPlayer = p1
		} else {
			panic("invalid turn")
		}

		state, err := handleMove(activePlayer, activeBoard, opponentPlayer, opponentBoard)

		if err != nil {
			activePlayer.ShowError(err)
			continue
		}

		if state == Won {
			activePlayer.Win()
			break
		} else if state == Continue {
			game.turn.next()
		} else {
			panic("invalid game state")
		}
	}
}

func handleMove(activePlayer player.Player, activeBoard *board.Board, opponentPlayer player.Player, opponentBoard *board.Board) (GameState, error) {
	move := activePlayer.NextMove(*activeBoard, opponentBoard.Active)

	if move == player.MoveDraw {
		return draw(activePlayer, activeBoard)
	} else if move == player.MoveAttack {
		return attack(activeBoard, &opponentBoard.Active)
	}

	return Error, fmt.Errorf("cannot handle move %d", move)
}

func draw(active player.Player, board *board.Board) (GameState, error) {
	card, err := board.DrawCard()
	if err != nil {
		return Error, fmt.Errorf("cannot handle draw move: %v", err)
	}

	result := active.DrawCard(card)

	if result == player.SetActive {
		err = board.SetActive(card)
	} else if result == player.Discard {
		err = board.DiscardCard(card)
	} else if result == player.AddHelper {
		addHelper(active, board, card)
	}

	if err != nil {
		return Error, fmt.Errorf("cannot handle draw move: %v", err)
	}

	return Continue, nil
}

func addHelper(playe player.Player, boar *board.Board, card *card.Card) {
	var err error

	for {
		index := playe.UpdateHelpers(boar.Helpers)
		if index < 0 {
			err = boar.AddHelper(card)
		} else {
			err = boar.ReplaceHelper(card, index)
		}

		if err != nil {
			playe.ShowError(err)
		} else {
			return
		}
	}
}

func attack(activeBoard *board.Board, opponent *board.ActiveCard) (GameState, error) {
	if err := activeBoard.Active.AttackOpponent(opponent); err != nil {
		return Error, err
	}

	if opponent.Life <= 0 {
		return Won, nil
	}

	return Continue, nil

}

type GameState int

const (
	Error    GameState = -1
	Continue GameState = 0
	Won      GameState = 1
)
