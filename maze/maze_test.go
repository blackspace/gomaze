package maze

import (
	"testing"
	"log"
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
	for i:=30;i<50;i++{
		m := BuildMaze(i, 0)
		ma := BuildMazeArea(i, 0)

		if !m.Equal(ma) {
			log.Println(i)
			t.Fail()
		}
	}

}

