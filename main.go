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
	player  *Object
	objects []*Object
)

type Object struct {
	layer int
	x, y  int
	char  string
	color string
}

func (obj *Object) move(dx, dy int) {
	/*move is method for handling objetcs movement;
	it receives pointer to object, and adds arguments to object values*/
	obj.x += dx
	obj.y += dy
}

func (obj *Object) draw(layer, x, y int, ch string) {
	/*It is method that prints Objects
	on specified positions on specified layer*/
	blt.Layer(layer)
	blt.Print(x, y, ch)
}

func (obj *Object) clear(layer, x, y, w, h int) {
	/*It is method that clears Objects position on specified layer*/
	blt.Layer(layer)
	blt.ClearArea(x, y, w, h)
}

func handleKeys(key int) {
	/*Function handleKeys allows to control player character
	by reading input from main loop*/
	if key == blt.TK_UP {
		player.y--
	} else if key == blt.TK_DOWN {
		player.y++
	} else if key == blt.TK_LEFT {
		player.x--
	} else if key == blt.TK_RIGHT {
		player.x++
	}
}

func loopOver() {
	/*Function loopOver is main loop of the game.*/
	for {
		blt.Refresh()
		key := blt.Read()
		for i := 0; i < len(objects); i++ {
			n := objects[i]
			n.clear(n.layer, n.x, n.y, 1, 1)
		}
		if key == blt.TK_CLOSE || key == blt.TK_ESCAPE {
			break
		} else {
			handleKeys(key)
		}
		for j := 0; j < len(objects); j++ {
			n := objects[j]
			n.draw(n.layer, n.x, n.y, n.char)
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
	blt.Clear()
	player = &Object{0, windowSizeX / 2, windowSizeY / 2, "@", "white"}
	objects = append(objects, player)
}
