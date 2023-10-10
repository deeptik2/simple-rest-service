package services

import "fmt"

// If you want to mock, use an interface
type pingServiceInterface interface {
	PingService() (string, error)
}

type PingServicesStruct struct{}

var PingServiceVar pingServiceInterface = PingServicesStruct{}

func (service PingServicesStruct) PingService() (string, error) {
	fmt.Println("connecting to an external 3rd party service")
	return "pong", nil
}
