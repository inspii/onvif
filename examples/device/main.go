package main

import (
	"fmt"
	"github.com/inspii/onvif"
)

func main() {
	camera := onvif.NewOnvifDevice("192.168.5.82") // OK

	camera.Auth("admin", "Phr@12345")

	err := camera.Initialize()
	if err != nil {
		fmt.Println("camera.Initialize error:", err.Error())
		return
	}

	di, err := camera.Device.GetDeviceInformation()
	fmt.Println(di)
	users, err := camera.Device.GetUsers()
	fmt.Println(users)
}
