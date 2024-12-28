package advent2024

import (
	"io"
	"net/http"
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
