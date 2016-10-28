package maze

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)


type PointStack struct {
	_data *arraystack.Stack
}

func NewPointStack() *PointStack {
	return &PointStack{_data:arraystack.New()}
}


func (p *PointStack)Push(x,y int) {
	p._data.Push(PointToInt64(x,y))
}

func (p *PointStack)Pop() (x,y int) {
	if v,ok:=p._data.Pop();ok{
		x,y=Int64ToPoint(v.(int64))

		return
	} else {
		panic("This point stack is empty.")
	}
}

func (p *PointStack)Last() (x,y int,ok bool){
	v,ok:=p._data.Peek()
	x,y=Int64ToPoint(v.(int64))

	return
}

func (p *PointStack)HasPoint(x,y int) bool {
	temp:=PointToInt64(x,y)
	it := p._data.Iterator()
	for it.Next() {
		_, v := it.Index(), it.Value()

		if temp==v {
			return true
		}
	}

	return false
}

func (p *PointStack)Values() (result []Point) {
 	for _,v:=range p._data.Values() {
		x,y:=Int64ToPoint(v.(int64))
		result=append(result,Point{x,y})
	}

	return
}
