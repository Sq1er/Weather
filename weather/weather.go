package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sq1er/weather/geo"
)

var ErrWrongFormat = errors.New("wrong_format")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error_url")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error_http")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error_readbody")
	}
	return string(body), nil
}
