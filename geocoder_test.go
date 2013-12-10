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
	city, err := Coords(-34.905944, -56.191556)

	t.Log(city)
	if err != nil {
		t.Errorf("Not able to geocode:", err)
	}

	assert.Equal(city.Name, "Montevideo", "Unexpected coordinates for Montevideo")
}

func TestPossibleCityName(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)
	city, err := City("Colonia del sacramento, Uruguay")

	if err != nil {
		t.Errorf("Not able to geocode")
	}

	assert.NotNil(city.Name, "Missing city name")
	assert.Equal(city.Name, "Colonia Del Sacramento", "Wrong city name")
}

func TestNonExistentCity(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)

	_, err := City("Balsd234")
	assert.NotNil(err, "Should Fail")
}
