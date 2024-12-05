package oxtel

import "fmt"

func boolToInt(input bool) int16 {
	if input {
		return 1
	}

	return 0
}

func intToBool(val int) (bool, error) {
	if val == 0 {
		return false, nil
	}
	if val == 1 {
		return true, nil
	}

	return false, &BaseError{
		Message: "Unable to convert integer to boolean",
	}

}

func buildMediaTallies(data MediaTallies) string {
	msg := ""
	if data.Unused1 {
		msg += "1"
	} else {
		msg += "0"
	}

	if data.Unused2 {
		msg += "1"
	} else {
		msg += "0"
	}

	if data.Unused3 {
		msg += "1"
	} else {
		msg += "0"
	}

	if data.Unused4 {
		msg += "1"
	} else {
		msg += "0"
	}

	if data.Unused5 {
		msg += "1"
	} else {
		msg += "0"
	}

	if data.Images {
		msg += "1"
	} else {
		msg += "0"
	}

	return msg
}

func buildAudioProgramCommand(program AudioProgram) (string, error) {
	if program.Channel1 > 15 {
		return "", &InvalidAudioProfileError{
			BaseError: BaseError{
				Message: "Channel 1 must be less than 16",
			},
		}
	}
	msg := fmt.Sprintf("jAL%02x%02x%02x%02x", program.Sdi, program.AudioProgram, program.JungerPreset, program.Channel1)

	if program.Channel2 != nil {
		if *program.Channel2 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 2 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, *program.Channel2)
	}
	if program.Channel3 != nil {
		if *program.Channel3 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 3 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel3)
	}
	if program.Channel4 != nil {
		if *program.Channel4 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 4 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel4)
	}
	if program.Channel5 != nil {
		if *program.Channel5 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 5 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel5)
	}
	if program.Channel6 != nil {
		if *program.Channel6 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 6 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel6)
	}
	if program.Channel7 != nil {
		if *program.Channel7 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 7 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel7)
	}
	if program.Channel8 != nil {
		if *program.Channel8 > 15 {
			return "", &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel 8 must be less than 16",
				},
			}
		}
		msg = fmt.Sprintf("%s%02x", msg, program.Channel8)
	}
	return msg, nil

}

func MakeTwoChannelMask() ChannelMask {
	return ChannelMask{
		Channel1:  true,
		Channel2:  true,
		Channel3:  false,
		Channel4:  false,
		Channel5:  false,
		Channel6:  false,
		Channel7:  false,
		Channel8:  false,
		Channel9:  false,
		Channel10: false,
		Channel11: false,
		Channel12: false,
		Channel13: false,
		Channel14: false,
		Channel15: false,
		Channel16: false,
	}
}

func MakeFourChannelMask() ChannelMask {
	return ChannelMask{
		Channel1:  true,
		Channel2:  true,
		Channel3:  true,
		Channel4:  true,
		Channel5:  false,
		Channel6:  false,
		Channel7:  false,
		Channel8:  false,
		Channel9:  false,
		Channel10: false,
		Channel11: false,
		Channel12: false,
		Channel13: false,
		Channel14: false,
		Channel15: false,
		Channel16: false,
	}
}

func MakeEightChannelMask() ChannelMask {
	return ChannelMask{
		Channel1:  true,
		Channel2:  true,
		Channel3:  true,
		Channel4:  true,
		Channel5:  true,
		Channel6:  true,
		Channel7:  true,
		Channel8:  true,
		Channel9:  false,
		Channel10: false,
		Channel11: false,
		Channel12: false,
		Channel13: false,
		Channel14: false,
		Channel15: false,
		Channel16: false,
	}
}

func MakeSixteenChannelMask() ChannelMask {
	return ChannelMask{
		Channel1:  true,
		Channel2:  true,
		Channel3:  true,
		Channel4:  true,
		Channel5:  true,
		Channel6:  true,
		Channel7:  true,
		Channel8:  true,
		Channel9:  true,
		Channel10: true,
		Channel11: true,
		Channel12: true,
		Channel13: true,
		Channel14: true,
		Channel15: true,
		Channel16: true,
	}
}

func buildChannelMask(mask ChannelMask) string {
	var outMask int16
	outMask |= boolToInt(mask.Channel1)
	outMask |= boolToInt(mask.Channel2) << 1
	outMask |= boolToInt(mask.Channel3) << 2
	outMask |= boolToInt(mask.Channel4) << 3
	outMask |= boolToInt(mask.Channel5) << 4
	outMask |= boolToInt(mask.Channel6) << 5
	outMask |= boolToInt(mask.Channel7) << 6
	outMask |= boolToInt(mask.Channel8) << 7
	outMask |= boolToInt(mask.Channel9) << 8
	outMask |= boolToInt(mask.Channel10) << 9
	outMask |= boolToInt(mask.Channel11) << 10
	outMask |= boolToInt(mask.Channel12) << 11
	outMask |= boolToInt(mask.Channel13) << 12
	outMask |= boolToInt(mask.Channel14) << 13
	outMask |= boolToInt(mask.Channel15) << 14
	outMask |= boolToInt(mask.Channel16) << 15

	return fmt.Sprintf("%04X", outMask)
}

// BuildSessionLocks takes in booleans for the mixer and each graphic layer to calculate the bitwise representation
// of the locks for SetSessionLocks.
func BuildSessionLocks(mixer bool, layer0 bool, layer1 bool, layer2 bool, layer3 bool, layer4 bool, layer5 bool, layer6 bool, layer7 bool) int32 {
	var locks int32
	if mixer {
		locks |= 0x00000001
	}
	if layer0 {
		locks |= 0x00000100
	}
	if layer1 {
		locks |= 0x00000200
	}
	if layer2 {
		locks |= 0x00000400
	}
	if layer3 {
		locks |= 0x00000800
	}
	if layer4 {
		locks |= 0x00001000
	}
	if layer5 {
		locks |= 0x00002000
	}
	if layer6 {
		locks |= 0x00004000
	}
	if layer7 {
		locks |= 0x00008000
	}

	return locks
}
