package geocoder

import "testing"

func TestCity(t *testing.T) {
        g := new(Geocoder)
        _, err := g.City("Montevideo")

        if err != nil {
                t.Errorf("Not able to geocode")
        }
}
