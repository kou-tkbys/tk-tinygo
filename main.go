package main

// TinyGo:pico
import (
	"image/color"
	"machine"
	"time"
	"tk-tinygo/ssd1306"
)

var period uint64 = 1e9 / 500

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	display := ssd1306.NewDisplay(machine.I2C0)
	x := int16(0)
	y := int16(0)
	deltaX := int16(1)
	deltaY := int16(1)
	for {
		pixel := display.GetPixel(x, y)
		c := color.RGBA{255, 255, 255, 255}
		if pixel {
			c = color.RGBA{0, 0, 0, 255}
		}
		display.SetPixel(x, y, c)
		display.Display()

		x += deltaX
		y += deltaY

		if x == 0 || x == 127 {
			deltaX = -deltaX
		}

		if y == 0 || y == 31 {
			deltaY = -deltaY
		}
		time.Sleep(1 * time.Millisecond)
	}
}
