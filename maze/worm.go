package maze

import "errors"

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Worm struct {
	Maze                 * Maze
	current_x, current_y int
	PointSet             *PointSet
}

func NewWorm(m *Maze) *Worm {
	return &Worm{Maze:m,PointSet:NewPointSet()}
}

func (w *Worm)GetInMaze(x,y int) {
	w.current_x =x
	w.current_y =y
	w.PointSet.Add(x,y)
}

func (w *Worm)CurrentPosition() (x,y int) {
	return w.current_x,w.current_y
}

func (w *Worm)UpCell() *Cell {
	return w.Maze.Get(w.current_x,w.current_y -1)
}

func (w *Worm)DownCell()  *Cell  {
	return w.Maze.Get(w.current_x,w.current_y +1)
}

func (w *Worm)LeftCell() *Cell {
	return w.Maze.Get(w.current_x -1,w.current_y)
}

func (w *Worm)RightCell() *Cell {
	return w.Maze.Get(w.current_x +1,w.current_y)
}

func (w *Worm)Up()  {
	cell:=w.Maze.Get(w.current_x,w.current_y)
	next_cell :=w.UpCell()

	if next_cell !=nil  {
		cell.EraseTop()
		next_cell.EraseBottom()
		w.current_y =w.current_y -1
		w.PointSet.Add(w.current_x,w.current_y -1)

	} else {
		panic(errors.New("If Up,It will go out the maze"))
	}
}

func (w *Worm)Down()  {
	cell:=w.Maze.Get(w.current_x,w.current_y)
	next_cell:=w.DownCell()

	if next_cell !=nil  {
		cell.EraseBottom()
		next_cell.EraseTop()
		w.current_y =w.current_y +1
		w.PointSet.Add(w.current_x,w.current_y +1)
	} else {
		panic(errors.New("If Down,It will go out the maze"))
	}

}

func (w *Worm)Left()  {
	cell:=w.Maze.Get(w.current_x,w.current_y)
	next_cell:=w.LeftCell()

	if next_cell !=nil  {
		cell.EraseLeft()
		next_cell.EraseRight()
		w.current_x =w.current_x -1
		w.PointSet.Add(w.current_x -1,w.current_y)
	} else {
		panic(errors.New("If Left,It will go out the maze"))
	}
}


func (w *Worm)Right()  {
	cell:=w.Maze.Get(w.current_x,w.current_y)
	next_cell:=w.RightCell()

	if next_cell !=nil {
		cell.EraseRight()
		next_cell.EraseLeft()
		w.current_x =w.current_x +1
		w.PointSet.Add(w.current_x +1,w.current_y)
	} else {
		panic(errors.New("If Right,It will go out the maze"))
	}
}

