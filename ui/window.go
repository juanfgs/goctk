package ui


type Window struct {
	Widget
	Box
	Title string
}

func NewWindow(title string, x, y, width, height int) Window {
	window := Window{}

	window.Widget = Widget{}
	window.Widget.SetPositionX(x)
	window.Widget.SetPositionY(y)


	window.Box =  NewBox(x, y, width, height)
	window.SetTitle(title)

	return window
}

func (this *Window) SetTitle(title string) {
	this.Draw()
	this.Title = title
	this.xPrint(this.GetHorizontalCenter()-1, this.Widget.PositionY,  "╣")
	this.xPrint(this.GetHorizontalCenter(), this.Widget.PositionY,  title)
	this.xPrint(this.GetHorizontalCenter()+len(title), this.Widget.PositionY,  "╠")
	this.Flush()
}
