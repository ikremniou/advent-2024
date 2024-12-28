package advent2024

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getTaskInputAsString(url string, session string) string {
	var request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Cookie", "session="+session)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(bodyBytes)
}

func getInputIntFieldsOfSize(url string, session string, columnsNumber int) [][]int {
	var taskInput = getTaskInputAsString(url, session)

	var split = strings.Fields(taskInput)
	var columns = make([][]int, columnsNumber)
	for i := 0; i < columnsNumber; i++ {
		columns[i] = make([]int, len(split)/columnsNumber)
	}

	for i := 0; i < len(split); i = i + 2 {
		for j := 0; j < columnsNumber; j++ {
			columns[j][i/columnsNumber], _ = strconv.Atoi(split[i+j])
		}
	}

	return columns
}

func getInputIntFieldsAsRows(taskInput string) [][]int {
	var split = strings.Split(taskInput, "\n")
	var rows = make([][]int, 0, len(split))

	for i := 0; i < len(split); i += 1 {
		var row = strings.Fields(split[i])
		if len(row) == 0 {
			continue
		}

		var currRow = make([]int, len(row))
		for j := 0; j < len(row); j++ {
			currRow[j], _ = strconv.Atoi(row[j])
		}

		rows = append(rows, currRow)
	}

	return rows
}
