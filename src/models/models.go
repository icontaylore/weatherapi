package models

import "time"

type Config struct {
	Host     string
	Port     string
	DBname   string
	DBuser   string
	Password string
}

type Weather struct {
	CityName    string
	WeatherDate time.Time
	Temperature int
	Humidity    int
	Pressure    int
	Description string
}
