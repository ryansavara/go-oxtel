package oxtel

type UnsolicitedMessage struct {
	Raw string
}

type KeyerPositionTally struct {
	UnsolicitedMessage
	Layer     OxtelLayer
	Direction OxtelDirection
}

type ImageLoadTally struct {
	UnsolicitedMessage
	Layer    OxtelLayer
	Template string
}

type ImagePreloadTally struct {
	UnsolicitedMessage
	Layer    OxtelLayer
	Template string
}

type MediaTally struct {
	UnsolicitedMessage
	MediaType MediaTallies
	Action    OxtelMediaTallies
	Filename  string
}

type PlayStateTally struct {
	UnsolicitedMessage
	Layer OxtelLayer
	State OxtelPlayStateTally
}

type VideoTally struct {
	UnsolicitedMessage
	MixerInput   uint8
	Layer0       OxtelDirection
	Layer1       OxtelDirection
	MixerASource OxtelVideoSource
	MixerBSource OxtelVideoSource
	Unused1      uint8
	Unused2      uint8
}

type AudioProfileTally struct {
	UnsolicitedMessage
	Source  OxtelAudioSource
	Profile uint8
}

type LockTally struct {
	UnsolicitedMessage
	SessionLocks   LocksResponse
	PermanentLocks LocksResponse
}

type ExternalIOSourceChangedTally struct {
	UnsolicitedMessage
	IODirection     OxtelExternalIODirection
	IOId            OxtelExternalIOId
	IOType          OxtelExternalIOType
	ConfigurationId uint8
	State           uint8
}

type ExternalIODynamicConfigChangedTally struct {
	UnsolicitedMessage
	IOType         OxtelExternalIOType
	IOId           OxtelExternalIOId
	IODirection    OxtelExternalIODirection
	LocalInterface *string
	IPAddress      *string
	Port           *uint32
	IPAddress2     *string
	Port2          *uint32
	SDPFileName    *string
}
