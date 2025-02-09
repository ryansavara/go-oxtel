package oxtel

import (
	"fmt"
	"strconv"
	"strings"
)

// LoadImage loads a template onto the specific layer. If another template is loaded on the specified layer, then the
// current template will be unloaded before the new template is loaded.
// If a 'non-existent' filename is specified in the templateName parameter, the layer is unloaded, however the preloaded
// template is not unloaded.
//
// templateName does not need the file system path, as the file system path is loaded from the Effects configuration
// section in SystemManager. Example: a.png
//
// HTML templates must contain the directory name along with the template filename separated by a forward slash. Example: a/a.html
//
// An unsolicited response will be sent if tallies are enabled with the format "Y9%l %s" where %l is the layer and %s is the filename.
func (o *Oxtel) LoadImage(layer OxtelLayer, templateName string) error {
	msg := fmt.Sprintf("R0%x%s", layer, templateName)

	return o.sendCommand(msg)
}

// LoadImage_AsString returns the command string used to load a template onto the specific layer. If another template is
// loaded on the specified layer, then the current template will be unloaded before the new template is loaded.
//
// For use with scheduled commands.
func LoadImage_AsString(layer OxtelLayer, templateName string) string {
	return fmt.Sprintf("R0%x%s", layer, templateName)
}

// EnquireLoadImage queries the template file that is currently loaded into the specified layer.
// When no template is loaded on the specified layer, the filename will be ">Empty<".
//
// For HTML templates, the filename will contain the directory name, slash, and template name. Example: a/a.html
//
// Response is a LayerTemplateResponse.
func (o *Oxtel) EnquireLoadImage(layer OxtelLayer) (LayerTemplateResponse, error) {
	val, err := o.sendCommandExpectResponse("R0", fmt.Sprintf("%x", layer))
	if err != nil {
		return LayerTemplateResponse{}, err
	}

	returnLayer, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return LayerTemplateResponse{}, err
	}

	return LayerTemplateResponse{
		Layer:    OxtelLayer(returnLayer),
		Filename: val[1:],
	}, nil
}

// EnquireLoadImage_AsString returns the command string used to query the template file that is currently loaded into
// the specified layer.
//
// For use with scheduled commands.
func EnquireLoadImage_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("R0%x", layer)
}

// PreloadImage preloads a template on the specified layer. Once the preload is complete, the template can be swapped with
// the on-air template in a frame-accurate manner.
//
// A preloaded template is loaded in an off-air location. When the preload is complete and a subsequent LoadImage command is
// issued on the same layer and filename, the preloaded template is immediately swapped onto the on-air location.
// If a different valid filename is requested in the LoadImage command, then the preloaded template is cleared.
//
// Load times are dependent on the template size and load on the system. Automation should take care to make sure the adequate
// time is allowed for the template to be completely preloaded before sending the LoadImage command.
// It is recommended that automation monitor the ImagePreloadTally (YA) to know when the preload completes.
//
// An unsolicited response will be sent if tallies are enabled with the format "YA%l %s" where %l is the layer and %s is the filename.
func (o *Oxtel) PreloadImage(layer OxtelLayer, templateName string) error {
	msg := fmt.Sprintf("R7%x%s", layer, templateName)

	return o.sendCommand(msg)
}

// PreloadImage_AsString returns the command string used to preload a template on the specified layer. Once the preload is
// complete, the template can be swapped with the on-air template in a frame-accurate manner.
//
// For use with scheduled commands.
func PreloadImage_AsString(layer OxtelLayer, templateName string) string {
	return fmt.Sprintf("R7%x%s", layer, templateName)
}

// EnquirePreloadImage queries the template file that is currently preloaded into the specified layer.
//
// Response is a LayerTemplateResponse.
func (o *Oxtel) EnquirePreloadImage(layer OxtelLayer) (LayerTemplateResponse, error) {
	val, err := o.sendCommandExpectResponse("R7", fmt.Sprintf("%x", layer))
	if err != nil {
		return LayerTemplateResponse{}, err
	}

	returnLayer, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return LayerTemplateResponse{}, err
	}

	return LayerTemplateResponse{
		Layer:    OxtelLayer(returnLayer),
		Filename: val[1:],
	}, nil
}

// EnquirePreloadImage_AsString returns the command string used to query the template file that is currently preloaded
// into the specified layer.
//
// For use with scheduled commands.
func EnquirePreloadImage_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("R7%x", layer)
}

// EraseStore unloads the template from the specified layer.
func (o *Oxtel) EraseStore(layer OxtelLayer) error {
	msg := fmt.Sprintf("A%x", layer)

	return o.sendCommand(msg)
}

// EraseStore_AsString returns the command string used to unload the template from the specified layer.
//
// For use with scheduled commands.
func EraseStore_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("A%x", layer)
}

// SetImagePosition sets the position of the loaded template relative to the origin. The origin (x=0, y=0) is defined in
// the upper left-hand corner of the screen.
//
// Positive values move the template right and down. Negative values move the template left and up.
func (o *Oxtel) SetImagePosition(layer OxtelLayer, xOffset uint16, yOffset uint16) error {
	msg := fmt.Sprintf("G%x %x %x", layer, xOffset, yOffset)

	return o.sendCommand(msg)
}

// SetImagePosition_AsString returns the command string used to set the position of the loaded template relative to the origin.
// The origin (x=0, y=0) is defined in the upper left-hand corner of the screen.
//
// For use with scheduled commands.
func SetImagePosition_AsString(layer OxtelLayer, xOffset uint16, yOffset uint16) string {
	return fmt.Sprintf("G%x %x %x", layer, xOffset, yOffset)
}

// EnquireImagePosition submits an inquiry on the position of the loaded template relative to the origin.
//
// Response is a ImagePositionResponse with units of pixels.
func (o *Oxtel) EnquireImagePosition(layer OxtelLayer) (ImagePositionResponse, error) {
	val, err := o.sendCommandExpectResponse("G", fmt.Sprintf("%x", layer))
	if err != nil {
		return ImagePositionResponse{}, err
	}

	data := strings.Split(val, " ")

	returnLayer, err := strconv.ParseUint(data[0], 16, 8)
	if err != nil {
		return ImagePositionResponse{}, err
	}

	xoffset, err := strconv.ParseInt(data[1], 16, 0)
	if err != nil {
		return ImagePositionResponse{}, err
	}

	yoffset, err := strconv.ParseInt(data[2], 16, 0)
	if err != nil {
		return ImagePositionResponse{}, err
	}

	return ImagePositionResponse{
		Layer:   OxtelLayer(returnLayer),
		XOffset: int(xoffset),
		YOffset: int(yoffset),
	}, nil

}

// EnquireImagePosition_AsString returns the command string used to submit an inquiry on the position of the loaded template
// relative to the origin.
//
// For use with scheduled commands.
func EnquireImagePosition_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("G%x", layer)
}
