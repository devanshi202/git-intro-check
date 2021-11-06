package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showNews(w fyne.Window){
	// a:= app.New()
	// w:= a.NewWindow("News App")
	// w.Resize(fyne.NewSize(400,400))
	URL:= fmt.Sprintf("https://gnews.io/api/v4/top-headlines?token=5dcae75049c183ac470c5cf29c795189&lang=en&country=in&max=50")

	res, err:= http.Get(URL)

	if err!=nil{
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err:= ioutil.ReadAll(res.Body)

	if err!=nil{
		fmt.Println(err)
	}

	//app heading
	label1:=canvas.NewText("News App", color.Black)
	label1.Alignment = fyne.TextAlignCenter
    label1.TextStyle = fyne.TextStyle{Bold: true}

	//app icon
	img:= canvas.NewImageFromFile("news.png")
	img.FillMode=canvas.ImageFillOriginal


	news, _:=UnmarshalNews(body)

	fmt.Println(news)
	var count int =0
	// total number of articles
	label2:= widget.NewLabel(fmt.Sprintf("No. Of Articles:%d", news.TotalArticles))

	label3:= widget.NewLabel(fmt.Sprintf("%s", news.Articles[count].Title))
	label3.TextStyle=fyne.TextStyle{Bold:true}
	label3.Wrapping=fyne.TextWrapBreak

	entry1:= widget.NewLabel(fmt.Sprintf("%s", news.Articles[count].Description))
	entry1.Wrapping=fyne.TextWrapBreak

	btn:= widget.NewButton("Next", func ()  {
		count++
		label3.Text= news.Articles[count].Title
		entry1.Text= news.Articles[count].Description
		label3.Refresh()
        entry1.Refresh()
	})

	e:= container.NewVBox(label1, label3, entry1, btn)
	e.Resize(fyne.NewSize(300, 300))

	c:= container.NewBorder(img, label2, nil, nil, e)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, c))
	w.Show()


}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()


func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	TotalArticles int64     `json:"totalArticles"`
	Articles      []Article `json:"articles"`     
}

type Article struct {
	Title       string `json:"title"`      
	Description string `json:"description"`
	Content     string `json:"content"`    
	URL         string `json:"url"`        
	Image       string `json:"image"`      
	PublishedAt string `json:"publishedAt"`
	Source      Source `json:"source"`     
}

type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"` 
}

