package ssd1306

import (
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ssd1306"
)

func NewDisplay(bus drivers.I2C) *ssd1306.Device {
	display := ssd1306.NewI2C(bus)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  32,
	})
	display.ClearDisplay()
	return &display
}
