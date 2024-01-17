package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type ResponseNationalize struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

func GetCountry(name string) (string, error) {
	if name == "" {
		return "", ErrorEmptyName
	}
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api.nationalize.io got HTTP status code %d", resp.StatusCode)
	}

	var response ResponseNationalize
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	var country string
	country = findMaxProbCountry(response.Country).CountryID
	if country == "" {
		return "", ErrorEmptyCountry
	}
	return country, nil
}

func findMaxProbCountry(countries []Country) *Country {
	maxProb := -1.0
	maxProbCountry := &Country{}

	for i, country := range countries {
		if country.Probability > maxProb {
			maxProb = country.Probability
			maxProbCountry = &countries[i]
		}
	}

	fmt.Printf("Got response with ID of the country with the highest probability: %s\n", maxProbCountry.CountryID)
	return maxProbCountry
}
