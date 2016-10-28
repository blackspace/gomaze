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
	temp:=int64(x)<<32+int64(y)
	p._data.Push(temp)
}

func (p *PointStack)Pop() (x,y int) {
	if v,ok:=p._data.Pop();ok{
		temp:=v.(int64)

		tx:=temp>>32
		ty:=temp&0xffffffff

		return int(tx),int(ty)
	} else {
		panic("This point stack is empty.")
	}
}

func (p *PointStack)Last() (x,y int,ok bool){
	v,ok:=p._data.Peek()
	temp:=v.(int64)
	x=int(temp>>32)
	y=int(temp&0xffffffff)

	return

}

func (p *PointStack)HasPoint(x,y int) bool {
	temp:=int64(x)<<32+int64(y)
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
		temp:=v.(int64)
		x:=int(temp>>32)
		y:=int(temp&0xffffffff)

		result=append(result,Point{x,y})
	}

	return
}
