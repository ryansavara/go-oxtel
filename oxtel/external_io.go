package oxtel

import (
	"fmt"
	"strconv"
	"strings"
)

// EnquireNumberOfExternalIOConfigurations returns the number of external IO configurations for the specified type and direction.
// Use this command to determine the number of configurations. To gain access to the complete configuration, use this
// command and the EnquireExternalIOConfiguration command.
//
// Note: Since 2022-6 and 2022-6/2022-7 share the same entries in the configuration file, this command will return the same value
// for both Type 1 (2022-6) and Type 4 (2022-6/2022-7). If it is important to know if 2022-7 is configured, then the response
// returned by the EnquireExternalIOConfiguration command can be used to determine if the second leg is defined.
//
// Note: A separate Type is not specified for 2110/2022-7 as those settings are configured in the 2110 Session Description
// Protocol (SDP) files.
func (o *Oxtel) EnquireNumberOfExternalIOConfigurations(ioType OxtelExternalIOType,
	direction OxtelExternalIODirection) (NumberOfExternalIOConfigurationsResponse, error) {
	val, err := o.sendCommandExpectResponse("hXNC", fmt.Sprintf("%02x%02x", ioType, direction))
	if err != nil {
		return NumberOfExternalIOConfigurationsResponse{}, err
	}

	outType, err := strconv.ParseUint(string(val[:2]), 16, 8)
	if err != nil {
		return NumberOfExternalIOConfigurationsResponse{}, err
	}

	dir, err := strconv.ParseUint(string(val[2:4]), 16, 8)
	if err != nil {
		return NumberOfExternalIOConfigurationsResponse{}, err
	}

	num, err := strconv.ParseUint(string(val[4:]), 16, 8)
	if err != nil {
		return NumberOfExternalIOConfigurationsResponse{}, err
	}

	return NumberOfExternalIOConfigurationsResponse{
		IOType:            OxtelExternalIOType(outType),
		IODirection:       OxtelExternalIODirection(dir),
		NumConfigurations: uint8(num),
	}, nil
}

// EnquireNumberOfExternalIOConfigurations_AsString returns the command string used to return the number of external IO configurations for the specified type and direction.
//
// For use with scheduled commands.
func EnquireNumberOfExternalIOConfigurations_AsString(ioType OxtelExternalIOType,
	direction OxtelExternalIODirection) string {
	return "hXNC" + fmt.Sprintf("%02x%02x", ioType, direction)
}

// EnquireExternalIOConfiguration returns the external IO configuration for the specified type, direction, and index.
// The index is a value less than the number returned from the EnquireNumberOfExternalIOConfigurations command.
//
// For SDI, the configuration information is internal and not user configurable.
//
// For 2022-6 and 2110, the configuration information is the same information configured via the SystemManager.
//
// Note: The configuration information for Type 1 (2022-6) and Type 4 (2022-6/2022-7) share the same configuration
// file entries. If Type 4 is requested, but the second leg for 2022-7 is not configured, IPAddress2 and Port2 will nil.
//
// Response is an ExternalIOConfigurationResponse. Every response will have a IOType, IODirection, Index, ConfigurationID,
// and Name. Depending on what IOType is queried, the other fields may be nil.
func (o *Oxtel) EnquireExternalIOConfiguration(ioType OxtelExternalIOType, direction OxtelExternalIODirection, index uint8) (ExternalIOConfigurationResponse, error) {
	val, err := o.sendCommandExpectResponse("hXNC", fmt.Sprintf("%02x%02x%02x", ioType, direction, index))
	if err != nil {
		return ExternalIOConfigurationResponse{}, err
	}

	outType, err := strconv.ParseUint(string(val[:2]), 16, 8)
	if err != nil {
		return ExternalIOConfigurationResponse{}, err
	}

	dir, err := strconv.ParseUint(string(val[2:4]), 16, 8)
	if err != nil {
		return ExternalIOConfigurationResponse{}, err
	}

	outIndex, err := strconv.ParseUint(string(val[4:6]), 16, 8)
	if err != nil {
		return ExternalIOConfigurationResponse{}, err
	}

	configId, err := strconv.ParseUint(string(val[6:8]), 16, 8)
	if err != nil {
		return ExternalIOConfigurationResponse{}, err
	}

	parts := strings.Split(val, ",")
	name := parts[0][8:]

	var localInterface *string
	var ipAddress *string
	var port *uint32
	var ipAddress2 *string
	var port2 *uint32
	var sdfFileName *string

	if outType != OXTEL_EXT_IO_TYPE_SDI {
		localInterface = &parts[1]
	}
	if outType == OXTEL_EXT_IO_TYPE_2022_6 || outType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
		ipAddress = &parts[2]
		tmpPort, err := strconv.ParseUint(parts[3], 10, 32)
		if err != nil {
			return ExternalIOConfigurationResponse{}, err
		}
		w := uint32(tmpPort)
		port = &w
	}
	if outType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
		ipAddress2 = &parts[4]
		tmpPort, err := strconv.ParseUint(parts[5], 10, 32)
		if err != nil {
			return ExternalIOConfigurationResponse{}, err
		}
		w := uint32(tmpPort)
		port2 = &w
	}

	if outType == OXTEL_EXT_IO_TYPE_2110 {
		sdfFileName = &parts[2]
	}

	return ExternalIOConfigurationResponse{
		IOType:          OxtelExternalIOType(outType),
		IODirection:     OxtelExternalIODirection(dir),
		Index:           uint8(outIndex),
		ConfigurationId: uint8(configId),
		Name:            name,
		LocalInterface:  localInterface,
		IPAddress:       ipAddress,
		Port:            port,
		IPAddress2:      ipAddress2,
		Port2:           port2,
		SDPFileName:     sdfFileName,
	}, nil
}

// EnquireExternalIOConfiguration returns the command string used to return the external IO configuration for the specified type, direction, and index.
//
// For use with scheduled commands.
func EnquireExternalIOConfiguration_AsString(ioType OxtelExternalIOType, direction OxtelExternalIODirection, index uint8) string {
	return "hXNC" + fmt.Sprintf("%02x%02x%02x", ioType, direction, index)
}

// SetExternalIOSource sets the source for the specified External IO. The External IO is defined by the direction of the IO id.
//
// The _configId_ parameter is the static configuration ID specified in the SystemManager. "0" is always the Dynamic Config,
// which is not specified in the SystemManager.
func (o *Oxtel) SetExternalIOSource(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType, configId uint8, forceRestart bool) error {
	return o.sendCommand(fmt.Sprintf("hXS%02x%02x%02x%02x%02x", direction, ioId, ioType, configId, boolToInt(forceRestart)))
}

// SetExternalIOSource_AsString returns the command string used to set the source for the specified External IO. The External IO is defined by the direction of the IO id.
//
// For use with scheduled commands.
func SetExternalIOSource_AsString(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType, configId uint8, forceRestart bool) string {
	return fmt.Sprintf("hXS%02x%02x%02x%02x%02x", direction, ioId, ioType, configId, boolToInt(forceRestart))
}

// EnquireExternalIOSource gets the source for the specified External IO. the External IO is defined by the direction of the IO id.
//
// Response is a ExternalIOSourceResponse. State is always 0.
func (o *Oxtel) EnquireExternalIOSource(direction OxtelExternalIODirection, ioId OxtelExternalIOId) (ExternalIOSourceResponse, error) {
	val, err := o.sendCommandExpectResponse("hXS", fmt.Sprintf("%02x%02x", direction, ioId))
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	outDir, err := strconv.ParseUint(string(val[:2]), 16, 8)
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	outIOId, err := strconv.ParseUint(string(val[2:4]), 16, 8)
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	outType, err := strconv.ParseUint(string(val[4:6]), 16, 8)
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	configId, err := strconv.ParseUint(string(val[6:8]), 16, 8)
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	state, err := strconv.ParseUint(string(val[8:10]), 16, 8)
	if err != nil {
		return ExternalIOSourceResponse{}, err
	}

	return ExternalIOSourceResponse{
		IODirection:     OxtelExternalIODirection(outDir),
		IOId:            OxtelExternalIOId(outIOId),
		IOType:          OxtelExternalIOType(outType),
		ConfigurationId: uint8(configId),
		State:           uint8(state),
	}, nil
}

// EnquireExternalIOSource_AsString returns the command string used to get the source for the specified External IO. the External IO is defined by the direction of the IO id.
//
// For use with scheduled commands.
func EnquireExternalIOSource_AsString(direction OxtelExternalIODirection, ioId OxtelExternalIOId) string {
	return "hXS" + fmt.Sprintf("%02x%02x", direction, ioId)
}

// SetExternalIODynamicConfiguration sets the dynamic configuration for the specified External IO. The External IO is
// defined by the direction and IO Id.
//
// For SDI, this command is not applicable.
//
// 2022-6 does not use ipAddress2, port2, or sdpFileName.
// 2022-6/2022-7 does not use sdpFileName.
// 2110 does not use ipAddress, port, ipAddress2, or port2.
func (o *Oxtel) SetExternalIODynamicConfiguration(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType, flags uint16, localInterface string, ipAddress *string, port *string, ipAddress2 *string, port2 *string, sdpFileName *string) error {

	if ioType == OXTEL_EXT_IO_TYPE_2022_6 {
		if ipAddress == nil || port == nil {
			return &InvalidParametersError{
				BaseError: BaseError{
					Message: "Parameters ipAddress and port are required",
				},
			}
		}

		return o.sendCommand(fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s,%s", direction, ioId, ioType, flags, localInterface, *ipAddress, *port))
	} else if ioType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
		if ipAddress == nil || port == nil || ipAddress2 == nil || port2 == nil {
			return &InvalidParametersError{
				BaseError: BaseError{
					Message: "Parameters ipAddress, port, ipAddress2, and port2 are required",
				},
			}
		}

		return o.sendCommand(fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s,%s,%s,%s", direction, ioId, ioType, flags, localInterface, *ipAddress, *port, *ipAddress2, *port2))
	} else if ioType == OXTEL_EXT_IO_TYPE_2110 {
		if sdpFileName == nil {
			return &InvalidParametersError{
				BaseError: BaseError{
					Message: "Parameter sdpFileName is required",
				},
			}
		}

		return o.sendCommand(fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s", direction, ioId, ioType, flags, localInterface, *sdpFileName))
	}

	return &InvalidParametersError{
		BaseError: BaseError{
			Message: "SDI does not apply to this command",
		},
	}
}

// SetExternalIODynamicConfiguration_AsString returns the command used to set the dynamic configuration for the specified External IO. The External IO is
//
// For use with scheduled commands.
func SetExternalIODynamicConfiguration_AsString(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType, flags uint16, localInterface string, ipAddress *string, port *string, ipAddress2 *string, port2 *string, sdpFileName *string) string {

	if ioType == OXTEL_EXT_IO_TYPE_2022_6 {
		return fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s,%s", direction, ioId, ioType, flags, localInterface, *ipAddress, *port)
	} else if ioType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
		return fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s,%s,%s,%s", direction, ioId, ioType, flags, localInterface, *ipAddress, *port, *ipAddress2, *port2)
	} else if ioType == OXTEL_EXT_IO_TYPE_2110 {
		return fmt.Sprintf("hXDC%02x%02x%02x%02x,%s,%s", direction, ioId, ioType, flags, localInterface, *sdpFileName)
	}

	return "hXDC"
}

// EnquireExternalIODynamicConfiguration gets the dynamic configuration for the specified External IO. The External IO is defined
// by the direction and IO Id.
//
// For SDI, this command is not applicable.
//
// Response is an ExternalIODynamicConfigurationResponse. Not all fields will be used.
func (o *Oxtel) EnquireExternalIODynamicConfiguration(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType) (ExternalIODynamicConfigurationResponse, error) {
	val, err := o.sendCommandExpectResponse("hXDC", fmt.Sprintf("%02x%02x%02x", direction, ioId, ioType))
	if err != nil {
		return ExternalIODynamicConfigurationResponse{}, err
	}

	outDir, err := strconv.ParseUint(string(val[:2]), 16, 8)
	if err != nil {
		return ExternalIODynamicConfigurationResponse{}, err
	}

	outId, err := strconv.ParseUint(string(val[2:4]), 16, 8)
	if err != nil {
		return ExternalIODynamicConfigurationResponse{}, err
	}

	outType, err := strconv.ParseUint(string(val[4:6]), 16, 8)
	if err != nil {
		return ExternalIODynamicConfigurationResponse{}, err
	}

	parts := strings.Split(val, ",")
	localInterface := parts[1]
	var ipAddress *string
	var port *uint32
	var ipAddress2 *string
	var port2 *uint32
	var sdpFileName *string

	if outType == OXTEL_EXT_IO_TYPE_2022_6 {
		ipAddress = &parts[2]
		tmpPort, err := strconv.ParseUint(parts[3], 10, 32)
		if err != nil {
			return ExternalIODynamicConfigurationResponse{}, err
		}
		w := uint32(tmpPort)
		port = &w
	}
	if outType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
		ipAddress = &parts[2]
		tmpPort, err := strconv.ParseUint(parts[3], 10, 32)
		if err != nil {
			return ExternalIODynamicConfigurationResponse{}, err
		}
		w := uint32(tmpPort)
		port = &w

		ipAddress2 = &parts[4]
		tmpPort2, err := strconv.ParseUint(parts[5], 10, 32)
		if err != nil {
			return ExternalIODynamicConfigurationResponse{}, err
		}

		x := uint32(tmpPort2)
		port2 = &x
	}

	if outType == OXTEL_EXT_IO_TYPE_2110 {
		sdpFileName = &parts[2]
	}

	return ExternalIODynamicConfigurationResponse{
		IOType:         OxtelExternalIOType(outType),
		IODirection:    OxtelExternalIODirection(outDir),
		IOId:           OxtelExternalIOId(outId),
		IPAddress:      ipAddress,
		Port:           port,
		IPAddress2:     ipAddress2,
		Port2:          port2,
		SDPFileName:    sdpFileName,
		LocalInterface: &localInterface,
	}, nil
}

// EnquireExternalIODynamicConfiguration_AsString returns the command used to get the dynamic configuration for the specified External IO. The External IO is defined
//
// For use with scheduled commands.
func EnquireExternalIODynamicConfiguration_AsString(direction OxtelExternalIODirection, ioId OxtelExternalIOId, ioType OxtelExternalIOType) string {
	return "hXDC" + fmt.Sprintf("%02x%02x%02x", direction, ioId, ioType)
}

// EnableExternalIOTally enables/disables the External IO Tally for the connection on which the command was received.
//
// When enabled, unsolicited tallies will be sent so that the client can record the initial sate.
func (o *Oxtel) EnableExternalIOTally(enable bool) error {
	return o.sendCommand(fmt.Sprintf("hXIOT%01x", boolToInt(enable)))
}

// EnableExternalIOTally_AsString returns the command string used to enable/disable the External IO Tally for the connection on which the command was received.
//
// For use with scheduled commands.
func EnableExternalIOTally_AsString(enable bool) string {
	return fmt.Sprintf("hXIOT%01x", boolToInt(enable))
}

// EnquireExternalIOTally queries if external IO tally is enabled for the connection on which the command was received.
//
// Response is a boolean.
func (o *Oxtel) EnquireExternalIOTally() (bool, error) {
	val, err := o.sendCommandExpectResponse("hXIOT", "")
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(val)
}

// EnquireExternalIOTally_AsString returns the command string used to query if external IO tally is enabled for the connection on which the command was received.
//
// For use with scheduled commands.
func EnquireExternalIOTally_AsString() string {
	return "hXIOT"
}

// EnquireExternalIOSupported returns whether or not the External IO feature is supported.
//
// Response is a boolean representing if External IO is supported.
func (o *Oxtel) EnquireExternalIOSupported() (bool, error) {
	val, err := o.sendCommandExpectResponse("hEXTIO", "")
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(val)
}

// EnquireExternalIOSupported_AsString returns the command string used to return whether or not the External IO feature is supported.
//
// For use with scheduled commands.
func EnquireExternalIOSupported_AsString() string {
	return "hEXTIO"
}

// EnquireExternalInputs returns a list of the external inputs supported by the channel. This command can be used to dynamically
// determine the external inputs that are currently supported for the channel configuration.
//
// Returns an ExternalInputsResponse. The VideoSourceId is the same value as defined for the SelectMixerInputCommand
func (o *Oxtel) EnquireExternalInputs() (ExternalInputsResponse, error) {
	val, err := o.sendCommandExpectResponse("hXIN", "")
	if err != nil {
		return ExternalInputsResponse{}, err
	}

	parts := strings.Split(val, ";")
	num, err := strconv.ParseUint(parts[0], 16, 8)
	if err != nil {
		return ExternalInputsResponse{}, err
	}

	var inputs []ExternalInput

	for i := 1; i < len(parts); i += 2 {
		sourceId, err := strconv.ParseUint(parts[i+1], 10, 8)
		if err != nil {
			return ExternalInputsResponse{}, err
		}

		inputs = append(inputs, ExternalInput{
			Name:          parts[i],
			VideoSourceId: uint8(sourceId),
		})
	}

	return ExternalInputsResponse{
		NumberOfInputs: uint8(num),
		ExternalInputs: inputs,
	}, nil
}

// EnquireExternalInputs_AsString returns the command string used to return a list of the external inputs supported by the channel. This command can be used to dynamically
//
// For use with scheduled commands.
func EnquireExternalInputs_AsString() string {
	return "hXIN"
}

// EnquireOutputs returns a list of the outputs supported for the channel. This command can be used to dynamically determine
// the outputs that are currently supported for the channel configuration.
//
// Returns an ExternalOutputsResponse.
func (o *Oxtel) EnquireOutputs() (ExternalOutputsResponse, error) {
	val, err := o.sendCommandExpectResponse("hOUT", "")
	if err != nil {
		return ExternalOutputsResponse{}, err
	}

	parts := strings.Split(val, ";")
	num, err := strconv.ParseUint(parts[0], 16, 8)
	if err != nil {
		return ExternalOutputsResponse{}, err
	}

	var outputs []ExternalOutput

	for i := 1; i < len(parts); i += 2 {
		outId, err := strconv.ParseUint(parts[i+1], 10, 8)
		if err != nil {
			return ExternalOutputsResponse{}, err
		}

		outputs = append(outputs, ExternalOutput{
			Name: parts[i],
			Id:   OxtelExternalIOId(outId),
		})
	}

	return ExternalOutputsResponse{
		NumberOfOutputs: uint8(num),
		ExternalOutputs: outputs,
	}, nil
}

// EnquireOutputs_AsString returns the command string used to return a list of the outputs supported for the channel. This command can be used to dynamically determine
//
// For use with scheduled commands.
func EnquireOutputs_AsString() string {
	return "hOUT"
}
