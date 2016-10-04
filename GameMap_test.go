package main

import (
	"testing"
)

const SIZE = 5

func TestEmptyMap(t *testing.T) {
	gameMap := initGameMap(SIZE, SIZE)

	if !isMapEmpty(gameMap) {
		t.Error("GameMap must be empty, but it is:", gameMap)
	}

	gameMap.Update()
	if !isMapEmpty(gameMap) {
		t.Error("GameMap must be empty, but it is:", gameMap)
	}
}


func TestCellBecomeAlive(t *testing.T) {
	gameMap := initGameMap(SIZE, SIZE)

	gameMap.SetValue(2, 1, true)
	gameMap.SetValue(1, 3, true)
	gameMap.SetValue(3, 3, true)
	//00000
	//00100
	//00!00
	//01010
	//00000

	gameMap.Update()
	//00000
	//00000
	//00100
	//00000
	//00000
	if !gameMap.GetValue(2, 2) {
		t.Error("Cell 2,3 must be alive, but it doesn't:", gameMap)
	}
}

func TestSimpleOscillator(t *testing.T) {
	gameMap := initGameMap(SIZE, SIZE)

	gameMap.SetValue(1, 2, true)
	gameMap.SetValue(2, 2, true)
	gameMap.SetValue(3, 2, true)
	//00000
	//00000
	//01110
	//00000
	//00000


	gameMap.Update()
	//00000
	//00100
	//00100
	//00100
	//00000
	if !gameMap.GetValue(2, 1) ||
		!gameMap.GetValue(2, 2) ||
		!gameMap.GetValue(2, 3) {
		t.Error("Oscillator must be vertical, but it doesn't:", gameMap)
	}
}

func isMapEmpty(gameMap *GameMap) bool {
	for i := 0; i < len(gameMap.cellMap); i ++ {
		for j := 0; j < len(gameMap.cellMap[0]); j++ {
			cellAlive := gameMap.GetValue(i, j)
			if cellAlive {
				return false
			}
		}
	}
	return true
}
