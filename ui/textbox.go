package ui

type TextBox struct {
	Widget
	Value string
}


func NewTextBox(defaultValue string) TextBox{
	tb := TextBox{}
	tb.Value = defaultValue
	return tb
}

func (this TextBox) Draw(){

	this.xPrint(this.PositionX, this.PositionY, this.Value )

	this.Flush()
}
