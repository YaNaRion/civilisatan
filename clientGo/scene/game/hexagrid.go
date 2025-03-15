package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type HexagoneTile struct {
	Center rl.Vector2
	Color  rl.Color
}

type HexaGrid struct {
	OriginPos  rl.Vector2
	Center     rl.Vector2
	Grid       [][]HexagoneTile
	RowsLen    int
	ColomnsLen int
	sides      int32
	radius     float32
	Width      float32
	Height     float32
}

func setupGrid() *HexaGrid {
	circumradius := float32(50)
	inradisu := float32(math.Sqrt(3)/2) * float32(circumradius)
	grid := HexaGrid{
		OriginPos: rl.Vector2{
			X: 0,
			Y: 0,
		},
		RowsLen:    5,
		ColomnsLen: 5,
		Grid:       make([][]HexagoneTile, 5, 5),
		sides:      6,
		radius:     circumradius, // De l'hexagone
		Width:      2 * inradisu,
		Height:     2 * circumradius,
	}
	grid.PopulateGrid()
	grid.Center = grid.Grid[2][2].Center
	return &grid
}

func (g *HexaGrid) PopulateGrid() {
	for i := range g.RowsLen {
		g.Grid[i] = make([]HexagoneTile, g.ColomnsLen)
		for j := range g.ColomnsLen {
			if j%2 == 0 {
				g.Grid[i][j] = HexagoneTile{
					Center: rl.Vector2{
						X: g.OriginPos.X + g.Width + (float32(i) * g.Width),
						Y: g.OriginPos.Y + (g.Height / 2) + (float32(j) * g.Height * 3 / 4),
					},
				}
			} else {
				g.Grid[i][j] = HexagoneTile{
					Center: rl.Vector2{
						X: g.OriginPos.X + (g.Width / 2) + (float32(i) * g.Width),
						Y: g.OriginPos.Y + (g.Height / 2) + (float32(j) * g.Height * 3 / 4),
					},
				}
			}
		}
	}
}

func (g *HexaGrid) DrawGrid() {
	for i, row := range g.Grid {
		for j, tile := range row {
			if (i == 0 && j == 0) || (i == 4 && j == 0) || (i == 0 && j == 1) ||
				(i == 0 && j == 3) || (i == 0 && j == 4) || (i == 4 && j == 4) {
				continue
			}
			rl.DrawPoly(tile.Center, g.sides, g.radius, 30.0, rl.DarkBlue)
			rl.DrawPolyLines(tile.Center, g.sides, g.radius, 30.0, rl.Black)
		}
	}
}
