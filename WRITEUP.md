# r/roguelikedev

## Introduction.

`first draft`

What do you need to know?

First of all, I'm not experienced Go developer. I suppose that my code will be not idiomatic, against good practices, etc. My primary language is Python. I want to seize the opportunity, and use architecture I know well - roguebasin python+libtcod roguelike tutorial - to learn something about Go. Although, my idea is *sameness* - I want to keep my code as close to original one as possible. But why use BearLibTerminal instead of *canonical* libtcod in that case? Because I want to, deal with it.

It's not supposed to be tutorial. Just some writeups.

## Part 0 - set environment up.

`first draft`

### You need some components to start. 

`first draft`

The first one is [Go programming language](https://golang.org/). I will use 1.8.3 version. Developers promises that there were no compatibility breaking changes during Go 1 development, but they removed one package from standard library recently, so... Be careful. Go is available for Windows, Linux, OSX and FreeBSD.  
You need to have set some environmental variables: GOROOT and GOPATH (your workspace). If you need help, google is full of materials.

The second is modern C compiler. It's necessary because BearLibTerminal use CGO (method of calling C code into Go). For Windows, there is variety of compilers with own quirks. TDM-GCC is safe bet. If you are on Linux, probably you have modern GCC already installed.

The third is [BearLibTerminal library](https://bitbucket.org/cfyzium/bearlibterminal). You'd need BearLibTerminal.go and BearLibTerminal.h. The latter is C header needed by CGO. You need BearLibTerminal go and h files into your GOPATH; .dll and .lib, or .dylib, or .so in lib directory of your GCC.

Architecture (32 or 64 bit) must be the same for all components.

### What's different?

`first draft`

Well, it's easier to say what Python and Go have in common, because these languages is really different. But regarding to tutorial structure, most important difference is about OOP. Python is language that uses objects extensively, `everything is object`, and supports large amount of OOP mechanics, like inheritance and polymorphism. Go has no classes, using interfaces and types instead.

Libtcod is, say, *full* roguelike toolkit. It provides display handling, FOV and pathfinding algorithms, etc.  
BearLibTerminal is all about display, but handles it in different manner; for example, there is native support for True / Open Type Fonts, while libtcod relies on bitmaps.

### Let's print some dummy code.

`first draft`

I mean, it's time to check if we set environment properly.

You need font file in `media` directory. Start by creating empty `main.go` file. Now, you root folder should looks like (windows example) that:    

    base_directory
    |
    |-media
    | |-Lato-Heavy.ttf
    |
    |-BearTerminal.dll
    |-BearLibTerminal.lib
    |-main.go

Every Go file needs starts by keyword `package` and have `main` function that it place when application execution starts. It's possible to add `init` function that starts before `main`, but it's option.

BearLibTerminal need to be initialized by Open console and Set properties. Syntax is a bit weird, because it uses strings for setting variables. The base variables of console are: width, height, window title, font, and font size.

So, we'll start with base skeleton of file with blt window set, so

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

    func loopOver() {
        ...
    }

    func main() {
        /*Main starts main loop and closes terminal if loop breaks*/
        loopOver()
        blt.Close()
    }

    func init() {
        /*It's app initialization.
        Starts by setting blt console properties*/
        blt.Open() //open blt console
        //convert given ints into strings for blt
        sizeX, sizeY := strconv.Itoa(windowSizeX), strconv.Itoa(windowSizeY)
        //then set string values for blt functions
        size := "size=" + sizeX + "x" + sizeY
        title := "title='" + gameTitle + "'"
        window := "window: " + size + "," + title
        fontSize := "size=" + strconv.Itoa(baseFontSize)
        font := "font: " + baseFont + ", " + fontSize
        blt.Set(window + "; " + font)
        blt.Clear()
}

We'll add dummy main loop later, now I'd like to describe code above. 

So, code starts with package name declaration. `main` is, well, mail file.  
Imports are grouped by type - first one is package from standard library (`strconv` stands for string conversion), the second one is bearlibterminal renamed to `blt` for that document.  
I stored blt console datas as constants. BearLib uses strings for settings so I could declare windowSizeX as `"30"` instead of int, but I think that declaration is more readable.  
Decided to use func `init` to set blt console.  
First, we need to convert all variables into strings, as I wrote earlier. We are using `strconvItoa` for converting integers into strings.  
`:=` stands for short variable declaration. In that case, there is no type declaration. That syntax element is possible to use only in scope of function.  
So, we are assigning blt-specified string syntax into local variables, then use them into blt.Set function.  
Full, hardcoded version of that Set would look like  
`blt.Set("window: size=30x15, title='r/roguelikedev'; font: media/Lato-Heavy.ttf, size=20")`  
As last step of `init` we are clearing whole blt.Console, just to be sure.  
BearLibTerminal use system of Layers. `Clear()` flushes them all.  
After `init` function, `main` starts automatically. For now, it just starts main loop then, when that loop breaks, closes BearLib console.

`loopOver`, ie main game loop, is really simple for now.

    func loopOver() {
        /*Function loopOver is main loop of the game.*/
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

Starts by printing "Hello, r/roguelikedev!" text, then starts potentially infinite loop.  
BearLibTerminal doesn't refresh own console automatically, you need to call Refresh() every time when you want to do it. Next, we are checking for events. TK_CLOSE means triggering 'X' on the window bar. In that case, loop breaks, and `blt.Close` from func `main` executes, closing console.

#### at the end...

`first draft`

As you could notice, Go has own quirks. For example, where are all semicolons? They are hidden. It means that they are present for compiler, but you don't need to write them as element of your code. Another small strange thing is Go style guide. Everything is written as mixedCase, even variables. But why sometimes first letter of function name is small and sometimes big? It's matter of export. `fooBar` is accessible only from scope of that document; `FooBar` is exported function and it is accessible from other files.