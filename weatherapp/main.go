package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const apiKey = "15375295d16277f654cdfd08c691686a"
const weatherAPIURL = "http://api.openweathermap.org/data/2.5/weather"

type WeatherData struct {
	Name      string `json:"name"`
	Main      Main   `json:"main"`
	Weather   []Weather `json:"weather"`
}

type Main struct {
	Temperature float64 `json:"temp"`
	Pressure    int     `json:"pressure"`
	Humidity    int     `json:"humidity"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Welcome to the Weather Application!")

	city := getUserInput("Enter the city name: ")
	units := getUserInput("Choose units (metric/imperial): ")

	url := fmt.Sprintf("%s?q=%s&units=%s&appid=%s", weatherAPIURL, city, units, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nWeather in %s:\n", data.Name)
	fmt.Printf("Temperature: %.2f°C\n", data.Main.Temperature)
	fmt.Printf("Weather: %s\n", data.Weather[0].Description)
	fmt.Printf("Pressure: %d hPa\n", data.Main.Pressure)
	fmt.Printf("Humidity: %d%%\n", data.Main.Humidity)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
