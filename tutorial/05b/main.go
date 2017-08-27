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
	"math"
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
	maxMonsters  = 3
	gameTitle    = "r/roguelikedev"
	baseFont     = "media/Lato-Heavy.ttf"
	baseFontSize = 10
	fovRays      = 360
	fovLength    = 5
	fovStep      = 3
)

var (
	player  *Object
	objects []*Object
	board   [][]*Tile
	sinBase = []float64{
		0.00000, 0.01745, 0.03490, 0.05234, 0.06976, 0.08716, 0.10453,
		0.12187, 0.13917, 0.15643, 0.17365, 0.19081, 0.20791, 0.22495, 0.24192,
		0.25882, 0.27564, 0.29237, 0.30902, 0.32557, 0.34202, 0.35837, 0.37461,
		0.39073, 0.40674, 0.42262, 0.43837, 0.45399, 0.46947, 0.48481, 0.50000,
		0.51504, 0.52992, 0.54464, 0.55919, 0.57358, 0.58779, 0.60182, 0.61566,
		0.62932, 0.64279, 0.65606, 0.66913, 0.68200, 0.69466, 0.70711, 0.71934,
		0.73135, 0.74314, 0.75471, 0.76604, 0.77715, 0.78801, 0.79864, 0.80902,
		0.81915, 0.82904, 0.83867, 0.84805, 0.85717, 0.86603, 0.87462, 0.88295,
		0.89101, 0.89879, 0.90631, 0.91355, 0.92050, 0.92718, 0.93358, 0.93969,
		0.94552, 0.95106, 0.95630, 0.96126, 0.96593, 0.97030, 0.97437, 0.97815,
		0.98163, 0.98481, 0.98769, 0.99027, 0.99255, 0.99452, 0.99619, 0.99756,
		0.99863, 0.99939, 0.99985, 1.00000, 0.99985, 0.99939, 0.99863, 0.99756,
		0.99619, 0.99452, 0.99255, 0.99027, 0.98769, 0.98481, 0.98163, 0.97815,
		0.97437, 0.97030, 0.96593, 0.96126, 0.95630, 0.95106, 0.94552, 0.93969,
		0.93358, 0.92718, 0.92050, 0.91355, 0.90631, 0.89879, 0.89101, 0.88295,
		0.87462, 0.86603, 0.85717, 0.84805, 0.83867, 0.82904, 0.81915, 0.80902,
		0.79864, 0.78801, 0.77715, 0.76604, 0.75471, 0.74314, 0.73135, 0.71934,
		0.70711, 0.69466, 0.68200, 0.66913, 0.65606, 0.64279, 0.62932, 0.61566,
		0.60182, 0.58779, 0.57358, 0.55919, 0.54464, 0.52992, 0.51504, 0.50000,
		0.48481, 0.46947, 0.45399, 0.43837, 0.42262, 0.40674, 0.39073, 0.37461,
		0.35837, 0.34202, 0.32557, 0.30902, 0.29237, 0.27564, 0.25882, 0.24192,
		0.22495, 0.20791, 0.19081, 0.17365, 0.15643, 0.13917, 0.12187, 0.10453,
		0.08716, 0.06976, 0.05234, 0.03490, 0.01745, 0.00000, -0.01745, -0.03490,
		-0.05234, -0.06976, -0.08716, -0.10453, -0.12187, -0.13917, -0.15643,
		-0.17365, -0.19081, -0.20791, -0.22495, -0.24192, -0.25882, -0.27564,
		-0.29237, -0.30902, -0.32557, -0.34202, -0.35837, -0.37461, -0.39073,
		-0.40674, -0.42262, -0.43837, -0.45399, -0.46947, -0.48481, -0.50000,
		-0.51504, -0.52992, -0.54464, -0.55919, -0.57358, -0.58779, -0.60182,
		-0.61566, -0.62932, -0.64279, -0.65606, -0.66913, -0.68200, -0.69466,
		-0.70711, -0.71934, -0.73135, -0.74314, -0.75471, -0.76604, -0.77715,
		-0.78801, -0.79864, -0.80902, -0.81915, -0.82904, -0.83867, -0.84805,
		-0.85717, -0.86603, -0.87462, -0.88295, -0.89101, -0.89879, -0.90631,
		-0.91355, -0.92050, -0.92718, -0.93358, -0.93969, -0.94552, -0.95106,
		-0.95630, -0.96126, -0.96593, -0.97030, -0.97437, -0.97815, -0.98163,
		-0.98481, -0.98769, -0.99027, -0.99255, -0.99452, -0.99619, -0.99756,
		-0.99863, -0.99939, -0.99985, -1.00000, -0.99985, -0.99939, -0.99863,
		-0.99756, -0.99619, -0.99452, -0.99255, -0.99027, -0.98769, -0.98481,
		-0.98163, -0.97815, -0.97437, -0.97030, -0.96593, -0.96126, -0.95630,
		-0.95106, -0.94552, -0.93969, -0.93358, -0.92718, -0.92050, -0.91355,
		-0.90631, -0.89879, -0.89101, -0.88295, -0.87462, -0.86603, -0.85717,
		-0.84805, -0.83867, -0.82904, -0.81915, -0.80902, -0.79864, -0.78801,
		-0.77715, -0.76604, -0.75471, -0.74314, -0.73135, -0.71934, -0.70711,
		-0.69466, -0.68200, -0.66913, -0.65606, -0.64279, -0.62932, -0.61566,
		-0.60182, -0.58779, -0.57358, -0.55919, -0.54464, -0.52992, -0.51504,
		-0.50000, -0.48481, -0.46947, -0.45399, -0.43837, -0.42262, -0.40674,
		-0.39073, -0.37461, -0.35837, -0.34202, -0.32557, -0.30902, -0.29237,
		-0.27564, -0.25882, -0.24192, -0.22495, -0.20791, -0.19081, -0.17365,
		-0.15643, -0.13917, -0.12187, -0.10453, -0.08716, -0.06976, -0.05234,
		-0.03490, -0.01745, -0.00000,
	}
	cosBase = []float64{
		1.00000, 0.99985, 0.99939, 0.99863, 0.99756, 0.99619, 0.99452,
		0.99255, 0.99027, 0.98769, 0.98481, 0.98163, 0.97815, 0.97437, 0.97030,
		0.96593, 0.96126, 0.95630, 0.95106, 0.94552, 0.93969, 0.93358, 0.92718,
		0.92050, 0.91355, 0.90631, 0.89879, 0.89101, 0.88295, 0.87462, 0.86603,
		0.85717, 0.84805, 0.83867, 0.82904, 0.81915, 0.80902, 0.79864, 0.78801,
		0.77715, 0.76604, 0.75471, 0.74314, 0.73135, 0.71934, 0.70711, 0.69466,
		0.68200, 0.66913, 0.65606, 0.64279, 0.62932, 0.61566, 0.60182, 0.58779,
		0.57358, 0.55919, 0.54464, 0.52992, 0.51504, 0.50000, 0.48481, 0.46947,
		0.45399, 0.43837, 0.42262, 0.40674, 0.39073, 0.37461, 0.35837, 0.34202,
		0.32557, 0.30902, 0.29237, 0.27564, 0.25882, 0.24192, 0.22495, 0.20791,
		0.19081, 0.17365, 0.15643, 0.13917, 0.12187, 0.10453, 0.08716, 0.06976,
		0.05234, 0.03490, 0.01745, 0.00000, -0.01745, -0.03490, -0.05234, -0.06976,
		-0.08716, -0.10453, -0.12187, -0.13917, -0.15643, -0.17365, -0.19081,
		-0.20791, -0.22495, -0.24192, -0.25882, -0.27564, -0.29237, -0.30902,
		-0.32557, -0.34202, -0.35837, -0.37461, -0.39073, -0.40674, -0.42262,
		-0.43837, -0.45399, -0.46947, -0.48481, -0.50000, -0.51504, -0.52992,
		-0.54464, -0.55919, -0.57358, -0.58779, -0.60182, -0.61566, -0.62932,
		-0.64279, -0.65606, -0.66913, -0.68200, -0.69466, -0.70711, -0.71934,
		-0.73135, -0.74314, -0.75471, -0.76604, -0.77715, -0.78801, -0.79864,
		-0.80902, -0.81915, -0.82904, -0.83867, -0.84805, -0.85717, -0.86603,
		-0.87462, -0.88295, -0.89101, -0.89879, -0.90631, -0.91355, -0.92050,
		-0.92718, -0.93358, -0.93969, -0.94552, -0.95106, -0.95630, -0.96126,
		-0.96593, -0.97030, -0.97437, -0.97815, -0.98163, -0.98481, -0.98769,
		-0.99027, -0.99255, -0.99452, -0.99619, -0.99756, -0.99863, -0.99939,
		-0.99985, -1.00000, -0.99985, -0.99939, -0.99863, -0.99756, -0.99619,
		-0.99452, -0.99255, -0.99027, -0.98769, -0.98481, -0.98163, -0.97815,
		-0.97437, -0.97030, -0.96593, -0.96126, -0.95630, -0.95106, -0.94552,
		-0.93969, -0.93358, -0.92718, -0.92050, -0.91355, -0.90631, -0.89879,
		-0.89101, -0.88295, -0.87462, -0.86603, -0.85717, -0.84805, -0.83867,
		-0.82904, -0.81915, -0.80902, -0.79864, -0.78801, -0.77715, -0.76604,
		-0.75471, -0.74314, -0.73135, -0.71934, -0.70711, -0.69466, -0.68200,
		-0.66913, -0.65606, -0.64279, -0.62932, -0.61566, -0.60182, -0.58779,
		-0.57358, -0.55919, -0.54464, -0.52992, -0.51504, -0.50000, -0.48481,
		-0.46947, -0.45399, -0.43837, -0.42262, -0.40674, -0.39073, -0.37461,
		-0.35837, -0.34202, -0.32557, -0.30902, -0.29237, -0.27564, -0.25882,
		-0.24192, -0.22495, -0.20791, -0.19081, -0.17365, -0.15643, -0.13917,
		-0.12187, -0.10453, -0.08716, -0.06976, -0.05234, -0.03490, -0.01745,
		-0.00000, 0.01745, 0.03490, 0.05234, 0.06976, 0.08716, 0.10453, 0.12187,
		0.13917, 0.15643, 0.17365, 0.19081, 0.20791, 0.22495, 0.24192, 0.25882,
		0.27564, 0.29237, 0.30902, 0.32557, 0.34202, 0.35837, 0.37461, 0.39073,
		0.40674, 0.42262, 0.43837, 0.45399, 0.46947, 0.48481, 0.50000, 0.51504,
		0.52992, 0.54464, 0.55919, 0.57358, 0.58779, 0.60182, 0.61566, 0.62932,
		0.64279, 0.65606, 0.66913, 0.68200, 0.69466, 0.70711, 0.71934, 0.73135,
		0.74314, 0.75471, 0.76604, 0.77715, 0.78801, 0.79864, 0.80902, 0.81915,
		0.82904, 0.83867, 0.84805, 0.85717, 0.86603, 0.87462, 0.88295, 0.89101,
		0.89879, 0.90631, 0.91355, 0.92050, 0.92718, 0.93358, 0.93969, 0.94552,
		0.95106, 0.95630, 0.96126, 0.96593, 0.97030, 0.97437, 0.97815, 0.98163,
		0.98481, 0.98769, 0.99027, 0.99255, 0.99452, 0.99619, 0.99756, 0.99863,
		0.99939, 0.99985, 1.00000,
	}
)

type Object struct {
	layer  int
	x, y   int
	char   string
	name   string
	color  string
	blocks bool
}

type Tile struct {
	explored    bool
	blocked     bool
	blocksSight bool
}

type Rect struct {
	x, y int
	w, h int
}

func (obj *Object) move(dx, dy int) {
	/* move is method for handling objects movement;
	   it receives pointer to object, then checks cell for blocked field,
	   and adds arguments to object values if tile is passable*/
	if isBlocked(obj.x+dx, obj.y+dy) == false {
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
	/*center is method that gets center cell of room*/
	centerX := (room.x + (room.x + room.h)) / 2
	centerY := (room.y + (room.y + room.w)) / 2
	return centerX, centerY
}

func (room *Rect) intersect(other *Rect) bool {
	/*intersect is method that checks by coordinates comparison
	if rooms (room and other) are not overlapping;
	returns true or false*/
	cond1 := (room.x <= other.x+other.w)
	cond2 := (room.x+room.w >= other.x)
	cond3 := (room.y <= other.y+other.h)
	cond4 := (room.y+room.h >= other.y)
	return (cond1 && cond2 && cond3 && cond4)
}

func min(a, b int) int {
	/*Function min returns smaller of two integers*/
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	/*Function max returns bigger of two integers*/
	if a > b {
		return a
	}
	return b
}

func round64(value, rounding float64, places int) float64 {
	/*Function round64 rounds float64 values (value) to specified
	number of digits (places) using given point-of-rounding-up (rounding)*/
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)
	var round float64
	if value > 0 {
		if div >= rounding {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div > rounding {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}
	return round / pow
}

func round64ToInt(value float64) int {
	/*Function round64ToInt gets float64 value, uses round64 function,
	then returns new value converted to integer*/
	a := round64(value, 0.5, 0)
	return int(a)
}

func randIntRange(a, b int) int {
	/*func randIntRange returns random integer withing specified range;
	uses rand.Intn(n) from standard library that returns [0, n)*/
	return rand.Intn(b-a) + a
}

func isBlocked(x, y int) bool {
	/*Function isBlocked checks if map cell is blocked by wall or object;
	returns true if cell's blocked field is set to true;
	iterates through objects slice and
	returns true if object's blocks field is set to true and its coordinates
	matches function arguments;
	otherwise, returns false*/
	if board[x][y].blocked == true {
		return true
	}
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		if obj.blocks == true && obj.x == x && obj.y == y {
			return true
		}
	}
	return false
}

func placeObjects(room *Rect) {
	/*Function placeObjects places monsters within room borders;
	it gets random number of monsters to place, then for every monster:
	- draws x, y coordinates
	- decides type of monster
	- adds monster to objects slice*/
	numMonsters := rand.Intn(maxMonsters + 1)
	var monster *Object
	for i := 0; i < numMonsters; i++ {
		x := randIntRange(room.x+1, room.x+room.w)
		y := randIntRange(room.y+1, room.y+room.h)
		if isBlocked(x, y) == false {
			if rand.Intn(100+1) <= 80 {
				monster = &Object{1, x, y, "o", "orc", "dark green", true}
			} else {
				monster = &Object{1, x, y, "T", "troll", "darker green", true}
			}
			objects = append(objects, monster)
		}
	}
}

func createRoom(room *Rect) {
	/*Function createRoom uses Rect struct for
	marking specific area as passable;
	takes initial [x][y]cell and width, height of room,
	then iterates through map*/
	for x := room.x + 1; x < room.x+room.w; x++ {
		for y := room.y + 1; y < room.y+room.h; y++ {
			board[x][y].blocked = false
			board[x][y].blocksSight = false
		}
	}
}

func horizontalTunnel(x1, x2, y int) {
	/*Function horizontalTunnel carves passable area
	from x1 to x2 on y row*/
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		board[x][y].blocked = false
		board[x][y].blocksSight = false
	}
}

func verticalTunnel(y1, y2, x int) {
	/*Function verticalTunnel carves passable area
	from y1 to y2 on x column*/
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		board[x][y].blocked = false
		board[x][y].blocksSight = false
	}
}

func makeMap() {
	/*Function makeMap creates dungeon map by:
	- creating empty 2d array then filling it by Tiles;
	- creating new room that doesn't overlap other rooms;
	- connects rooms using tunnels*/
	var rooms []*Rect
	newMap := make([][]*Tile, mapSizeX)
	for i := range newMap {
		newMap[i] = make([]*Tile, mapSizeY)
	}
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			newMap[x][y] = &Tile{false, true, true}
		}
	}
	board = newMap
	numRooms := 0
	for i := 0; i < maxRooms; i++ {
		w := randIntRange(roomMinSize, roomMaxSize)
		h := randIntRange(roomMinSize, roomMaxSize)
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
				if rand.Intn(1+1) == 1 {
					horizontalTunnel(prevX, newX, prevY)
					verticalTunnel(prevY, newY, newX)
				} else {
					verticalTunnel(prevY, newY, prevX)
					horizontalTunnel(prevX, newX, newY)
				}
				placeObjects(newRoom)
			}
			rooms = append(rooms, newRoom)
			numRooms++
		}
	}
}

func castRays() {
	/*func castRays is simple raycasting function for turning tiles to explored;
	it cast (fovRays / fovStep) rays (bigger fovStep means faster but
	more error-prone raycasting) from player to coordinates in fovLength range;
	source of algorithm:
	http://www.roguebasin.com/index.php?title=Raycasting_in_python [20170712]*/
	for i := 0; i < fovRays; i += fovStep {
		rayX := sinBase[i]
		rayY := cosBase[i]
		x := float64(player.x)
		y := float64(player.y)
		board[round64ToInt(x)][round64ToInt(y)].explored = true
		for j := 0; j < fovLength; j++ {
			x -= rayX
			y -= rayY
			if x < 0 || y < 0 || x > windowSizeX-1 || y > windowSizeY-1 {
				break
			}
			bx, by := round64ToInt(x), round64ToInt(y)
			board[bx][by].explored = true
			if board[bx][by].blocked {
				break
			}
		}
	}
}

func isInFOV(sx, sy, tx, ty int) bool {
	/*checks if target (tx, ty) is in fov of source (sx, sy);
	returns true if tx, ty == sx, sy; otherwise, it casts (fovRays / fovStep)
	rays (bigger fovStep means faster but more error-prone algorithm)
	from source to tiles in fovLength range; stops if cell is blocked;
	source of algorithm:
	http://www.roguebasin.com/index.php?title=Raycasting_in_python [20170712]*/
	if sx == tx && sy == ty {
		return true
	}
	for i := 0; i < fovRays; i += fovStep {
		rayX := sinBase[i]
		rayY := cosBase[i]
		x := float64(sx)
		y := float64(sy)
		for j := 0; j < fovLength; j++ {
			x -= rayX
			y -= rayY
			if x < 0 || y < 0 || x > windowSizeX-1 || y > windowSizeY-1 {
				break
			}
			bx, by := round64ToInt(x), round64ToInt(y)
			if bx == tx && by == ty {
				return true
			}
			if board[bx][by].blocked {
				break
			}
		}
	}
	return false
}

func renderAll() {
	/*Function renderAll handles display;
	clears all layers of blt console, initializes raycasting,
	and sets current layer to the bottom one;
	draws floors and walls with regard to board[x][y] *Tile, then
	use (obj *Object) draw() method with list of game objects*/
	blt.Clear()
	castRays()
	blt.Layer(0)
	for y := 0; y < mapSizeY; y++ {
		for x := 0; x < mapSizeX; x++ {
			if board[x][y].explored == true {
				if isInFOV(player.x, player.y, x, y) {
					if board[x][y].blocked == true {
						txt := "[color=colorLightWall]#"
						blt.Print(x, y, txt)
					} else {
						txt := "[color=colorLightGround]."
						blt.Print(x, y, txt)
					}
				} else {
					if board[x][y].blocked == true {
						txt := "[color=colorDarkWall]#"
						blt.Print(x, y, txt)
					} else {
						txt := "[color=colorDarkGround]."
						blt.Print(x, y, txt)
					}
				}
			}
		}
	}
	for j := 0; j < len(objects); j++ {
		n := objects[j]
		if isInFOV(player.x, player.y, n.x, n.y) {
			n.draw()
		}
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
	rand.Seed(time.Now().Unix())
	blt.Open()
	sizeX, sizeY := strconv.Itoa(windowSizeX), strconv.Itoa(windowSizeY)
	size := "size=" + sizeX + "x" + sizeY
	title := "title='" + gameTitle + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(baseFontSize)
	font := "font: " + baseFont + ", " + fontSize
	blt.Set(window + "; " + font)
	blt.Set("palette: colorLightWall = #826E32, colorDarkWall = #000064, " +
		"colorLightGround = #C8B432, colorDarkGround = #323296")
	blt.Clear()
	player = &Object{2, 0, 0, "@", "player", "white", true}
	objects = append(objects, player)
	makeMap()
}
