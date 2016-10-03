package main

type GameMap struct {
	cellMap      [][]bool
	changeEvents []ChangeEvent
}

func (gameMap *GameMap) GetSize() (int, int) {
	height := len(gameMap.cellMap)
	if height == 0 {
		panic("Height must not be 0")
	}

	width := len(gameMap.cellMap[0])
	if width == 0 {
		panic("Width must not be 0")
	}

	return height, width
}

func (gameMap *GameMap) SetValue(height, width int, value bool) {
	gameMap.cellMap[height][width] = value
}

func (gameMap *GameMap) GetValue(height, width int) bool {
	return gameMap.cellMap[height][width]
}

func (gameMap *GameMap) Update() {
	height, width := gameMap.GetSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			f(i, j, gameMap)
		}
	}
}

func (gameMap *GameMap) DoForEveryCell(f func (height, width int, gameMap *GameMap)) {
	height, width := gameMap.GetSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			f(i, j, gameMap)
		}
	}
}

type ChangeEvent struct {
	x,y int
	newValue bool
}