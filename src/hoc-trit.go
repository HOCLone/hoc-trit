package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	status  [3][3]uint8
	current uint8
}

func (g *Game) winner() uint8 {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		first := g.status[line[0][0]][line[0][1]]
		if first != 0 &&
			first == g.status[line[1][0]][line[1][1]] &&
			first == g.status[line[2][0]][line[2][1]] {
			return first
		}
	}

	return 0
}

func winprt(winner uint8, screen *ebiten.Image) {
	if winner == 1 {
		ebitenutil.DebugPrint(screen, "X player win")
	} else {
		ebitenutil.DebugPrint(screen, "O player win")
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.status[y/200][x/200] = g.current
		if g.current == 1 {
			g.current = 2
		} else {
			g.current = 1
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if w := g.winner(); w != 0 {
		winprt(w, screen)
		return
	}

	vector.StrokeLine(screen, 200, 0, 200, 600, 1, color.White, true)
	vector.StrokeLine(screen, 400, 0, 400, 600, 1, color.White, true)
	vector.StrokeLine(screen, 0, 200, 600, 200, 1, color.White, true)
	vector.StrokeLine(screen, 0, 400, 600, 400, 1, color.White, true)

	for i, row := range g.status {
		for j, item := range row {
			if item == 1 {
				vector.StrokeLine(screen, float32(j*200), float32(i*200), float32(j*200+200), float32(i*200+200), float32(1), color.White, true)
				vector.StrokeLine(screen, float32(j*200+200), float32(i*200), float32(j*200), float32(i*200+200), float32(1), color.White, true)
			} else if item == 2 {
				vector.StrokeCircle(screen, float32(j*200+100), float32(i*200+100), 100, 1, color.White, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle(">>> Press ESC to exit <<<")
	if err := ebiten.RunGame(&Game{status: [3][3]uint8{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, current: 1}); err != nil {
		log.Fatal(err)
	}
}
