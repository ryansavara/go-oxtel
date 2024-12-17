package oxtel

import "fmt"

// FadeKeyer fades the specified graphic keying layer in the specified number of fields.
// If a prior fade is not finished when a reverse command is received, the fade transition will reverse direction and
// continue at the same rate.
//
// The direction specifies the direciton of the fade transition.
//
// The rate parameter specifies the number of fields (interlaced) or frames (progressive) in the fade transition.
// When not specified, the value set using the SetTransitionDuration command will be used. The default value is 0.
//
// An unsolicited response will be transmitted as a OxtelKeyerPositionTally
func (o *Oxtel) FadeKeyer(layer OxtelLayer, direction OxtelDirection, rate *uint16) error {
	msg := ""
	if rate == nil {
		msg = fmt.Sprintf("1%x %d", layer, direction)
	} else {
		if *rate > 999 {
			return &InvalidDurationError{
				BaseError: BaseError{
					Message: "Rate must be less than 1000",
				},
			}
		}
		msg = fmt.Sprintf("1%x %d %x", layer, direction, rate)
	}

	return o.sendCommand(msg)
}

// FadeKeyer_AsString returns the command string used to fade the specified graphic keying layer in the specified number of fields.
//
// For use with scheduled commands.
func FadeKeyer_AsString(layer OxtelLayer, direction OxtelDirection, rate *uint16) string {
	if rate == nil {
		return fmt.Sprintf("1%x %d", layer, direction)
	} else {
		return fmt.Sprintf("1%x %d %x", layer, direction, rate)
	}
}

// CutKeyer cuts the specified graphic keying layer up or down. This command is the same as the FadeKeyer command executed
// with a rate of 0.
//
// The direction specifies the direciton of the fade transition.
//
// An unsolicited response will be transmitted as a OxtelKeyerPositionTally
func (o *Oxtel) CutKeyer(layer OxtelLayer, direction OxtelDirection) error {
	msg := fmt.Sprintf("3%x %d", layer, direction)

	return o.sendCommand(msg)
}

// CutKeyer_AsString returns the command string used to cut the specified graphic keying layer up or down. This command is
// the same as the FadeKeyer command executed with a rate of 0.
//
// For use with scheduled commands.
func CutKeyer_AsString(layer OxtelLayer, direction OxtelDirection) string {
	return fmt.Sprintf("3%x %d", layer, direction)
}

// SetTransitionDuration sets the keyer fade duration for the specified layer. This value is used by the FadeKeyer command
// when the rate parameter is not specified.
//
// The rate parameter specifies the number of fields (interlaced) or frames (progressive).
//
// An unsolicited response will be transmitted as a OxtelKeyerPositionTally
func (o *Oxtel) SetTransitionDuration(layer OxtelLayer, rate uint16) error {
	if rate > 999 {
		return &InvalidRateError{
			BaseError: BaseError{
				Message: "Rate must be less than 1000 fields/frames",
			},
		}
	}
	msg := fmt.Sprintf("B%x 1 %x", layer, rate)

	return o.sendCommand(msg)
}

// SetTransitionDuration_AsString returns the command string used to set the keyer fade duration for the specified layer.
// This value is used by the FadeKeyer command when the rate parameter is not specified.
//
// For use with scheduled commands.
func SetTransitionDuration_AsString(layer OxtelLayer, rate uint16) string {
	return fmt.Sprintf("B%x 1 %x", layer, rate)
}

// SetFaderAngle sets the keyer fader for the specified layer to an absolute level.
func (o *Oxtel) SetFaderAngle(layer OxtelLayer, angle uint16) error {
	if angle > 512 {
		return &InvalidAngleError{
			BaseError: BaseError{
				Message: "Angle must be less than 513",
			},
		}
	}
	msg := fmt.Sprintf("@%x 1 %x", layer, angle)

	return o.sendCommand(msg)
}

// SetFaderAngle_AsString returns the command string used to set the keyer fader for the specified layer to an absolute level.
//
// For use with scheduled commands.
func SetFaderAngle_AsString(layer OxtelLayer, angle uint16) string {
	return fmt.Sprintf("@%x 1 %x", layer, angle)
}
