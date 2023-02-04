// https://cdn-shop.adafruit.com/datasheets/ht16K33v110.pdf
package ht16k33

import (
	"tinygo.org/x/drivers"
)

type Device struct {
	bus     drivers.I2C
	Address uint8
}

// New creates a new ht16k33 connection. The I2C bus must already be configured.
// This function only creates the Device object, it does not touch the device.
func New(bus drivers.I2C) *Device {
	return &Device{
		bus:     bus,
		Address: REG_BASE_HT_ADDR,
	}
}

// Configure sets up the device for communication
func (d *Device) Configure() {
	// FIXME:このタイミングですべてのLEDが点灯した状態になる。
	d.i2cWrite(REG_SS, COM_SS_NORMAL)               // Wakeup
	d.i2cWrite(REG_DSP, COM_DSP_ON|COM_DSP_NOBLINK) // Display on and no blinking
	d.i2cWrite(REG_RIS, COM_RIS_OUT)                // INT pin works as row output
	d.i2cWrite(REG_DIM, COM_DIM_16)                 // Brightness set to max

	// FIXME:最後に点灯状態を消灯時状態にする
}

func (d *Device) i2cWrite(reg, val uint8) error {
	return d.bus.WriteRegister(d.Address, reg, []uint8{val})
}
