package main

import (
	"bytes"
	"sync"
	"runtime"
)

var WORKERS_COUNT int = 0

// Game map.
// Coordinates formatted as:
// 0123
// 1
// 2
// 3
type GameMap struct {
	cellMap      [][]bool
	changeEvents []*ChangeEvent
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

	return width, height
}

func (gameMap *GameMap) SetValue(width, height int, value bool) {
	gameMap.cellMap[height][width] = value
}

func (gameMap *GameMap) GetValue(width, height int) bool {
	return gameMap.cellMap[height][width]
}

//simple linear algorithm
func (gameMap *GameMap) Update() {
	_, height := gameMap.GetSize()

	rowsChan := make(chan int, height)
	changesChan := make(chan *ChangeEvent)
	finished := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(height) //every row to be handled
	go waitForWorkFinished(finished, &wg)

	createRowWorkers(gameMap, rowsChan, changesChan, &wg)
	go addRowsWorks(height, rowsChan)

	select {
	case changeEvent := <-changesChan:
		gameMap.addChange(changeEvent)
	case <-finished:
		break
	}

	gameMap.performAllChanges()
}

func waitForWorkFinished(finished chan bool, wg *sync.WaitGroup) {
	wg.Wait()
	finished <- true
}

func addRowsWorks(height int, rowsChan chan int) {
	for rowNum := 0; rowNum < height; rowNum++ {
		rowsChan <- rowNum
	}
}

func createRowWorkers(gameMap *GameMap, rowsChan chan int, changesChan chan *ChangeEvent, wg *sync.WaitGroup) {
	for i := 0; i < getWorkersCount(); i++ {
		go rowWorker(gameMap, rowsChan, changesChan, wg)
	}
}

func getWorkersCount() int {
	if WORKERS_COUNT == 0 {
		WORKERS_COUNT = runtime.NumCPU()
	}

	return WORKERS_COUNT
}

func rowWorker(gameMap *GameMap, rowsChan chan int, changesChan chan *ChangeEvent, wg *sync.WaitGroup) {
	for rowNumber := range rowsChan {
		handleRow(gameMap, rowNumber, changesChan)
		wg.Done()
	}
}

func handleRow(gameMap *GameMap, rowNum int, changesChan chan *ChangeEvent) {
	row := gameMap.cellMap[rowNum]
	rowLength := len(row)

	for xCoordinate := 0; xCoordinate < rowLength; xCoordinate++ {
		handleCell(gameMap, xCoordinate, rowNum, changesChan)
	}
}

func handleCell(gameMap *GameMap, xCoordinate int, rowNum int, changesChan chan *ChangeEvent) {
	aliveNeighboursCount := gameMap.getAliveNeighboursCountEff(xCoordinate, rowNum)

	cellAlive := gameMap.GetValue(xCoordinate, rowNum)
	if cellAlive {
		if !cellContinueToLive(aliveNeighboursCount) {
			changesChan <- &ChangeEvent{xCoordinate, rowNum, false}
		}
	} else {
		if aliveNeighboursCount == 3 {
			changesChan <- &ChangeEvent{xCoordinate, rowNum, true}
		}
	}
}

func (gameMap *GameMap) performAllChanges() {
	for _, changeEvent := range gameMap.changeEvents {
		gameMap.SetValue(changeEvent.width, changeEvent.height, changeEvent.newValue)
	}
	gameMap.changeEvents = nil
}

func (gameMap *GameMap) addChange(changeEvent *ChangeEvent) {
	//TODO concurrent safety to be added
	changeEvents := gameMap.changeEvents
	changeEvents = append(changeEvents, changeEvent)
	gameMap.changeEvents = changeEvents
}

func cellContinueToLive(aliveNeighboursCount int) bool {
	return aliveNeighboursCount == 2 || aliveNeighboursCount == 3
}

func (gameMap *GameMap) getAliveNeighboursCountEff(width, height int) int {
	aliveNeighbours := 0

	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width - 1, height - 1))
	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width - 1, height))
	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width - 1, height + 1))

	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width, height - 1))
	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width, height + 1))

	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width + 1, height - 1))
	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width + 1, height))
	aliveNeighbours += oneIfTrue(gameMap.isNeighbourAlive(width + 1, height + 1))

	return aliveNeighbours
}

func oneIfTrue(is bool) int {
	if is {
		return 1
	} else {
		return 0
	}
}

func (gameMap *GameMap) isNeighbourAlive(width, height int) bool {
	mapWidth, mapHeight := gameMap.GetSize()

	possibleHeight := checkPossibleValue(mapHeight, height)
	possibleWidth := checkPossibleValue(mapWidth, width)

	if possibleHeight && possibleWidth {
		return gameMap.GetValue(width, height)
	}

	return false
}

//obsolete code
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

//obsolete code
func (gameMap *GameMap) getCellNeighbours(width, height int) []bool {
	var neighbours []bool

	neighbours = gameMap.addNeighbour(neighbours, width - 1, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width - 1, height)
	neighbours = gameMap.addNeighbour(neighbours, width - 1, height + 1)

	neighbours = gameMap.addNeighbour(neighbours, width, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width, height + 1)

	neighbours = gameMap.addNeighbour(neighbours, width + 1, height - 1)
	neighbours = gameMap.addNeighbour(neighbours, width + 1, height)
	neighbours = gameMap.addNeighbour(neighbours, width + 1, height + 1)

	return neighbours
}

//obsolete code
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
	return checkValue >= 0 && checkValue < basicValue
}

func (gameMap *GameMap) DoForEveryCell(f func(width, height int, gameMap *GameMap)) {
	width, height := gameMap.GetSize()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			f(i, j, gameMap)
		}
	}
}

func (gameMap *GameMap) String() string {
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