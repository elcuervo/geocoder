package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Coordinates struct {
	Lat, Lng float64
}

type City struct {
	Name        string
	Coordinates Coordinates
}

type Response struct {
	Status  string
	Results []*Result
}

type Result struct {
	Addresses []*Address `json:"address_components"`
	Geometry  Geometry
}

type Address struct {
	Name  string `json:"long_name"`
	Types []string
}

type Geometry struct {
	Location Coordinates
}

const api = "http://maps.googleapis.com/maps/api/geocode/json"

type Geocoder struct{}

func (g *Geocoder) City(address string) (*City, error) {
	safe_address := url.QueryEscape(address)
	uri := fmt.Sprintf("%s?sensor=false&address=%s", api, safe_address)
	return g.fetch(uri)
}

func (g *Geocoder) Coords(lat float64, lng float64) (*City, error) {
	uri := fmt.Sprintf("%s?sensor=false&latlng=%f,%f", api, lat, lng)
	return g.fetch(uri)
}

func (g *Geocoder) possibleCityName(results []*Result) string {
	types := []string{
		"locality",
		"sublocality",
		"administrative_area_level_3",
		"administrative_area_level_2"}

	for _, possible := range types {
		for _, result := range results {
			for _, address := range result.Addresses {
				for _, city_type := range address.Types {
					if possible == city_type {
					        return address.Name
					}
				}
			}
		}
	}

	return ""
}

func (g *Geocoder) fetch(uri string) (*City, error) {
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

	city := &City{
		g.possibleCityName(res.Results),
		res.Results[0].Geometry.Location}

	return city, nil
}
