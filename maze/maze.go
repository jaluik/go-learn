package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		panic(err)
	}
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, err = fmt.Fscanf(file, "%d", &maze[i][j])
			if err != nil {
				panic(err)
			}
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(adder point) point {
	return point{p.i + adder.i, p.j + adder.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for index := range steps {
		steps[index] = make([]int, len(maze[index]))
	}

	q := []point{start}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			break
		}
		for _, val := range dirs {
			next := cur.add(val)
			value, b := next.at(maze)
			if !b || value == 1 {
				continue
			}
			val, b1 := next.at(steps)
			if !b1 || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			currentVal, _ := cur.at(steps)
			steps[next.i][next.j] = currentVal + 1
			q = append(q, next)
		}

	}
	return steps

}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	steps := walk(maze, point{0, 0}, point{5, 4})

	fmt.Println()

	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
}
