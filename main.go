package main

import (
	"encoding/json"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"io/ioutil"
	"net/http"
)

func main() {
	region := getUserRegion()
	var coordinates Coordinates = getUserCoordinates(region)
	var temperatureInformation TemperatureInformation = getWheatherInformations(coordinates)

	temperature := temperatureInformation.Hourly.Temperature2M[0]

	temperaturex := fmt.Sprintf("%v", temperature)
	printValue := string(fmt.Sprintf("" + temperaturex + " C"))
	myFigure := figure.NewFigure(printValue, "", true)
	myFigure.Print()

}

func getUserRegion() string {
	response, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // defer islem sonunda calisir

	var locationInformation LocationApiResponse
	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &locationInformation)

	if err != nil {
		panic(err)
	}

	return locationInformation.Region
}

func getUserCoordinates(region string) Coordinates {
	response, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + region + "&limit=1")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var coordinatesInformation CoordinatesInformationResults

	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &coordinatesInformation)

	latitude := coordinatesInformation.Results[0].Latitude
	longitude := coordinatesInformation.Results[0].Longitude

	region += region
	return Coordinates{Latitude: latitude, Longitude: longitude}
}

func getWheatherInformations(coordinates Coordinates) TemperatureInformation {
	response, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=" + fmt.Sprintf("%v", coordinates.Latitude) + "&longitude=" + fmt.Sprintf(
		"%v", coordinates.Longitude) + "&hourly=temperature_2m")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var temperatureInformation TemperatureInformation

	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &temperatureInformation)

	return temperatureInformation
}
