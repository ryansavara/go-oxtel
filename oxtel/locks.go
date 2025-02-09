package oxtel

import (
	"fmt"
	"strconv"
)

// SetSessionLocks sets the session lock for the specified item for this connection.
//
// A session lock is a mechanism that Oxtel clients can use to prevent specific items from being changed via the Oxtel protocol.
// Any Oxtel client can enable a Session Lock. When one client has a Session Lock on a defined item, a client that does not have the
// same Session Lock on the item will be prevented from making changes.
// All clients with a Session Lock on an item will be allowed to make changes.
//
// A Session Lock is a connection-based meaning that when the Oxtel connection is closed, the lock is released.
//
// Items that can be locked with a Session Lock include: Mixer, Graphic Layer 0-7.
//
// Use BuildSessionLocks to get the locks bitwise value.
func (o *Oxtel) SetSessionLocks(locks int32) error {
	return o.sendCommand(fmt.Sprintf("hSL%08x", locks))
}

// SetSessionLocks_AsString returns the command string used to set the session lock for the specified item for this connection.
//
// For use with scheduled commands.
func SetSessionLocks_AsString(locks int32) string {
	return fmt.Sprintf("hSL%08x", locks)
}

// EnquireSessionLocks queries the connection's session locks.
//
// Response is a LocksResponse.
func (o *Oxtel) EnquireSessionLocks() (LocksResponse, error) {
	val, err := o.sendCommandExpectResponse("hSL", "")
	if err != nil {
		return LocksResponse{}, err
	}

	result, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		return LocksResponse{}, err
	}

	mixer := (result & 0x00000001) != 0
	layer0 := (result & 0x00000100) != 0
	layer1 := (result & 0x00000200) != 0
	layer2 := (result & 0x00000400) != 0
	layer3 := (result & 0x00000800) != 0
	layer4 := (result & 0x00001000) != 0
	layer5 := (result & 0x00002000) != 0
	layer6 := (result & 0x00004000) != 0
	layer7 := (result & 0x00008000) != 0

	return LocksResponse{
		Mixer:  mixer,
		Layer0: layer0,
		Layer1: layer1,
		Layer2: layer2,
		Layer3: layer3,
		Layer4: layer4,
		Layer5: layer5,
		Layer6: layer6,
		Layer7: layer7,
	}, nil

}

// EnquireSessionLocks_AsString returns the command string used to query the connection's session locks.
//
// For use with scheduled commands.
func EnquireSessionLocks_AsString() string {
	return "hSL"
}

// EnquireGlobalSessionLocks queries the connection's session locks.
//
// Response is a LocksResponse.
func (o *Oxtel) EnquireGlobalSessionLocks() (LocksResponse, error) {
	val, err := o.sendCommandExpectResponse("hGSL", "")
	if err != nil {
		return LocksResponse{}, err
	}

	result, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		return LocksResponse{}, err
	}

	mixer := (result & 0x00000001) != 0
	layer0 := (result & 0x00000100) != 0
	layer1 := (result & 0x00000200) != 0
	layer2 := (result & 0x00000400) != 0
	layer3 := (result & 0x00000800) != 0
	layer4 := (result & 0x00001000) != 0
	layer5 := (result & 0x00002000) != 0
	layer6 := (result & 0x00004000) != 0
	layer7 := (result & 0x00008000) != 0

	return LocksResponse{
		Mixer:  mixer,
		Layer0: layer0,
		Layer1: layer1,
		Layer2: layer2,
		Layer3: layer3,
		Layer4: layer4,
		Layer5: layer5,
		Layer6: layer6,
		Layer7: layer7,
	}, nil

}

// EnquireGlobalSessionLocks_AsString returns the command string used to query the connection's session locks.
//
// For use with scheduled commands.
func EnquireGlobalSessionLocks_AsString() string {
	return "hGSL"
}

// SetPermanentLocks sets the permanent lock for the specified item.
//
// Use BuildSessionLocks to get the lock bitwise value.
func (o *Oxtel) SetPermanentLocks(locks int32) error {
	return o.sendCommand(fmt.Sprintf("hPL%08x", locks))
}

// SetPermanentLocks_AsString returns the command string used to set the permanent lock for the specified item.
//
// For use with scheduled commands.
func SetPermanentLocks_AsString(locks int32) string {
	return fmt.Sprintf("hPL%08x", locks)
}

// EnquirePermanentLocks queries the permanent locks.
//
// Response is a LocksResponse.
func (o *Oxtel) EnquirePermanentLocks() (LocksResponse, error) {
	val, err := o.sendCommandExpectResponse("hPL", "")
	if err != nil {
		return LocksResponse{}, err
	}

	result, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		return LocksResponse{}, err
	}

	mixer := (result & 0x00000001) != 0
	layer0 := (result & 0x00000100) != 0
	layer1 := (result & 0x00000200) != 0
	layer2 := (result & 0x00000400) != 0
	layer3 := (result & 0x00000800) != 0
	layer4 := (result & 0x00001000) != 0
	layer5 := (result & 0x00002000) != 0
	layer6 := (result & 0x00004000) != 0
	layer7 := (result & 0x00008000) != 0

	return LocksResponse{
		Mixer:  mixer,
		Layer0: layer0,
		Layer1: layer1,
		Layer2: layer2,
		Layer3: layer3,
		Layer4: layer4,
		Layer5: layer5,
		Layer6: layer6,
		Layer7: layer7,
	}, nil
}

// EnquirePermanentLocks_AsString returns the command string used to query the permanent locks.
//
// For use with scheduled commands.
func EnquirePermanentLocks_AsString() string {
	return "hPL"
}

// EnableOxtelLockTally enables or disables the Oxtel Lock Tally for the connection on which the command was received.
// When enabled, the Oxtel Lock Tally will be transmitted so the client can record the initial state.
//
// An unsolicited response will be transmitted for each item and its current lock status.
func (o *Oxtel) EnableOxtelLockTally(enable bool) error {
	return o.sendCommand(fmt.Sprintf("hOLT%01x", boolToInt(enable)))
}

// EnableOxtelLockTally_AsString returns the command string used to enable or disable the Oxtel Lock Tally for the connection on which the command was received.
//
// For use with scheduled commands.
func EnableOxtelLockTally_AsString(enable bool) string {
	return fmt.Sprintf("hOLT%01x", boolToInt(enable))
}

// EnquireOxtelLockTally queries the enable/disable state of the Oxtel Lock tally.
//
// Response is a boolean representing if oxtel lock tally is enabled.
func (o *Oxtel) EnquireOxtelLockTally() (bool, error) {
	val, err := o.sendCommandExpectResponse("hOLT", "")
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(val)
}

// EnquireOxtelLockTally_AsString returns the command string used to query the enable/disable state of the Oxtel Lock tally.
//
// For use with scheduled commands.
func EnquireOxtelLockTally_AsString() string {
	return "hOLT"
}
