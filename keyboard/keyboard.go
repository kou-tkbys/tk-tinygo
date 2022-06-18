package keyboard

import (
	"machine"
)

const (
	_KEY_UP   uint8 = 0
	_KEY_DOWN uint8 = 1
)

type keyState = uint8

const (
	up keyState = iota
	down
)

type Keyboard struct {
	rowPins   []machine.Pin
	rowCnt    uint8
	colPins   []machine.Pin
	colCnt    uint8
	keyStates [][]keyState
}

func New(rps, cps []machine.Pin) *Keyboard {
	rpsLen := len(rps)
	cpsLen := len(cps)
	return &Keyboard{
		rowPins:   rps,
		rowCnt:    uint8(rpsLen),
		colPins:   cps,
		colCnt:    uint8(cpsLen),
		keyStates: make([][]uint8, rpsLen, cpsLen),
	}
}
