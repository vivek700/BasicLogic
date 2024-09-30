package main

import "fmt"

func isValidSudoku(b [][]string) bool {

	//checking rows
	for _, value := range b {
		rowMap := make(map[string]int)
		for _, v := range value {
			if _, ok := rowMap[v]; !ok {
				rowMap[v] = 1
			} else {
				if v == "." {
					rowMap[v] += 1
				} else {
					return false
				}
			}
		}
	}

	//checking columns

	for i := range b {
		colMap := make(map[string]int)

		for j := range b {
			if _, ok := colMap[b[j][i]]; !ok {
				colMap[b[j][i]] = 1
			} else {
				if b[j][i] == "." {
					colMap[b[j][i]] += 1

				} else {
					return false
				}
			}

		}
	}

	//checking 3x3

	for startrow := 0; startrow < 9; startrow += 3 {
		for startcol := 0; startcol < 9; startcol += 3 {
			squareMap := make(map[string]int)
			for row := startrow; row < startrow+3; row++ {
				for col := startcol; col < startcol+3; col++ {
					if _, ok := squareMap[b[row][col]]; !ok {
						squareMap[b[row][col]] = 1

					} else {
						if b[row][col] == "." {
							squareMap[b[row][col]] += 1
						} else {
							return false
						}
					}
				}
			}

		}

	}

	return true
}

func main() {

	board := [][]string{{"1", "2", ".", ".", "3", ".", ".", ".", "."},
		{"4", ".", ".", "5", ".", ".", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", ".", "3"},
		{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
		{".", ".", ".", "8", ".", "3", ".", ".", "5"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", ".", "2", ".", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "8"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"}}

	board1 := [][]string{{"1", "2", ".", ".", "3", ".", ".", ".", "."},
		{"4", ".", ".", "5", ".", ".", ".", ".", "."},
		{".", "9", "1", ".", ".", ".", ".", ".", "3"},
		{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
		{".", ".", ".", "8", ".", "3", ".", ".", "5"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", ".", "2", ".", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "8"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"}}

	fmt.Println(isValidSudoku(board))
	fmt.Println(isValidSudoku(board1))

}
