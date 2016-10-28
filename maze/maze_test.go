package maze

import (
	"testing"
)

func BenchmarkMakeMaze(b *testing.B) {
	for i:=0;i<b.N;i++ {
		BuildMaze(100,0)
	}
}

