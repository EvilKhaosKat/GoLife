package main

import (
	"bytes"
)

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
	width := len(gameMap.cellMap)
	if width == 0 {
		panic("Width must not be 0")
	}

	height := len(gameMap.cellMap[0])
	if height == 0 {
		panic("Height must not be 0")
	}

	return width, height
}

func (gameMap *GameMap) SetValue(width, height int, value bool) {
	gameMap.cellMap[width][height] = value
}

func (gameMap *GameMap) GetValue(width, height int) bool {
	return gameMap.cellMap[width][height]
}

//simple linear algorithm
func (gameMap *GameMap) Update() {
	width, height := gameMap.GetSize()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
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
		gameMap.SetValue(changeEvent.width, changeEvent.height, changeEvent.newValue)
	}
}

func (gameMap *GameMap) addChange(width, height int, newValue bool) {
	//TODO concurrent safety to be added
	changeEvents := gameMap.changeEvents
	changeEvents = append(changeEvents, ChangeEvent{width, height, newValue})
	gameMap.changeEvents = changeEvents
}

func cellContinueToLive(aliveNeighboursCount int) bool {
	return aliveNeighboursCount == 2 || aliveNeighboursCount == 3
}

func (gameMap *GameMap) getAliveNeighboursCount(width, height int) int {
	neighbours := gameMap.getCellNeighbours(width, height)

	aliveNeighboursCount := 0

	for _, cellAlive := range neighbours {
		if cellAlive {
			aliveNeighboursCount++
		}
	}

	return aliveNeighboursCount
}

func (gameMap *GameMap) getCellNeighbours(width, height int) []bool {
	var neighbours []bool

	neighbours = gameMap.addNeighbour(neighbours, width-1, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width-1, height)
	neighbours = gameMap.addNeighbour(neighbours, width-1, height + 1)

	neighbours = gameMap.addNeighbour(neighbours, width, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width, height + 1)

	neighbours = gameMap.addNeighbour(neighbours, width + 1, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width + 1, height)
	neighbours = gameMap.addNeighbour(neighbours, width + 1, height + 1)

	return neighbours
}
func (gameMap *GameMap) addNeighbour(neighbours []bool, width, height int) []bool {
	mapWidth, mapHeight := gameMap.GetSize()

	possibleHeight := checkPossibleValue(mapHeight, height)
	possibleWidth := checkPossibleValue(mapWidth, width)

	if possibleHeight && possibleWidth {
		neighbours = append(neighbours, gameMap.GetValue(width, height)) //TODO not efficient copying of array
	}

	return neighbours
}
func checkPossibleValue(basicValue int, checkValue int) bool {
	return checkValue > 0 && checkValue < basicValue
}

func (gameMap *GameMap) DoForEveryCell(f func(width, height int, gameMap *GameMap)) {
	height, width := gameMap.GetSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			f(i, j, gameMap)
		}
	}
}

func (gameMap *GameMap)String() string{
	var buffer bytes.Buffer

	width, height := gameMap.GetSize()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			cellAlive := gameMap.GetValue(i, j)
			if cellAlive {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
		buffer.WriteByte(13)
		buffer.WriteByte(10)
	}

	return buffer.String()
}

type ChangeEvent struct {
	width, height int
	newValue      bool
}