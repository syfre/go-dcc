// package pi provides a Raspberry Pi driver for go-dcc.
//
// for L298
//   pi17 => EnA
//   pi27 => In1
//   pi22 => In2
//
//   L298
//     EnA   In1 In2
//      0     x   x   OFF
//      1     1   0   ON 
// Note that the Raspberry Pi needs to be equipped with an additional booster
// circuit in order to send the signal to the tracks. See the README.md for
// more information.

package dccpi

import (
	"fmt"
	"os"

	rpio "github.com/stianeikeland/go-rpio"
)

// GPIO Outputs for the Raspberry PI DCC encoder
var (
	BrakeGPIO  rpio.Pin = 17
	SignalGPIO rpio.Pin = 27
	NSignalGPIO rpio.Pin = 22
)

func init() {
	err := rpio.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot initialize GPIO: "+err.Error()+".\n")
		return
	}
	BrakeGPIO.Output()
	BrakeGPIO.Pull(rpio.PullUp)
	SignalGPIO.Output()
	NSignalGPIO.Output()
}

type DCCPi struct {
}

func NewDCCPi() (*DCCPi, error) {
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	return &DCCPi{}, nil
}

func (pi *DCCPi) Low() {
	SignalGPIO.Low()
	NSignalGPIO.High();
}

func (pi *DCCPi) High() {
	SignalGPIO.High()
	NSignalGPIO.Low()
}

func (pi *DCCPi) TracksOff() {
	BrakeGPIO.Low()
}

func (pi *DCCPi) TracksOn() {
	BrakeGPIO.High()
}
