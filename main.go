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
	blt "bearlibterminal"
	"strconv"
)

const (
	window_x_size  = 30
	window_y_size  = 15
	game_title     = "r/roguelikedev"
	base_font      = "media/Lato-Heavy.ttf"
	base_font_size = 20
)

func loopOver() {
	blt.Print(4, 7, "Hello, r/roguelikedev!")
	for {
		blt.Refresh()
		key := blt.Read()
		if key != blt.TK_CLOSE {
			continue
		} else {
			break
		}
	}
}

func main() {
	loopOver()
	blt.Close()
}

func init() {
	/*app initialization
	starts by setting blt console properties*/
	blt.Open()
	size_x, size_y := strconv.Itoa(window_x_size), strconv.Itoa(window_y_size)
	size := "size=" + size_x + "x" + size_y
	title := "title='" + game_title + "'"
	window := "window: " + size + "," + title
	font_size := "size=" + strconv.Itoa(base_font_size)
	font := "font: " + base_font + ", " + font_size
	blt.Set(window + "; " + font)
	blt.Clear()
}
