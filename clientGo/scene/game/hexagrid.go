package game

import (
	"client/window"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type HexagoneSide struct {
	StartingPoint rl.Vector2
	EndingPoint   rl.Vector2
	color         *rl.Color
	IsRender      bool
}

type HexagoneTile struct {
	Center rl.Vector2
	Color  rl.Color
	Sides  []*HexagoneSide
}

type HexaGrid struct {
	OriginPos  rl.Vector2
	Center     rl.Vector2
	Grid       [][]*HexagoneTile
	RowsLen    int
	ColomnsLen int
	sides      int32
	radius     float32
	Width      float32
	Height     float32
	OffSet     float32
}

func setupGrid() *HexaGrid {
	circumradius := float32(50)
	inradisu := float32(math.Sqrt(3)/2) * float32(circumradius)
	grid := HexaGrid{
		OriginPos: rl.Vector2{
			X: float32(window.SCREEN_WIDTH) / 3,
			Y: float32(window.SCREEN_HEIGHT) / 3,
		},
		RowsLen:    5,
		ColomnsLen: 5,
		Grid:       make([][]*HexagoneTile, 5),
		sides:      6,
		radius:     circumradius, // De l'hexagone
		Width:      2 * inradisu,
		Height:     2 * circumradius,
		OffSet:     10,
	}
	grid.PopulateGrid()
	grid.Center = grid.Grid[2][2].Center
	return &grid
}

func (g *HexaGrid) PopulateGrid() {
	for i := range g.RowsLen {
		g.Grid[i] = make([]*HexagoneTile, g.ColomnsLen)
		for j := range g.ColomnsLen {
			if (i == 0 && j == 0) || (i == 4 && j == 0) || (i == 0 && j == 1) ||
				(i == 0 && j == 3) || (i == 0 && j == 4) || (i == 4 && j == 4) {
				continue
			}
			var center rl.Vector2
			if j%2 == 0 {
				center = rl.Vector2{
					X: g.OriginPos.X + g.Width + (float32(i) * g.Width),
					Y: g.OriginPos.Y + (g.Height / 2) + (float32(j) * g.Height * 3 / 4),
				}
			} else {
				center = rl.Vector2{
					X: g.OriginPos.X + (g.Width / 2) + (float32(i) * g.Width),
					Y: g.OriginPos.Y + (g.Height / 2) + (float32(j) * g.Height * 3 / 4),
				}
			}
			tile := HexagoneTile{
				Color:  rl.DarkBlue,
				Center: center,
			}
			tile.CreateSides(g.Width, g.Height)
			g.Grid[i][j] = &tile
		}
	}
	g.Grid[0][0] = nil
}

func (g *HexagoneTile) DrawRoute(activePlayerCol rl.Color) {
	for _, side := range g.Sides {
		thinkness := float32(10)
		if rl.CheckCollisionPointLine(
			rl.GetMousePosition(),
			side.StartingPoint,
			side.EndingPoint,
			int32(thinkness-2),
		) && rl.IsMouseButtonPressed(rl.MouseButtonRight) && side.color == nil {
			side.IsRender = !side.IsRender
			side.color = &activePlayerCol
		}

		if side.IsRender {
			rl.DrawLineEx(
				side.StartingPoint,
				side.EndingPoint,
				float32(thinkness),
				*side.color,
			)
		}
	}
}

func (g *HexaGrid) DrawGrid(activePlayerCol rl.Color) {
	for _, row := range g.Grid {
		for _, tile := range row {
			if tile == nil {
				continue
			}
			tile.DrawRoute(activePlayerCol)
			rl.DrawPoly(tile.Center, g.sides, g.radius-2, 30.0, rl.DarkBlue)
		}
	}
}

func (h *HexagoneTile) CreateSides(width float32, height float32) {
	var edge []rl.Vector2
	edge = append(edge, rl.Vector2{
		X: 0,
		Y: -height / 2,
	})
	edge = append(edge, rl.Vector2{
		X: width / 2,
		Y: -height / 4,
	})
	edge = append(edge, rl.Vector2{
		X: width / 2,
		Y: height / 4,
	})
	edge = append(edge, rl.Vector2{
		X: 0,
		Y: height / 2,
	})
	edge = append(edge, rl.Vector2{
		X: -width / 2,
		Y: height / 4,
	})
	edge = append(edge, rl.Vector2{
		X: -width / 2,
		Y: -height / 4,
	})
	edge = append(edge, edge[0])

	h.Sides = make([]*HexagoneSide, 6)
	for i := range h.Sides {
		h.Sides[i] = &HexagoneSide{
			StartingPoint: rl.Vector2Add(h.Center, edge[i]),
			EndingPoint:   rl.Vector2Add(h.Center, edge[i+1]),
			IsRender:      false,
		}
	}
}
