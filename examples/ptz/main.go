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

	camera.PTZ.Client.SetLogger(func(message string) {
		fmt.Println(message)
	}, func(message string) {
		fmt.Println(message)
	})
	resp, err := camera.PTZ.ContinuousMove("Profile_1", -0.5, 0, 0)
	fmt.Println(resp)
}
