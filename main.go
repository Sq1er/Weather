package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sq1er/weather/geo"
	"github.com/sq1er/weather/weather"
)

func main() {
	fmt.Println("___Погода___")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("Ошибка определения местоположения:", err)
		os.Exit(1)
	}
	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(weatherData)
}
