package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type WeatherInformations struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Rain struct {
		} `json:"rain"`
		Snow struct {
			ThreeH float64 `json:"3h"`
		} `json:"snow"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
	} `json:"city"`
}

func EnvLoad() {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/kshiva1126/weather/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ParseJsonReceivedAndExecute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes := ([]byte)(body)
	weather := new(WeatherInformations)

	if err = json.Unmarshal(jsonBytes, weather); err != nil {
		fmt.Println("Sorry, the city is not found. \nPlease kindly confirm it.")
		return
	}

	place_name := weather.City.Name
	place_temp := strconv.FormatFloat(weather.List[0].Main.Temp, 'f', 1, 64)
	place_status := weather.List[0].Weather[0].Description

	fmt.Println("The current temperature in " + place_name + " is " + place_temp + "° C and the sky is " + place_status + ".\nHave a nice day!")
}
