package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

type Device struct {
	id          uint `gorm:"primaryKey"`
	deviceId    string
	deviceType  string
	coordinates string
	status      string
	timezone    string
}

func deviceInfoTransform(rawDevices map[string]interface{}) []Device {
	var devices []Device
	for _, v := range rawDevices {
		var newDevice Device
		if reflect.ValueOf(v).Kind() == reflect.Map {
			val := reflect.ValueOf(v)
			for _, key := range val.MapKeys() {
				strKey := fmt.Sprintf("%v", key.Interface())
				strVal := fmt.Sprintf("%v", val.MapIndex(key))
				switch strKey {
				case "id":
					newDevice.deviceId = strVal
					break
				case "status":
					newDevice.status = strVal
					break
				case "timezone":
					newDevice.timezone = strVal
					break
				case "coordinates":
					newDevice.coordinates = strVal
					break
				case "type":
					newDevice.deviceType = strVal
					break
				default:
					break
				}
			}
			devices = append(devices, newDevice)
		}
	}
	return devices
}

func insertDevices(devices []Device) {
	fmt.Println(len(devices))
	dsn := "awair:awair@tcp(127.0.0.1:3306)/awair?charset=utf8mb4&parseTime=True&loc=Local"
	db, errConnectDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errConnectDB != nil {
		panic("Cannot open connection to database")
	}
	db.Table("devices").AutoMigrate(&Device{})
	db.CreateInBatches(devices, 100)
}
