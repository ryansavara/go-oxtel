package oxtel

import (
	"fmt"
	"strconv"
)

// StartAnimation starts/resumes the animation of the specified layer.
func (o *Oxtel) StartAnimation(layer OxtelLayer) error {
	msg := fmt.Sprintf("S0%x", layer)

	return o.sendCommand(msg)
}

// StartAnimation_AsString returns the command string used to start/resume the animation of the specified layer.
func StartAnimation_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("S0%x", layer)
}

// StopAnimation stops the animation of the specified player.
//
// If immediate is true, the animation halts immediately at the current frame. If it is set to false, the animation
// completes before stopping.
func (o *Oxtel) StopAnimation(layer OxtelLayer, immediate bool) error {
	msg := fmt.Sprintf("S1%x%x", layer, boolToInt(immediate))

	return o.sendCommand(msg)
}

// StopAnimation_AsString returns the command string used to stop the animation of the specified player.
//
// For use with scheduled commands.
func StopAnimation_AsString(layer OxtelLayer, immediate bool) string {
	return fmt.Sprintf("S1%x%x", layer, boolToInt(immediate))
}

// SelectionAnimationFrame sets the template animation frame on the specified layer.
//
// This command is only applicable to HTML template. For HTML templates authored with the main timeline, this
// command will automatically work.
//
// For templates authored completely in JavaScript, it is the developer's responsibility to implement this functionality
// by overriding necessary JavaScript functions.
func (o *Oxtel) SelectionAnimationFrame(layer OxtelLayer, frame uint32) error {
	msg := fmt.Sprintf("S2%x%04x", layer, frame)

	return o.sendCommand(msg)
}

// SelectionAnimationFrame_AsString returns the command string used to set the template animation frame on the specified layer.
//
// For use with scheduled commands.
func SelectionAnimationFrame_AsString(layer OxtelLayer, frame uint32) string {
	return fmt.Sprintf("S2%x%04x", layer, frame)
}

// RestartAnimation restarts the template animation from the beginning on the specified layer.
//
// This command is only applicable to HTML template. For HTML templates authored with the main timeline, this command
// will automatically work.
//
// For templates authored completely in JavaScript, it is the developer's responsibility to implement this functionality
// by overriding necessary JavaScript functions.
func (o *Oxtel) RestartAnimation(layer OxtelLayer) error {
	msg := fmt.Sprintf("S4%x", layer)

	return o.sendCommand(msg)
}

// RestartAnimation_AsString returns the command string used to restart the template animation from the beginning on
// the specified layer.
//
// For use with scheduled commands.
func RestartAnimation_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("S4%x", layer)
}

// EnablePlayStateTally enables/disables the template play state tally for the connection on which the command was received.
//
// When first enabled, tallies for all layers are returned so that the client can record the initial state of each layer.
//
// An unsolicited response will be transmitted as a OxtelPlayStateTally.
func (o *Oxtel) EnablePlayStateTally(enable bool) error {
	msg := fmt.Sprintf("YS%x", boolToInt(enable))

	return o.sendCommand(msg)
}

// EnablePlayStateTally_AsString returns the command used to enable/disable the template play state tally for the
//
//	connection on which the command was received.
//
// For use with scheduled commands.
func EnablePlayStateTally_AsString(enable bool) string {
	return fmt.Sprintf("YS%x", boolToInt(enable))
}

// EnquirePlayStateTally queries the enable/disable state of the Play State tally.
//
// Response is a boolean representing if play state tally is enabled.
func (o *Oxtel) EnquirePlayStateTally() (bool, error) {
	val, err := o.sendCommandExpectResponse("YS", "")
	if err != nil {
		return false, err
	}
	retval, err := strconv.ParseBool(val)
	return retval, err
}

// EnquirePlayStateTally_AsString returns the command string used to query the enable/disable state of the Play State tally.
//
// For use with scheduled commands.
func EnquirePlayStateTally_AsString() string {
	return "YS"
}
