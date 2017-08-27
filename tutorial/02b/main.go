/*
Copyright (c) 2017 Tomasz "VedVid" Nowakowski ( v.v.roguelike@gmail.com )

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

import (
	"strconv"

	blt "bearlibterminal"
)

const (
	windowSizeX  = 80
	windowSizeY  = 50
	mapSizeX     = windowSizeX
	mapSizeY     = windowSizeY - 5
	gameTitle    = "r/roguelikedev"
	baseFont     = "media/Lato-Heavy.ttf"
	baseFontSize = 10
)

var (
	player  *Object
	objects []*Object
	board   [][]*Tile
)

type Object struct {
	layer int
	x, y  int
	char  string
	color string
}

type Tile struct {
	blocked      bool
	blocks_sight bool
}

func (obj *Object) move(dx, dy int) {
	/* move is method for handling objects movement;
	   it receives pointer to object, then checks cell for blocked field,
	   and adds arguments to object values if tile is passable*/
	if board[obj.x+dx][obj.y+dy].blocked == false {
		obj.x += dx
		obj.y += dy
	}
}

func (obj *Object) draw() {
	/*draw is method that prints Objects
	on specified positions on specified layer*/
	blt.Layer(obj.layer)
	ch := "[color=" + obj.color + "]" + obj.char
	blt.Print(obj.x, obj.y, ch)
}

func (obj *Object) clear() {
	/*clear is method that clears area starting from coords on specific layer*/
	blt.Layer(obj.layer)
	blt.ClearArea(obj.x, obj.y, 1, 1)
}

func makeMap() {
	/*Function makeMap creates dungeon map by
	creating empty 2d array then filling it by Tiles*/
	newMap := make([][]*Tile, mapSizeX)
	for i := range newMap {
		newMap[i] = make([]*Tile, mapSizeY)
	}
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			if y == 0 || y == mapSizeY-1 || x == 0 || x == mapSizeX-1 {
				newMap[x][y] = &Tile{true, true}
			} else {
				newMap[x][y] = &Tile{false, false}
			}
		}
	}
	board = newMap
}

func renderAll() {
	/*Function renderAll handles display;
	clears all layers of blt console and sets current layer to the bottom one;
	draws floors and walls with regard to board[x][y] *Tile, then
	use (obj *Object) draw() method with list of game objects*/
	blt.Clear()
	blt.Layer(0)
	for y := 0; y < mapSizeY; y++ {
		for x := 0; x < mapSizeX; x++ {
			if board[x][y].blocked == true {
				txt := "[color=colorDarkWall]#"
				blt.Print(x, y, txt)
			} else {
				txt := "[color=colorDarkGround]."
				blt.Print(x, y, txt)
			}
		}
	}
	for j := 0; j < len(objects); j++ {
		n := objects[j]
		n.draw()
	}
}

func handleKeys(key int) {
	/*Function handleKeys allows to control player character
	by reading input from main loop*/
	if key == blt.TK_UP {
		player.move(0, -1)
	} else if key == blt.TK_DOWN {
		player.move(0, 1)
	} else if key == blt.TK_LEFT {
		player.move(-1, 0)
	} else if key == blt.TK_RIGHT {
		player.move(1, 0)
	}
}

func loopOver() {
	/*Function loopOver is main loop of the game.*/
	for {
		renderAll()
		blt.Refresh()
		key := blt.Read()
		for i := 0; i < len(objects); i++ {
			n := objects[i]
			n.clear()
		}
		if key == blt.TK_CLOSE || key == blt.TK_ESCAPE {
			break
		} else {
			handleKeys(key)
		}
	}
}

func main() {
	/*Function main initializes main loop;
	when loop breaks, closes blt console.*/
	loopOver()
	blt.Close()
}

func init() {
	/*It's app initialization.
	Starts by setting blt console properties.*/
	blt.Open()
	sizeX, sizeY := strconv.Itoa(windowSizeX), strconv.Itoa(windowSizeY)
	size := "size=" + sizeX + "x" + sizeY
	title := "title='" + gameTitle + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(baseFontSize)
	font := "font: " + baseFont + ", " + fontSize
	blt.Set(window + "; " + font)
	blt.Set("palette: colorDarkWall = #000064, colorDarkGround = #323296")
	blt.Clear()
	player = &Object{1, mapSizeX / 2, mapSizeY / 2, "@", "white"}
	npc := &Object{0, mapSizeX/2 - 5, mapSizeY / 2, "@", "yellow"}
	objects = append(objects, player, npc)
	makeMap()
}
