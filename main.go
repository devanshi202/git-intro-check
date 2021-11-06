package main

import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("Virtual OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var homeBtn fyne.Widget
var img fyne.CanvasObject
var panelContent *fyne.Container


func main(){
	img = canvas.NewImageFromFile("wallpaper.jpg")

	btn1=widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon() ,func ()  {
		showCalc()
	})
	btn2=widget.NewButtonWithIcon("Weather App", theme.InfoIcon() ,func ()  {
		showWeather(myWindow)
	})
	btn3=widget.NewButtonWithIcon("Audio App", theme.RadioButtonIcon() ,func ()  {
		showAudio()
	})
	btn4=widget.NewButtonWithIcon("Text Editor", theme.FileIcon() ,func ()  {
		showText()
	})
	btn5=widget.NewButtonWithIcon("News App", theme.MailComposeIcon() ,func ()  {
		showNews(myWindow)
	})
	btn6=widget.NewButtonWithIcon("Gallery App", theme.StorageIcon() ,func ()  {
		showGallery(myWindow)
	})
	homeBtn=widget.NewButtonWithIcon("Home", theme.HomeIcon() ,func ()  {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(
		container.NewGridWithColumns(7,
			homeBtn,
			btn1,
			btn2,
			btn3,
			btn4,
			btn5,
			btn6,
		),
	)

	myWindow.Resize(fyne.NewSize(1250, 700))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	myWindow.ShowAndRun()
}