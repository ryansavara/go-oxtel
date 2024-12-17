package oxtel

type FileInfoResponse struct {
	Exists   bool
	Filename string
}

type FileQueryResponse struct {
	EndOfDir bool
	Filename string
}

type ExtendedFileInfoResponse struct {
	Exists          bool
	XPosition       int64
	YPosition       int64
	Width           int64
	Height          int64
	Clip            int64
	Gain            int64
	Transparency    int64
	ImageType       int64
	Frames          int64
	AnimationMode   int64
	LoadTime        int64
	AssociatedAudio int64
	Filename        string
}

type ValidateTemplateResponse struct {
	Filename      string
	FileExists    bool
	MissingAssets int16
}

type LayerTemplateResponse struct {
	Layer    OxtelLayer
	Filename string
}

type ImagePositionResponse struct {
	Layer   OxtelLayer
	XOffset int
	YOffset int
}

type MediaTallies struct {
	Unused1 bool
	Unused2 bool
	Unused3 bool
	Unused4 bool
	Unused5 bool
	Images  bool
}

type MixerInputResponse struct {
	Input  OxtelMixerInput
	Source OxtelVideoSource
}

type MixModeResponse struct {
	TransitionType uint8
	ABMixRate      uint16
	WipeSoftness   uint8
	ABMixAngle     uint16
	VFadeColor     uint8
}

type ColorGeneratorResponse struct {
	Unit  uint8
	Red   uint8
	Green uint8
	Blue  uint8
}

type AudioProfileResponse struct {
	Source  OxtelAudioSource
	Profile uint8
}

type AudioProgram struct {
	Sdi          uint8
	AudioProgram OxtelAudioProgram
	JungerPreset OxtelJungerPreset
	Channel1     uint8
	Channel2     *uint8
	Channel3     *uint8
	Channel4     *uint8
	Channel5     *uint8
	Channel6     *uint8
	Channel7     *uint8
	Channel8     *uint8
}

type AudioLoudnessConfigurationResponse struct {
	Channel1 bool
	Channel2 bool
	Channel3 bool
	Channel4 bool
	Channel5 bool
	Channel6 bool
	Channel7 bool
	Channel8 bool
}

type AudioLoudnessResponse struct {
	Channel1     bool
	Channel2     bool
	Channel3     bool
	Channel4     bool
	Channel5     bool
	Channel6     bool
	Channel7     bool
	Channel8     bool
	Sdi          *uint8
	AudioProgram *OxtelAudioProgram
	JungerPreset *OxtelJungerPreset
}

type AudioABFollowVideoABResponse struct {
	Unused1 uint8
	Enabled bool
}

type ChannelMask struct {
	Channel1  bool
	Channel2  bool
	Channel3  bool
	Channel4  bool
	Channel5  bool
	Channel6  bool
	Channel7  bool
	Channel8  bool
	Channel9  bool
	Channel10 bool
	Channel11 bool
	Channel12 bool
	Channel13 bool
	Channel14 bool
	Channel15 bool
	Channel16 bool
}

type AudioGainResponse struct {
	Source      OxtelAudioSource
	ChannelMask string
	Gain        int8
}

type LatencyResponse struct {
	Source  OxtelLatencySource
	Latency int8
}

type SystemStatusResponse struct {
	SystemMode        uint8
	VersionHigh       uint16
	VersionLow        uint16
	VideoStandard     OxtelVideoStandard
	PreviewSource     uint16
	FadeRateDSK1      uint16
	FadeRateDSK2      uint16
	FTBRateDSK1       uint16
	FTBRateDSK2       uint16
	SystemNotAccessed uint8
}

type VideoLayerStatusResponse struct {
	LayerFaderAngle uint16
	LayerFTBAngle   uint16
	Unused1         uint16
	Unused2         uint16
	Unused3         uint16
}

type CommandAvailabilityResponse struct {
	CommandByte1 byte
	CommandByte2 byte
	Supported    bool
}

type SlaveLayerStatusResponse struct {
	Layer0State bool
	Layer1State bool
	Layer2State bool
	Layer3State bool
	Layer4State bool
	Layer5State bool
	Layer6State bool
	Layer7State bool
	Unused      int32
}

type FullVersionNumberResponse struct {
	Major       int
	Minor       int
	Patch       int
	Branch      int
	BuildNumber string
	AsString    string
}

type CurrentTimeResponse struct {
	FieldRate OxtelFieldRate
	Hours     uint8
	Minutes   uint8
	Seconds   uint8
	Frames    uint8
}

type LocksResponse struct {
	Mixer  bool
	Layer0 bool
	Layer1 bool
	Layer2 bool
	Layer3 bool
	Layer4 bool
	Layer5 bool
	Layer6 bool
	Layer7 bool
}

type NumberOfExternalIOConfigurationsResponse struct {
	IOType            OxtelExternalIOType
	IODirection       OxtelExternalIODirection
	NumConfigurations uint8
}

type ExternalIOConfigurationResponse struct {
	IOType          OxtelExternalIOType
	IODirection     OxtelExternalIODirection
	Index           uint8
	ConfigurationId uint8
	Name            string
	LocalInterface  *string
	IPAddress       *string
	Port            *uint32
	IPAddress2      *string
	Port2           *uint32
	SDPFileName     *string
}

type ExternalIOSourceResponse struct {
	IODirection     OxtelExternalIODirection
	IOId            OxtelExternalIOId
	IOType          OxtelExternalIOType
	ConfigurationId uint8
	State           uint8
}

type ExternalIODynamicConfigurationResponse struct {
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

type ExternalInputsResponse struct {
	NumberOfInputs uint8
	ExternalInputs []ExternalInput
}

type ExternalInput struct {
	Name          string
	VideoSourceId uint8
}

type ExternalOutputsResponse struct {
	NumberOfOutputs uint8
	ExternalOutputs []ExternalOutput
}

type ExternalOutput struct {
	Name string
	Id   OxtelExternalIOId
}
