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
	"math/rand"
	"strconv"
	"time"

	blt "bearlibterminal"
)

const (
	windowSizeX  = 80
	windowSizeY  = 50
	mapSizeX     = windowSizeX
	mapSizeY     = windowSizeY - 5
	roomMaxSize  = 10
	roomMinSize  = 6
	maxRooms     = 30
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

type Rect struct {
	x, y int
	w, h int
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

func (room *Rect) center() (cx, cy int) {
	centerX := (room.x + (room.x + room.h)) / 2
	centerY := (room.y + (room.y + room.w)) / 2
	return centerX, centerY
}

func (room *Rect) intersect(other *Rect) (overlap bool) {
	cond1 := (room.x <= other.x+other.w)
	cond2 := (room.x+room.w >= other.x)
	cond3 := (room.y <= other.y+other.h)
	cond4 := (room.y+room.h >= other.y)
	return (cond1 && cond2 && cond3 && cond4)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func createRoom(room *Rect) {
	/*Function createRoom uses Rect struct for
	marking specific area as passable;
	takes initial [x][y]cell and width, height of room,
	then iterates through map*/
	for x := room.x + 1; x < room.x+room.w; x++ {
		for y := room.y + 1; y < room.y+room.h; y++ {
			board[x][y].blocked = false
			board[x][y].blocks_sight = false
		}
	}
}

func horizontalTunnel(x1, x2, y int) {
	/*Function horizontalTunnel carves passable area
	from x1 to x2 on y row*/
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		board[x][y].blocked = false
		board[x][y].blocks_sight = false
	}
}

func verticalTunnel(y1, y2, x int) {
	/*Function verticalTunnel carves passable area
	from y1 to y2 on x column*/
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		board[x][y].blocked = false
		board[x][y].blocks_sight = false
	}
}

func makeMap() {
	switch n := rand.Intn(2); n {
	case 0:
		mapBSP()
	case 1:
		mapMazeRuins()
	}
}

func mapBSP() {
	/*Function makeMap creates dungeon map by:
	- creating empty 2d array then filling it by Tiles;
	- creating new room that doesn't overlap other rooms;
	- connetcs rooms using tunnels*/
	var rooms []*Rect
	newMap := make([][]*Tile, mapSizeX)
	for i := range newMap {
		newMap[i] = make([]*Tile, mapSizeY)
	}
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			newMap[x][y] = &Tile{true, true}
		}
	}
	board = newMap
	numRooms := 0
	for i := 0; i < maxRooms; i++ {
		w := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		h := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		x := rand.Intn(mapSizeX - w - 1)
		y := rand.Intn(mapSizeY - h - 1)
		newRoom := &Rect{x, y, w, h}
		failed := false
		for j := 0; j < len(rooms); j++ {
			otherRoom := rooms[j]
			if newRoom.intersect(otherRoom) == true {
				failed = true
				break
			}
		}
		if failed == false {
			createRoom(newRoom)
			newX, newY := newRoom.center()
			if numRooms == 0 {
				player.x = newX
				player.y = newY
			} else {
				prevX, prevY := rooms[numRooms-1].center()
				if rand.Intn(1) == 1 {
					horizontalTunnel(prevX, newX, prevY)
					verticalTunnel(prevY, newY, newX)
				} else {
					verticalTunnel(prevY, newY, prevX)
					horizontalTunnel(prevX, newX, newY)
				}
			}
			rooms = append(rooms, newRoom)
			numRooms++
		}
	}
}

func mapMazeRuins() {
	newMap := make([][]*Tile, mapSizeX)
	for i := range newMap {
		newMap[i] = make([]*Tile, mapSizeY)
	}
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			newMap[x][y] = &Tile{true, true}
		}
	}
	board = newMap
	flip := [2]string{"up", "left"}
	for x := 2; x < mapSizeX-1; x++ {
		for y := 2; y < mapSizeY-1; y++ {
			newFlip := flip[rand.Intn(len(flip))]
			if newFlip == "up" {
				board[x][y-1].blocked = false
				board[x][y-1].blocks_sight = false
			} else if newFlip == "left" {
				board[x-1][y].blocked = false
				board[x-1][y].blocks_sight = false
			}
		}
	}

}

func renderAll() {
	/*Function renderAll handles display;
	draws floors and walls with regard to board[x][y] *Tile, then
	use (obj *Object) draw() method with list of game objects*/
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
		renderAll()
	}
}

func main() {
	/*Function main initializes main loop;
	when loop breaks, closes blt console.*/
	renderAll()
	loopOver()
	blt.Close()
}

func init() {
	/*It's app initialization.
	Starts by setting blt console properties.*/
	rand.Seed(time.Now().Unix())
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
