package advent2024

import "fmt"

func AdventDay2Part1() {
	// const taskInput = `
	// 	9 6 4 2 1
	// 	40 42 44 47 49 50 48
	// `
	var taskInput = getTaskInputAsString(DAY2INPUT, SESSION)

	var numberOfSafeReports = 0
	var rows = getInputIntFieldsAsRows(taskInput)

	for i := 0; i < len(rows); i += 1 {
		var isSafe = true
		var prevNumber = rows[i][0]
		var isAsc = rows[i][0] < rows[i][1]

		for j := 1; j < len(rows[i]) && isSafe; j += 1 {
			if isAsc {
				if rows[i][j] <= prevNumber || rows[i][j] > prevNumber+3 {
					isSafe = false
				}
			} else {
				if rows[i][j] >= prevNumber || rows[i][j] < prevNumber-3 {
					isSafe = false
				}
			}

			prevNumber = rows[i][j]
		}

		if isSafe {
			numberOfSafeReports += 1
		}
	}

	fmt.Println(numberOfSafeReports)
}

func AdventDay2Part2() {
	// const taskInput = `
	// 	7 6 4 2 1
	// 	1 2 7 8 9
	// 	9 7 6 2 1
	// 	1 3 2 4 5
	// 	8 6 4 4 1
	// 	1 3 6 7 9
	//  	40 30 43 46 49 50 52
	//  	20 40 43 46 49 50 52
	// `

	var taskInput = getTaskInputAsString(DAY2INPUT, SESSION)
	var numberOfSafeReportsNew = 0
	var rows = getInputIntFieldsAsRows(taskInput)

	for i := 0; i < len(rows); i += 1 {

		var calculateNumberOfProblems = func(currRow []int) bool {
			var isSafeLocal = true
			var prevNumber = currRow[0]
			var isAsc = currRow[0] < currRow[1]

			for j := 1; j < len(currRow) && isSafeLocal; j += 1 {
				if isAsc {
					if currRow[j] <= prevNumber || currRow[j] > prevNumber+3 {
						isSafeLocal = false
					}
				} else {
					if currRow[j] >= prevNumber || currRow[j] < prevNumber-3 {
						isSafeLocal = false
					}
				}

				prevNumber = currRow[j]
			}

			return isSafeLocal
		}

		var isSafe = false
		for j := 0; j < len(rows[i]) && !isSafe; j++ {
			partWithout1 := append([]int{}, rows[i][:j]...)
			partWithout1 = append(partWithout1, rows[i][j+1:]...)

			isSafe = calculateNumberOfProblems(partWithout1)
		}

		if isSafe {
			numberOfSafeReportsNew += 1
		}
	}

	fmt.Println(numberOfSafeReportsNew)
}
