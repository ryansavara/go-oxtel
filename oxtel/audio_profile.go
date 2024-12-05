package oxtel

import (
	"fmt"
	"strconv"
)

// SetAudioProfile sets the audio profile for the specified audio source.
// Valid Audio Profile values are 0 (use the default from SystemManager), 1-15. The Audio Profiles are configured in SystemManager.
func (o *Oxtel) SetAudioProfile(source OxtelAudioSource, profile uint8) error {
	if profile > 16 {
		return &InvalidAudioProfileError{
			BaseError: BaseError{
				Message: "Audio Profile must be less than 17",
			},
		}
	}

	msg := fmt.Sprintf("jAP%x%02x", source, profile)

	return o.sendCommand(msg)
}

// SetAudioProfile_AsString returns the command string used set the audio profile for the specified audio source.
//
// For use with scheduled commands.
func SetAudioProfile_AsString(source OxtelAudioSource, profile uint8) string {
	return fmt.Sprintf("jAP%x%02x", source, profile)
}

// EnquireAudioProfile queries the audio profile for the specified audio source.
//
// Valid Audio Profile values are 0 (use the default from SystemManager), 1-16. The Audio Profiles are configured in SystemManager.
//
// Response is an AudioProfileResponse.
func (o *Oxtel) EnquireAudioProfile(source OxtelAudioSource) (AudioProfileResponse, error) {
	val, err := o.sendCommandExpectResponse("jAP", fmt.Sprintf("%x", source))
	if err != nil {
		return AudioProfileResponse{}, err
	}

	audioSource, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return AudioProfileResponse{}, err
	}

	profile, err := strconv.ParseUint(string(val[1:]), 16, 8)

	if err != nil {
		return AudioProfileResponse{}, err
	}

	return AudioProfileResponse{
		Source:  OxtelAudioSource(audioSource),
		Profile: uint8(profile),
	}, err
}

// EnquireAudioProfile_AsString returns the command string used to query the audio profile for the specified audio source.
//
// For use with scheduled commands.
func EnquireAudioProfile_AsString(source OxtelAudioSource) string {
	return fmt.Sprintf("jAP%x", source)
}

// EnableAudioProfileTallies enables/disables VideoTally, ImageLoadTally, ImagePreloadTally, KeyerPositionTally on the
// connection on which the command was received.
func (o *Oxtel) EnableAudioProfileTallies(enable bool) error {
	msg := fmt.Sprintf("jAT%x", boolToInt(enable))

	return o.sendCommand(msg)
}

// EnableAudioProfileTallies_AsString returns the command string used to enable/disable VideoTally, ImageLoadTally,
// ImagePreloadTally, KeyerPositionTally on the connection on which the command was received.
//
// For use with scheduled commands.
func EnableAudioProfileTallies_AsString(enable bool) string {
	return fmt.Sprintf("jAT%x", boolToInt(enable))
}

// EnquireAudioProfileTallies queries the state of the Video Tallies set using the EnableVideoTallies command.
//
// Response is a boolean representing if audio profile tallies are enabled.
func (o *Oxtel) EnquireAudioProfileTallies() (bool, error) {
	val, err := o.sendCommandExpectResponse("jAT", "")
	if err != nil {
		return false, err
	}

	retval, err := strconv.ParseBool(val)
	return retval, err
}

// EnquireAudioProfileTallies_AsString returns the command string used to query the state of the Video Tallies set
// using the EnableVideoTallies command.
//
// For use with scheduled commands.
func EnquireAudioProfileTallies_AsString() string {
	return "jAT"
}
