package engine

type turn int

const (
	turnOne = 0
	turnTwo = 1
)

func (t *turn) next() {
	*t = (*t + 1) % 2
}
