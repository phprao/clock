package main

import (
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/robfig/cron/v3"
)

var a fyne.App

func main() {
	DemoClock()
}

func DemoClock() {
	a = app.New()
	cr := cron.New()
	// 周一到周五，下午14点40分
	if _, err := cr.AddFunc("40 14 * * 1-5", func() {
		CreateWindow()
	}); err != nil {
		log.Fatal(err)
	}

	cr.Start()

	a.Run()
}

func CreateWindow() {
	st := time.Now()

	w := a.NewWindow("Clock")
	w.Resize(fyne.NewSize(320, 180))
	w.CenterOnScreen()

	img := canvas.NewImageFromResource(resourceClockJpg)
	img.FillMode = canvas.ImageFillOriginal
	content1 := container.New(layout.NewCenterLayout(), container.New(layout.NewGridWrapLayout(fyne.NewSize(80, 80)), img))

	clock := canvas.NewText("", color.Black)
	content2 := container.New(layout.NewCenterLayout(), clock)
	clock.Text = time.Now().Format("15:04:05")
	clock.TextSize = 20
	clock.Refresh()

	empty := canvas.NewText("", color.Black)

	w.SetContent(container.New(layout.NewVBoxLayout(), content1, empty, content2))

	w.Show()

	go func() {
		for range time.Tick(time.Second) {
			clock.Text = time.Now().Format("15:04:05")
			clock.Refresh()
			if time.Now().Sub(st) >= 5*time.Minute {
				w.Hide()
				break
			}
		}
	}()
}
