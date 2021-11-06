package main

import (
	// "fmt"
    "io/ioutil"
     "log"
	 "strings"
	 "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func showGallery(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(400, 300)) 
	// var picsArr []string
	root_dir := "C:\\Users\\Devanshi Sharma\\OneDrive\\Desktop\\go_gallery"
	files, err := ioutil.ReadDir(root_dir)
	if err != nil {
        log.Fatal(err)
    }
	tabs := container.NewAppTabs()
	for _, f := range files {
		// fmt.Println(f.Name(), f.IsDir())
		if !f.IsDir(){
			extension := strings.Split(f.Name(), ".")[1]
			if extension == "png" || extension == "jpg"{
				image := canvas.NewImageFromFile(root_dir+"\\"+f.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(f.Name(), image),)
			}
		}
	}
	//viewing image in fyne using canvas:
		// img := canvas.NewImageFromFile(picsArr[0])
		
	// tabs := container.NewAppTabs(
	// 	// for i:=0; i<len(picsArr); i++{ // for loop won't worl here 
	// 	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
			
	// 	// }
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// )
	// so we will make for loop and append images in tabs
	// for i:=0; i<len(picsArr); i++{ // for loop won't worl here 
	// tabs.Append(container.NewTabItem("image", canvas.NewImageFromFile(picsArr[i])),)
			
	// }

	//few properties of tabs:
		// container.TabLocationTop,
		// container.TabLocationBottom,
		// container.TabLocationLeading,
		// container.TabLocationTrailing

	//get tabs vertically aligned
		//tabs.SetTabLocation(container.TabLocationLeading)

	//some properties to adjust pics in canvas:
		//img1.FillMode = canvas.ImageFillOriginal
		//image := image.NewAlpha(image.Rectangle{image.Point{0, 0}, image.Point{100, 100}})
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs))
	w.Show()
}

