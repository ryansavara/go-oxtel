package oxtel

import "fmt"

// OverrideSDIInputColorSpace overrides the color space for a given SDI input. Once the command is issued,
// it will stick until another OverrideSDIInputColorSpace command or the MCS is reconfigured.
//
// If the MCS is reconfigured, it will revert back to the SystemManager settings.
func (o *Oxtel) OverrideSDIInputColorSpace(input OxtelMixerInput, colorSpace OxtelColorSpace) error {
	return o.sendCommand(fmt.Sprintf("hCSI%01x%01x", input, colorSpace))
}

// OverrideSDIInputColorSpace_AsString returns the command string used to override the color space for a given SDI input. Once the command is issued,
//
// For use with scheduled commands.
func OverrideSDIInputColorSpace_AsString(input OxtelMixerInput, colorSpace OxtelColorSpace) string {
	return fmt.Sprintf("hCSI%01x%01x", input, colorSpace)
}

// ChangeKantarWatermarkingChannelName changes the Kantar watermarking channel name.
// The name must be a valid name contained in the system's Kantar audience license.
//
// Note: If Kantar is set to Mirror, then any OxtelKantarOutput will select both.
func (o *Oxtel) ChangeKantarWatermarkingChannelName(output OxtelKantarOutput, startingChannel uint8, audienceName string) error {
	if startingChannel > 15 {
		return &InvalidAudioChannelError{
			BaseError: BaseError{
				Message: "Starting Channel must be less than 16",
			},
		}
	}

	return o.sendCommand(fmt.Sprintf("hKWM%01x%02x%s", output, startingChannel, audienceName))
}

// ChangeKantarWatermarkingChannelName_AsString returns the command string used to change the Kantar watermarking channel name.
//
// For use with scheduled commands.
func ChangeKantarWatermarkingChannelName_AsString(output OxtelKantarOutput, startingChannel uint8, audienceName string) string {
	return fmt.Sprintf("hKWM%01x%02x%s", output, startingChannel, audienceName)
}
