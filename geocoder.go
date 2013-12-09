package geocoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Location struct {
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	Country     string      `json:"country"`
	Coordinates Coordinates `json:"coordinates"`
}

type Response struct {
	Status  string
	Results []*Result
}

type Result struct {
	Addresses        []*Address `json:"address_components"`
	FormattedAddress string     `json:"formatted_address"`
	Geometry         Geometry
}

type Address struct {
	Name  string `json:"long_name"`
	Types []string
}

type Geometry struct {
	Location Coordinates
}

const api = "http://maps.googleapis.com/maps/api/geocode/json"

func City(address string) (*Location, error) {
	safe_address := url.QueryEscape(address)
	uri := fmt.Sprintf("%s?sensor=false&address=%s", api, safe_address)
	return fetch(uri)
}

func Coords(lat float64, lng float64) (*Location, error) {
	uri := fmt.Sprintf("%s?sensor=false&latlng=%f,%f", api, lat, lng)
	return fetch(uri)
}

type ReducedCity struct {
	Name    string
	Address string
	Country string
}

func possibleCityName(results []*Result) *ReducedCity {
	city := new(ReducedCity)
	types := []string{
		"country",
		"locality",
		"sublocality",
		"administrative_area_level_3",
		"administrative_area_level_2",
		"administrative_area_level_1"}

	for _, possible := range types {
		for _, result := range results {
			for _, current_address := range result.Addresses {
				for _, city_type := range current_address.Types {
					if possible == city_type {
						if possible == "country" {
							city.Country = current_address.Name
						} else if city.Name == "" {
							city.Address = result.FormattedAddress
							city.Name = current_address.Name
						}
					}
				}
			}
		}
	}

	return city
}

func fetch(uri string) (*Location, error) {
	fmt.Print(uri + "\n")
	data, err := http.Get(uri)

	if err != nil {
		panic(err)
	}

	defer data.Body.Close()

	res := new(Response)
	err = json.NewDecoder(data.Body).Decode(res)

	if err != nil {
		panic(err)
	}

	city := possibleCityName(res.Results)

	if city.Name == "" {
		return nil, errors.New("City not found")
	}

	out := &Location{
		Name:        city.Name,
		Address:     city.Address,
		Country:     city.Country,
		Coordinates: res.Results[0].Geometry.Location}

	return out, nil

}
