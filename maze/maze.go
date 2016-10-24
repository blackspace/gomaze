package maze

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Maze struct {
	Cells [][]Cell

}

func NewMaze(n int) (m *Maze) {
	m=&Maze{}

	m.Cells=make([]([]Cell),n,n)
	for i:=0;i<n;i++ {
		m.Cells[i]=make([]Cell,n,n)
		for j:=0;j<n;j++ {
			m.Cells[i][j]=Cell{255}
		}
	}
	return
}

func (m *Maze)Len() int {
	return len(m.Cells)
}

func (m *Maze)Get(x,y int) *Cell{
	if x+1>m.Len() || y+1>m.Len() {
		return nil
	}
	if x<0 || y<0 {
		return nil
	}
	return &m.Cells[x][y]
}

func (m *Maze)Has(x,y int) bool {
	if x+1>m.Len() || y+1>m.Len() {
		return false
	}
	if x<0 || y<0 {
		return false
	}
	return true
}


func (m *Maze)GetFirstClosedCell() (x,y int,ok bool){
	for y:=0;y<len(m.Cells);y++ {
		for x:=0;x<len(m.Cells);x++ {
			cell:=m.Get(x,y)
			if cell.IsClosed() {
				return x,y,true
			}
		}
	}

	return
}

func (m *Maze)LeftBottom() (x,y int) {
	return 0,m.Len()-1
}

func (m *Maze)RightTop() (x,y int) {
	return m.Len()-1,0
}

func (m *Maze)IsOpen(x,y,direction int) bool {
	cell:=m.Get(x,y)
	switch direction {
	case UP:
		next_cell:=m.Get(x,y-1)

		if next_cell!=nil {
			if cell.HasTop() || next_cell.HasBottom() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant UP"))
		}
	case DOWN:
		next_cell:=m.Get(x,y+1)

		if next_cell!=nil {
			if cell.HasBottom() || next_cell.HasTop() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant DOWN"))
		}
	case LEFT:
		next_cell:=m.Get(x-1,y)

		if next_cell!=nil {
			if cell.HasLeft() || next_cell.HasRight() {
				return false
			} else {
				return true
			}

		} else {
			panic(errors.New("Cant LEFT"))
		}
	case RIGHT:
		next_cell:=m.Get(x+1,y)

		if next_cell!=nil {
			if cell.HasRight() || next_cell.HasLeft() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant RIGHT"))
		}
	}

	panic(errors.New("The direction must be one of UP DOWN LEFT RIGHT"))

}

func (m *Maze)GetJoiningPointSet(cx,cy int,ps *PointSet){
	ps.Add(cx,cy)

	if m.Has(cx,cy-1) && m.IsOpen(cx,cy,UP) && !ps.HasPoint(cx,cy-1)  {
		m.GetJoiningPointSet(cx,cy-1,ps)

	}

	if m.Has(cx,cy+1) && m.IsOpen(cx,cy,DOWN) && !ps.HasPoint(cx,cy+1)  {
		m.GetJoiningPointSet(cx,cy+1,ps)

	}

	if m.Has(cx-1,cy) && m.IsOpen(cx,cy,LEFT) && !ps.HasPoint(cx-1,cy)  {
		m.GetJoiningPointSet(cx-1,cy,ps)

	}

	if m.Has(cx+1,cy) && m.IsOpen(cx,cy,RIGHT) && !ps.HasPoint(cx+1,cy)  {
		m.GetJoiningPointSet(cx+1,cy,ps)
	}
}

func (m *Maze)HasClosed() (result int) {
	for y:=0;y<len(m.Cells);y++ {
		for x:=0;x<len(m.Cells);x++ {
			if m.Cells[x][y].IsClosed() {
				result++
			}
		}
	}
	return
}

func (m *Maze)Draw(r *sdl.Renderer,w int) {
	n:=m.Len()

	for y:=0;y<n;y++ {
		for x:=0;x<n;x++ {
			cell:=m.Get(x,y)

			location_x:=x*w
			location_y:=y*w

			if cell.HasTop() {
				r.DrawLine(location_x,location_y,location_x+w,location_y)
			}

			if cell.HasBottom() {
				r.DrawLine(location_x,location_y+w,location_x+w,location_y+w)
			}

			if cell.HasLeft() {
				r.DrawLine(location_x,location_y,location_x,location_y+w)
			}

			if cell.HasRight() {
				r.DrawLine(location_x+w,location_y,location_x+w,location_y+w)
			}
		}

	}
}

func (m *Maze)FindPath(x0,y0,x1,y1 int,path * PointStack)  {
	if x0==x1 && y0==y1 {
		path.Push(x0,y0)
		return
	}

	path.Push(x0,y0)

	if m.Get(x0,y0-1)!=nil&&m.IsOpen(x0,y0,UP)&&!path.HasPoint(x0,y0-1) {
		m.FindPath(x0,y0-1,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok && lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0,y0+1)!=nil&&m.IsOpen(x0,y0,DOWN)&&!path.HasPoint(x0,y0+1) {
		m.FindPath(x0,y0+1,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok &&lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0+1,y0)!=nil&&m.IsOpen(x0,y0,RIGHT)&&!path.HasPoint(x0+1,y0) {
		m.FindPath(x0+1,y0,x1,y1,path)

		if lx,ly,ok:=path.Last();ok && lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0-1,y0)!=nil&&m.IsOpen(x0,y0,LEFT)&&!path.HasPoint(x0-1,y0){
		m.FindPath(x0-1,y0,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok &&lx==x1 && ly==y1 {
			return
		}

	}

	path.Pop()

	return
}

func BuildMaze(w int,r int) * Maze {
	rand_generator := rand.New(rand.NewSource(int64(r)))
	mm :=NewMaze(w)

	var ps * PointSet
	var need_reget_joining_point_set bool


	RESTART:
	x,y,ok:=mm.GetFirstClosedCell()

	if !ok {
		return mm
	}

	ww := NewWorm(mm)
	ww.GetInMaze(x,y)

	need_reget_joining_point_set=true

	for {
		next_act_0 :=make([]int,0,4)
		next_act_final :=make([]int,0,4)
		if ww.UpCell()!=nil && !mm.IsOpen(ww.current_x,ww.current_y,UP) {
			next_act_0 =append(next_act_0,UP)
		}

		if ww.DownCell()!=nil && !mm.IsOpen(ww.current_x,ww.current_y,DOWN) {
			next_act_0 =append(next_act_0,DOWN)
		}

		if ww.LeftCell()!=nil && !mm.IsOpen(ww.current_x,ww.current_y,LEFT)  {
			next_act_0 =append(next_act_0,LEFT)
		}

		if ww.RightCell()!=nil && !mm.IsOpen(ww.current_x,ww.current_y,RIGHT) {
			next_act_0 =append(next_act_0,RIGHT)
		}


		if need_reget_joining_point_set {
			ps = NewPointSet()
			mm.GetJoiningPointSet(ww.current_x, ww.current_y, ps)
		}

		for _,a:=range next_act_0 {
			switch a {
			case UP:
				if !ps.HasPoint(ww.current_x, ww.current_y - 1) {
					next_act_final = append(next_act_final, UP)
				}
			case DOWN:
				if !ps.HasPoint(ww.current_x, ww.current_y + 1) {
					next_act_final = append(next_act_final, DOWN)
				}
			case LEFT:
				if !ps.HasPoint(ww.current_x - 1, ww.current_y) {
					next_act_final = append(next_act_final, LEFT)
				}
			case RIGHT:
				if !ps.HasPoint(ww.current_x + 1, ww.current_y) {
					next_act_final = append(next_act_final, RIGHT)
				}
			}
		}

		if len(next_act_final)!=0 {
			switch next_act_final[rand_generator.Intn(len(next_act_final))]{
			case UP:
				if !(ww.UpCell().IsClosed()) {
					need_reget_joining_point_set=true
				} else {
					need_reget_joining_point_set=false
				}
				ww.Up()
			case DOWN:
				if !(ww.DownCell().IsClosed()) {
					need_reget_joining_point_set=true
				} else {
					need_reget_joining_point_set=false
				}
				ww.Down()
			case LEFT:
				if !(ww.LeftCell().IsClosed()) {
					need_reget_joining_point_set=true
				} else {
					need_reget_joining_point_set=false
				}
				ww.Left()
			case RIGHT:
				if !(ww.RightCell().IsClosed()) {
					need_reget_joining_point_set=true
				} else {
					need_reget_joining_point_set=false
				}
				ww.Right()
			}
			ps.Add(ww.current_x, ww.current_y)
		} else {
			goto RESTART
		}
	}

	return mm
}



