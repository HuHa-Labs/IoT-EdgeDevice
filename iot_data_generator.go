package main

import (
	"encoding/json"
	"log"
	"math/rand"
)

type IotData struct {
	Temperature      float32 `json:"temp_F"`
	Humidity         float32 `json:"humidity"`
	ElectricityUsage float32 `json:"electricityUsage_watt"`
	MotionDetected   bool    `json:"motionDetected"`
}

func generateTemperature() float32 {
	return (rand.Float32() * 51.8) + (73.4 - 51.8)
}

func generateHumidity() float32 {
	return (rand.Float32() * 0.55) + (0.66 - 0.55)
}

func generateElectricityUsage() float32 {
	return (rand.Float32() * 3000) + (10000 - 3000)
}

func generateMotionDetected() bool {
	return rand.Intn(2) == 1
}

func GetIotData() []byte {
	newIotData := IotData{
		generateTemperature(),
		generateHumidity(),
		generateElectricityUsage(),
		generateMotionDetected(),
	}
	buffer, err := json.Marshal(newIotData)
	if err != nil {
		log.Fatalf("error marshaling JSON: %v\n", err)
	}
	return buffer
}
