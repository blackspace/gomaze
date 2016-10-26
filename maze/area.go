package maze

type Area struct {
	PointSets []*PointSet
}

func NewArea() *Area{
	return &Area{PointSets:make([]*PointSet,0,1<<10)}
}

func (a *Area)FindByPoint(x,y int) *PointSet {
	for i:=0;i<len(a.PointSets);i++ {
		if a.PointSets[i].HasPoint(x,y) {
			return a.PointSets[i]
		}
	}
	return nil
}

func (a *Area)HasArea(ps *PointSet) bool {
	for i:=0;i<len(a.PointSets);i++ {
		if ps==a.PointSets[i] {
			return true
		}
	}
	return false
}

func (a *Area)AddArea(ps *PointSet) {
	a.PointSets=append(a.PointSets,ps)
}



