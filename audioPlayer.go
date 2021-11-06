package main

import (
	"time"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var format beep.Format
var isPaused bool = false
var streamer beep.StreamSeekCloser

func showAudio(){
	// a:= app.New()
	// w:= a.NewWindow("Audio Player")
	// w.Resize(fyne.NewSize(400, 400))

	w := myApp.NewWindow("Calc")
	w.Resize(fyne.NewSize(500, 280))

	img:= canvas.NewImageFromFile("music.png")
	img.FillMode=canvas.ImageFillOriginal


		play:=widget.NewButtonWithIcon("play", theme.MediaPlayIcon(), func ()  {
			if(isPaused==false){
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(streamer)
			}else{
				speaker.Unlock()
				isPaused=false
			}
		})

		pause:=widget.NewButtonWithIcon("pause", theme.MediaPauseIcon(), func ()  {
			isPaused=true
			speaker.Lock()
		})

		stop:=widget.NewButtonWithIcon("stop", theme.MediaStopIcon(), func ()  {
			speaker.Close()
			// or speaker.clear()
		})

	label:= widget.NewLabel("Audio MP3")
	label.Alignment=fyne.TextAlignCenter

	label2:= widget.NewLabel("Play MP3")
	label2.Alignment=fyne.TextAlignCenter

	fileToBePlayed:= widget.NewButton("BROWSE", func ()  {
		fd:= dialog.NewFileOpen(
			func (fr fyne.URIReadCloser, _ error){
				streamer, format, _= mp3.Decode(fr)

				label2.Text = fr.URI().Name()
				label2.Refresh()
			}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
		fd.Show()
	})

	c:= container.NewVBox(label, img, fileToBePlayed, label2, 
		container.NewHBox(
			play,
			pause,
			stop,
		),
	)

	// e:= container.NewBorder(img, nil ,nil, nil, c)

	w.SetContent(container.NewBorder(homeBtn, nil, nil, nil, c),)
	w.Show()
}
