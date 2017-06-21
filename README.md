# This game is part of r/RoguelikeDev Does The Complete Roguelike Tutorial game jam hosted by [aaron_ds](https://www.reddit.com/user/aaron_ds) on [roguelikedev](www.reddit/com/r/roguelikedev)

![RoguelikeDev Does the Complete Roguelike Tutorial Event Logo](https://i.imgur.com/ksc9EW3.png)

## Basic info

I'm using [Go programming language](https://golang.org/) and [BearLibTerminal](https://bitbucket.org/cfyzium/bearlibterminal/overview) library.

## Compilation

1. Get and install Go programming language and GCC / MinGW (architecture version must match, ie 64-bit Go needs 64-bit GCC)
2. Download [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) (again, architecture version must match)
3. Put BearLibTerminal's .so, or .dll and .lib, or .dylib into GCC/lib directory, and into game sources directory
4. Execute $go build main.go