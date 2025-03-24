package game

import (
	"client/window"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Corner int

const (
	PlaceVillage Corner = iota
	UpgradeTownCenter
)

type HexagoneCorner struct {
	Type     Corner
	Center   rl.Vector2
	IsRender bool
	color    *rl.Color
}

type HexagoneSide struct {
	StartingPoint rl.Vector2
	EndingPoint   rl.Vector2
	color         *rl.Color
	IsRender      bool
	Neighbor      []*HexagoneSide
	Corner        []*HexagoneCorner
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
	Corner     []*HexagoneCorner
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

func (g *HexaGrid) ResetGrid() {
	g.Grid = make([][]*HexagoneTile, 5)
	g.Routes = nil
	g.Corner = nil
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
	g.addNeighbor()
}

func (g *HexaGrid) addNeighbor() {
	for _, route := range g.Routes {
		for _, neighbor := range g.Routes {
			if rl.Vector2Equals(route.StartingPoint, neighbor.StartingPoint) ||
				rl.Vector2Equals(route.StartingPoint, neighbor.EndingPoint) ||
				rl.Vector2Equals(route.EndingPoint, neighbor.StartingPoint) ||
				rl.Vector2Equals(route.EndingPoint, neighbor.EndingPoint) {
				route.Neighbor = append(route.Neighbor, neighbor)
			}
		}
	}
}

func (g *HexaGrid) DrawCorner(activePlayerCol *rl.Color) {
	for _, corner := range g.Corner {
		if rl.CheckCollisionPointCircle(rl.GetMousePosition(), corner.Center, 10) &&
			rl.IsMouseButtonPressed(rl.MouseButtonLeft) &&
			corner.color == nil {
			corner.color = activePlayerCol
			corner.IsRender = !corner.IsRender
		}

		if corner.IsRender {
			rl.DrawCircle(int32(corner.Center.X), int32(corner.Center.Y), 5, *corner.color)
		}
	}
}

func (g *HexaGrid) DrawRoute(activePlayerCol *rl.Color) {
	for _, side := range g.Routes {
		thinkness := float32(5)
		if rl.CheckCollisionPointLine(
			rl.GetMousePosition(),
			side.StartingPoint,
			side.EndingPoint,
			int32(thinkness-2),
		) && rl.IsMouseButtonPressed(rl.MouseButtonRight) && side.color == nil {
			side.IsRender = !side.IsRender
			side.color = activePlayerCol
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

func (g *HexaGrid) DrawGrid(activePlayerCol *rl.Color) {
	for _, row := range g.Grid {
		for _, tile := range row {
			if tile == nil {
				continue
			}
			rl.DrawPoly(tile.Center, g.sides, g.radius-2, 30.0, rl.DarkBlue)
		}
	}
	g.DrawRoute(activePlayerCol)
	g.DrawCorner(activePlayerCol)
}
