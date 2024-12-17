package oxtel

import "fmt"

// UpdateTextField updates the text in the specified text field.
//
// flags is bitwise of type OxtelUpdateTextFieldFlag.
//
// If the flag is set to OXTEL_UPDATE_TEXT_FIELD_RENDER, then the specified text string will update on screen.
// If the OXTEL_UPDATE_TEXT_FIELD_RENDER flag is not set, then changes do not appear on screen until the next call to
// RenderBox or UpdateTextField with the render flag set.
//
// If the OXTEL_UPDATE_TEXT_FIELD_APPEND flag is set, then the specified text string is appended to the existing text.
// This allows for long text strings to be defined over several command packets.
func (o *Oxtel) UpdateTextField(layer OxtelLayer, field uint8, flags OxtelUpdateTextFieldFlag, text string) error {
	if field > 254 {
		return &InvalidFieldError{
			BaseError: BaseError{
				Message: "Field must be less than 255",
			},
		}
	}

	msg := fmt.Sprintf("Z0%x%02x%x%s", layer, field, flags, text)

	return o.sendCommand(msg)
}

// UpdateTextField_AsString returns the command string used to update the text in the specified text field.
//
// For use with scheduled commands.
func UpdateTextField_AsString(layer OxtelLayer, field uint8, flags OxtelUpdateTextFieldFlag, text string) string {
	return fmt.Sprintf("Z0%x%02x%x%s", layer, field, flags, text)
}

// UpdatePreloadedTextField updates the text in the specified text field for a preloaded template.
//
// flags is bitwise of type OxtelUpdateTextFieldFlag.
//
// If the flag is set to OXTEL_UPDATE_TEXT_FIELD_RENDER, then the specified text string will update on screen.
// If the OXTEL_UPDATE_TEXT_FIELD_RENDER flag is not set, then changes do not appear on screen until the next call to
// RenderBox or UpdateTextField with the render flag set.
//
// If the OXTEL_UPDATE_TEXT_FIELD_APPEND flag is set, then the specified text string is appended to the existing text.
// This allows for long text strings to be defined over several command packets.
func (o *Oxtel) UpdatePreloadedTextField(layer OxtelLayer, field uint8, flags OxtelUpdateTextFieldFlag, text string) error {
	if field > 254 {
		return &InvalidFieldError{
			BaseError: BaseError{
				Message: "Field must be less than 255",
			},
		}
	}

	msg := fmt.Sprintf("hZ0%x%02x%x%s", layer, field, flags, text)

	return o.sendCommand(msg)
}

// UpdatePreloadedTextField_AsString returns the command string used to update the text in the specified text field for
// a preloaded template.
//
// For use with scheduled commands.
func UpdatePreloadedTextField_AsString(layer OxtelLayer, field uint8, flags OxtelUpdateTextFieldFlag, text string) string {
	return fmt.Sprintf("hZ0%x%02x%x%s", layer, field, flags, text)
}

// RenderBox updates the specified text field
//
// If field is 0xFF (255), all fields are updated.
func (o *Oxtel) RenderBox(layer OxtelLayer, field uint8) error {
	msg := fmt.Sprintf("Z0%x%02x", layer, field)

	return o.sendCommand(msg)
}

// RenderBox_AsString returns the command string used to update the specified text field
//
// For use with scheduled commands.
func RenderBox_AsString(layer OxtelLayer, field uint8) string {
	return fmt.Sprintf("Z0%x%02x", layer, field)
}

// ChangeImage changes the image associated with a text field to be replaced with another image on disk.
//
// The new settings take effect when the RenderBox command is issued.
func (o *Oxtel) ChangeImage(layer OxtelLayer, field uint8, fileName string) error {
	if field > 254 {
		return &InvalidFieldError{
			BaseError: BaseError{
				Message: "Field must be less than 255",
			},
		}
	}

	msg := fmt.Sprintf("Z4%x%02x%s", layer, field, fileName)

	return o.sendCommand(msg)
}

// ChangeImage_AsString returns the command string used to change the image associated with a text field to be
// replaced with another image on disk.
//
// For use with scheduled commands.
func ChangeImage_AsString(layer OxtelLayer, field uint8, fileName string) string {
	return fmt.Sprintf("Z4%x%02x%s", layer, field, fileName)
}

// StopTextFieldAnimation stops an animation from playing in the specified text field on the specified layer.
//
// This command behaves similarly to the StopAnimation command; however, it includes a text field number parameter
//  to allow finer control of the text field in the template.
// If immediate is true, the animation halts immediately at the current frame. If it is set to false, the animation
// completes before stopping.
func (o *Oxtel) StopTextFieldAnimation(layer OxtelLayer, field uint8, immediate bool) error {
	if field > 254 {
		return &InvalidFieldError{
			BaseError: BaseError{
				Message: "Field must be less than 255",
			},
		}
	}

	msg := fmt.Sprintf("Zf%x%02x%x", layer, field, boolToInt(immediate))

	return o.sendCommand(msg)
}

// StopTextFieldAnimation_AsString returns the command string used to stop an animation from playing in the specified
// text field on the specified layer.
//
// For use with scheduled commands.
func StopTextFieldAnimation_AsString(layer OxtelLayer, field uint8, immediate bool) string {
	return fmt.Sprintf("Zf%x%02x%x", layer, field, boolToInt(immediate))
}

// PauseRestartStrap pauses/restarts the specified text field on the specified layer.

// If the restart is false, then the pause command is sent instead.
func (o *Oxtel) PauseRestartStrap(layer OxtelLayer, field uint8, restart bool) error {
	if field > 254 {
		return &InvalidFieldError{
			BaseError: BaseError{
				Message: "Field must be less than 255",
			},
		}
	}

	msg := fmt.Sprintf("Zg%x%02x%x", layer, field, boolToInt(restart))

	return o.sendCommand(msg)
}

// PauseRestartStrap_AsString returns the command string used to pause/restart the specified text field on the specified layer.
//
// For use with scheduled commands.
func PauseRestartStrap_AsString(layer OxtelLayer, field uint8, restart bool) string {
	return fmt.Sprintf("Zg%x%02x%x", layer, field, boolToInt(restart))
}
