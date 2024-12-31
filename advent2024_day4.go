package advent2024

import (
	"bytes"
	"fmt"
)

const XMAS = "XMAS"
const SAMX = "SAMX"

func followXmasSnowflake(dataRows []string, i, j int, total *int) {
	if j < len(dataRows[i])-3 && dataRows[i][j:j+4] == XMAS {
		*total += 1
	}

	if j > 2 && dataRows[i][j-3:j+1] == SAMX {
		*total += 1
	}

	if i < len(dataRows)-3 {
		var down = []byte{dataRows[i][j], dataRows[i+1][j], dataRows[i+2][j], dataRows[i+3][j]}
		if bytes.Equal(down, []byte(XMAS)) {
			*total += 1
		}
	}

	if i > 2 {
		var up = []byte{dataRows[i][j], dataRows[i-1][j], dataRows[i-2][j], dataRows[i-3][j]}
		if bytes.Equal(up, []byte(XMAS)) {
			*total += 1
		}
	}

	if i < len(dataRows[i])-3 && j < len(dataRows)-3 {
		var diagonalDownRight = []byte{dataRows[i][j], dataRows[i+1][j+1], dataRows[i+2][j+2], dataRows[i+3][j+3]}
		if bytes.Equal(diagonalDownRight, []byte(XMAS)) {
			*total += 1
		}
	}

	if i < len(dataRows[i])-3 && j > 2 {
		var diagonalDownLeft = []byte{dataRows[i][j], dataRows[i+1][j-1], dataRows[i+2][j-2], dataRows[i+3][j-3]}
		if bytes.Equal(diagonalDownLeft, []byte(XMAS)) {
			*total += 1
		}
	}

	if i > 2 && j > 2 {
		var diagonalUpLeft = []byte{dataRows[i][j], dataRows[i-1][j-1], dataRows[i-2][j-2], dataRows[i-3][j-3]}
		if bytes.Equal(diagonalUpLeft, []byte(XMAS)) {
			*total += 1
		}
	}

	if i > 2 && j < len(dataRows)-3 {
		var diagonalUpRight = []byte{dataRows[i][j], dataRows[i-1][j+1], dataRows[i-2][j+2], dataRows[i-3][j+3]}
		if bytes.Equal(diagonalUpRight, []byte(XMAS)) {
			*total += 1
		}
	}
}

func isMas(masOrNot []byte) bool {
	if bytes.Equal(masOrNot, []byte(XMAS[1:4])) {
		return true
	}

	if bytes.Equal(masOrNot, []byte(SAMX[0:3])) {
		return true
	}

	return false
}

func followMasSnowflake(dataRows []string, i, j int, total *int) {
	// we found A at this point
	if i > 0 && j > 0 && i < len(dataRows)-1 && j < len(dataRows[i])-1 {
		if isMas([]byte{dataRows[i-1][j-1], dataRows[i][j], dataRows[i+1][j+1]}) &&
			isMas([]byte{dataRows[i-1][j+1], dataRows[i][j], dataRows[i+1][j-1]}) {
			*total += 1
		}
	}
}

func AdventDay4Part1() {
	// find XMAS
	// const inputData = `
	// 	MMMSXXMASM
	// 	MSAMXMSMSA
	// 	AMXSXMAAMM
	// 	MSAMASMSMX
	// 	XMASAMXAMM
	// 	XXAMMXXAMA
	// 	SMSMSASXSS
	// 	SAXAMASAAA
	// 	MAMMMXMMMM
	// 	MXMXAXMASX
	// `
	var inputData = getTaskInputAsString(DAY4INPUT, SESSION)
	var dataRows = getInputFieldsAsStrings(inputData)
	var total = 0

	for i := 0; i < len(dataRows); i++ {
		for j := 0; j < len(dataRows[i]); j++ {
			var currentRune = dataRows[i][j]

			if currentRune == 'X' {
				followXmasSnowflake(dataRows, i, j, &total)
			}
		}
	}

	fmt.Println(total)
}

func AdventDay4Part2() {
	// find MAS
	// const inputData = `
	// 	MMMSXXMASM
	// 	MSAMXMSMSA
	// 	AMXSXMAAMM
	// 	MSAMASMSMX
	// 	XMASAMXAMM
	// 	XXAMMXXAMA
	// 	SMSMSASXSS
	// 	SAXAMASAAA
	// 	MAMMMXMMMM
	// 	MXMXAXMASX
	// `
	var inputData = getTaskInputAsString(DAY4INPUT, SESSION)
	var dataRows = getInputFieldsAsStrings(inputData)
	var total = 0

	for i := 0; i < len(dataRows); i++ {
		for j := 0; j < len(dataRows[i]); j++ {
			var currentRune = dataRows[i][j]

			if currentRune == 'A' {
				followMasSnowflake(dataRows, i, j, &total)
			}
		}
	}

	fmt.Println(total)
}
