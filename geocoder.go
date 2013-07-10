package geocoder

import (
	"encoding/json"
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

func possibleCityName(results []*Result) (string, string) {
	types := []string{
		"locality",
		"sublocality",
		"administrative_area_level_3",
		"administrative_area_level_2",
		"administrative_area_level_1"}

	for _, possible := range types {
		for _, result := range results {
			for _, address := range result.Addresses {
				for _, city_type := range address.Types {
					if possible == city_type {
						return address.Name, result.FormattedAddress
					}
				}
			}
		}
	}

	return "Unknown", "Unknown"
}

func fetch(uri string) (*Location, error) {
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

	city_name, address := possibleCityName(res.Results)

	city := &Location{
		Name:        city_name,
		Address:     address,
		Coordinates: res.Results[0].Geometry.Location}

	return city, nil
}
