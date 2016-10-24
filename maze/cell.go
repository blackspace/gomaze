package maze

type Cell struct {
	_data byte
}

func NewCell(d byte) *Cell {
	return &Cell{d}
}

func (c *Cell)HasTop()  bool {
	return c._data&0x08==8
}


func (c *Cell)HasRight()  bool {
	return c._data&0x4 ==4
}

func (c *Cell)HasBottom()  bool {
	return c._data&0x2==2
}

func (c *Cell)HasLeft()  bool {
	return c._data&0x1 == 1
}

func (c *Cell)EraseLeft() {
	c._data =c._data&0xfe
}

func (c *Cell)EraseRight() {
	c._data =c._data&0xfb
}


func (c *Cell)EraseTop() {
	c._data =c._data&0xf7
}

func (c *Cell)EraseBottom() {
	c._data =c._data&0xfd
}

func (c *Cell)IsClosed() bool {
	return c._data==0xff || c._data==0x0f
}


