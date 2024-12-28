package advent2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AdventDay1Part1() {
	var taskInput = getTaskInputAsString(DAY1INPUT, SESSION)

	var split = strings.Fields(taskInput)
	var leftAdventList = make([]int, len(split)/2)
	var rightAdventList = make([]int, len(split)/2)

	for i := 0; i < len(split); i = i + 2 {
		leftAdventList[i/2], _ = strconv.Atoi(split[i])
		rightAdventList[i/2], _ = strconv.Atoi(split[i+1])
	}

	sort.Ints(leftAdventList)
	sort.Ints(rightAdventList)

	realAbsInts := func(x, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}

	var pathSum = 0
	for i := 0; i < len(leftAdventList); i += 1 {
		pathSum += realAbsInts(leftAdventList[i], rightAdventList[i])
	}

	fmt.Println(pathSum)
}

func AdventDay1Part2() {
	// var left, right = []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}
	var columns = getInputIntFieldsOfSize(DAY1INPUT, SESSION, 2)
	var left, right = columns[0], columns[1]

	var occurMap = make(map[int]int)
	for i := 0; i < len(right); i += 1 {
		occurMap[right[i]]++
	}

	var occurSum = 0
	for i := 0; i < len(left); i += 1 {
		occurSum += occurMap[left[i]] * left[i]
	}

	fmt.Println(occurSum)
}
