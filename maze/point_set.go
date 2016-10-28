package maze

import (
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
)

type PointSet struct {
	_data *treeset.Set
}

func NewPointSet() *PointSet {
	return &PointSet{_data:treeset.NewWith(utils.Int64Comparator)}
}

func (p *PointSet)Add(x,y int) {
	temp:=int64(x)<<32+int64(y)
	if !p.HasPoint(x,y) {
		p._data.Add(temp)
	}
}


func (p *PointSet)HasPoint(x,y int) bool {
	temp:=int64(x)<<32+int64(y)
	return p._data.Contains(temp)
}


func (p *PointSet)Count() int {
	return p._data.Size()
}

func (p *PointSet)Join(jp *PointSet) {
	jp._data.Each(func(index int,value interface{}){
		x:=int((value.(int64))>>32)
		y:=int((value.(int64)&0xffffffff))
		p.Add(x,y)
	})
}



