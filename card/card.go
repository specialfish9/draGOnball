package card

type Card struct {
	Name   string
	Attack int
	Life   int
	Effect
}

func New(name string, attack int, life int, effect Effect) *Card {
	return &Card{
		Name:   name,
		Attack: attack,
		Life:   life,
		Effect: effect,
	}
}

type Effect struct {
	AttackIncr  int
	DefenceIncr int
}
