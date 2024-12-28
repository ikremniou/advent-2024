package advent2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AdventTask1Part1() {
	fmt.Println("Advent 2024 Task 1")
	var taskInput = getTaskInputAsString("https://adventofcode.com/2024/day/1/input",
		"")

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
	for i := 0; i < len(leftAdventList); i++ {
		pathSum += realAbsInts(leftAdventList[i], rightAdventList[i])
	}

	fmt.Println(pathSum)
}
