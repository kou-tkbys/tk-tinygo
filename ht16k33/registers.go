package ht16k33

// "address" is base address 0-7 which becomes 11100xxx = E0-E7
const (
	REG_BASE_HT_ADDR uint8 = 0x70
)

//
const (
	REG_DDAP uint8 = 0x00 // Display data Address pointer: 0000xxxx
	REG_KDAP uint8 = 0x40 // Key data Address pointer
	REG_SS   uint8 = 0x20 // System setup(20H)
	REG_RIS  uint8 = 0xA0 // ROW/INT set(A0H)
	REG_DSP  uint8 = 0x80 // Display setup(80H)
	REG_DIM  uint8 = 0xef // Dimming set(EFH)
)

// Commands
const (
	COM_SS_STANDBY uint8 = 0b00000000 // System setup - oscillator in standby mode
	COM_SS_NORMAL  uint8 = 0b00000001 // System setup - oscillator in normal mode

	COM_IFAP uint8 = 0b01100000 // Read status of INT flag

	COM_DSP_OFF       uint8 = 0b00000000 // Display setup - display off
	COM_DSP_ON        uint8 = 0b00000001 // Display setup - display on
	COM_DSP_NOBLINK   uint8 = 0b00000000 // Display setup - no blink
	COM_DSP_BLINK2HZ  uint8 = 0b00000010 // Display setup - 2hz blink
	COM_DSP_BLINK1HZ  uint8 = 0b00000100 // Display setup - 1hz blink
	COM_DSP_BLINK05HZ uint8 = 0b00000110 // Display setup - 0.5hz blink

	COM_RIS_OUT  uint8 = 0b00000000 // Set INT as row driver output
	COM_RIS_INTL uint8 = 0b00000001 // Set INT as int active low
	COM_RIS_INTH uint8 = 0b00000011 // Set INT as int active high

	COM_DIM    uint8 = 0b11100000 // Dimming set
	COM_DIM_1  uint8 = 0b00000000 // Dimming set - 1/16
	COM_DIM_2  uint8 = 0b00000001 // Dimming set - 2/16
	COM_DIM_3  uint8 = 0b00000010 // Dimming set - 3/16
	COM_DIM_4  uint8 = 0b00000011 // Dimming set - 4/16
	COM_DIM_5  uint8 = 0b00000100 // Dimming set - 5/16
	COM_DIM_6  uint8 = 0b00000101 // Dimming set - 6/16
	COM_DIM_7  uint8 = 0b00000110 // Dimming set - 7/16
	COM_DIM_8  uint8 = 0b00000111 // Dimming set - 8/16
	COM_DIM_9  uint8 = 0b00001000 // Dimming set - 9/16
	COM_DIM_10 uint8 = 0b00001001 // Dimming set - 10/16
	COM_DIM_11 uint8 = 0b00001010 // Dimming set - 11/16
	COM_DIM_12 uint8 = 0b00001011 // Dimming set - 12/16
	COM_DIM_13 uint8 = 0b00001100 // Dimming set - 13/16
	COM_DIM_14 uint8 = 0b00001101 // Dimming set - 14/16
	COM_DIM_15 uint8 = 0b00001110 // Dimming set - 15/16
	COM_DIM_16 uint8 = 0b00001111 // Dimming set - 16/16
)
