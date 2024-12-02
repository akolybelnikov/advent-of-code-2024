package utils

import (
	"fmt"
	"io"
	"net/http"
)

func ReadHTTP(year, day int, session string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session = fmt.Sprintf("session=%s", session)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(bodyBytes))
	}

	return string(bodyBytes), nil
}
