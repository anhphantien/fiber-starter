package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpService struct{}

func (s HttpService) Get() (map[string]any, error) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "<url>", nil)
	req.Header = http.Header{
		"Authorization": {"Bearer <token>"},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buffer, _ := io.ReadAll(res.Body)
	fmt.Println(string(buffer))
	data := map[string]any{}
	err = json.Unmarshal([]byte(string(buffer)), &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}
