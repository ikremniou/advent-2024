package advent2024_test

import (
	advent2024 "ikremniou/advent-2024"
	"testing"
)

func TestAdvent2024Task1(t *testing.T) {
	t.Run("Advent 2024 Day 1 part 1", func(t *testing.T) {
		advent2024.AdventDay1Part1()
	})

	t.Run("Advent 2024 Day 1 part 2", func(t *testing.T) {
		advent2024.AdventDay1Part2()
	})
}

func TestAdvent2024Day2(t *testing.T) {
	t.Run("Advent 2024 Day 2 part 1", func(t *testing.T) {
		advent2024.AdventDay2Part1()
	})

	t.Run("Advent 2024 Day 2 part 2", func(t *testing.T) {
		advent2024.AdventDay2Part2()
	})
}

func TestAdvent2024Day3(t *testing.T) {
	t.Run("Advent 2024 Day 3 part 1", func(t *testing.T) {
		advent2024.AdventDay3Part1()
	})

	t.Run("Advent 2024 Day 3 part 2", func(t *testing.T) {
		advent2024.AdventDay3Part2()
	})
}
