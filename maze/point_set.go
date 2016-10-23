package maze

type PointSet struct {
	_data []int
}

func NewPointSet() *PointSet {
	return &PointSet{_data:make([]int,0,1000)}
}

func (p *PointSet)Add(x,y int) {
	if !p.HasPoint(x,y) {
		p._data =append(p._data,x,y)
	}
}


func (p *PointSet)HasPoint(x,y int) bool {
	for i:=0;i<len(p._data);i=i+2 {
		px:=p._data[i]
		py:=p._data[i+1]

		if px==x && py==y {
			return true
		}
	}

	return false
}