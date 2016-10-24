package maze

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