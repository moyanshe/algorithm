package main

import "fmt"

func main() {
	var matrix = [][]int{{-5}}
	fmt.Println(findNumberIn2DArray(matrix, -10))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	columns := len(matrix[0])
	row := rows - 1
	column := 0

	for row >= 0 && column < columns { //如果将columns替换成len(matrix[0]),那效率会降低不少
		if target > matrix[row][column] {
			row--
		} else if target > matrix[row][column] {
			column++
		} else {
			return true
		}
	}
	return false
}
