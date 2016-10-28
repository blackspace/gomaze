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
	if !p.HasPoint(x,y) {
		p._data.Add(XYToInt64(x,y))
	}
}


func (p *PointSet)HasPoint(x,y int) bool {
	return p._data.Contains(XYToInt64(x,y))
}


func (p *PointSet)Count() int {
	return p._data.Size()
}

func (p *PointSet)Join(jp *PointSet) {
	jp._data.Each(func(index int,value interface{}){
		x,y:= Int64ToXY(value.(int64))
		p.Add(x,y)
	})
}



