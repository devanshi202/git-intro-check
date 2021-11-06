package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate" //to import this search arbitrary expression evaluation in go on google, first gothub file -> copy ling from github -> in the terminal type go get and paste the copied link -> once its added paste it in import
)

func showCalc() {
	// a := app.New()
	// w := a.NewWindow("calculator")

	output:=""
	input:=widget.NewLabel(output)
	historyStr := ""
	history := widget.NewLabel(historyStr)
	var historyArr [] string;
	isHistoryClicked := false

	historyBtn:=widget.NewButton("history", func() {
		if isHistoryClicked{
			historyStr=""
		}else{
			for i:=len(historyArr)-1 ; i>=0; i--{
				historyStr += historyArr[i];
				historyStr += "\n"
			}
		}
		isHistoryClicked = !isHistoryClicked
		history.SetText(historyStr)
	})	
	backBtn:=widget.NewButton("back", func() {
		if len(output)!=0{
			output = output[0 : len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn:=widget.NewButton("clear", func() {
		output=""
		input.SetText(output)
	})
	open:=widget.NewButton("(", func() {
		output += "("
		input.SetText(output)
	})
	close:=widget.NewButton(")", func() {
		output += ")"
		input.SetText(output)
	})
	divide:=widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	seven:=widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	eight:=widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	nine:=widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	multiply:=widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	four:=widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	five:=widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	six:=widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	minus:=widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	one:=widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	two:=widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	three:=widget.NewButton("3", func() {
		output += "3"
		input.SetText(output)
	})
	add:=widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	zero:=widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	dot:=widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	equals:=widget.NewButton("=", func() {
		//these are the functions used from the library we imported from github (arbitrary expression evaluation in go github file)

		expression, err := govaluate.NewEvaluableExpression(output);//this will check if the expression is correct or not, if conrrect will give the expression if not will give the error
		if err == nil{
			result, err := expression.Evaluate(nil); //map is passed in paranthesis but there is not need in this condition; map for 89+4+a+b(example by sir)

			if err==nil{
				//output = result // we need to convert it into string
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64) // this is used to convert the expression into string
				strToAppend := output + " = "+ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			}else{
				output= "error"
			}
		}else{
			output= "error"
		}
		input.SetText(output)
	
	})
	
	
	// hello := widget.NewLabel("Hello Fyne!")
	calcContainer:= container.NewVBox(
		container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					historyBtn,
					backBtn,
				),
				container.NewGridWithColumns(4,
					clearBtn,
					open,
					close,
					divide,
				),
				container.NewGridWithColumns(4,
					seven,
					eight,
					nine,
					multiply,
				),
				container.NewGridWithColumns(4,
					four,
					five,
					six,
					minus,
				),
				container.NewGridWithColumns(4,
					one,
					two,
					three,
					add,
				),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
						zero,
						dot,
					),
					equals,
				),
			),
		),
	)
	
	w := myApp.NewWindow("Calc")
	w.Resize(fyne.NewSize(500, 280))


	w.SetContent(
		container.NewBorder(homeBtn, nil, nil, nil, calcContainer),
	)

	w.Show()
}