package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var trianglePath [][]int
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&trianglePath)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	sumTrianglePath := getSumFromTrianglePath(trianglePath)

	fmt.Println("sumTrianglePath:: ", sumTrianglePath)
}

func getSumFromTrianglePath(pathTriangle [][]int) int {

	for row := len(pathTriangle) - 2; row >= 0; row-- {
		for col := 0; col < len(pathTriangle[row]); col++ {
			// อัปเดตค่าปัจจุบันให้เป็นค่ามากสุดที่ไปได้
			pathTriangle[row][col] += getMaxNumber(pathTriangle[row+1][col], pathTriangle[row+1][col+1])
		}
	}

	totalSumTrianglePath := pathTriangle[0][0]

	return totalSumTrianglePath
}

func getMaxNumber(input1, input2 int) int {

	if input1 > input2 {

		return input1
	}

	return input2
}
