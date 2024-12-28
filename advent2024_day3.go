package advent2024

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func AdventDay3Part1() {
	// const inputData = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	var inputData = getTaskInputAsString(DAY3INPUT, SESSION)
	var mulRegex = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	var allSubs = mulRegex.FindAllStringSubmatch(inputData, -1)
	var result = 0

	for i := 0; i < len(allSubs); i += 1 {
		left, _ := strconv.Atoi(allSubs[i][1])
		right, _ := strconv.Atoi(allSubs[i][2])

		result += left * right
	}

	fmt.Println(result)
}

func AdventDay3Part2() {
	// var inputData = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)udo()n?mul(8,5))"
	var inputData = getTaskInputAsString("https://adventofcode.com/2024/day/3/input", SESSION)

	var shouldLoop = true

	for shouldLoop {
		beforeNot, afterNot, foundNot := strings.Cut(inputData, "don't()")
		if foundNot {
			_, afterDo, _ := strings.Cut(afterNot, "do()")
			inputData = beforeNot + afterDo
		} else {
			inputData = beforeNot
		}

		if !foundNot {
			shouldLoop = false
		}
	}

	fmt.Println(inputData)

	var mulRegex = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	var allSubs = mulRegex.FindAllStringSubmatch(inputData, -1)
	var result = 0

	for i := 0; i < len(allSubs); i += 1 {
		left, _ := strconv.Atoi(allSubs[i][1])
		right, _ := strconv.Atoi(allSubs[i][2])

		result += left * right
	}

	fmt.Println(result)
}
