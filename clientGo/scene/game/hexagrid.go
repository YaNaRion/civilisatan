package game

import (
	"client/window"
	"log"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HexagoneSide struct {
	StartingPoint rl.Vector2
	EndingPoint   rl.Vector2
	color         *rl.Color
	IsRender      bool
	// neighbor      []*HexagoneSide
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
	Routes     []*HexagoneSide
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
			tile := &HexagoneTile{
				Color:  rl.DarkBlue,
				Center: center,
			}
			g.Grid[i][j] = tile
			g.CreateSides(g.Width, g.Height, tile, i, j)
		}
	}
	g.Grid[0][0] = nil
}

func (g *HexaGrid) DrawRoute(activePlayerCol rl.Color) {
	for _, side := range g.Routes {
		thinkness := float32(5)
		log.Println("DANS DRAW ROUTE")
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
			rl.DrawPoly(tile.Center, g.sides, g.radius-2, 30.0, rl.DarkBlue)
		}
	}
	g.DrawRoute(activePlayerCol)
}

func (g *HexaGrid) CreateSides(width float32, height float32, tile *HexagoneTile, i int, j int) {
	var edge []rl.Vector2

	// 0
	edge = append(edge, rl.Vector2{
		X: -width / 2,
		Y: height / 4,
	})

	// 1
	edge = append(edge, rl.Vector2{
		X: -width / 2,
		Y: -height / 4,
	})

	// 2
	edge = append(edge, rl.Vector2{
		X: 0,
		Y: -height / 2,
	})

	// 3
	edge = append(edge, rl.Vector2{
		X: width / 2,
		Y: -height / 4,
	})

	// 4
	edge = append(edge, rl.Vector2{
		X: width / 2,
		Y: height / 4,
	})

	// 5
	edge = append(edge, rl.Vector2{
		X: 0,
		Y: height / 2,
	})

	for i := range 3 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[i]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[i+1]),
			IsRender:      false,
		})
	}

	if i == 3 && j == 0 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})
	}

	if i == 4 && j == 1 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})
	}

	if i == 4 && j == 2 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})

		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})

	}

	if i == 4 && j == 3 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})

		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
	}

	if i == 3 && j == 4 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})

		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
	}

	if i == 2 && j == 4 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
	}

	if i == 1 && j == 4 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
	}

	if i == 0 && j == 2 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
	}

	if i == 1 && j == 3 {
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
	}
}
