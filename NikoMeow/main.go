package main

import (
	"embed"
	"io"
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
)

//go:embed assets/*
var assets embed.FS

var meowFiles = []string{
	"cat_1.ogg",
	"cat_2.ogg",
	"cat_3.ogg",
}

func playRandomMeow() {
	fileName := meowFiles[rand.Intn(len(meowFiles))]

	f, err := assets.Open("assets/sounds/" + fileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		log.Println("Decode error:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func main() {
	a := app.New()
	w := a.NewWindow("Niko :3")
	w.Resize(fyne.NewSize(200, 250))

	f, err := assets.Open("assets/images/niko.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	niko := fyne.NewStaticResource("niko.png", data)
	w.SetIcon(niko)

	// img := canvas.NewImageFromReader(f, "niko.png")
	img := canvas.NewImageFromResource(niko)
	img.ScaleMode = canvas.ImageScalePixels
	img.FillMode = canvas.ImageFillContain

	btn := widget.NewButton(":3", func() {
		go playRandomMeow()
	})
	btn.Resize(fyne.NewSize(btn.MinSize().Height, btn.MinSize().Height))

	content := container.NewBorder(nil, container.NewHBox(layout.NewSpacer(), btn, layout.NewSpacer()), nil, nil, img)
	w.SetContent(content)

	w.ShowAndRun()
}
