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
// 任意のI2Cバスと通信するデバイスオブジェクトを生成します。
// デバイスオブジェクトを生成するだけで、実デバイスへタッチしていません。
func New(bus drivers.I2C) Device {
	return Device{
		bus:     bus,
		Address: REG_BASE_HT_ADDR,
	}
}

// Configure sets up the device for communication
func (d *Device) Configure() {
	// FIXME:このタイミングですべてのLEDが点灯した状態になる。
	d.bus.WriteRegister(d.Address, REG_SS, []uint8{COM_SS_NORMAL})                 // Wakeup
	d.bus.WriteRegister(d.Address, REG_DSP, []uint8{COM_DSP_ON | COM_DSP_NOBLINK}) // Display on and no blinking
	d.bus.WriteRegister(d.Address, REG_RIS, []uint8{COM_RIS_OUT})                  // INT pin works as row output
	d.bus.WriteRegister(d.Address, REG_DIM, []uint8{COM_DIM_16})                   // Brightness set to max
}
