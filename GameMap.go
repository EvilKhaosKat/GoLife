package main

type GameMap struct {
	cellMap      [][]bool
	changeEvents []ChangeEvent
}

func (gameMap *GameMap) getSize() (int, int) {
	height := len(gameMap)
	if height == 0 {
		panic("Height must not be 0")
	}

	width := len(gameMap[0])
	if width == 0 {
		panic("Width must not be 0")
	}

	return height, width
}

func (gameMap *GameMap) setValue(height, width int, value bool) {
	gameMap.cellMap[height][width] = value
}

type ChangeEvent struct {
	x,y int
	newValue bool
}