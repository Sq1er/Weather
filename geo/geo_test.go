package geo_test

import (
	"testing"

	"github.com/sq1er/weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "Orel1"
	// Act - выполняем функцию
	_, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
