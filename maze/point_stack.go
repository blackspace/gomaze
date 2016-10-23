package maze

import "errors"

type PointStack struct {
	_data []int
}

func NewPointStack() *PointStack {
	return &PointStack{_data:make([]int,0,1000)}
}


func (p *PointStack)Push(x,y int) {
	p._data=append(p._data,x,y)
}

func (p *PointStack)Pop() (x,y int) {
	x=p._data[len(p._data)-2]
	y=p._data[len(p._data)-1]
	p._data=p._data[:len(p._data)-2]
	return
}

func (p *PointStack)Last() (x,y int,ok bool){
	if len(p._data)<2 {
		return -10000,-10000,false
	} else {
		x=p._data[len(p._data)-2]
		y=p._data[len(p._data)-1]
		return	x,y,true
	}

}

func (p *PointStack)HasPoint(x,y int) bool {
	for i:=0;i<len(p._data);i=i+2 {
		px:=p._data[i]
		py:=p._data[i+1]

		if px==x && py==y {
			return true
		}
	}

	return false
}

func (p *PointStack)Index(i int) (x,y int) {
	if len(p._data)<2*i+1 {
		panic(errors.New("i is too bigger"))
	}

	x=p._data[2*i]
	y=p._data[2*i+1]

	return
}

func (p *PointStack)Count() int {
	return len(p._data)/2
}