package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/gdamore/tcell"
	"time"
)

type GameMap struct {
	CellAutoMap [][]bool
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	//gameMap := initGameMap(termbox.Size())
	//fmt.Println(gameMap)
	termbox.SetCell(5, 10, '⏣', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(1, 2, '⏺', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(10, 5, '⏹', termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()

	time.Sleep(time.Second*3)
}

func initGameMap(width, height int) GameMap {
	cellAutoMap := make([][]bool, height)

	for i := 0; i < height; i++ {
		cellAutoMap[i] = make([]bool, width)
	}

	return GameMap{cellAutoMap}
}