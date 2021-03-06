package irq

import "arch/cortexm/nvic"

const (
	WinWdg nvic.IRQ = iota
	PVD
	TampStamp
	RTCWkup
	Flash
	RCC
	Ext0
	Ext1
	Ext2
	Ext3
	Ext4
	DMA1Stream0
	DMA1Stream1
	DMA1Stream2
	DMA1Stream3
	DMA1Stream4
	DMA1Stream5
	DMA1Stream6
	ADC
	CAN1Tx
	CAN1Rx0
	CAN1Rx1
	CAN1SCE
	Ext9_5
	Tim1BrkTim9
	Tim1UpTim10
	Tim1TrgComTim11
	Tim1CC
	Tim2
	Tim3
	Tim4
	I2C1Ev
	I2C1Er
	I2C2Ev
	I2C2Er
	SPI1
	SPI2
	USART1
	USART2
	USART3
	Ext15_10
	RTCAlarm
	OTGFSWkup
	Tim8BrkTim12
	Tim8UpTim13
	Tim8TrgComTim14
	Tim8CC
	DMA1Stream7
	FSMC
	SDIO
	Tim5
	SPI3
	UART4
	UART5
	Tim6DAC
	Tim7
	DMA2Stream0
	DMA2Stream1
	DMA2Stream2
	DMA2Stream3
	DMA2Stream4
	Eth
	EthWkup
	CAN2Tx
	CAN2Rx0
	CAN2Rx1
	CAN2SCE
	OTGFS
	DMA2Stream5
	DMA2Stream6
	DMA2Stream7
	USART6
	I2C3Ev
	I2C3Er
	OTGHSEP1In
	OTGHSWkup
	OTGHS
	DCMI
	CRYP
	HashRng
	FPU
)
