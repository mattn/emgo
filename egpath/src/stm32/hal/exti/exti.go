package exti

import (
	"stm32/hal/gpio"

	"stm32/hal/raw/exti"
)

// Lines is bitmask that represents EXTI input lines.
type Lines uint32

const (
	L0 Lines = 1 << iota
	L1
	L2
	L3
	L4
	L5
	L6
	L7
	L8
	L9
	L10
	L11
	L12
	L13
	L14
	L15

	PVD    Lines = 1 << 16 // Programmable Voltage Detector output.
	RTCALR Lines = 1 << 17 // Real Time Clock Alarm event.
)

// Connect connects lines to corresponding pins of GPIO port. After reset lines
// 0..15 are connected to GPIO port A.
//
// Connect enables AFIO/SYSCFG clock before configuration and disables it before
// return.
func (lines Lines) Connect(port *gpio.Port) {
	if lines >= L15<<1 {
		panic("exti: can not connect lines to GPIO port")
	}
	exticrEna()
	p := uint32(port.Num())
	var n int
	for lines != 0 {
		if lines&0xf != 0 {
			r := exticr(n)
			if lines&1 != 0 {
				r.StoreBits(0x000f, p)
			}
			if lines&2 != 0 {
				r.StoreBits(0x00f0, p<<4)
			}
			if lines&4 != 0 {
				r.StoreBits(0x0f00, p<<8)
			}
			if lines&8 != 0 {
				r.StoreBits(0xf000, p<<12)
			}
		}
		lines = lines >> 4
		n++
	}
	exticrDis()
}

// RiseTrigEnabled returns lines that have rising edge detection enabled.
func RiseTrigEnabled() Lines {
	return Lines(exti.EXTI.RTSR.U32.Load())
}

// EnableRiseTrig enables rising edge detection for lines.
func (lines Lines) EnableRiseTrig() {
	exti.EXTI.RTSR.U32.SetBits(uint32(lines))
}

// DisableRiseTrig disables rising edge detection for lines.
func (lines Lines) DisableRiseTrig() {
	exti.EXTI.RTSR.U32.ClearBits(uint32(lines))
}

// FallTrigEnabled returns lines that have falling edge detection enabled.
func FallTrigEnabled() Lines {
	return Lines(exti.EXTI.FTSR.U32.Load())
}

// EnableFallTrig enables falling edge detection for lines.
func (lines Lines) EnableFallTrig() {
	exti.EXTI.FTSR.U32.SetBits(uint32(lines))
}

// DisableFallTrig disables falling edge detection for lines.
func (lines Lines) DisableFallTrig() {
	exti.EXTI.FTSR.U32.ClearBits(uint32(lines))
}

// Trigger allows to generate interrupt/event request by software. Interrupt
// pending flag on the line is set only when interrupt generation is enabled
// for this line.
func (lines Lines) Trigger() {
	exti.EXTI.SWIER.U32.Store(uint32(lines))
}

// IntEnabled returns lines that have interrupt generation enabled.
func IntEnabled() Lines {
	return Lines(exti.EXTI.IMR.U32.Load())
}

// EnableInt enables interrupt generation by lines.
func (lines Lines) EnableInt() {
	exti.EXTI.IMR.U32.SetBits(uint32(lines))
}

// DisableInt disable interrupt generation by lines.
func (lines Lines) DisableInt() {
	exti.EXTI.IMR.U32.ClearBits(uint32(lines))
}

// EventEnabled returns lines that have event generation enabled.
func EventEnabled() Lines {
	return Lines(exti.EXTI.EMR.Load())
}

// EnableEvent enables event generation by lines.
func (lines Lines) EnableEvent() {
	exti.EXTI.EMR.U32.SetBits(uint32(lines))
}

// DisableEvent disable event generation by lines.
func (lines Lines) DisableEvent() {
	exti.EXTI.EMR.U32.ClearBits(uint32(lines))
}

// Pending returns lines that have pending interrupt flag set.
func Pending() Lines {
	return Lines(exti.EXTI.PR.U32.Load())
}

// ClearPending clears pending interrupt flag for lines.
func (l Lines) ClearPending() {
	exti.EXTI.PR.U32.Store(uint32(l))
}
