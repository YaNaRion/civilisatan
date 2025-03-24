package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *HexaGrid) CreateSides(width float32, height float32, tile *HexagoneTile, i int, j int) {
	var edge []rl.Vector2
	// 0 En bas à gauche
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
		endingPoint := rl.Vector2Add(tile.Center, edge[i+1])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[i]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})

		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 3 && j == 0 {
		endingPoint := rl.Vector2Add(tile.Center, edge[4])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 4 && j == 1 {
		endingPoint := rl.Vector2Add(tile.Center, edge[4])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 4 && j == 2 {
		endingPoint := rl.Vector2Add(tile.Center, edge[4])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[5])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 4 && j == 3 {
		endingPoint := rl.Vector2Add(tile.Center, edge[4])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})

		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[5])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 3 && j == 4 {
		endingPoint := rl.Vector2Add(tile.Center, edge[4])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[3]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[4]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[5])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[0])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 2 && j == 4 {
		endingPoint := rl.Vector2Add(tile.Center, edge[5])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[0])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   endingPoint,
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 1 && j == 4 {
		endingPoint := rl.Vector2Add(tile.Center, edge[5])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[4]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[5]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})

		endingPoint = rl.Vector2Add(tile.Center, edge[0])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 0 && j == 2 {
		endingPoint := rl.Vector2Add(tile.Center, edge[0])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})

		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}

	if i == 1 && j == 3 {
		endingPoint := rl.Vector2Add(tile.Center, edge[0])
		g.Routes = append(g.Routes, &HexagoneSide{
			StartingPoint: rl.Vector2Add(tile.Center, edge[5]),
			EndingPoint:   rl.Vector2Add(tile.Center, edge[0]),
			IsRender:      false,
		})
		g.Corner = append(g.Corner, &HexagoneCorner{
			Center:   endingPoint,
			IsRender: false,
		})
	}
}
