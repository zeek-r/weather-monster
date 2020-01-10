package request

import (
	"bytes"
	"net/http"
)

func Request(body []byte, method, url string) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
