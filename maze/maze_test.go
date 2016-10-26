package maze

import "testing"

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
	m:=BuildMaze(40,0)
	ma:=BuildMazeArea(40,0)

	if !m.Equal(ma) {
		t.Fail()
	}
}
