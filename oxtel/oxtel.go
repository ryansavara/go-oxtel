package oxtel

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type Oxtel struct {
	address     string
	port        uint16
	conn        net.Conn
	reader      bufio.Reader
	RxMessages  chan string
	lastCommand string
	Unsolicited chan interface{}
}

func NewOxtel(address string, port uint16) *Oxtel {
	return &Oxtel{
		address:     address,
		port:        port,
		conn:        nil,
		RxMessages:  make(chan string),
		Unsolicited: make(chan interface{}),
	}
}

func (o *Oxtel) Connect() error {
	address := fmt.Sprintf("%s:%d", o.address, o.port)

	c, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	o.conn = c
	o.reader = *bufio.NewReader(c)
	go o.rxLoop()
	return nil
}

func (o *Oxtel) Disconnect() error {
	if o.conn == nil {
		return nil
	}
	close(o.RxMessages)
	close(o.Unsolicited)
	err := o.conn.Close()

	if err == io.EOF {
		err = nil
	}
	o.conn = nil
	return err
}

func (o *Oxtel) rxLoop() {
	for {
		line, err := o.reader.ReadString(':')
		if err != nil {
			if err == io.EOF {
				break
			}
			o.Disconnect()
			break
		}

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			o.handleMessage(line)
		}
	}
}

func (o *Oxtel) handleMessage(message string) {
	if len(o.lastCommand) > 0 && message[:len(o.lastCommand)] == o.lastCommand {
		o.RxMessages <- message
	} else {
		cleanMessage := message[:len(message)-1]
		if message[0] == '3' {
			data := cleanMessage[1:]
			layer, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse layer from KeyerPositionTally")
			}
			direction, err := strconv.ParseUint(string(data[2]), 10, 8)
			if err != nil {
				panic("unable to parse direction from KeyerPositionTally")
			}

			o.Unsolicited <- KeyerPositionTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				Layer:     OxtelLayer(layer),
				Direction: OxtelDirection(direction),
			}
		} else if message[0:2] == "Y9" {
			data := cleanMessage[2:]
			layer, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse layer from ImageLoadTally")
			}

			filename := data[1:]

			o.Unsolicited <- ImageLoadTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				Layer:    OxtelLayer(layer),
				Template: filename,
			}
		} else if message[0:2] == "YA" {
			data := cleanMessage[2:]
			layer, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse layer from ImagePreloadTally")
			}

			filename := data[1:]

			o.Unsolicited <- ImagePreloadTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				Layer:    OxtelLayer(layer),
				Template: filename,
			}
		} else if message[0:2] == "YB" {
			data := cleanMessage[2:]

			u1, err := strconv.ParseBool(string(data[0]))
			if err != nil {
				panic("unable to parse unused1 from MediaTallies")
			}

			u2, err := strconv.ParseBool(string(data[1]))
			if err != nil {
				panic("unable to parse unused2 from MediaTallies")
			}

			u3, err := strconv.ParseBool(string(data[2]))
			if err != nil {
				panic("unable to parse unused3 from MediaTallies")
			}

			u4, err := strconv.ParseBool(string(data[3]))
			if err != nil {
				panic("unable to parse unused4 from MediaTallies")
			}

			u5, err := strconv.ParseBool(string(data[4]))
			if err != nil {
				panic("unable to parse unused5 from MediaTallies")
			}

			images, err := strconv.ParseBool(string(data[5]))
			if err != nil {
				panic("unable to parse images from MediaTallies")
			}

			action, err := strconv.ParseUint(string(data[6]), 16, 8)
			if err != nil {
				panic("unable to parse action from MediaTallies")
			}

			filename := data[7:]

			o.Unsolicited <- MediaTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				MediaType: MediaTallies{
					Unused1: u1,
					Unused2: u2,
					Unused3: u3,
					Unused4: u4,
					Unused5: u5,
					Images:  images,
				},
				Action:   OxtelMediaTallies(action),
				Filename: filename,
			}
		} else if message[:2] == "YS" {
			data := cleanMessage[2:]
			layer, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse layer from PlayStateTally")
			}

			state, err := strconv.ParseUint(string(data[1]), 16, 8)
			if err != nil {
				panic("unable to parse state from PlayStateTally")
			}

			o.Unsolicited <- PlayStateTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				Layer: OxtelLayer(layer),
				State: OxtelPlayStateTally(state),
			}
		} else if message[:2] == "Y6" {
			data := cleanMessage[2:]
			abmix, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse a/b mix from VideoTally")
			}
			layer0, err := strconv.ParseUint(string(data[1]), 16, 8)
			if err != nil {
				panic("unable to parse layer 0 keyer from VideoTally")
			}
			layer1, err := strconv.ParseUint(string(data[2]), 16, 8)
			if err != nil {
				panic("unable to parse layer 1 keyer from VideoTally")
			}
			mixerA, err := strconv.ParseUint(string(data[3]), 16, 8)
			if err != nil {
				panic("unable to parse layer 0 keyer from VideoTally")
			}
			mixerB, err := strconv.ParseUint(string(data[4]), 16, 8)
			if err != nil {
				panic("unable to parse layer 0 keyer from VideoTally")
			}
			u1, err := strconv.ParseUint(string(data[5:7]), 16, 8)
			if err != nil {
				panic("unable to parse layer 0 keyer from VideoTally")
			}
			u2, err := strconv.ParseUint(string(data[7:9]), 16, 8)
			if err != nil {
				panic("unable to parse layer 0 keyer from VideoTally")
			}

			o.Unsolicited <- VideoTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				MixerInput:   uint8(abmix),
				Layer0:       OxtelDirection(layer0),
				Layer1:       OxtelDirection(layer1),
				MixerASource: OxtelVideoSource(mixerA),
				MixerBSource: OxtelVideoSource(mixerB),
				Unused1:      uint8(u1),
				Unused2:      uint8(u2),
			}
		} else if message[:3] == "jAY" {
			data := cleanMessage[3:]
			source, err := strconv.ParseUint(string(data[0]), 16, 8)
			if err != nil {
				panic("unable to parse source from AudioProfileTally")
			}
			profile, err := strconv.ParseUint(string(data[1:3]), 16, 8)
			if err != nil {
				panic("unable to parse profile from AudioProfileTally")
			}

			o.Unsolicited <- AudioProfileTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				Source:  OxtelAudioSource(source),
				Profile: uint8(profile),
			}
		} else if message[:4] == "hOLY" {
			data := cleanMessage[4:]
			session, err := strconv.ParseUint(string(data[:8]), 16, 32)
			if err != nil {
				panic("unable to parse session locks from LockTally")
			}
			permanent, err := strconv.ParseUint(string(data[8:16]), 16, 32)
			if err != nil {
				panic("unable to parse permanent lock from LockTally")
			}

			sMixer := (session & 0x00000001) != 0
			sLayer0 := (session & 0x00000100) != 0
			sLayer1 := (session & 0x00000200) != 0
			sLayer2 := (session & 0x00000400) != 0
			sLayer3 := (session & 0x00000800) != 0
			sLayer4 := (session & 0x00001000) != 0
			sLayer5 := (session & 0x00002000) != 0
			sLayer6 := (session & 0x00004000) != 0
			sLayer7 := (session & 0x00008000) != 0

			pMixer := (permanent & 0x00000001) != 0
			pLayer0 := (permanent & 0x00000100) != 0
			pLayer1 := (permanent & 0x00000200) != 0
			pLayer2 := (permanent & 0x00000400) != 0
			pLayer3 := (permanent & 0x00000800) != 0
			pLayer4 := (permanent & 0x00001000) != 0
			pLayer5 := (permanent & 0x00002000) != 0
			pLayer6 := (permanent & 0x00004000) != 0
			pLayer7 := (permanent & 0x00008000) != 0

			o.Unsolicited <- LockTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				SessionLocks: LocksResponse{
					Mixer:  sMixer,
					Layer0: sLayer0,
					Layer1: sLayer1,
					Layer2: sLayer2,
					Layer3: sLayer3,
					Layer4: sLayer4,
					Layer5: sLayer5,
					Layer6: sLayer6,
					Layer7: sLayer7,
				},
				PermanentLocks: LocksResponse{
					Mixer:  pMixer,
					Layer0: pLayer0,
					Layer1: pLayer1,
					Layer2: pLayer2,
					Layer3: pLayer3,
					Layer4: pLayer4,
					Layer5: pLayer5,
					Layer6: pLayer6,
					Layer7: pLayer7,
				},
			}
		} else if message[:4] == "hXSY" {
			data := cleanMessage[4:]
			direction, err := strconv.ParseUint(string(data[:2]), 16, 8)
			if err != nil {
				panic("unable to parse direction from ExternalIOSourceChangedTally")
			}
			ioID, err := strconv.ParseUint(string(data[2:4]), 16, 8)
			if err != nil {
				panic("unable to parse io id from ExternalIOSourceChangedTally")
			}
			ioType, err := strconv.ParseUint(string(data[4:6]), 16, 8)
			if err != nil {
				panic("unable to parse io type from ExternalIOSourceChangedTally")
			}
			configId, err := strconv.ParseUint(string(data[6:8]), 16, 8)
			if err != nil {
				panic("unable to parse config id from ExternalIOSourceChangedTally")
			}
			state, err := strconv.ParseUint(string(data[8:10]), 16, 8)
			if err != nil {
				panic("unable to parse state from ExternalIOSourceChangedTally")
			}

			o.Unsolicited <- ExternalIOSourceChangedTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				IODirection:     OxtelExternalIODirection(direction),
				IOId:            OxtelExternalIOId(ioID),
				IOType:          OxtelExternalIOType(ioType),
				ConfigurationId: uint8(configId),
				State:           uint8(state),
			}
		} else if message[:5] == "hXDCY" {
			data := cleanMessage[5:]
			direction, err := strconv.ParseUint(string(data[:2]), 16, 8)
			if err != nil {
				panic("unable to parse direction from ExternalIODynamicConfigChangedTally")
			}
			ioId, err := strconv.ParseUint(string(data[2:4]), 16, 8)
			if err != nil {
				panic("unable to parse io id from ExternalIODynamicConfigChangedTally")
			}
			ioType, err := strconv.ParseUint(string(data[4:6]), 16, 8)
			if err != nil {
				panic("unable to parse io type from ExternalIODynamicConfigChangedTally")
			}

			parts := strings.Split(data, ",")

			var localInterface *string
			var ipAddress *string
			var port *uint32
			var ipAddress2 *string
			var port2 *uint32
			var sdfFileName *string

			if ioType != OXTEL_EXT_IO_TYPE_SDI {
				localInterface = &parts[1]
			}
			if ioType == OXTEL_EXT_IO_TYPE_2022_6 || ioType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
				ipAddress = &parts[2]
				tmpPort, err := strconv.ParseUint(parts[3], 10, 32)
				if err != nil {
					panic("unable to parse port from ExternalIODynamicConfigChangedTally")
				}
				w := uint32(tmpPort)
				port = &w
			}
			if ioType == OXTEL_EXT_IO_TYPE_2022_6_2022_7 {
				ipAddress2 = &parts[4]
				tmpPort, err := strconv.ParseUint(parts[5], 10, 32)
				if err != nil {
					panic("unable to parse port2 from ExternalIODynamicConfigChangedTally")
				}
				w := uint32(tmpPort)
				port2 = &w
			}

			if ioType == OXTEL_EXT_IO_TYPE_2110 {
				sdfFileName = &parts[2]
			}

			o.Unsolicited <- ExternalIODynamicConfigChangedTally{
				UnsolicitedMessage: UnsolicitedMessage{
					Raw: cleanMessage,
				},
				IODirection:    OxtelExternalIODirection(direction),
				IOId:           OxtelExternalIOId(ioId),
				IOType:         OxtelExternalIOType(ioType),
				LocalInterface: localInterface,
				IPAddress:      ipAddress,
				Port:           port,
				IPAddress2:     ipAddress2,
				Port2:          port2,
				SDPFileName:    sdfFileName,
			}
		} else {
			unsolictedMsg := UnsolicitedMessage{
				Raw: message,
			}
			o.Unsolicited <- unsolictedMsg
		}
	}
}

func (o *Oxtel) sendCommand(cmd string) error {
	escapedCmd := strings.ReplaceAll(cmd, "\\", "\\5C")
	escapedCmd = strings.ReplaceAll(escapedCmd, "|", "\\7C")
	escapedCmd = strings.ReplaceAll(escapedCmd, ";", "\\3B")
	escapedCmd = strings.ReplaceAll(escapedCmd, ":", "\\3A")

	escapedCmd += ":"
	cmdBytes := []byte(escapedCmd)

	_, err := o.conn.Write(cmdBytes)
	if err != nil {
		if err == io.EOF {
			o.Disconnect()
		}
	}
	return err
}

func (o *Oxtel) sendCommandExpectResponse(cmd string, data string) (string, error) {
	o.lastCommand = cmd
	err := o.sendCommand(cmd + data)
	if err != nil && err != io.EOF {
		return "", err
	}

	timeout := time.After(5 * time.Second)
	for {
		select {
		case unsolicited := <-o.RxMessages:
			if unsolicited[:len(cmd)] == cmd {
				return unsolicited[len(cmd) : len(unsolicited)-1], nil
			}
		case <-timeout:
			return "", &TimeoutError{
				BaseError: BaseError{
					Message: "Timed out waiting for response",
				},
			}
		}
	}
}
