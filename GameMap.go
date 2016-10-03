package main

type GameMap struct {
	cellMap      [][]bool
	changeEvents []ChangeEvent
}

func (gameMap *GameMap) getSize() (int, int) {
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

func (gameMap *GameMap) setValue(height, width int, value bool) {
	gameMap.cellMap[height][width] = value
}

func (gameMap *GameMap) getValue(height, width int) bool {
	return gameMap.cellMap[height][width]
}

func (gameMap *GameMap) doForEveryCell(f func (height, width int, gameMap *GameMap)) {
	height, width := gameMap.getSize()

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