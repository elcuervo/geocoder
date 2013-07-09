package geocoder

import "testing"

func TestCity(t *testing.T) {
	g := new(Geocoder)
	city, err := g.City("Montevideo")

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	expected := Coordinates{-34.8836111, -56.1819444}

	if city.Coordinates != expected {
		t.Error("Unexpected coordinates for Montevideo")
	}

}

func TestCoords(t *testing.T) {
	g := new(Geocoder)
	city, err := g.Coords(-34.8836111, -56.1819444)

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	if city.Name != "Montevideo" {
		t.Errorf("Unexpected city name:(%s) for Montevideo coordinates", city.Name)
	}
}
