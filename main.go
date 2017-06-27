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

func handleKeys(key int) {
	/*Function handleKeys allows to control player character*/
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

func loopOver() {
	/*Function loopOver is main loop of the game.*/
	blt.Print(4, 7, "Hello, r/roguelikedev!")
	for {
		blt.Refresh()
		key := blt.Read()
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
	blt.Clear()
}
