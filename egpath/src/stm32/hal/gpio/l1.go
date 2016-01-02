// +build l1xx_md l1xx_mdp l1xx_hd l1xx_xl

package gpio

import (
	"mmio"
	"unsafe"

	"stm32/hal/raw/mmap"
	"stm32/hal/raw/rcc"
)

const (
	VeryLow Speed = 0 // 400 kHz (CL = 50 pF, VDD > 2.7 V)
	Low     Speed = 1 //   2 MHz (CL = 50 pF, VDD > 2.7 V)
	Medium  Speed = 2 //  10 MHz (CL = 50 pF, VDD > 2.7 V)
	High    Speed = 3 //  50 MHz (CL = 50 pF, VDD > 2.7 V)
)

const (
	TIM2 = AF1

	TIM3 = AF2
	TIM4 = AF2
	TIM5 = AF2

	TIM9  = AF3
	TIM10 = AF3
	TIM11 = AF3

	I2C1 = AF4
	I2C2 = AF4

	SPI1 = AF5
	SPI2 = AF5

	SPI3 = AF6

	USART1 = AF7
	USART2 = AF7
	USART3 = AF7

	UART4 = AF8
	UART5 = AF8

	USB = AF10

	LCD = AF11

	FSMC = AF12

	RI = AF14
)

var (
	A = (*Port)(unsafe.Pointer(mmap.GPIOA_BASE))
	B = (*Port)(unsafe.Pointer(mmap.GPIOB_BASE))
	C = (*Port)(unsafe.Pointer(mmap.GPIOC_BASE))
	D = (*Port)(unsafe.Pointer(mmap.GPIOD_BASE))
	E = (*Port)(unsafe.Pointer(mmap.GPIOE_BASE))
	F = (*Port)(unsafe.Pointer(mmap.GPIOF_BASE))
	G = (*Port)(unsafe.Pointer(mmap.GPIOG_BASE))
)

func pnum(p *Port) int {
	return int(uintptr(unsafe.Pointer(p))-mmap.AHBPERIPH_BASE) / 0x400
}

func enr() *mmio.U32   { return &rcc.RCC.AHBENR.U32 }
func lpenr() *mmio.U32 { return &rcc.RCC.AHBLPENR.U32 }
func rstr() *mmio.U32  { return &rcc.RCC.AHBRSTR.U32 }