package oxtel

import (
	"fmt"
	"strconv"
)

// SetAudioLoudness reconfigures or creates a new Loudness instance on the specified audio channels.
//
// Multiple SetAudioLoudness commands can be use to set multiple loudness groups within the output stream.
// SetAudioLoudness may be combined with AddScheduledCommand to occur at a specific time.
//
// If the Junger Preset argument is set to zero, loudness is turned off for the specified audio program that resides
// on the specified channels(s).
// SetAudioLoudness resets the Junger library processing history, if the new settings match the previous ones.
//
// For Stereo unused audio channels Ch3 - Ch6 are not specified. Also note that only the first channel is required.
// If only the first channel is specified for multi-channel program, the remaining channels will be assigned sequentially.
//
// The second audio program is optional.
func (o *Oxtel) SetAudioLoudness(firstProgram AudioProgram, secondProgram *AudioProgram) error {
	msg, err := buildAudioProgramCommand(firstProgram)
	if err != nil {
		return err
	}

	msg2 := ""
	if secondProgram != nil {
		msg2, err = buildAudioProgramCommand(*secondProgram)
		if err != nil {
			return err
		}
	}

	return o.sendCommand(fmt.Sprintf("%s%s", msg, msg2))
}

// SetAudioLoudness_AsString returns the command string used to reconfigure or create a new Loudness instance on the
// specified audio channels.
//
// For use with scheduled commands.
func SetAudioLoudness_AsString(firstProgram AudioProgram, secondProgram *AudioProgram) string {
	msg, _ := buildAudioProgramCommand(firstProgram)

	msg2 := ""
	if secondProgram != nil {
		msg2, _ = buildAudioProgramCommand(*secondProgram)
	}

	return fmt.Sprintf("%s%s", msg, msg2)
}

// GetAudioLoudness returns all audio channels that have a loudness program configured.
//
// Optionally specify the channel number to get the loudness setting on a specific channel.
//
// If loundness is not licensed, this will always return "jAL0000".
//
// Response is a AudioLoudnessResponse. Sdi, AudioProgram, and JungarProgram are only populated when a channel is specified.
func (o *Oxtel) GetAudioLoudness(sdi uint8, channel *uint8) (AudioLoudnessResponse, error) {
	if channel == nil {
		// Return all audio channels that have loundess configured.
		val, err := o.sendCommandExpectResponse("jAL", fmt.Sprintf("%02x", sdi))
		if err != nil {
			return AudioLoudnessResponse{}, err
		}

		ch1 := false
		ch2 := false
		ch3 := false
		ch4 := false
		ch5 := false
		ch6 := false
		ch7 := false
		ch8 := false

		for i := 0; i < len(val); i += 2 {
			num, err := strconv.ParseUint(string(val[i:i+2]), 16, 8)
			if err != nil {
				return AudioLoudnessResponse{}, err
			}

			if num == 1 {
				ch1 = true
			}
			if num == 2 {
				ch2 = true
			}
			if num == 3 {
				ch3 = true
			}
			if num == 4 {
				ch4 = true
			}
			if num == 5 {
				ch5 = true
			}
			if num == 6 {
				ch6 = true
			}
			if num == 7 {
				ch7 = true
			}
			if num == 8 {
				ch8 = true
			}
		}

		return AudioLoudnessResponse{
			Channel1: ch1,
			Channel2: ch2,
			Channel3: ch3,
			Channel4: ch4,
			Channel5: ch5,
			Channel6: ch6,
			Channel7: ch7,
			Channel8: ch8,
		}, nil
	} else {
		if *channel > 8 {
			return AudioLoudnessResponse{}, &InvalidAudioProfileError{
				BaseError: BaseError{
					Message: "Channel must be less than 8",
				},
			}
		}

		val, err := o.sendCommandExpectResponse("jAL", fmt.Sprintf("%02x%02x", sdi, *channel))
		if err != nil {
			return AudioLoudnessResponse{}, err
		}

		outSdi, err := strconv.ParseUint(string(val[0]), 16, 8)
		if err != nil {
			return AudioLoudnessResponse{}, err
		}
		sdiPtr := &outSdi
		sdiUint8 := uint8(*sdiPtr)

		program, err := strconv.ParseUint(string(val[1:3]), 16, 8)
		if err != nil {
			return AudioLoudnessResponse{}, err
		}
		pgmPtr := &program
		programUint8 := uint8(*pgmPtr)

		preset, err := strconv.ParseUint(string(val[3:5]), 16, 8)
		if err != nil {
			return AudioLoudnessResponse{}, err
		}
		presetPtr := &preset
		presetUint8 := uint8(*presetPtr)

		ch1 := false
		ch2 := false
		ch3 := false
		ch4 := false
		ch5 := false
		ch6 := false
		ch7 := false
		ch8 := false

		for i := 5; i < len(val); i += 2 {
			num, err := strconv.ParseUint(string(val[i:i+2]), 16, 8)
			if err != nil {
				return AudioLoudnessResponse{}, err
			}

			if num == 1 {
				ch1 = true
			}
			if num == 2 {
				ch2 = true
			}
			if num == 3 {
				ch3 = true
			}
			if num == 4 {
				ch4 = true
			}
			if num == 5 {
				ch5 = true
			}
			if num == 6 {
				ch6 = true
			}
			if num == 7 {
				ch7 = true
			}
			if num == 8 {
				ch8 = true
			}
		}

		return AudioLoudnessResponse{
			Channel1:     ch1,
			Channel2:     ch2,
			Channel3:     ch3,
			Channel4:     ch4,
			Channel5:     ch5,
			Channel6:     ch6,
			Channel7:     ch7,
			Channel8:     ch8,
			AudioProgram: (*OxtelAudioProgram)(&programUint8),
			JungerPreset: (*OxtelJungerPreset)(&presetUint8),
			Sdi:          (*uint8)(&sdiUint8),
		}, nil
	}
}

// GetAudioLoudness_AsString returns the command string used to return all audio channels that have a loudness program configured.
//
// For use with scheduled commands.
func GetAudioLoudness_AsString(sdi uint8, channel *uint8) string {
	if channel == nil {
		return fmt.Sprintf("jAL%02x", sdi)
	} else {
		return fmt.Sprintf("jAL%02x%02x", sdi, *channel)
	}
}

// DisableAudioLoudness stops all loundness processing.
func (o *Oxtel) DisableAudioLoudness(sdi uint8) error {
	if sdi > 2 {
		return &InvalidSdiError{
			BaseError: BaseError{
				Message: "SDI must be less than 3",
			},
		}
	}

	msg := fmt.Sprintf("jALP%02x", sdi)
	return o.sendCommand(msg)
}

// DisableAudioLoudness_AsString returns the commands string used to stop all loundness processing.
//
// For use with scheduled commands.
func DisableAudioLoudness_AsString(sdi uint8) string {
	return fmt.Sprintf("jALP%02x", sdi)
}

// EnableAudioLoudness resumes loudness processing.
//
// If audio loudness was not disabled, this command has no effect.
func (o *Oxtel) EnableAudioLoudness(sdi uint8) error {
	if sdi > 2 {
		return &InvalidSdiError{
			BaseError: BaseError{
				Message: "SDI must be less than 3",
			},
		}
	}

	msg := fmt.Sprintf("jALR%02x", sdi)
	return o.sendCommand(msg)
}

// EnableAudioLoudness_AsString returns the command string used to resume loudness processing.
//
// For use with scheduled commands.
func EnableAudioLoudness_AsString(sdi uint8) string {
	return fmt.Sprintf("jALR%02x", sdi)
}

// ChangeAudioLoudnessProfile sets the Junger audio Loundness profile for the specified audio program.
//
// Audio profiles are set up in SystemManager. There can be up to 16 loudness profiles defined.
// Only one profile is in effect at a time, and profiles other than 1, can only be selected with the SetAudioLoudness command.
func (o *Oxtel) ChangeAudioLoudnessProfile(sdi uint8, profile uint8) error {
	if sdi > 2 {
		return &InvalidSdiError{
			BaseError: BaseError{
				Message: "SDI must be less than 3",
			},
		}
	}

	if profile > 16 {
		return &InvalidAudioProfileError{
			BaseError: BaseError{
				Message: "Audio Profile must be less than 17",
			},
		}
	}

	msg := fmt.Sprintf("jALA%02x%02x", sdi, profile)

	return o.sendCommand(msg)
}

// ChangeAudioLoudnessProfile_AsString returns the command string used to set the Junger audio Loundness profile
// for the specified audio program.
//
// For use with scheduled commands.
func ChangeAudioLoudnessProfile_AsString(sdi uint8, profile uint8) string {
	return fmt.Sprintf("jALA%02x%02x", sdi, profile)
}

// GetAudioLoudnessProfile returns the current Audio Loudness Profile for the specified program.
//
// Response is a uint8 representing the current profile.
func (o *Oxtel) GetAudioLoudnessProfile(sdi uint8) (uint8, error) {
	if sdi > 2 {
		return 0, &InvalidSdiError{
			BaseError: BaseError{
				Message: "SDI must be less than 3",
			},
		}
	}

	val, err := o.sendCommandExpectResponse("jALA", fmt.Sprintf("%02x", sdi))
	if err != nil {
		return 0, err
	}

	retval, err := strconv.ParseUint(val, 16, 8)
	return uint8(retval), err
}

// GetAudioLoudnessProfile_AsString returns the command string used to return the current Audio Loudness Profile
// for the specified program.
//
// For use with scheduled commands.
func GetAudioLoudnessProfile_AsString(sdi uint8) string {
	return fmt.Sprintf("jALA%02x", sdi)
}

// GetLoudnessLicenseStatus queries the current loudness license status.
//
// Response is a boolean representing if loudness is licensed.
func (o *Oxtel) GetLoudnessLicenseStatus() (bool, error) {
	val, err := o.sendCommandExpectResponse("jALL", "")
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(val[1]))
}

// GetLoudnessLicenseStatus_AsStrinc returns the command string used to query the current loudness license status.
//
// For use with scheduled commands.
func GetLoudnessLicenseStatus_AsString() string {
	return "jALL"

}
