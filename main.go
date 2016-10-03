package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/gdamore/tcell"
	// "time"
	// "os"
	"fmt"
	"math/rand"
	"time"
)

const DELAY_MS = 200



func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go launchLife()

	//wait for esc or ctrl+q pressed, and then exit
	loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc ||
				ev.Key == termbox.KeyCtrlQ {
				break loop
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func launchLife() {
	gameMap := initGameMap(termbox.Size())
	fmt.Println(gameMap)


	termbox.SetCell(5, 10, '⏣', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(1, 2, '⏺', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(10, 5, '⏹', termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}

func initGameMap(width, height int) *GameMap {
	cellAutoMap := make([][]bool, height)

	for i := 0; i < height; i++ {
		cellAutoMap[i] = make([]bool, width)
	}

	return &GameMap{cellMap: cellAutoMap}
}

func fillMapRandomValues(gameMap *GameMap) {
	rand.Seed(time.Now().UTC().UnixNano())

	height, width := gameMap.getSize()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			gameMap.setValue(i, j, getRandomBoolValue())
		}
	}
}

func getRandomBoolValue() bool {
	randomValue := rand.Intn(2) //0 or 1
	return randomValue == 0
}

func printGameMap(gameMap *GameMap) {
	
}