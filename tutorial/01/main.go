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
	windowSizeX  = 30
	windowSizeY  = 15
	gameTitle    = "r/roguelikedev"
	baseFont     = "media/Lato-Heavy.ttf"
	baseFontSize = 20
)

var (
	playerX = windowSizeX / 2
	playerY = windowSizeY / 2
)

type Object struct {
	layer int
	x, y  int
	char  string
	color string
}

func (obj *Object) draw() {
	/* draw is method that prints Objects
	   on specified positions on specified layer.*/
	blt.Layer(obj.layer)
	ch := "[color=" + obj.color + "]" + obj.char
	blt.Print(obj.x, obj.y, ch)
}

func (obj *Object) clear() {
	/* clear is method that clears area starting from coords on specific layer.*/
	blt.Layer(obj.layer)
	blt.ClearArea(obj.x, obj.y, 1, 1)
}

func (obj *Object) move(dx, dy int) {
	/* move is method for handling objects movement;
	   it receives pointer to object, and adds arguments to object values.*/
	obj.x += dx
	obj.y += dy
}

func handleKeys(key int) {
	/*Function handleKeys allows to control player character
	by reading input from main loop*/
	if key == blt.TK_UP {
		playerY--
	} else if key == blt.TK_DOWN {
		playerY++
	} else if key == blt.TK_LEFT {
		playerX--
	} else if key == blt.TK_RIGHT {
		playerX++
	}
}

func printPlayer() {
	/*Function printPlayer clears old player position on its layer
	then prints player symbol on new coords*/
	blt.Layer(0)
	blt.ClearArea(0, 0, windowSizeX, windowSizeY)
	blt.Print(playerX, playerY, "@")
}

func loopOver() {
	/* Function loopOver is main loop of the game.*/
	printPlayer()
	for {
		blt.Refresh()
		key := blt.Read()
		if key == blt.TK_CLOSE || key == blt.TK_ESCAPE {
			break
		} else {
			handleKeys(key)
			printPlayer()
		}
	}
}

func main() {
	/* Function main initializes main loop;
	   when loop breaks, closes blt console.*/
	loopOver()
	blt.Close()
}

func init() {
	/* Function init is app initialization.
	   Sets BearLibTerminal console properties.*/
	blt.Open()
	sizeX, sizeY := strconv.Itoa(windowSizeX), strconv.Itoa(windowSizeY)
	size := "size=" + sizeX + "x" + sizeY
	title := "title='" + gameTitle + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(baseFontSize)
	font := "font: " + baseFont + ", " + fontSize
	blt.Set(window + "; " + font)
	blt.Clear()
}
