package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/gdamore/tcell"
	// "time"
	// "os"
	"math/rand"
	"time"
)

const DELAY_MS = 150 * time.Millisecond

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
	fillMapRandomValues(gameMap)
	printGameMap(gameMap)

	timer := time.NewTimer(DELAY_MS)

	for {
		gameMap.Update()

		<-timer.C //wait until timer expire, usually longer than updateTheWorld
		timer.Reset(DELAY_MS)

		printGameMap(gameMap)
	}

	//termbox.SetCell(5, 10, '⏣', termbox.ColorWhite, termbox.ColorBlack)
	//termbox.SetCell(1, 2, '⏺', termbox.ColorWhite, termbox.ColorBlack)
	//termbox.SetCell(10, 5, '⏹', termbox.ColorWhite, termbox.ColorBlack)
	//termbox.Flush()
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

	width, height := gameMap.GetSize()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			gameMap.SetValue(i, j, getRandomBoolValue())
		}
	}
}

func getRandomBoolValue() bool {
	randomValue := rand.Intn(10)
	return randomValue == 0
}

func printGameMap(gameMap *GameMap) {
	width, height  := gameMap.GetSize()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			cellAlive := gameMap.GetValue(i, j)
			printGameMapCell(i, j, cellAlive)
		}
	}

	termbox.Flush()
}

func printGameMapCell(height, width int, cellAlive bool) {
	if cellAlive {
		termbox.SetCell(height, width, '█', termbox.ColorWhite, termbox.ColorBlack)
	} else {
		termbox.SetCell(height, width, '█', termbox.ColorBlack, termbox.ColorWhite)
	}
}