//weather app api -> json data-> apna data
//quick start will be used to read json data (it will make an structure that we will use to implement this task)
//"http://api.openweathermap.org/data/2.5/weather?q=delhi&APPID=41dcf0299a4288361ffd37668541e49d"
package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http" // package for requesting on api

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showWeather(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(400, 300))
		
	
	//api work
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=delhi&APPID=5146ed6c14e69fee3ee21165db937031")

	if err!=nil{
		fmt.Println(err)
	}

	defer res.Body.Close() //res me jo bhi body aegi usko close jisse uspe kaam ho sake
	//using ioutil to read the body
	body, err:= ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err)
	}
	//dropdown selector
	combo := widget.NewSelect([]string{"delhi", "mumbai"}, func(value string) {
		log.Println("Select set to", value)
	})
	//img
	img := canvas.NewImageFromFile("weather app img.png")
	img.FillMode = canvas.ImageFillOriginal
	//the api link in http.get gives json data but in fyne we need an structure to read that data which we will get by quicktype
		//paste api link on browser
		//paste the json data in quickfix -> openquicktype -> paste data -> pick json -> pick go -> copy paste the code provided by quickfix below the main func
		//declare variable and equal it to unmarshalweather function
	weather,err := UnmarshalWeather(body)
	
	label1:= canvas.NewText("weather details", color.Black)
	label1.TextStyle = fyne.TextStyle{Bold:true}


	label2 := canvas.NewText(fmt.Sprintf("country%s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("humidity%.2f", weather.Main.Humidity), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("temp_max%.2f", weather.Main.TempMax), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("temp_min%.2f", weather.Main.TempMin), color.Black)
	label6 := canvas.NewText(fmt.Sprintf("wind_speed%.2f", weather.Wind.Speed), color.Black)

	weatherContainer:= container.NewVBox(
		label1,
		img,
		combo,
		label2,
		label3,
		label4,
		label5,
		label6,
		container.NewGridWithColumns(1,),
	)

	w.SetContent(
		container.NewBorder(panelContent, nil, nil , nil, weatherContainer),
	)
	w.Show()
}
 // for creating dropdown (useful when making this app dynamic)
 //combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		//log.Println("Select set to", value)
	//})


	//TASK:
	//for making it dynamic change the country in the api link to the one selected by the userand then get the data by calling it(important)




// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed int64 `json:"speed"`
	Deg   int64 `json:"deg"`  
}
