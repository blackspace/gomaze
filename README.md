*Life's Path Is a Maze, Not a Straight Line.*

![a 10x10 maze](https://raw.githubusercontent.com/blackspace/gomaze/master/samples/maze_10x10.png)
![a 150x150 maze](https://raw.githubusercontent.com/blackspace/gomaze/master/samples/maze_150x150.png)

About Installing
---------------------------------------
First because the gomaze draws the maze by the sdl2,you should install 
go-sdl2 by following the installing instructions from the [go-sdl2](https://github.com/veandco/go-sdl2) page.

I take [gods](https://github.com/emirpasic/gods) as the data struction of point set,so then 

`
go get github.com/emirpasic/gods
`

Exampels
------------------------------------

To make 100 9x9 mazes:

```
go run maze_nxn.go
```

Or,to make one nxn maze:


```
go run maze_single.go
```
