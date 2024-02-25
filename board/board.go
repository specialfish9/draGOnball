package board

import (
	"dragonball/card"
	"fmt"
)

const MaxHelpers = 3

type ActiveCard struct {
	Card   *card.Card
	Life   int
	Attack int
}

type Board struct {
	Active  ActiveCard
	Helpers []*card.Card
	Deck    []*card.Card
	deckTop int
	Dead    []*card.Card
}

func NewBoard(cards []*card.Card) *Board {
	return &Board{
		Deck: cards,
	}
}

func (b *Board) DrawCard() (*card.Card, error) {
	if len(b.Deck) <= b.deckTop {
		return nil, fmt.Errorf("empty deck")
	}
	card := b.Deck[b.deckTop]
	b.deckTop++
	return card, nil
}

func (b *Board) DiscardCard(card *card.Card) error {
	if card == nil {
		return fmt.Errorf("cannot discard nil card")
	}

	b.Dead = append(b.Dead, card)
	return nil
}

func (b *Board) SetActive(card *card.Card) error {
	if b.Active.Card != nil {
		if err := b.DiscardCard(b.Active.Card); err != nil {
			return fmt.Errorf("cannot set card as active: %v", err)
		}
	}

	b.Active = ActiveCard{
		Card:   card,
		Life:   card.Life,
		Attack: card.Attack,
	}

	for _, v := range b.Helpers {
		b.Active.ApplyEffect(v.Effect)
	}

	return nil
}

func (b *Board) AddHelper(card *card.Card) error {
	if len(b.Helpers) >= MaxHelpers {
		return fmt.Errorf("too many helpers!")
	}

	b.Helpers = append(b.Helpers, card)
	b.Active.ApplyEffect(card.Effect)

	return nil
}

func (b *Board) ReplaceHelper(card *card.Card, repl int) error {
	if repl >= len(b.Helpers) || repl < 0 {
		return fmt.Errorf("invalid replace index")
	}

	if err := b.RemoveHelper(b.Helpers[repl]); err != nil {
		return fmt.Errorf("cannot replace helper: %v", err)
	}
	if err := b.AddHelper(card); err != nil {
		return fmt.Errorf("cannot replace helper: %v", err)
	}

	return nil
}

func (b *Board) RemoveHelper(card *card.Card) error {
	for i, h := range b.Helpers {
		if h == card {
			b.Helpers = append(b.Helpers[:i], b.Helpers[i+1:]...)
			b.Active.RemoveEffect(h.Effect)
			if err := b.DiscardCard(h); err != nil {
				return fmt.Errorf("cannot remove helper: %v", err)
			}
			return nil
		}
	}

	return fmt.Errorf("this card is not an helper!")
}

func (a *ActiveCard) AttackOpponent(opponent *ActiveCard) error {
	if a.Card == nil {
		return fmt.Errorf("cannot Attack without an active card!")
	}

	opponent.Life -= a.Attack

	if opponent.Life < 0 {
		opponent.Life = 0
	}

	return nil
}

func (a *ActiveCard) ApplyEffect(effect card.Effect) {
	a.Attack += effect.AttackIncr
	a.Life += effect.DefenceIncr

	if a.Attack < 0 {
		a.Attack = 0
	}

	if a.Life < 0 {
		a.Life = 0
	}
}

func (a *ActiveCard) RemoveEffect(effect card.Effect) {
	a.Attack -= effect.AttackIncr
	a.Life -= effect.DefenceIncr

	if a.Attack < 0 {
		a.Attack = 0
	}

	if a.Life < 0 {
		a.Life = 0
	}
}
