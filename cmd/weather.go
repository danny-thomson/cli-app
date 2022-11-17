/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var Version = "0.0.3"

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:               "weather [flags] <argument>",
	Short:             "Get current weather data",
	Long:              `Get the current weather data by providing the name of the city`,
	Run:               GetWeatherData,
	Version:           Version,
	DisableAutoGenTag: true,
}

type Weather struct {
	Weather []WeatherElement `json:"weather"`
	Main    Main             `json:"main"`
	Name    string           `json:"name"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type WeatherElement struct {
	Description string `json:"description"`
}

var client http.Client
var city string

func GetWeatherData(cmd *cobra.Command, args []string) {
	var (
		weather Weather
		apiKey  string = "441a189b27a79ef5bfd0b7e21869d645"
	)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=metric&q=%s", apiKey, city)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(respdata, &weather)
	if err != nil {
		panic(err)
	}

	if weather.Name == "" {
		fmt.Println("Not a valid name")
		return
	}

	fmt.Printf("Weather report for %q at %v\n", weather.Name, time.Now().Format("01-02-2006 03:04:05 PM"))
	fmt.Printf("Current Tempature: %v%c celcius \n", weather.Main.Temp, byte(176))
	fmt.Printf("Current Weather Description: %v \n\n", weather.Weather[0].Description)
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "Name of the city")
	weatherCmd.MarkFlagRequired("city")
	// err := doc.GenMarkdownTree(weatherCmd, "./docs")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = doc.GenYamlTree(weatherCmd, "./docs")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
