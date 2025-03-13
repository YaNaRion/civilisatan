package grid

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HexagoneTile struct {
	center rl.Vector2
	color  rl.Color
}

type HexaGrid struct {
	originPos  rl.Vector2
	grid       [][]HexagoneTile
	rowsLen    int
	colomnsLen int
	sides      int32
	radius     float32
	width      float32
	height     float32
}

func SetupGrid(origin rl.Vector2) *HexaGrid {

	circumradius := float32(50)
	inradisu := float32(math.Sqrt(3)/2) * float32(circumradius)
	grid := HexaGrid{
		originPos:  origin,
		rowsLen:    5,
		colomnsLen: 5,
		grid:       make([][]HexagoneTile, 5, 5),
		sides:      6,
		radius:     circumradius, // De l'hexagone
		width:      2 * inradisu,
		height:     2 * circumradius,
	}
	grid.PopulateGrid()

	return &grid
}

func (g *HexaGrid) PopulateGrid() {
	for i := range g.rowsLen {
		g.grid[i] = make([]HexagoneTile, g.colomnsLen)
		for j := range g.colomnsLen {
			if j%2 == 0 {
				g.grid[i][j] = HexagoneTile{
					center: rl.Vector2{
						X: g.originPos.X + g.width + (float32(i) * g.width),
						Y: g.originPos.Y + (g.height / 2) + (float32(j) * g.height * 3 / 4),
					},
				}
			} else {
				g.grid[i][j] = HexagoneTile{
					center: rl.Vector2{
						X: g.originPos.X + (g.width / 2) + (float32(i) * g.width),
						Y: g.originPos.Y + (g.height / 2) + (float32(j) * g.height * 3 / 4),
					},
				}
			}
		}
	}
}

func (g *HexaGrid) DrawGrid() {
	for i, row := range g.grid {
		for j, tile := range row {
			if (i == 0 && j == 0) || (i == 4 && j == 0) || (i == 0 && j == 1) ||
				(i == 0 && j == 3) || (i == 0 && j == 4) || (i == 4 && j == 4) {
				continue
			}
			rl.DrawPoly(tile.center, g.sides, g.radius, 30.0, rl.DarkBlue)
			rl.DrawPolyLines(tile.center, g.sides, g.radius, 30.0, rl.Black)
		}
	}
}
