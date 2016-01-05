package ui

import(
	"github.com/nsf/termbox-go"
)
type Widget struct {
	PositionX      int
	PositionY      int
	Width  int
	Height int
	Color termbox.Attribute
	BgColor termbox.Attribute
	Children []Drawable
}

type Drawable interface {
	Draw()
	DrawAll()
}

func (this Widget) DrawAll(){
	for _, w := range this.Children {
		w.Draw()
	}
}

func (this *Widget) SetPositionX( x int) {
	this.PositionX = x
}

func (this *Widget) SetPositionY( y int) {
	this.PositionY = y
}

func (this *Widget) SetColor( newcolor termbox.Attribute){
	this.Color = newcolor
}

func (this *Widget) SetBgColor( newcolor termbox.Attribute){
	this.BgColor = newcolor
}

func (this *Widget) AddChild( widget Drawable) {
	this.Children = append( this.Children, widget)
}

func (this Widget) Flush() {
	termbox.Flush()
}

func (this Widget) xPrint(x, y int,  msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, this.Color, this.BgColor)
		x++
	}
}

func (this Widget) yPrint(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, this.Color, this.BgColor)
		y++
	}
}


func repeatChar(char string, times int ) string{
	var result string = ""

	for i := 0; i < times; i++ {
		result = result + char
	}

	return result
}
