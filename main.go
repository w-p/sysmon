package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"

	"github.com/w-p/sysmon/pkg"
)

const (
	interval = 1
)

var ticker *time.Ticker
var quit chan bool

func main() {
	systray.Run(setup, exit)
}

func setup() {
	ticker = time.NewTicker(interval * time.Second)
	quit = make(chan bool)

	q := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-q.ClickedCh
		systray.Quit()
	}()

	update()
	go func() {
		for {
			select {
			case <-ticker.C:
				go func() { update() }()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func update() {
	stats := pkg.GenerateStats(interval * time.Second)
	icon := pkg.Render(stats)
	title := fmt.Sprintf("%.2f ms", stats["ping"])

	systray.SetIcon(icon)
	systray.SetTitle(title)
}

func exit() {
	quit <- true
	fmt.Println("quitting...")
}
