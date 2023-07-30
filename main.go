package slack

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	_ "net/url"
	"os"
	"strings"
)

func SendInfo(message string) (string, error) {
	message = "ℹ️ " + message + " ℹ️ \r\n"
	return SendSlackAlert(message)
}
func SendWarning(message string) (string, error) {
	message = "⚠️ " + message + " ⚠️ \r\n"
	return SendSlackAlert(message)
}
func SendAlert(message string) (string, error) {
	message = "‼️ " + message + " ‼️ \r\n"
	return SendSlackAlert(message)
}

func SendSlackAlert(message string) (string, error) {
	var url string

	url = os.Getenv("SLACK_NOTIFICATION")
	if len(url) == 0 {
		return "", errors.New("SLACK_NOTIFICATION environment variable is not set")

	}

	if !strings.HasPrefix(url, "https://hooks.slack.com/services/") {
		url = "https://hooks.slack.com/services/" + url
	}

	method := "POST"

	s := fmt.Sprintf(`
	{
	  "text": " %s"
	}
	`, message)

	payload := strings.NewReader(s)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
