package geocoder

import (
        "fmt"
        "encoding/json"
        "net/http"
        "net/url"
)

type Coordinates struct {
        Lat, Lng float64
}

type Response struct {
        Status  string
        Results []*Result `json:"address_components"`
}

type Result struct {
        Address  string   `json:"formatted_address"`
        Geometry Geometry
}

func (r *Result) Coordinates() Coordinates {
        return r.Geometry.Location
}

type Geometry struct {
        Location Coordinates
}

const api = "http://maps.googleapis.com/maps/api/geocode/json"

type Geocoder struct{}

func (g *Geocoder) City(address string) (*Result, error) {
        safe_address := url.QueryEscape(address)
        uri := fmt.Sprintf("%s?sensor=false&address=%s", api, safe_address)
        data, err := http.Get(uri)

        if err != nil {
                panic(err)
        }

        defer data.Body.Close()

        response := new(Response)
        err = json.NewDecoder(data.Body).Decode(response)

        if err != nil {
                panic(err)
        }

        return response.Results[0], nil
}
