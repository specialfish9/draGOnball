package player

import (
	"bufio"
	"dragonball/board"
	"dragonball/card"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Human struct {
	deck []*card.Card
	name string
}

var _ Player = (*Human)(nil)

func NewHuman(name string, deck []*card.Card) *Human {
	return &Human{
		name: name,
		deck: deck,
	}
}

func (h *Human) GetDeck() []*card.Card {
	return h.deck
}

func (h *Human) NextMove(board board.Board, enemy board.ActiveCard) Move {
	fmt.Println("  ╔═Your turn═════════════╗")
	fmt.Printf("  ║   %-20s║\n", h.name)
	fmt.Println("  ╚═══════════════════════╝")
	printBoard(&board, enemy)
	if len(board.Helpers) > 0 {
		fmt.Println("Here's your helpers:")
		printHelpers(board.Helpers)
	}
	for {
		fmt.Println("  ╔═Choose═════════════╗")
		fmt.Println("  ║ [1] Draw a card!   ║")
		if board.Active.Card != nil {
			fmt.Println("  ║ [2] Attack!        ║")
		}
		fmt.Println("  ╚════════════════════╝")
		value := read()
		if value == "1" {
			return MoveDraw
		} else if value == "2" {
			return MoveAttack
		} else {
			fmt.Printf("Invalid option '%s'\n", value)
		}
	}
}

func (h *Human) DrawCard(card *card.Card) DrawMove {
	fmt.Println("Here is your card:")
	printCard(card)
	for {
		fmt.Println("  ╔═Choose═════════════╗")
		fmt.Println("  ║ Here is your card. ║")
		fmt.Println("  ║ [1] Set as active  ║")
		fmt.Println("  ║ [2] Add as helper  ║")
		fmt.Println("  ║ [3] Discard!       ║")
		fmt.Println("  ╚════════════════════╝")
		value := read()
		if value == "1" {
			return SetActive
		} else if value == "2" {
			return AddHelper
		} else if value == "3" {
			return Discard
		} else {
			fmt.Printf("Invalid option '%s'\n", value)
		}
	}

}

func (h *Human) Win() {
	fmt.Println(h.name, " won!")
}

func (h *Human) ShowError(err error) {
	fmt.Println("Error: ", err.Error())
}

func (h *Human) UpdateHelpers(helpers []*card.Card) int {
	if len(helpers) == 0 {
		return -1
	}
	fmt.Println("Here's your helpers:")
	printHelpers(helpers)

	for {
		fmt.Println("  ╔═Choose════════════════════════╗")
		fmt.Println("  ║ What do you want to do?       ║")
		fmt.Println("  ║ [1] Add new helper            ║")
		fmt.Println("  ║ [2] Replace existing helper   ║")
		fmt.Println("  ╚═══════════════════════════════╝")
		value := read()
		if value == "1" {
			return -1
		} else if value == "2" {
			for {
				fmt.Println("Which helper do you want to replace? (0,1,2...)")
				value = read()
				res, err := strconv.Atoi(value)
				if err != nil || res < 0 {
					fmt.Println("Nope, try again")
				} else {
					return res
				}
			}
		}
	}

}
func printBoard(b *board.Board, a board.ActiveCard) {
	var display [10]string
	if b.Active.Card == nil {
		display[0] += fmt.Sprintf("  ╔═Active card════════════════╗")
		display[1] += fmt.Sprintf("  ║           <NONE>           ║")
		display[2] += fmt.Sprintf("  ╚════════════════════════════╝")
		display[3] += fmt.Sprintf("                                ")
		display[4] += fmt.Sprintf("                                ")
	} else {
		display[0] += fmt.Sprintf("  ╔═Active card═══════════╗")
		display[1] += fmt.Sprintf("  ║       Name: %-10s║", b.Active.Card.Name)
		display[2] += fmt.Sprintf("  ║     Attack: %-10d║", b.Active.Attack)
		display[3] += fmt.Sprintf("  ║       Life: %-10d║", b.Active.Life)
		display[4] += fmt.Sprintf("  ╚═══════════════════════╝")
	}

	if a.Card == nil {
		display[0] += fmt.Sprintf("  ╔═Enemy card═════════════════╗")
		display[1] += fmt.Sprintf("  ║           <NONE>           ║")
		display[2] += fmt.Sprintf("  ╚════════════════════════════╝")
		display[3] += fmt.Sprintf("                                ")
		display[4] += fmt.Sprintf("                                ")
	} else {
		display[0] += fmt.Sprintf("  ╔═Enemy card════════════╗")
		display[1] += fmt.Sprintf("  ║       Name: %-10s║", a.Card.Name)
		display[2] += fmt.Sprintf("  ║     Attack: %-10d║", a.Attack)
		display[3] += fmt.Sprintf("  ║       Life: %-10d║", a.Life)
		display[4] += fmt.Sprintf("  ╚═══════════════════════╝")
	}

	for _, s := range display {
		if s != "" {
			fmt.Println(s)
		}
	}
}

func printHelpers(helpers []*card.Card) {
	var display [5]string
	for _, c := range helpers {
		display[0] += fmt.Sprintf("  ╔═Helper card═══════════╗")
		display[1] += fmt.Sprintf("  ║       Name: %-10s║", c.Name)
		display[2] += fmt.Sprintf("  ║     Attack: +%-9d║", c.Effect.AttackIncr)
		display[3] += fmt.Sprintf("  ║       Life: +%-9d║", c.Effect.DefenceIncr)
		display[4] += fmt.Sprintf("  ╚═══════════════════════╝")
	}

	for _, s := range display {
		fmt.Println(s)
	}
}

func printCard(c *card.Card) {
	fmt.Println("  ╔═══════════════════════╗")
	fmt.Printf("  ║%6s%-11s%6s║\n", "", c.Name, "")
	fmt.Printf("  ║%23s║\n", "")
	fmt.Printf("  ║%6sAttack:    %-6d%s║\n", "", c.Attack, "")
	fmt.Printf("  ║%6sLife:      %-6d%s║\n", "", c.Life, "")
	fmt.Printf("  ║%6sEffect:%10s║\n", "", "")
	fmt.Printf("  ║%6s- Attack: +%-3d%3s║\n", "", c.Effect.AttackIncr, "")
	fmt.Printf("  ║%6s- Life:   +%-3d%3s║\n", "", c.Effect.DefenceIncr, "")
	fmt.Println("  ╚═══════════════════════╝")
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("error reading from stdin!")
	}
	return strings.Trim(text, "\n ")
}
