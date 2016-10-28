package maze

type Point struct {
	X,Y int
}

func Int64ToPoint(v int64) (x,y int) {
	x=int(v>>32)
	y=int(v&0xffffffff)

	return x,y
}

func PointToInt64(x,y int) int64 {
	return int64(x)<<32+int64(y)
}






