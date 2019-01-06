package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ikester/blinkt"
)

// need to find which state we're in on startup to trigger the light.

func off(bl blinkt.Blinkt) {
	bl.Clear()
	bl.Show()
}

func red(bl blinkt.Blinkt) {
	bl.SetAll(255, 0, 0)
	bl.Show()
}

func green(bl blinkt.Blinkt) {
	bl.SetAll(0, 255, 0)
	bl.Show()
}

func blue(bl blinkt.Blinkt) {
	bl.SetAll(0, 0, 255)
	bl.Show()
}

func main() {
	cP := flag.String("c", "o", "color/off: r=red, g=green, b=blue, o=off")
	bP := flag.Float64("b", 1.0, "brightness: 0.0-1.0")
	flag.Parse()
	if *cP != "o" && *cP != "r" && *cP != "g" && *cP != "b" {
		fmt.Println("-c must be one of o, r, g, b")
		os.Exit(1)
	}
	if *bP < 0.0 || *bP > 1.0 {
		fmt.Println("-b must be in range [0.0, 1.0]")
	}
	switch *cP {
	case "o":
		bl := blinkt.NewBlinkt(*bP)
		bl.Setup()
		off(bl)
	case "r":
		bl := blinkt.NewBlinkt(*bP)
		bl.Setup()
		red(bl)
	case "g":
		bl := blinkt.NewBlinkt(*bP)
		bl.Setup()
		green(bl)
	case "b":
		bl := blinkt.NewBlinkt(*bP)
		bl.Setup()
		blue(bl)
	}
}
