package oxtel

import (
	"strconv"
)

// EnquireTemperature is implemented but always returns 0 for the temperature.
//
// Response is a float64 of the temperature, which is always 0.0
func (o *Oxtel) EnquireTemperature() (float64, error) {
	val, err := o.sendCommandExpectResponse("X0", "")
	if err != nil {
		return 0x00000, err
	}
	retval, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0x00000, err
	}

	return retval, err
}

// EnquireTemperature_AsString returns the command used to get the temperature.
//
// For use with scheduled commands.
func EnquireTemperature_AsString() string {
	return "X0"
}
