package main

import (
	"fmt"
	"os"
)
type point struct {
	i, j int
}
var  dirs =  [4] point {
	// 上 左 下 右
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}
func readMazeContent(name string) [] []int  {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	var row, col int
	// read & format the first line of maze
	fmt.Fscanf(file,"%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze{
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
func (p point) add(r point) point  {
	return point{p.i + r.i, p.j + r.j}
}
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j< 0 || p.j>= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start point, end point) [][] int  {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Queue := []point{start}
	for {

		current := Queue[0]
		Queue = Queue[1:]

		if current == end{
			break;
		}
		for _, dir := range dirs {
			next := current.add(dir)
			// 判断是否在迷宫中，是否越界 || 下一个是否不通
			value, ok := next.at(maze)
			if !ok || value == 1 {
				continue
			}
			if next == start {
				continue
			}
			// 是否已经走过
			value, ok = next.at(steps)
			if  !ok || value != 0{
				continue
			}
			curSteps, _ := current.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Queue = append(Queue, next)
		}
	}
	return steps;
}
func main() {
	maze := readMazeContent("./maze")
	for _, row := range maze {
		for _, v := range row {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps{
		for _, val := range row{
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
