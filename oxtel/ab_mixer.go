package oxtel

import (
	"fmt"
	"strconv"
)

// CutToA cuts the A/B mixer immediately to the A input of the mixer.
func (o *Oxtel) CutToA() error {
	return o.sendCommand("U0")
}

// CutToA_AsString returns the command string used to cut the A/B mixer immediately to the A input of the mixer.
//
// For use with scheduled commands.
func CutToA_AsString() string {
	return "U0"
}

// CutToB cuts the A/B mixer immediately to the B input of the mixer.
func (o *Oxtel) CutToB() error {
	return o.sendCommand("U1")
}

// CutToB_AsString returns the command string used to cut the A/B mixer immediately to the B input of the mixer.
//
// For use with scheduled commands.
func CutToB_AsString() string {
	return "U1"
}

// FadeToA fades the A/B mixer to the A input of the mixer over the specified number of fields (interlaces) or frames (progressive).
func (o *Oxtel) FadeToA(duration uint16) error {
	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000 fields/frames",
			},
		}
	}

	msg := fmt.Sprintf("U2%03x", duration)

	return o.sendCommand(msg)
}

// FadeToA_AsString returns the command string used to fade the A/B mixer to the A input of the mixer over the specified number
// of fields (interlaces) or frames (progressive).
//
// For use with scheduled commands.
func FadeToA_AsString(duration uint16) string {
	return fmt.Sprintf("U2%03x", duration)

}

// FadeToB fades the A/B mixer to the B input of the mixer over the specified number of fields (interlaces) or frames (progressive).
func (o *Oxtel) FadeToB(duration uint16) error {
	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000 fields/frames",
			},
		}
	}
	msg := fmt.Sprintf("U3%03x", duration)

	return o.sendCommand(msg)
}

// FadeToB_AsString returns the command string used to fade the A/B mixer to the B input of the mixer over the specified number
// of fields (interlaces) or frames (progressive).
//
// For use with scheduled commands.
func FadeToB_AsString(duration uint16) string {
	return fmt.Sprintf("U3%03x", duration)

}

// CutAB cuts the A/B mixer to the opposite input form the one currently visible.
func (o *Oxtel) CutAB() error {
	return o.sendCommand("U4")
}

// CutAB_AsString returns the command string used to cut the A/B mixer to the opposite input form the one currently visible.
//
// For use with scheduled commands.
func CutAB_AsString() string {
	return "U4"
}

// FadeAB fades the A/B mixer to the opposite input from the one currently visible over the specified number of
// fields (interlaces) or frames (progressive).
func (o *Oxtel) FadeAB(duration uint16) error {
	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000 fields/frames",
			},
		}
	}

	msg := fmt.Sprintf("U5%03x", duration)

	return o.sendCommand(msg)
}

// FadeAB_AsString returns the command string used to fade the A/B mixer to the opposite input from the one currently visible
// over the specified number of fields (interlaces) or frames (progressive).
//
// For use with scheduled commands.
func FadeAB_AsString(duration uint16) string {
	return fmt.Sprintf("U5%03x", duration)
}

// SetTransitionType selects the A/B mixer transition type used in the FadeToA, FadeToB, FadeAB and AsymmetricTransition commands.
func (o *Oxtel) SetTransitionType(transitionType OxtelTransitionType) error {
	msg := fmt.Sprintf("U6%02x", transitionType)

	return o.sendCommand(msg)
}

// SetTransitionType_AsString returns the command string used to select the A/B mixer transition type used in the
// FadeToA, FadeToB, FadeAB and AsymmetricTransition commands.
//
// For use with scheudled commands.
func SetTransitionType_AsString(transitionType OxtelTransitionType) string {
	return fmt.Sprintf("U6%02x", transitionType)
}

// AsymmetricVFadeAB instructs the A/B mixer to V-Fade from the current input to the other through the V-Fade color (black)
// over the specified number of fields (interlaces) or frames (progressive).
func (o *Oxtel) AsymmetricVFadeAB(downDuration uint16, upDuration uint16) error {
	if downDuration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "downDuration must be less than 1000 fields/frames",
			},
		}
	}

	if upDuration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "upDuration must be less than 1000 fields/frames",
			},
		}
	}

	msg := fmt.Sprintf("U8%03x%03x", downDuration, upDuration)

	return o.sendCommand(msg)
}

// AsymmetricVFadeAB_AsString returns the command string used to instruct the A/B mixer to V-Fade from the current input
// to the other through the V-Fade color (black) over the specified number of fields (interlaces) or frames (progressive).
//
// For use with scheduled commands.
func AsymmetricVFadeAB_AsString(downDuration uint16, upDuration uint16) string {
	return fmt.Sprintf("U8%03x%03x", downDuration, upDuration)
}

// SetAbsoluteMix sets the A/B mix position to the specified absolute value.
//
// The absolute mix value can range from 0 (A = 100%, B = 0%) to 512 (A = 0%, B = 100%).
func (o *Oxtel) SetAbsoluteMix(mix uint16) error {
	if mix > 512 {
		return &InvalidMixError{
			BaseError: BaseError{
				Message: "Absolute mix must be less than 513",
			},
		}
	}

	msg := fmt.Sprintf("U9%03x", mix)

	return o.sendCommand(msg)
}

// SetAbsoluteMix_AsString returns the command string used to set the A/B mix position to the specified absolute value.
//
// For use with scheduled commands.
func SetAbsoluteMix_AsString(mix uint16) string {
	return fmt.Sprintf("U9%03x", mix)
}

// AsymmetricTransition performs an A/B mixer transition to the destination specified.
//
// The transitionType used is defined by the SetTransitionType command. For Cut transitions, the duration parameter is ignored.
func (o *Oxtel) AsymmetricTransition(destination OxtelMixerInput, downDuration uint16, upDuration uint16) error {
	if downDuration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "DownDuration must be less than 1000 fields/frames",
			},
		}
	}

	if upDuration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "PpDuration must be less than 1000 fields/frames",
			},
		}
	}

	msg := fmt.Sprintf("UA%x%03x%03x", destination, downDuration, upDuration)

	return o.sendCommand(msg)
}

// AsymmetricTransition_AsString returns the command string used to perform an A/B mixer transition to the destination specified.
//
// For use with scheduled commands.
func AsymmetricTransition_AsString(destination OxtelMixerInput, downDuration uint16, upDuration uint16) string {
	return fmt.Sprintf("UA%x%03x%03x", destination, downDuration, upDuration)
}

// FadeToSpecificPosition transitions the A/B mixer to the specified position over the given number of fields (interlaced)
// or frames (progressive).
//
// The position can be from 0x000 to 0x200 and all values in-between.
func (o *Oxtel) FadeToSpecificPosition(destination uint16, duration uint16) error {
	if destination > 512 {
		return &InvalidMixError{
			BaseError: BaseError{
				Message: "Destination must be less than 513 fields/frames",
			},
		}
	}

	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000 fields/frames",
			},
		}
	}

	msg := fmt.Sprintf("UC%03x%03x", destination, duration)

	return o.sendCommand(msg)
}

// FadeToSpecificPosition_AsString returns the command string used to transition the A/B mixer to the specified position
// over the given number of fields (interlaced) or frames (progressive).
//
// For use with scheduled commands.
func FadeToSpecificPosition_AsString(destination uint16, duration uint16) string {
	return fmt.Sprintf("UC%03x%03x", destination, duration)
}

// SelectMixerInput routes video sources into the A and B input of the A/B mixer.
//
// The source parameter depends on the system and its configuration.
//
// An optional third argument is accepted as a Harmonic extension. This argument overrides the aspect ratio conversion
// of the video (typically the down-conversion at the SD simulcast output).
// The ARC only applies when the fader is showing the given input.
// When no third argument is specified, the ARC is reset back to default.
// This ARC mode overrides all other settings (including ARC inferred from AFD and specified on the player at clip attach time).
// The is currently no way of querying the ARC mode.
func (o *Oxtel) SelectMixerInput(input OxtelMixerInput, source OxtelVideoSource, arc *OxtelARC) error {
	msg := ""
	if arc == nil {
		msg = fmt.Sprintf("UE %x %x", input, source)
	} else {
		msg = fmt.Sprintf("UE %x %x %x", input, source, *arc)
	}

	return o.sendCommand(msg)
}

// SelectMixerInput_AsString returns the command string used to route video sources into the A and B input of the A/B mixer.
//
// For use with scheduled commands.
func SelectMixerInput_AsString(input OxtelMixerInput, source OxtelVideoSource, arc *OxtelARC) string {
	msg := ""
	if arc == nil {
		msg = fmt.Sprintf("UE %x %x", input, source)
	} else {
		msg = fmt.Sprintf("UE %x %x %x", input, source, *arc)
	}

	return msg
}

// EnquireMixerInput queries for the current video source for the specified mixer input.
//
// Response is a MixerInputResponse.
func (o *Oxtel) EnquireMixerInput(input OxtelMixerInput) (MixerInputResponse, error) {
	val, err := o.sendCommandExpectResponse("UE", fmt.Sprintf(" %x", input))
	if err != nil {
		return MixerInputResponse{}, err
	}

	in, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return MixerInputResponse{}, err
	}
	retval, err := strconv.ParseUint(string(val[1]), 16, 8)
	if err != nil {
		return MixerInputResponse{}, err
	}

	return MixerInputResponse{
		Input:  OxtelMixerInput(in),
		Source: OxtelVideoSource(retval),
	}, nil
}

// EnquireMixerInput_AsString returns the command string used to query for the current video source for the specified mixer input.
//
// For use with scheduled commands.
func EnquireMixerInput_AsString(input OxtelMixerInput) string {
	return fmt.Sprintf("UE %x", input)
}

// EnquireMixMode returns the status of the A/B mixer parameters.
//
// Response is a MixModeResponse.
func (o *Oxtel) EnquireMixMode() (MixModeResponse, error) {
	val, err := o.sendCommandExpectResponse("Ua", "")
	if err != nil {
		return MixModeResponse{}, err
	}

	transitionType, err := strconv.ParseUint(val[:2], 16, 8)
	if err != nil {
		return MixModeResponse{}, err
	}
	abmixRate, err := strconv.ParseUint(val[2:5], 16, 16)
	if err != nil {
		return MixModeResponse{}, err
	}
	wipeSoftness, err := strconv.ParseUint(val[5:8], 16, 8)
	if err != nil {
		return MixModeResponse{}, err
	}
	abmixAngle, err := strconv.ParseUint(val[8:11], 16, 16)
	if err != nil {
		return MixModeResponse{}, err
	}
	vfadeColor, err := strconv.ParseUint(val[11:17], 16, 64)
	if err != nil {
		return MixModeResponse{}, err
	}

	return MixModeResponse{
		TransitionType: uint8(transitionType),
		ABMixRate:      uint16(abmixRate),
		WipeSoftness:   uint8(wipeSoftness),
		ABMixAngle:     uint16(abmixAngle),
		VFadeColor:     uint8(vfadeColor),
	}, err

}

// EnquireMixMode_AsString returns the command string used to return the status of the A/B mixer parameters.
//
// For use with scheduled commands.
func EnquireMixMode_AsString() string {
	return "Ua"
}

// SetColorGeneratorColor sets the color of the specified Color Generator unit to the specified RGB value.
func (o *Oxtel) SetColorGeneratorColor(colorGeneratorUnit uint8, red uint8, green uint8, blue uint8) error {
	msg := fmt.Sprintf("UZ%x%02x%02x%02x", colorGeneratorUnit, red, green, blue)

	return o.sendCommand(msg)
}

// SetColorGeneratorColor_AsString returns the string used to set the color of the specified Color Generator unit to
// the specified RGB value.
func SetColorGeneratorColor_AsString(colorGeneratorUnit uint8, red uint8, green uint8, blue uint8) string {
	return fmt.Sprintf("UZ%x%02x%02x%02x", colorGeneratorUnit, red, green, blue)
}

// EnquireColorGeneratorColor queries the color generator color from the specified Color Generator unit.
//
// Response is a ColorGeneratorResponse.
func (o *Oxtel) EnquireColorGeneratorColor(colorGeneratorUnit uint8) (ColorGeneratorResponse, error) {
	val, err := o.sendCommandExpectResponse("UZ", fmt.Sprintf("%d", colorGeneratorUnit))
	if err != nil {
		return ColorGeneratorResponse{}, err
	}

	unit, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return ColorGeneratorResponse{}, err
	}

	red, err := strconv.ParseUint(val[1:3], 16, 8)
	if err != nil {
		return ColorGeneratorResponse{}, err
	}

	green, err := strconv.ParseUint(val[3:5], 16, 8)
	if err != nil {
		return ColorGeneratorResponse{}, err
	}

	blue, err := strconv.ParseUint(val[5:7], 16, 8)
	if err != nil {
		return ColorGeneratorResponse{}, err
	}

	return ColorGeneratorResponse{
		Unit:  uint8(unit),
		Red:   uint8(red),
		Green: uint8(green),
		Blue:  uint8(blue),
	}, err

}

// EnquireColorGeneratorColor_AsString returns the command string used to query the color generator color from the
// specified Color Generator unit.
//
// For use with scheduled commands.
func EnquireColorGeneratorColor_AsString(colorGeneratorUnit uint8) string {
	return fmt.Sprintf("UZ%d", colorGeneratorUnit)
}

// EnableVideoTallies enables/disables the Video Tally, Image Video Tally, Image Preload Tally, and Keyer Position Tally on
// the connection on which the command was received.
func (o *Oxtel) EnableVideoTallies(enable bool) error {
	msg := fmt.Sprintf("Y6%x", boolToInt(enable))

	return o.sendCommand(msg)
}

// EnableVideoTallies_AsString returns the command string used to  enable/disable the Video Tally, Image Video Tally,
// Image Preload Tally, and Keyer Position Tally on the connection on which the command was received.
func EnableVideoTallies_AsString(enable bool) string {
	return fmt.Sprintf("Y6%x", boolToInt(enable))
}

// EnquireVideoTallies queries the state of the Video Tallies set using the EnableVideoTallies command.
//
// Response is a boolean representing if video tallies are enabled.
func (o *Oxtel) EnquireVideoTallies() (bool, error) {
	val, err := o.sendCommandExpectResponse("Y6", "")
	if err != nil {
		return false, err
	}

	retval, err := strconv.ParseBool(val)
	return retval, err
}

// EnquireVideoTallies_AsString returns the command string used to query the state of the Video Tallies set using
// the EnableVideoTallies command.
//
// For use with scheudled commands.
func EnquireVideoTallies_AsString() string {
	return "Y6"
}
