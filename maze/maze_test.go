package maze

import (
	"testing"
	"log"
	"sync"
)

func BenchmarkMakeMaze(b *testing.B) {
	for i:=0;i<b.N;i++ {
		BuildMaze(20,0)
	}
}

func BenchmarkMakeMazeArea(b *testing.B) {
	for i:=0;i<b.N;i++ {
		BuildMazeArea(20,0)
	}
}

func TestMazeArea(t *testing.T) {

	var m *Maze
	var ma *Maze

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		m = BuildMaze(100, 0)
		wg.Done()
	}()

	go func() {
		ma = BuildMazeArea(100, 0)
		wg.Done()
	}()

	wg.Wait()

	if !m.Equal(ma) {
		t.Failed()
	}


}

