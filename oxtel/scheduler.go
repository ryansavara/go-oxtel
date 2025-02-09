package oxtel

import (
	"fmt"
	"strconv"
)

// AddScheduledCommand schedules an automation command at the specified time in hours, minutes, seconds, and frames.
//
// Scheduled commands specify a timecode value. This value should be specified as the time at which the command should be _recognized_.
// i.e. Toxt + Tmcs before the command's results are seen on the SDI output. Toxt is the Oxtel latency in reference fields
// and Tmcs is the MCS latency in reference fields.
//
// Refer to the Timing Model section for more details about latencies.
//
// Each command available in this module has a "_AsString" counterpart that will return the command as a string for
// use with this command.
//
// Note: If a command is scheduled for more than 23 hours and 59 minutes in the future, then the command will be issued immediate.
func (o *Oxtel) AddScheduledCommand(hours uint8, minutes uint8, seconds uint8, frames uint, command string) error {
	return o.sendCommand(fmt.Sprintf("i0%02d%02d%02d%02d;%s", hours, minutes, seconds, frames, command))
}

// AddScheduledCommand_AsString returns the command string used to schedule an automation command at the specified time
// in hours, minutes, seconds, and frames.
//
// For use with scheduled commands.
func AddScheduledCommand_AsString(hours uint8, minutes uint8, seconds uint8, frames uint, command string) string {
	return fmt.Sprintf("i0%02d%02d%02d%02d;%s", hours, minutes, seconds, frames, command)
}

// DeleteAllScheduledCommands deletes all scheduled commands.
func (o *Oxtel) DeleteAllScheduledCommands() error {
	return o.sendCommand("i2")
}

// DeleteAllScheduledCommands_AsString returns the command used to delete all scheduled commands.
//
// For use with scheduled commands.
func DeleteAllScheduledCommands_AsString() string {
	return "i2"
}

// EnquireCurrentTime queries the current time as referenced to VITC.
func (o *Oxtel) EnquireCurrentTime() (CurrentTimeResponse, error) {
	val, err := o.sendCommandExpectResponse("ix", "")
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	rate, err := strconv.ParseUint(string(val[0]), 10, 8)
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	hours, err := strconv.ParseUint(string(val[1:3]), 10, 8)
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	minutes, err := strconv.ParseUint(string(val[3:5]), 10, 8)
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	seconds, err := strconv.ParseUint(string(val[5:7]), 10, 8)
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	frames, err := strconv.ParseUint(string(val[7:9]), 10, 8)
	if err != nil {
		return CurrentTimeResponse{}, err
	}

	return CurrentTimeResponse{
		FieldRate: OxtelFieldRate(rate),
		Hours:     uint8(hours),
		Minutes:   uint8(minutes),
		Seconds:   uint8(seconds),
		Frames:    uint8(frames),
	}, nil
}

// EnquireCurrentTime_AsString returns the command string used to query the current time as referenced to VITC.
//
// For use with scheduled commands.
func EnquireCurrentTime_AsString() string {
	return "ix"
}
