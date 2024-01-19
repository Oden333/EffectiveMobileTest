package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ResponseAgify struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func GetAge(name string) (int, error) {
	if name == "" {
		return -1, ErrorEmptyName
	}
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("api.agify.io got HTTP status code %d", resp.StatusCode)
	}

	var response ResponseAgify
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return -1, err
	}

	logrus.Info(fmt.Sprintf("Got response with Age: %d\n", response.Age))
	return response.Age, nil
}
