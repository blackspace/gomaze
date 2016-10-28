package maze

type Area struct {
	_pointSets []*PointSet
}

func NewArea() *Area{
	return &Area{_pointSets:make([]*PointSet,0,1<<10)}
}

func (a *Area)FindByPoint(x,y int) *PointSet {
	for i:=0;i<len(a._pointSets);i++ {
		if a._pointSets[i].HasPoint(x,y) {
			return a._pointSets[i]
		}
	}
	return nil
}

func (a *Area)HasArea(ps *PointSet) bool {
	for i:=0;i<len(a._pointSets);i++ {
		if ps==a._pointSets[i] {
			return true
		}
	}
	return false
}

func (a *Area)AddArea(ps *PointSet) {
	a._pointSets =append(a._pointSets,ps)
}

func (a *Area)RemoveArea(ps *PointSet) {
	var index int

	for index=0;index<a.Size();index++{
		if a._pointSets[index]==ps {
			break
		}
	}

	start:=a._pointSets[:index]
	end:=a._pointSets[index+1:]

	a._pointSets =append(start,end...)
}

func (a *Area)Size() int {
	return len(a._pointSets)
}



