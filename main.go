package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/gdamore/tcell"
	// "time"
	// "os"
	"fmt"
)

const DELAY_MS = 200

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

type ChangeEvent struct {
	x,y int
	newValue bool
}

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

}