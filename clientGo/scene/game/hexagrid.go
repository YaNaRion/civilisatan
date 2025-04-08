package game

import (
	"client/player"
	"client/window"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Corner int

const (
	Empty Corner = iota
	Village
	TownCenter
)

type HexagoneCorner struct {
	Type     Corner
	Center   rl.Vector2
	IsRender bool
	color    *rl.Color
	Route    []*HexagoneSide
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
	g.addDefaultRoute()
}

func (g *HexaGrid) addDefaultRoute() {
	g.Routes[0].color = &rl.Green
	g.Routes[0].IsRender = true
	g.Routes[1].color = &rl.Red
	g.Routes[1].IsRender = true
}

func (g *HexaGrid) addNeighbor() {
	for _, route := range g.Routes {
		for _, neighbor := range g.Routes {
			if rl.Vector2Equals(route.StartingPoint, neighbor.EndingPoint) ||
				rl.Vector2Equals(route.EndingPoint, neighbor.StartingPoint) ||
				rl.Vector2Equals(route.EndingPoint, neighbor.EndingPoint) ||
				rl.Vector2Equals(route.StartingPoint, neighbor.StartingPoint) {
				route.Neighbor = append(route.Neighbor, neighbor)
			}
		}
	}

	for _, corner := range g.Corner {
		for _, route := range g.Routes {
			if rl.Vector2Equals(corner.Center, route.StartingPoint) ||
				rl.Vector2Equals(corner.Center, route.EndingPoint) {
				corner.Route = append(corner.Route, route)
				route.Corner = append(route.Corner, corner)
			}
		}
	}
}

func (g *HexaGrid) isRoutePositionValid(route *HexagoneSide, active *player.Player) bool {
	for _, route := range route.Neighbor {
		if route.color == active.ColorTeam {
			return true
		}
	}
	return false
}

func (g *HexaGrid) isVillagePositionValid(
	corner *HexagoneCorner,
	activePlayer *player.Player,
) bool {
	var isValide bool
	for _, route := range corner.Route {
		for _, corner := range route.Corner {
			isValide = corner.Type == Empty
			if !isValide {
				return isValide
			}
		}
	}

	isValide = false
	for _, route := range corner.Route {
		if route.color == activePlayer.ColorTeam {
			isValide = true
		}
	}

	return isValide
}

func (g *HexaGrid) DrawRoute(activePlayer *player.Player) {
	if activePlayer != nil {
		for _, side := range g.Routes {
			thinkness := float32(5)
			if rl.CheckCollisionPointLine(
				rl.GetMousePosition(),
				side.StartingPoint,
				side.EndingPoint,
				int32(thinkness-2),
			) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) && side.color == nil &&
				activePlayer.Action == player.PlaceRoute && g.isRoutePositionValid(side, activePlayer) {
				side.IsRender = !side.IsRender
				side.color = activePlayer.ColorTeam
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
}

func (g *HexaGrid) DrawGrid(activePlayer *player.Player) {
	numberOnTile := []string{
		"6",
		"5",
		"9",
		"4",
		"3",
		"8",
		"10",
		"6",
		"7",
		"5",
		"9",
		"12",
		"3",
		"2",
		"10",
		"11",
		"11",
		"4",
		"8",
	}
	number := 0
	for _, row := range g.Grid {
		for _, tile := range row {
			if tile == nil {
				continue
			}
			rl.DrawPoly(tile.Center, g.sides, g.radius-2, 30.0, rl.DarkBlue)
			rl.DrawText(
				numberOnTile[number],
				int32(tile.Center.X),
				int32(tile.Center.Y),
				16,
				rl.Black,
			)
			number++
		}
	}

	g.DrawRoute(activePlayer)
	g.DrawCorner(activePlayer)
}

func (g *HexaGrid) DrawCorner(activePlayer *player.Player) {
	for _, corner := range g.Corner {
		if rl.CheckCollisionPointCircle(rl.GetMousePosition(), corner.Center, 10) &&
			rl.IsMouseButtonPressed(rl.MouseButtonLeft) {

			if activePlayer.Action == player.PlaceVillage && !corner.IsRender &&
				g.isVillagePositionValid(corner, activePlayer) {
				corner.Type = Village
				corner.color = activePlayer.ColorTeam
				corner.IsRender = !corner.IsRender
			}

			if activePlayer.Action == player.UpgradeTownCenter &&
				activePlayer.ColorTeam == corner.color {
				corner.Type = TownCenter
			}
		}

		if corner.IsRender {
			if corner.Type == Village {
				rl.DrawCircle(int32(corner.Center.X), int32(corner.Center.Y), 5, *corner.color)
			}

			if corner.Type == TownCenter {
				rl.DrawCircle(int32(corner.Center.X), int32(corner.Center.Y), 8, *corner.color)
			}
		}
	}
}
