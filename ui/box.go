package ui


type Box struct {
	PositionX      int
	PositionY      int
	Width  int
	Height int

	// Embedded types
	Widget

}

func NewBox(x, y, width, height int) Box {
	box := Box{}
	box.PositionX = x
	box.PositionY = y
	box.Width = width
	box.Height = height

	return box
}

func (this Box) Draw() {
	this.xPrint(this.PositionX, this.PositionY,  repeatChar("═", this.Width))
	this.yPrint(this.PositionX, this.PositionY,  repeatChar("╔", 1))
	this.yPrint(this.PositionX, this.PositionY+1,  repeatChar("║", this.Height))
	this.yPrint(this.PositionX+this.Width, this.PositionY-1,  repeatChar(" ╗", 1))
	this.yPrint(this.PositionX+this.Width, this.PositionY+1,  repeatChar("║", this.Height))
	this.yPrint(this.PositionX, this.PositionY+this.Height,  repeatChar("╚", 1))
	this.xPrint(this.PositionX+1, this.PositionY+this.Height,  repeatChar("═", this.Width))
	this.yPrint(this.PositionX+this.Width, this.PositionY+this.Height, repeatChar("╝", 1))
	this.Flush()
}



func (this Box) GetHorizontalCenter() int{
	return (this.PositionX + this.Width) / 2
}

func (this Box) GetVerticalCenter() int{
	return (this.PositionY + this.Height) / 2
}
