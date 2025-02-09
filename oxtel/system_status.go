package oxtel

import (
	"fmt"
	"strconv"
	"strings"
)

// EnquireLatency returns the latency, in reference frames, of the specified source to the SDI output.
//
// Refer to the Timing Model section for more details.
//
// Response is a LatencyResponse.
func (o *Oxtel) EnquireLatency(source OxtelLatencySource) (LatencyResponse, error) {
	val, err := o.sendCommandExpectResponse("hLAT", fmt.Sprintf("%x", source))
	if err != nil {
		return LatencyResponse{}, err
	}

	latencySource, err := strconv.ParseUint(string(val[0]), 16, 8)
	if err != nil {
		return LatencyResponse{}, err
	}
	latency, err := strconv.ParseInt(string(val[1]), 16, 8)
	if err != nil {
		return LatencyResponse{}, err
	}

	return LatencyResponse{
		Source:  OxtelLatencySource(latencySource),
		Latency: int8(latency),
	}, nil
}

// EnquireLatency_AsString returns the command string used to return the latency, in reference frames, of the specified source
// to the SDI output.
//
// For use with scheduled commands.
func EnquireLatency_AsString(source OxtelLatencySource) string {
	return fmt.Sprintf("hLAT%x", source)
}

// EnquireNumberOfGraphicLayers returns the number of graphic layers license for this channel.
//
// Response is an integer representing the number of licensed graphic layers.
func (o *Oxtel) EnquireNumberOfGraphicLayers() (int, error) {
	val, err := o.sendCommandExpectResponse("hNGL", "")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

// EnquireNumberOfGraphicLayers_AsString returns the command string used to return the number of graphic layers license
// for this channel.
//
// For use with scheduled commands.
func EnquireNumberOfGraphicLayers_AsString() string {
	return "hNGL"
}

// EnquireSystemStatus queries information about the system status.
//
// It is of limited use, but can be used by automation to poll the system periodically to make sure that the Harmonic product
// is still alive.
//
// Response is a SystemStatusResponse.
func (o *Oxtel) EnquireSystemStatus() (SystemStatusResponse, error) {
	val, err := o.sendCommandExpectResponse("M", "")
	if err != nil {
		return SystemStatusResponse{}, err
	}

	mode, err := strconv.ParseUint(string(val[0]), 10, 8)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	vhigh, err := strconv.ParseUint(string(val[1:4]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	vlow, err := strconv.ParseUint(string(val[4:7]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	standard, err := strconv.ParseUint(string(val[7]), 10, 8)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	preview, err := strconv.ParseUint(string(val[8:11]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	fade1, err := strconv.ParseUint(string(val[11:14]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	fade2, err := strconv.ParseUint(string(val[14:17]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	ftb1, err := strconv.ParseUint(string(val[17:20]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	ftb2, err := strconv.ParseUint(string(val[20:23]), 16, 16)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	access, err := strconv.ParseUint(string(val[23]), 16, 8)
	if err != nil {
		return SystemStatusResponse{}, err
	}

	return SystemStatusResponse{
		SystemMode:        uint8(mode),
		VersionHigh:       uint16(vhigh),
		VersionLow:        uint16(vlow),
		VideoStandard:     OxtelVideoStandard(standard),
		PreviewSource:     uint16(preview),
		FadeRateDSK1:      uint16(fade1),
		FadeRateDSK2:      uint16(fade2),
		FTBRateDSK1:       uint16(ftb1),
		FTBRateDSK2:       uint16(ftb2),
		SystemNotAccessed: uint8(access),
	}, nil

}

// EnquireSystemStatus_AsString returns the command string used to query information about the system status.
//
// For use with scheduled commands.
func EnquireSystemStatus_AsString() string {
	return "M"
}

// EnquireVideoLayerStatus returns status information about the specified layer.
//
// Only the Intuition SD/HD[+] format is supported which requires an additional parameter to identify the layer status to query.
//
// Response is a VideoLayerStatusResponse.
func (o *Oxtel) EnquireVideoLayerStatus(layer OxtelLayer) (VideoLayerStatusResponse, error) {
	val, err := o.sendCommandExpectResponse("N", fmt.Sprintf("%x", layer))
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	fadeAngle, err := strconv.ParseUint(string(val[:3]), 16, 16)
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	ftbAngle, err := strconv.ParseUint(string(val[3:6]), 16, 16)
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	u1, err := strconv.ParseUint(string(val[6:9]), 16, 16)
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	u2, err := strconv.ParseUint(string(val[9:12]), 16, 16)
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	u3, err := strconv.ParseUint(string(val[12:14]), 16, 16)
	if err != nil {
		return VideoLayerStatusResponse{}, err
	}

	return VideoLayerStatusResponse{
		LayerFaderAngle: uint16(fadeAngle),
		LayerFTBAngle:   uint16(ftbAngle),
		Unused1:         uint16(u1),
		Unused2:         uint16(u2),
		Unused3:         uint16(u3),
	}, nil
}

// EnquireVideoLayerStatus_AsString returns the command string used to return status information about the specified layer.
//
// For use with scheduled commands.
func EnquireVideoLayerStatus_AsString(layer OxtelLayer) string {
	return fmt.Sprintf("N%x", layer)
}

// EnquireCommandAvailability is used to determine if a particular Oxtel command is currently available on the Harmonic
// product for automation to use.
//
// Note: This is not completely implemented. Some commands allow for 1-4 bytes in the Command prefix.
//
// Response is a CommandAvailabilityResponse.
func (o *Oxtel) EnquireCommandAvailability(byte1 byte, byte2 byte) (CommandAvailabilityResponse, error) {
	val, err := o.sendCommandExpectResponse("X3", fmt.Sprintf("%c%c", byte1, byte2))
	if err != nil {
		return CommandAvailabilityResponse{}, err
	}
	cmdByte1 := val[0]
	cmdByte2 := val[1]
	supported, err := strconv.ParseBool(string(val[2]))
	if err != nil {
		return CommandAvailabilityResponse{}, err
	}

	return CommandAvailabilityResponse{
		CommandByte1: cmdByte1,
		CommandByte2: cmdByte2,
		Supported:    supported,
	}, nil

}

// EnquireCommandAvailability_AsString returns the command string used to determine if a particular Oxtel command is
// currently available on the Harmonic product for automation to use.
//
// For use with scheduled commands.
func EnquireCommandAvailability_AsString(byte1 byte, byte2 byte) string {
	return fmt.Sprintf("X3%c%c", byte1, byte2)
}

// EnquireSlaveLayerStatus is different than the Miranda implementation. It queries the current stat of the keyer for
// each layer (0x1 = fader angle is 0x200, 0x0 = fader angle is not 0x200).
//
// Response is a SlaveLayerStatusResponse
func (o *Oxtel) EnquireSlaveLayerStatus() (SlaveLayerStatusResponse, error) {
	val, err := o.sendCommandExpectResponse("XA", "")
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer0, err := strconv.ParseBool(string(val[0]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer1, err := strconv.ParseBool(string(val[1]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer2, err := strconv.ParseBool(string(val[2]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer3, err := strconv.ParseBool(string(val[3]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer4, err := strconv.ParseBool(string(val[4]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer5, err := strconv.ParseBool(string(val[5]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer6, err := strconv.ParseBool(string(val[6]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	layer7, err := strconv.ParseBool(string(val[7]))
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	unused, err := strconv.ParseInt(string(val[8:16]), 16, 32)
	if err != nil {
		return SlaveLayerStatusResponse{}, err
	}

	return SlaveLayerStatusResponse{
		Layer0State: layer0,
		Layer1State: layer1,
		Layer2State: layer2,
		Layer3State: layer3,
		Layer4State: layer4,
		Layer5State: layer5,
		Layer6State: layer6,
		Layer7State: layer7,
		Unused:      int32(unused),
	}, nil
}

// EnquireSlaveLayerStatus_AsString returns the command string used to query the current state of the keyer for each
// layer (0x1 = fader angle is 0x200, 0x0 = fader angle is not 0x200).
//
// For use with scheduled commands.
func EnquireSlaveLayerStatus_AsString() string {
	return "XA"
}

// EnquireFullVersionNumber returns the full software version number.
//
// Response is a FullVersionNumberResponse.
// In order, the integers represent Major, Minor, Patch, Branch, and Build Number.
func (o *Oxtel) EnquireFullVersionNumber() (FullVersionNumberResponse, error) {
	val, err := o.sendCommandExpectResponse("Xb", "")
	if err != nil {
		return FullVersionNumberResponse{}, err
	}
	parts := strings.Split(val, ".")
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return FullVersionNumberResponse{}, err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return FullVersionNumberResponse{}, err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return FullVersionNumberResponse{}, err
	}

	branch, err := strconv.Atoi(parts[3])
	if err != nil {
		return FullVersionNumberResponse{}, err
	}

	build := parts[4]

	return FullVersionNumberResponse{
		Major:       major,
		Minor:       minor,
		Patch:       patch,
		Branch:      branch,
		BuildNumber: build,
		AsString:    val,
	}, nil
}

// EnquireFullVersionNumber_AsString returns the command string used to return the full software version number.
//
// For use with scheduled commands.
func EnquireFullVersionNumber_AsString() string {
	return "Xb"
}

// EnquireProductName returns the product name.
//
// Response is a string representing the product name. Valid responses will be "ChannelPort", "Spectrum X", or "Electra X".
func (o *Oxtel) EnquireProductName() (string, error) {
	return o.sendCommandExpectResponse("Xn", "")
}

// EnquireProductName_AsString returns the command string used to return the product name.
//
// For use with scheduled commands.
func EnquireProductName_AsString() string {
	return "Xn"
}

// GetMediaPortName returns the globally unique name of the media port or channel of a Spectrum-X media deck.
//
// Response is a string of the globally unique name of the media port.
func (o *Oxtel) GetMediaPortName() (string, error) {
	return o.EnquireMediaPortName()
}

// EnquireMediaPortName returns the globally unique name of the media port or channel of a Spectrum-X media deck.
//
// Response is a string of the globally unique name of the media port.
func (o *Oxtel) EnquireMediaPortName() (string, error) {
	return o.sendCommandExpectResponse("hTN", "")
}

// EnquireMediaPortName_AsString returns the command string used to return the globally unique name of the media port
// or channel of a Spectrum-X media deck.
//
// For use with scheduled commands.
func EnquireMediaPortName_AsString() string {
	return "hTN"
}
