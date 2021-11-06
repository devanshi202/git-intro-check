package main

import (
	"io/ioutil"
	"strconv"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
) 

func showText(){
	// a:= app.New()
	var count int =1

	w:= myApp.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(400, 600))

	content:= container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Fyne Text Editor"),
		),
	)

	content.Add(
		widget.NewButton("Add More Files", func ()  {
			content.Add(widget.NewLabel("New File"+strconv.Itoa(count)))
			count++
		}),
	)

	input := widget.NewMultiLineEntry() // in place of new entry -> new multiline entry
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn:= widget.NewButton("Save", func ()  {
		saveDialogBox := dialog.NewFileSave(
			func( ur fyne.URIWriteCloser, _ error){
				textContent:= []byte(input.Text)
				ur.Write(textContent)
			},
			w,
		)
		count--
		saveDialogBox.SetFileName("New File"+strconv.Itoa(count)+".txt")
		saveDialogBox.Show()
	})

	openBtn := widget.NewButton("Open File", func(){
		openDialogBox:= dialog.NewFileOpen(
			func(rf fyne.URIReadCloser, _ error){
				readData,_ := ioutil.ReadAll(rf) //reading data from rf
				output:= fyne.NewStaticResource("New File", readData) //making a new static resource on our app to for the retrived from rf 
				viewData:= widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))
				
				w:=fyne.CurrentApp().NewWindow(
					string(output.StaticName)) //name of the window

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(400,400))

				w.Show()
			},w,
		)

		openDialogBox.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		openDialogBox.Show()

	})
 // TASK: open vali window me save button add krna hai that will edit the saved file with new changes(important)

	textContainer:= container.NewVBox(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	w.SetContent(container.NewBorder(homeBtn, nil, nil , nil, textContainer),)
	w.Show()
}
