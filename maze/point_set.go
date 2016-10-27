package maze

import "strconv"

type PointSet struct {
	_data []int64
}

func NewPointSet() *PointSet {
	return &PointSet{_data:make([]int64,0,1<<10)}
}

func (p *PointSet)Add(x,y int) {
	temp:=int64(x)<<32+int64(y)
	if !p.HasPoint(x,y) {
		p._data =append(p._data,temp)
	}
}


func (p *PointSet)HasPoint(x,y int) bool {
	temp:=int64(x)<<32+int64(y)
	for i:=0;i<len(p._data);i++ {
		if temp==p._data[i] {
			return true
		}
	}

	return false
}

func (p *PointSet)Index(i int) (x,y int) {
	temp:=p._data[i]

	x=int(temp>>32)
	y=int(temp&0xffffffff)

	return
}

func (p *PointSet)Count() int {
	return len(p._data)
}

func (p *PointSet)Join(jp *PointSet) {
	for i:=0;i<jp.Count();i++ {
		x,y:=jp.Index(i)
		p.Add(x,y)
	}
}

func (p *PointSet)Equal(tp *PointSet) bool {
	if len(p._data)!=len(tp._data) {
		return false
	}

	for i:=0;i<len(p._data);i++ {
		x,y:=p.Index(i)

		if !tp.HasPoint(x,y) {
			return false
		}
	}

	for i:=0;i<len(tp._data);i++ {
		x,y:=tp.Index(i)

		if !p.HasPoint(x,y) {
			return false
		}
	}

	return true
}

func (p *PointSet)String() (result string) {
	for i:=0;i<p.Count();i++ {
		x,y:=p.Index(i)
		result=result+"("+strconv.Itoa(x)+","+strconv.Itoa(y)+")"
	}
	return  result
}