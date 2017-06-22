This game is part of r/RoguelikeDev Does The Complete Roguelike Tutorial game jam hosted by /u/aaron_ds on www.reddit/com/r/roguelikedev

1. Basic info
I'm using Go programming language (www.golang.org) and BearLibTerminal library (https://bitbucket.org/cfyzium/bearlibterminal/).
My intention is to mimic roguebasin python+libtcod roguelike tutorial ( http://www.roguebasin.com/index.php?title=Complete_Roguelike_Tutorial,_using_python%2Blibtcod ) as close as possible.  
Unfortunately, I didn't manage to set up libtcod bindings.

2. Compilation
    - Get and install Go programming language and GCC / MinGW (architecture version must match, ie 64-bit Go needs 64-bit GCC)
    - Download BearLibTerminal from http://foo.wyrd.name/en:bearlibterminal (again, architecture version must match)
    - Put BearLibTerminal's .so, or .dll and .lib, or .dylib into GCC/lib directory, and into game sources directory
    - Copy BearLibTerminal.go and BearLibTerminal.h into your GOPATH.
    - Execute $go build main.go