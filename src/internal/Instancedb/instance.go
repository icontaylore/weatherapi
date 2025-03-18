package Instancedb

import (
	"bot_notif/src/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Instance struct {
	DB *pgxpool.Pool
}

func New(c *pgxpool.Pool) (*Instance, error) {
	maken := &Instance{DB: c}
	wg := sy
	if maken != nil {
		return maken, nil
	}
	return nil, errors.New("ошибка nil instance")
}

func (i *Instance) Start() {
	rows, err := i.DB.Query(context.Background(), "SELECT * FROM weather_cache;")
	if err != nil {
		log.Printf("ошибка с запросом, нету столбцов")
	}
	for rows.Next() {
		fmt.Println("2")
		weather := &models.Weather{}
		rows.Scan(&weather.CityName, &weather.WeatherDate, &weather.Temperature, &weather.Humidity,
			&weather.Pressure, &weather.Description)
		fmt.Println(weather)
	}
	i.DB.Close()
}
