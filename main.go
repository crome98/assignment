package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open devices.json file
	jsonFile, err := os.Open("devices.json")
	if err != nil {
		fmt.Println("Cannot open file")
	}
	// Close file
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var devices map[string]interface{}
	err2 := json.Unmarshal(byteValue, &devices)
	if err2 != nil {
		fmt.Println("Error occurs")
	}

	// Get all devices ids
	i := 0
	devicesIDs := make([]string, len(devices))
	for k := range devices {
		devicesIDs[i] = k
		i++
	}
	transformedDevice := deviceInfoTransform(devices)
	fmt.Println(transformedDevice[9].id)
	insertDevices(transformedDevice)
	fmt.Println(len(devices))
	fmt.Println(len(transformedDevice))

}
