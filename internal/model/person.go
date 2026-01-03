package model

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Person model
type Person struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	FullName    string `json:"full_name" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// ParseFIO splits Russian FIO
func (p *Person) ParseFIO() {
	parts := strings.Split(p.FullName, " ")
	if len(parts) >= 2 {
		p.FirstName = parts[1]
		p.LastName = parts[0]
	}
}

// EnrichData calls external services (mocked)
func (p *Person) EnrichData() error {
	// Age/Gender from agify.io, genderize.io, nationalize.io
	p.Age, _ = getAge(p.FirstName)
	p.Gender, _ = getGender(p.FirstName)
	p.Nationality, _ = getNationality(p.FirstName)
	return nil
}

func getAge(name string) (int, error) {
	resp, _ := http.Get("https://api.agify.io/?name=" + name)
	var data struct{ Age int }
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Age, nil
}

func getGender(name string) (string, error) {
	resp, _ := http.Get("https://api.genderize.io/?name=" + name)
	var data struct{ Gender string }
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Gender, nil
}

func getNationality(name string) (string, error) {
	resp, _ := http.Get("https://api.nationalize.io/?name=" + name)
	var data struct {
		Country []struct {
			CountryID   string
			Probability float64
		}
	}
	json.NewDecoder(resp.Body).Decode(&data)
	if len(data.Country) > 0 {
		return data.Country[0].CountryID, nil
	}
	return "", nil
}
