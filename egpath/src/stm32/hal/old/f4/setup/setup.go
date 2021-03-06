package setup

import (
	"stm32/oldhal/f4/clock"
	"stm32/oldhal/f4/flash"
)

var (
	SysClk  uint // System clock [Hz]
	AHBClk  uint // AHB clock [Hz]
	APB1Clk uint // APB1 clock [Hz]
	APB2Clk uint // APB2 clock [Hz]
)

// Performance setups MCU for best performance (Flash prefetch, I/D cache on
// minimum allowed Flash latency).
// osc is freqency of external resonator in MHz. Allowed values are multiples
// of 2 MHz, from 4 MHz to 26 MHz. Use 0 to select internal HSI oscilator as
// system clock source. mul is PLL multipler. Allowed values are multiples of
// 24, from 96 to 216). sdiv determines system clock frequency according to
// the formula: sysclk = 2 * mul / sdiv. Performance ensures that pheripheral
// clock used by SDIO, RNG and USBFS is equal to 48 MHz (required by USBFS).
func Performance(osc, mul, sdiv int) {
	// Reset clock subsystem
	clock.ResetCR()
	clock.ResetPLLCFGR()
	clock.ResetCFGR()
	clock.ResetCIR()

	// Set HSI as temporary system clock source.
	clock.EnableHSI()
	for !clock.HSIReady() {
	}
	clock.SetSysClock(clock.HSI)
	for clock.SysClock() != clock.HSI {
	}

	if osc != 0 && (osc < 4 || osc > 26 || osc&1 != 0) {
		panic("wrong frequency of external resonator")
	}
	if mul < 96 || mul > 216 || mul%24 != 0 {
		panic("wrong PLL multipler")
	}
	switch sdiv {
	case 2, 4, 6, 8:
		// OK.
	default:
		panic("wrong PLL divider for SysClk")
	}

	// HSE needs millisecondss to stabilize, so enable it now.
	if osc != 0 {
		clock.EnableHSE()
	}

	sysclk := 2 * mul / sdiv
	SysClk = uint(sysclk) * 1e6

	lat := (sysclk - 1) / 30
	flash.SetLatency(lat) // Requires supply voltage 2.7-3.6 V.
	flash.SetPrefetch(true)
	flash.SetICache(true)
	flash.SetDCache(true)

	// Be sure that flash latency is set before incrase frequency.
	for flash.Latency() != lat {
	}

	// Configure clocks for AHB, APB1, APB2 bus.
	// APB1Freq must be <= 42 MHz.
	// APB2Freq must be <= 84 MHz.
	var div1, div2 clock.APBDiv
	switch {
	case sysclk <= 42:
		div1 = clock.APBDiv1
		div2 = clock.APBDiv1
		APB1Clk = SysClk
		APB2Clk = SysClk
	case sysclk <= 84:
		div1 = clock.APBDiv2
		div2 = clock.APBDiv1
		APB1Clk = SysClk / 2
		APB2Clk = SysClk
	case sysclk <= 168:
		div1 = clock.APBDiv4
		div2 = clock.APBDiv2
		APB1Clk = SysClk / 4
		APB2Clk = SysClk / 2
	default:
		div1 = clock.APBDiv8
		div2 = clock.APBDiv4
		APB1Clk = SysClk / 8
		APB2Clk = SysClk / 4
	}
	AHBClk = SysClk
	clock.SetPrescalerAHB(clock.AHBDiv1)
	clock.SetPrescalerAPB1(div1)
	clock.SetPrescalerAPB2(div2)

	if osc != 0 {
		// Be sure that HSE is ready.
		for !clock.HSEReady() {
		}
		clock.SetPLLClock(clock.SrcHSE)
		clock.SetPLLInputDiv(osc / 2)
	} else {
		clock.SetPLLClock(clock.SrcHSI)
		clock.SetPLLInputDiv(16 / 2)
	}

	clock.SetMainPLLMul(mul)
	clock.SetMainPLLSysDiv(sdiv)
	clock.SetMainPLLPeriphDiv(mul / 24) // Produces 48 MHz required by USBFS.
	clock.EnableMainPLL()
	for !clock.MainPLLReady() {
	}

	// Set PLL as system clock source
	clock.SetSysClock(clock.PLL)
	for clock.SysClock() != clock.PLL {
	}
	if osc != 0 {
		clock.DisableHSI()
	}

	sysclkChanged()
}

// Performance168 setups MCU to work with 168 MHz clock.
// See Performance for description of osc.
func Performance168(osc int) {
	Performance(osc, 168, 2)
}

// Performance84 setups MCU to work with 84 MHz clock.
// See Performance for description of osc.
func Performance84(osc int) {
	Performance(osc, 168, 4)
}
