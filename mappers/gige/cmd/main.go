package main

import (
	"github.com/kubeedge/mappers-go/mapper-sdk-go/pkg/service"
	"github.com/kubeedge/mappers-go/mappers/gige/driver"
)

func main() {
	gd := &driver.GigEVisionDevice{}
	service.Bootstrap("GigEVision", gd)
}
