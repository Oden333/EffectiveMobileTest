package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseGenderize struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func GetGender(name string) (string, error) {
	if name == "" {
		return "", ErrorEmptyName
	}
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api.genderize.io got HTTP status code %d", resp.StatusCode)
	}

	var response ResponseGenderize
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	fmt.Printf("Got response with Gender: %s\n", response.Gender)
	return response.Gender, nil
}
