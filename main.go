package main

import (
	"fmt"
	"github.com/ikester/blinkt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"syscall"
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
	fmt.Println("starting up")
	// bl := blinkt.NewBlinkt(1.0)
	bl := blinkt.NewBlinkt(0.05)
	bl.Setup()

	c := cron.New()
	// c.AddFunc("* 0 4 * *", func() { // 4:00am
	// 	fmt.Println("turning red")
	// 	red(bl)
	// })
	c.AddFunc("* 15 7 * *", func() { // 7:00am
		fmt.Println("turning green")
		green(bl)
	})
	c.AddFunc("* 0 9 * *", func() { // 9:00am
		fmt.Println("turning off")
		off(bl)
	})
	// c.AddFunc("* 30 19 * *", func() { // 7:30pm
	// 	fmt.Println("turning blue")
	// 	blue(bl)
	// })
	// c.AddFunc("* 0 20 * *", func() { // 8:00pm
	// 	fmt.Println("turning red")
	// 	red(bl)
	// })
	// c.AddFunc("* 30 20 * *", func() { // 8:30pm
	// 	fmt.Println("turning off")
	// 	off(bl)
	// })
	c.Start()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM, os.Kill)
	<-sigquit

	fmt.Println("shutting down, turning off light")
	off(bl)
	os.Exit(0)
}
