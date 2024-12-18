package oxtel

import (
	"fmt"
	"strconv"
)

// SetAudioABMixerFadeRate sets the fade rate for audio mixes.
func (o *Oxtel) SetAudioABMixerFadeRate(duration uint16) error {
	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000",
			},
		}
	}

	msg := fmt.Sprintf("j31%03x", duration)

	return o.sendCommand(msg)
}

// SetAudioABMixerFadeRate_AsStrig returns the command string used to set the fade rate for audio mixes.
//
// For use with scheduled commands.
func SrtAudioABMixerFadeRate_AsString(duration uint16) string {
	return fmt.Sprintf("j31%03x", duration)
}

// AudioCutAB cuts audio between the A and B sources
func (o *Oxtel) AudioCutAB(destination OxtelMixerInput) error {
	msg := fmt.Sprintf("j40%x", destination)

	return o.sendCommand(msg)
}

// AudioCutAB_AsString returns the command string used to cut audio between the A and B sources
//
// For use with scheduled commands.
func AudioCutAB_AsString(destination OxtelMixerInput) string {
	return fmt.Sprintf("j40%x", destination)
}

// AudioFadeAB fades audio between the A and B sources using the duration specified by the SetAudioABMixerFadeRate command
// and the mode specified by the SetAudioABMixMode command.
func (o *Oxtel) AudioFadeAB(destination OxtelMixerInput) error {
	msg := fmt.Sprintf("j41%x", destination)

	return o.sendCommand(msg)
}

// AudioFadeAB_AsString returns the command string used to fade audio between the A and B sources using the duration specified
// by the SetAudioABMixerFadeRate command and the mode specified by the SetAudioABMixMode command.
//
// For use with scheduled commands.
func AudioFadeAB_AsString(destination OxtelMixerInput) string {
	return fmt.Sprintf("j41%x", destination)
}

// SetAudioABFollowVideoAB allows the audio A/B mixer to automatically follow the position of the video A/B mixer.
//
// When enabled, the audio A/B mixer will follow the video A/B mixer. When disabled, the audio A/B mixer must be
// controlled independently.
func (o *Oxtel) SetAudioABFollowVideoAB(enable bool) error {
	msg := fmt.Sprintf("j51%x", boolToInt(enable))

	return o.sendCommand(msg)
}

// SetAudioABFollowVideoAB_AsString returns the command string used to allow the audio A/B mixer to automatically
// follow the position of the video A/B mixer.
//
// For use with scheduled commands.
func SetAudioABFollowVideoAB_AsString(enable bool) string {
	return fmt.Sprintf("j51%x", boolToInt(enable))
}

// EnquireAudioABFollowVideoAB queries the current audio follow video settings.
//
// Response is an AudioABFollowVideoABResponse.
func (o *Oxtel) EnquireAudioABFollowVideoAB() (AudioABFollowVideoABResponse, error) {
	val, err := o.sendCommandExpectResponse("j74", "")
	if err != nil {
		return AudioABFollowVideoABResponse{}, err
	}

	unused, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return AudioABFollowVideoABResponse{}, err
	}
	enabled, err := strconv.ParseBool(string(val[1]))

	return AudioABFollowVideoABResponse{
		Unused1: uint8(unused),
		Enabled: enabled,
	}, err
}

// EnquireAudioABFollowVideoAB_AsString returns the command string used to query the current audio follow video settings.
//
// For use with scheduled commands.
func EnquireAudioABFollowVideoAB_AsString() string {
	return "j74"
}

// SetAudioABPosition sets the position of the audio A/B mixer.
func (o *Oxtel) SetAudioABPosition(mix uint16) error {
	if mix > 512 {
		return &InvalidMixError{
			BaseError: BaseError{
				Message: "Mix must be less than 513",
			},
		}
	}

	msg := fmt.Sprintf("ja%03x", mix)

	return o.sendCommand(msg)
}

// SetAudioABPosition_AsString returns the command string used to set the position of the audio A/B mixer.
//
// For use with scheduled commands.
func SetAudioABPosition_AsString(mix uint16) string {
	return fmt.Sprintf("ja%03x", mix)
}

// SetAudioABMixMode sets the audio mixer mode.
func (o *Oxtel) SetAudioABMixMode(mode OxtelAudioMixMode) error {
	msg := fmt.Sprintf("jb%x", mode)

	return o.sendCommand(msg)
}

// SetAudioABMixMode_AsString returns the command string used to set the audio mixer mode.
//
// For use with scheduled commands.
func SetAudioABMixMode_AsString(mode OxtelAudioMixMode) string {
	return fmt.Sprintf("jb%x", mode)
}

// AudioABAsymmetricTransition performs an asymmetric transition such as a cut-fade or fade-cut.
//
// rate1 defines the duration of the first half of the fade (time it takes to reach the midpoint).
//
// rate2 is the time from the midpoint to completeion.
//
// Both rates are measured in fields (interlaces) or frames (progressive).
func (o *Oxtel) AudioABAsymmetricTransition(direction OxtelMixerInput, rate1 uint16, rate2 uint16) error {
	msg := fmt.Sprintf("jc%x%03x%03x", direction, rate1, rate2)

	return o.sendCommand(msg)
}

// AudioABAsymmetricTransition_AsString returns the command string used to perform an asymmetric transition such as
// a cut-fade or fade-cut.
//
// For use with scheduled commands.
func AudioABAsymmetricTransition_AsString(direction OxtelMixerInput, rate1 uint16, rate2 uint16) string {
	return fmt.Sprintf("jc%x%03x%03x", direction, rate1, rate2)
}

// AudioABFadeToPosition fades to the specified position over the defined duration in fields (interlaces) or frames (progressive).
//
// This is similar to the AudioABPosition command but provides a fade to the specified position instead of a cut.
func (o *Oxtel) AudioABFadeToPosition(mix uint16, duration uint16) error {
	if mix > 512 {
		return &InvalidMixError{
			BaseError: BaseError{
				Message: "Mix must be less than 513",
			},
		}
	}

	if duration > 999 {
		return &InvalidDurationError{
			BaseError: BaseError{
				Message: "Duration must be less than 1000",
			},
		}
	}

	msg := fmt.Sprintf("jd%03x%03x", mix, duration)

	return o.sendCommand(msg)
}

// AudioABFadeToPosition_AsString returns the command string used to fade to the specified position over the defined
// duration in fields (interlaces) or frames (progressive).
//
// For use with scheduled commands.
func AudioABFadeToPosition_AsString(mix uint16, duration uint16) string {
	return fmt.Sprintf("jd%03x%03x", mix, duration)
}

// SetAudioGain will control the audio gain for each source or the mixed output.
//
// channelMask is a 16 bit channel mask for audio channels [16:1]. 0X00FF = channels 1-8.
//
// gain is optional. Setting gain to -100 sets audio to mute.
//
// Note that while you can safely set the gain on any channel mask combination for the 16 audio channels, you can only
// get the gain for a channel mask for multiple channels if all channels are set the same gain.
//
// Getting the gain for a single channel should always return the correct setting.
func (o *Oxtel) SetAudioGain(output OxtelAudioOutput, channelMask ChannelMask, gain *int8) error {
	mask := buildChannelMask(channelMask)
	if *gain < -100 || *gain > 30 {
		return &InvalidGainError{
			BaseError: BaseError{
				Message: "Gain must be between -100 and +30",
			},
		}
	}

	msg := fmt.Sprintf("jAG%02x%s%d", output, mask, gain)

	return o.sendCommand(msg)
}

// SetAudioGain_AsString returns the command string used to control the audio gain for each source or the mixed output.
//
// For use with scheduled commands.
func SetAudioGain_AsString(output OxtelAudioOutput, channelMask ChannelMask, gain *int8) string {
	mask := buildChannelMask(channelMask)

	return fmt.Sprintf("jAG%02x%s%d", output, mask, gain)
}

// EnquireAudioGain queries the audio gain status for the specified channel mask.
func (o *Oxtel) EnquireAudioGain(output OxtelAudioOutput, channelMask ChannelMask) (AudioGainResponse, error) {
	mask := buildChannelMask(channelMask)
	val, err := o.sendCommandExpectResponse("jAG", fmt.Sprintf("%02x%s", output, mask))
	if err != nil {
		return AudioGainResponse{}, err
	}

	source, err := strconv.ParseUint(string(val[:2]), 16, 8)
	if err != nil {
		return AudioGainResponse{}, err
	}

	channelMaskRaw, err := strconv.ParseInt(string(val[2:6]), 16, 16)
	if err != nil {
		return AudioGainResponse{}, err
	}
	outChannelMask := fmt.Sprintf("%016b", channelMaskRaw)

	gain, err := strconv.Atoi(string(val[6:]))

	return AudioGainResponse{
		Source:      OxtelAudioSource(source),
		ChannelMask: outChannelMask,
		Gain:        int8(gain),
	}, err
}

// EnquireAudioGain_AsString returns the command string used to query the audio gain status for the specified channel mask.
//
// For use with scheduled commands.
func EnquireAudioGain_AsString(output OxtelAudioOutput, channelMask ChannelMask) string {
	mask := buildChannelMask(channelMask)

	return fmt.Sprintf("jAG%02x%s", output, mask)
}
