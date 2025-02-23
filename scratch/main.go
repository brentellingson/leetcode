package main

import "fmt"

func main() {
	var grid [][]int
	grid = make([][]int, 5)
	grid[0] = make([]int, 3)
	fmt.Println(len(grid))
	fmt.Println(len(grid[0]))
	fmt.Println(max(1, 2, 3, 4, 5))
}
