package geocoder

import (
	"cgl.tideland.biz/asserts"
	"testing"
)

func TestCity(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)
	city, err := City("Montevideo")

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	expected := Coordinates{-34.8836111, -56.1819444}
	assert.Equal(city.Coordinates, expected, "Unexpected coordinates for Montevideo")
}

func TestCoords(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)
	city, err := Coords(-34.8836111, -56.1819444)

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	assert.Equal(city.Name, "Montevideo", "Unexpected coordinates for Montevideo")
}

func TestPossibleCityName(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)
	city, err := City("Colonia, Uruguay")

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	assert.NotNil(city.Name, "Missing city name")
	assert.Equal(city.Name, "Colonia", "Wrong city name")
}
