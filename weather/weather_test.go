package weather_test

import (
	"strings"
	"testing"

	"github.com/sq1er/weather/geo"
	"github.com/sq1er/weather/weather"
)

func TestGetWeather(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	expected := "Orel"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	// Act - выполняем функцию
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	// Assert - проверка результата с expected
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 101},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatheWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange - подготовка, expected результат, данные для функции
			expected := "Orel"
			geoData := geo.GeoData{
				City: expected,
			}
			// Act - выполняем функцию
			_, err := weather.GetWeather(geoData, tc.format)
			// Assert - проверка результата с expected
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
