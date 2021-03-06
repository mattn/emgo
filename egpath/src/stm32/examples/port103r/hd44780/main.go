// This example shows how to use PCF8574T + HD44780 combo.
//
// Connections:
// P0 --> RS
// P1 --> R/W
// P2 --> E
// P3 --> Backlight
// P4 <-> DB4
// P5 <-> DB5
// P6 <-> DB6
// P7 <-> DB7
//
// It seems that PCF8574T works up to 480 kHz (VCC = 5V, short cables, 16:9).
package main

import (
	"fmt"

	"hdc"

	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/rtc"

	"stm32/hal/i2c"
)

var twi = &i2c.Driver{Periph: i2c.I2C2}

func init() {
	system.Setup(8, 72/8, false)
	rtc.Setup(32768)

	gpio.B.EnableClock(true)
	port, pins := gpio.B, gpio.Pin10|gpio.Pin11

	cfg := gpio.Config{
		Mode:   gpio.Alt,
		Driver: gpio.OpenDrain,
	}
	port.Setup(pins, &cfg)
	twi.EnableClock(true)
	twi.Reset() // Mandatory!
	twi.Setup(&i2c.Config{Speed: 400e3})
	twi.SetIntMode(irq.I2C2_EV, irq.I2C2_ER)
	twi.Enable()
}

func checkErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Error: %v.\n", err)
	for {
	}
}

var (
	lcd = hdc.Display{
		Rows: 20, Cols: 4,
		ReadWriter: twi.NewMasterConn(0x27, i2c.ASRD),
		DS:         4,
		RS:         1 << 0, RW: 1 << 1, E: 1 << 2, AUX: 1 << 3,
	}
)

func main() {
	checkErr(lcd.Init())
	checkErr(lcd.SetDisplayMode(hdc.DisplayOn))
	checkErr(lcd.SetAUX())
	lcd.WriteString("Hello world!")
}

func twiISR() {
	twi.ISR()
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.RTCAlarm: rtc.ISR,
	irq.I2C2_EV:  twiISR,
	irq.I2C2_ER:  twiISR,
}
