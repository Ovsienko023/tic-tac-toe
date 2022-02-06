package board

import (
	"fmt"
	"github.com/inancgumus/screen"
)

func CleanImage() {
	screen.Clear()
}

type Board struct {
	cells [3][3]string
}

func (b *Board) New() {
	for indX, x := range b.cells {
		for indY, _ := range x {
			b.cells[indX][indY] = " "
		}
	}
}

func (b *Board) Update(index int, chip string) {
	switch index {
	case 1:
		b.cells[0][0] = chip
	case 2:
		b.cells[0][1] = chip
	case 3:
		b.cells[0][2] = chip
	case 4:
		b.cells[1][0] = chip
	case 5:
		b.cells[1][1] = chip
	case 6:
		b.cells[1][2] = chip
	case 7:
		b.cells[2][0] = chip
	case 8:
		b.cells[2][1] = chip
	case 9:
		b.cells[2][2] = chip
	}
}
func (b *Board) Draw() {
	fmt.Printf(" %s | %s | %s \n", b.cells[0][0], b.cells[0][1], b.cells[0][2])
	fmt.Println(" ---------")
	fmt.Printf(" %s | %s | %s \n", b.cells[1][0], b.cells[1][1], b.cells[1][2])
	fmt.Println(" ---------")
	fmt.Printf(" %s | %s | %s \n", b.cells[2][0], b.cells[2][1], b.cells[2][2])
}

func IsWin(b *Board, chip string) bool {
	if b.cells[0][0] == chip && b.cells[0][1] == chip && b.cells[0][2] == chip {
		return true
	}
	if b.cells[1][0] == chip && b.cells[1][1] == chip && b.cells[1][2] == chip {
		return true
	}
	if b.cells[2][0] == chip && b.cells[2][1] == chip && b.cells[2][2] == chip {
		return true
	}
	if b.cells[0][0] == chip && b.cells[1][0] == chip && b.cells[2][0] == chip {
		return true
	}
	if b.cells[0][1] == chip && b.cells[1][1] == chip && b.cells[2][1] == chip {
		return true
	}
	if b.cells[0][2] == chip && b.cells[1][2] == chip && b.cells[2][2] == chip {
		return true
	}
	if b.cells[0][0] == chip && b.cells[1][1] == chip && b.cells[2][2] == chip {
		return true
	}
	if b.cells[2][0] == chip && b.cells[2][2] == chip && b.cells[0][2] == chip {
		return true
	}
	return false
}
