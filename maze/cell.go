package maze

type Cell struct {
	Data byte
}

func NewCell(d byte) *Cell {
	return &Cell{d}
}

func (c *Cell)HasTop()  bool {
	side:=c.Data&0x08

	return side==8
}


func (c *Cell)HasRight()  bool {
	side:=c.Data&0x4

	return side==4
}

func (c *Cell)HasBottom()  bool {
	side:=c.Data&0x2

	return side==2
}

func (c *Cell)HasLeft()  bool {
	side:=c.Data&0x1

	return side==1
}

func (c *Cell)EraseLeft() {
	c.Data=c.Data&0xfe
}

func (c *Cell)EraseRight() {
	c.Data=c.Data&0xfb
}


func (c *Cell)EraseTop() {
	c.Data=c.Data&0xf7
}

func (c *Cell)EraseBottom() {
	c.Data=c.Data&0xfd
}

func (c *Cell)IsClosed() bool {
	return c.Data==0xff || c.Data==0x0f
}


