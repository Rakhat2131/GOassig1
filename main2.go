package main2

import (
	"fmt"
	"time"
)

type Observer interface {
	Update(float64, float64, float64)
}

type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}
type WeatherData struct {
	observers                       []Observer
	temperature, humidity, pressure float64
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range wd.observers {
		if observer == o {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
			break
		}
	}
}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.temperature = 20 + randFloat(-5, 5)
	wd.humidity = 50 + randFloat(-10, 10)
	wd.pressure = 1000 + randFloat(-20, 20)
	wd.NotifyObservers()
}

type CurrentConditionsDisplay struct {
	temperature, humidity float64
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f°F and %.2f%% humidity\n", ccd.temperature, ccd.humidity)
}
func randFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func main() {
	weatherData := &WeatherData{}
	currentConditionsDisplay := &CurrentConditionsDisplay{}
	weatherData.RegisterObserver(currentConditionsDisplay)

	for i := 0; i < 5; i++ {
		weatherData.MeasurementsChanged()
		time.Sleep(time.Second)
	}
}
