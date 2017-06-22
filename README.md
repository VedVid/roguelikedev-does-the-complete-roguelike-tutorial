# This game is part of r/RoguelikeDev Does The Complete Roguelike Tutorial game jam hosted by [aaron_ds](https://www.reddit.com/user/aaron_ds) on [roguelikedev](www.reddit/com/r/roguelikedev)

![RoguelikeDev Does the Complete Roguelike Tutorial Event Logo](https://i.imgur.com/ksc9EW3.png)

## Basic info

I'm using [Go programming language](https://golang.org/) and [BearLibTerminal](https://bitbucket.org/cfyzium/bearlibterminal/overview) library.  
My intention is to mimic [roguebasin python+libtcod roguelike tutorial](http://www.roguebasin.com/index.php?title=Complete_Roguelike_Tutorial,_using_python%2Blibtcod) as close as possible.  
Unfortunately, I didn't manage to set up libtcod bindings.

## Compilation

1. Get and install Go programming language and GCC / MinGW (architecture version must match, ie 64-bit Go needs 64-bit GCC)
2. Download [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) (again, architecture version must match)
3. Put BearLibTerminal's .so, or .dll and .lib, or .dylib into GCC/lib directory, and into game sources directory
4. Copy BearLibTerminal.go and BearLibTerminal.h into your GOPATH.
5. Execute $go build main.go

## Other participants that use Go

[jcerise](https://github.com/jcerise/roguelikedev-does-the-complete-roguelike-tutorial)  
[dqnx](https://github.com/dqnx/roguelikedev-does-the-complete-roguelike-tutorial)