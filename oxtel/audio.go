package oxtel

import (
	"fmt"
	"strconv"
)

// AudioPopSuppression disables pop suppression for two frames. This command is meant to be used in conjunction with an
// audio cut command to prevent audio dropouts when the clips are contiguous and played back-to-back. I.E. loop record or
// situations like a fade to A then Cut to A where the cut is to the already selected source where pop suppression
// is not necessary and will cause audio drops.
//
// This is a one-shot command that will disable the pop suppression for the two fields then re-enable pop suppression going
// forward. This command can select any source, but is probably only useful for the Player.
//
// Note that because this is a one-shot that resets to re-enable pop suppression immediately after two fields,
// sending a second AudioPopSuppression with enable = false is not necessary.
func (o *Oxtel) AudioPopSuppression(source OxtelAudioSource, enable bool) error {
	return o.sendCommand(fmt.Sprintf("hDP%01x%01x", source, boolToInt(enable)))
}

// AudioPopSuppression_AsString returns the command string used to enable/disable pop suppression for two frames.
//
// For use with scheduled commands.
func AudioPopSuppression_AsString(source OxtelAudioSource, enable bool) string {
	return fmt.Sprintf("hDP%01x%01x", source, boolToInt(enable))
}

// PauseResumeDolbyEncoder pauses/resumes an active Dolby Encoder. Pausing the Dolby Encoder will pass through the PCM audio.
// The Dolby Encoder continues to be active, but it's output is not used when in a paused state.
//
// Encoding will resume when it receives the resume command with no latencies.
func (o *Oxtel) PauseResumeDolbyEncoder(input OxtelMixerInput, pause bool) error {
	return o.sendCommand(fmt.Sprintf("hDE%01x%01x", input, boolToInt(pause)))
}

// PauseResumeDolbyEncoder_AsString returns the command string that is used to pause/resume an active Dolby Encoder.
//
// For use with scheduled commands.
func PauseResumeDolbyEncoder_AsString(input OxtelMixerInput, pause bool) string {
	return fmt.Sprintf("hDE%01x%01x", input, boolToInt(pause))
}

// SetDolbyEncoderProfile selects one of the predefined Dolby Encoder Profiles.
// There can be up to four Profiles defined in the SystemManager application.
func (o *Oxtel) SetDolbyEncoderProfile(input OxtelMixerInput, profile uint8) error {
	if profile > 4 {
		return &InvalidDolbyProfileError{
			BaseError: BaseError{
				Message: "Profile must be less than 5",
			},
		}
	}

	return o.sendCommand(fmt.Sprintf("hDA%01x%01x", input, profile))
}

// SetDolbyEncoderProfile_AsString returns the command string used to select one of the predefined Dolby Encoder Profiles.
//
// For use with scheduled commands.
func SetDolbyEncoderProfile_AsString(input OxtelMixerInput, profile uint8) string {
	return fmt.Sprintf("hDA%01x%01x", input, profile)
}

// EnquireDolbyEncoderProfile queries Dolby Encoder Profile used on the input.
//
// Response is an int.
func (o *Oxtel) EnquireDolbyEncoderProfile(input OxtelMixerInput) (uint64, error) {
	val, err := o.sendCommandExpectResponse("hDA", fmt.Sprintf("%01x", input))
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(val, 16, 0)
}

// EnquireDolbyEncoderProfile_AsString returns the command string used to query Dolby Encoder Profile used on the input.
//
// For use with scheduled commands.
func EnquireDolbyEncoderProfile_AsString(input OxtelMixerInput) string {
	return fmt.Sprintf("hDA%01x", input)
}
