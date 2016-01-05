package ui

import(
	"github.com/nsf/termbox-go"

)


type Ui struct {
	e error
}

func NewUi() Ui{
	ui := Ui{}
	ui.e = termbox.Init()
	if ui.e != nil {
		panic(ui.e)
	}

	defer termbox.Close()


loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyCtrlA:
				ui.AddUrlPanel()
				break
			}
		}
	}

	return ui
}

func (this Ui) Draw() {
	termbox.Flush()
}


func (this Ui) AddUrlPanel(){
	window := NewWindow("Add Url",10, 10, 90, 10)
	window.Box.SetColor(termbox.ColorCyan)
	window.Box.SetBgColor(termbox.ColorBlack)

	textBox := NewTextBox("foobar")
	textBox.SetPositionX(13)
	textBox.SetPositionY(11)
	textBox.SetBgColor(termbox.ColorBlack)	
	textBox.SetColor(termbox.ColorGreen)

	textBox2 := NewTextBox("foobar")
	textBox2.SetPositionX(13)
	textBox2.SetPositionY(12)
	textBox2.SetBgColor(termbox.ColorMagenta)	
	textBox2.SetColor(termbox.ColorGreen)

	window.AddChild(textBox)
	window.AddChild(textBox2)	
	window.DrawAll()
}


