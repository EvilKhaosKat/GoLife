package main

// Game map.
// Coordinates formatted as:
// 0123
// 1
// 2
// 3
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

//simple linear algorithm
func (gameMap *GameMap) Update() {
	height, width := gameMap.GetSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			aliveNeighboursCount := gameMap.getAliveNeighboursCount(i, j)

			cellAlive := gameMap.GetValue(i, j)
			if cellAlive {
				if !cellContinueToLive(aliveNeighboursCount) {
					gameMap.addChange(i, j, false)
				}
			} else {
				if aliveNeighboursCount == 3 {
					gameMap.addChange(i, j, true)
				}
			}
		}
	}

	gameMap.performAllChanges()
}

func (gameMap *GameMap) performAllChanges() {
	for _, changeEvent := range gameMap.changeEvents {
		gameMap.SetValue(changeEvent.height, changeEvent.width, changeEvent.newValue)
	}
}

func (gameMap *GameMap) addChange(height, width int, newValue bool) {
	//TODO concurrent safety to be added
	changeEvents := gameMap.changeEvents
	changeEvents = append(changeEvents, ChangeEvent{height, width, newValue})
	gameMap.changeEvents = changeEvents
}

func cellContinueToLive(aliveNeighboursCount int) bool {
	return aliveNeighboursCount == 2 || aliveNeighboursCount == 3
}

func (gameMap *GameMap) getAliveNeighboursCount(height, width int) int {
	neighbours := gameMap.getCellNeighbours(height, width)

	aliveNeighboursCount := 0

	for _, cellAlive := range neighbours {
		if cellAlive {
			aliveNeighboursCount++
		}
	}

	return aliveNeighboursCount
}

func (gameMap *GameMap) getCellNeighbours(height, width int) []bool {
	var neighbours []bool

	neighbours = gameMap.addNeighbour(neighbours, height - 1, width - 1)
	neighbours = gameMap.addNeighbour(neighbours, height - 1, width)
	neighbours = gameMap.addNeighbour(neighbours, height - 1, width + 1)

	neighbours = gameMap.addNeighbour(neighbours, height, width - 1)
	neighbours = gameMap.addNeighbour(neighbours, height, width + 1)

	neighbours = gameMap.addNeighbour(neighbours, height + 1, width - 1)
	neighbours = gameMap.addNeighbour(neighbours, height + 1, width)
	neighbours = gameMap.addNeighbour(neighbours, height + 1, width + 1)

	return neighbours
}
func (gameMap *GameMap)addNeighbour(neighbours []bool, height, width int) []bool {
	mapHeight, mapWidth := gameMap.GetSize()

	possibleHeight := checkPossibleValue(mapHeight, height)
	possibleWidth := checkPossibleValue(mapWidth, width)

	if possibleHeight && possibleWidth {
		neighbours = append(neighbours, gameMap.GetValue(height, width)) //TODO not efficient copying of array
	}

	return neighbours
}
func checkPossibleValue(basicValue int, checkValue int) bool {
	return checkValue > 0 && checkValue < basicValue
}

func (gameMap *GameMap) DoForEveryCell(f func(height, width int, gameMap *GameMap)) {
	height, width := gameMap.GetSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			f(i, j, gameMap)
		}
	}
}

type ChangeEvent struct {
	height, width int
	newValue      bool
}