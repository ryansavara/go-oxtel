package oxtel

type OxtelLayer uint8
type OxtelDirection uint8
type OxtelKeyerPositionTally uint8
type OxtelEnquireFileInfoResponse uint8
type OxtelMediaTallies uint8
type OxtelPlayStateTally uint8
type OxtelUpdateTextFieldFlag uint8
type OxtelTransitionType uint8
type OxtelMixerInput uint8
type OxtelVideoSource uint8
type OxtelARC uint8
type OxtelAudioSource uint8
type OxtelAudioProgram uint8
type OxtelJungerPreset uint8
type OxtelAudioMixMode uint8
type OxtelAudioOutput uint8
type OxtelLatencySource uint8
type OxtelVideoStandard uint8
type OxtelFieldRate uint8
type OxtelColorSpace uint8
type OxtelKantarOutput uint8
type OxtelExternalIOType uint8
type OxtelExternalIODirection uint8
type OxtelExternalIOId uint8

const (
	OXTEL_LAYER_0 OxtelLayer = 0
	OXTEL_LAYER_1 OxtelLayer = 1
	OXTEL_LAYER_2 OxtelLayer = 2
	OXTEL_LAYER_3 OxtelLayer = 3
	OXTEL_LAYER_4 OxtelLayer = 4
	OXTEL_LAYER_5 OxtelLayer = 5
	OXTEL_LAYER_6 OxtelLayer = 6
	OXTEL_LAYER_7 OxtelLayer = 7

	OXTEL_DIR_DOWN   OxtelDirection = 0x0
	OXTEL_DIR_UP     OxtelDirection = 0x1
	OXTEL_DIR_TOGGLE OxtelDirection = 0x2

	OXTEL_KEYER_TALLY_DOWN          OxtelKeyerPositionTally = 0x0
	OXTEL_KEYER_TALLY_UP            OxtelKeyerPositionTally = 0x1
	OXTEL_KEYER_TALLY_IN_TRANSITION OxtelKeyerPositionTally = 0x2

	OXTEL_FILE_DOES_NOT_EXIST OxtelEnquireFileInfoResponse = 0x0
	OXTEL_FILE_EXISTS         OxtelEnquireFileInfoResponse = 0x1

	OXTEL_MEDIA_DELETED  OxtelMediaTallies = 0x0
	OXTEL_MEDIA_ADDED    OxtelMediaTallies = 0x1
	OXTEL_MEDIA_MODIFIED OxtelMediaTallies = 0x2

	OXTEL_PLAY_STATE_TALLY_STOPPED OxtelPlayStateTally = 0x0
	OXTEL_PLAY_STATE_TALLY_PLAYING OxtelPlayStateTally = 0x1

	OXTEL_UPDATE_TEXT_FIELD_RENDER OxtelUpdateTextFieldFlag = 0x1
	OXTEL_UPDATE_TEXT_FIELD_APPEND OxtelUpdateTextFieldFlag = 0x2

	OXTEL_TRANSITION_TYPE_V_FADE OxtelTransitionType = 0x01
	OXTEL_TRANSITION_TYPE_X_FADE OxtelTransitionType = 0x03
	OXTEL_TRANSITION_TYPE_CUT    OxtelTransitionType = 0x05

	OXTEL_MIXER_A          OxtelMixerInput = 0x0
	OXTEL_MIXER_B          OxtelMixerInput = 0x1
	OXTEL_MIXER_IN_BETWEEN OxtelMixerInput = 0x2

	OXTEL_VIDEO_SOURCE_PLAYER_A OxtelVideoSource = 0x0
	OXTEL_VIDEO_SOURCE_EXT_IN_1 OxtelVideoSource = 0x1
	OXTEL_VIDEO_SOURCE_PLAYER_B OxtelVideoSource = 0x6
	OXTEL_VIDEO_SOURCE_EXT_IN_2 OxtelVideoSource = 0x7
	OXTEL_VIDEO_SOURCE_EXT_IN_3 OxtelVideoSource = 0x8
	OXTEL_VIDEO_SOURCE_EXT_IN_4 OxtelVideoSource = 0x9
	OXTEL_VIDEO_SOURCE_EXT_IN_5 OxtelVideoSource = 0xA
	OXTEL_VIDEO_SOURCE_EXT_IN_6 OxtelVideoSource = 0xB
	OXTEL_VIDEO_SOURCE_COLOR    OxtelVideoSource = 0xF

	OXTEL_ARC_DEFAULT    OxtelARC = 0x0
	OXTEL_ARC_ANAMORPHIC OxtelARC = 0x1
	OXTEL_ARC_14_9_CROP  OxtelARC = 0x2
	OXTEL_ARC_FULL       OxtelARC = 0x3
	OXTEL_ARC_LETTER     OxtelARC = 0x4
	OXTEL_ARC_PILLAR     OxtelARC = 0x5

	OXTEL_AUDIO_SOURCE_PLAYER_A  OxtelAudioSource = 0x0
	OXTEL_AUDIO_SOURCE_EXT_IN_1  OxtelAudioSource = 0x1
	OXTEL_AUDIO_SOURCE_PLAYER_B  OxtelAudioSource = 0x6
	OXTEL_AUDIO_SOURCE_EXT_IN_2  OxtelAudioSource = 0x7
	OXTEL_AUDIO_SOURCE_EXT_IN_3  OxtelAudioSource = 0x8
	OXTEL_AUDIO_SOURCE_EXT_IN_4  OxtelAudioSource = 0x9
	OXTEL_AUDIO_SOURCE_EXT_IN_5  OxtelAudioSource = 0xA
	OXTEL_AUDIO_SOURCE_EXT_IN_6  OxtelAudioSource = 0xB
	OXTEL_AUDIO_SOURCE_GFX_AUDIO OxtelAudioSource = 0xE

	OXTEL_AUDIO_PROGRAM_INVALID   OxtelAudioProgram = 0x0
	OXTEL_AUDIO_PROGRAM_MONO      OxtelAudioProgram = 0x1
	OXTEL_AUDIO_PROGRAM_STEREO    OxtelAudioProgram = 0x2
	OXTEL_AUDIO_PROGRAM_DOLBY_5_1 OxtelAudioProgram = 0x3
	OXTEL_AUDIO_PROGRAM_DOLBY_7_1 OxtelAudioProgram = 0x4

	// These are decimal
	OXTEL_JUNGER_PRESET_OFF          OxtelJungerPreset = 0
	OXTEL_JUNGER_PRESET_DEFAULT      OxtelJungerPreset = 1
	OXTEL_JUNGER_PRESET_UNIVERSAL    OxtelJungerPreset = 2
	OXTEL_JUNGER_PRESET_MODERATE     OxtelJungerPreset = 3
	OXTEL_JUNGER_PRESET_MOVIE        OxtelJungerPreset = 4
	OXTEL_JUNGER_PRESET_LIMIT        OxtelJungerPreset = 5
	OXTEL_JUNGER_PRESET_NEWS         OxtelJungerPreset = 6
	OXTEL_JUNGER_PRESET_INTERSTITIAL OxtelJungerPreset = 7

	OXTEL_AUDIO_MIX_MODE_X_FADE OxtelAudioMixMode = 0x0
	OXTEL_AUDIO_MIX_MODE_V_FADE OxtelAudioMixMode = 0x1

	OXTEL_AUDIO_OUTPUT_PLAYER_A     OxtelAudioOutput = 0x00
	OXTEL_AUDIO_OUTPUT_EXT_IN_1     OxtelAudioOutput = 0x01
	OXTEL_AUDIO_OUTPUT_PLAYER_B     OxtelAudioOutput = 0x06
	OXTEL_AUDIO_OUTPUT_EXT_IN_2     OxtelAudioOutput = 0x07
	OXTEL_AUDIO_OUTPUT_EXT_IN_3     OxtelAudioOutput = 0x08
	OXTEL_AUDIO_OUTPUT_EXT_IN_4     OxtelAudioOutput = 0x09
	OXTEL_AUDIO_OUTPUT_EXT_IN_5     OxtelAudioOutput = 0xA
	OXTEL_AUDIO_OUTPUT_EXT_IN_6     OxtelAudioOutput = 0xB
	OXTEL_AUDIO_OUTPUT_VOICE_OVER   OxtelAudioOutput = 0xE
	OXTEL_AUDIO_OUTPUT_MIXER_OUTPUT OxtelAudioOutput = 0x10

	OXTEL_LATENCY_SOURCE_MCS           OxtelLatencySource = 0
	OXTEL_LATENCY_SOURCE_OXTEL         OxtelLatencySource = 1
	OXTEL_LATENCY_SOURCE_EXT_KEY_FILL  OxtelLatencySource = 3
	OXTEL_LATENCY_SOURCE_2022_6_INPUT  OxtelLatencySource = 4
	OXTEL_LATENCY_SOURCE_2022_6_OUTPUT OxtelLatencySource = 5

	OXTEL_VIDEO_STANDARD_PAL        OxtelVideoStandard = 0
	OXTEL_VIDEO_STANDARD_NTSC       OxtelVideoStandard = 1
	OXTEL_VIDEO_STANDARD_1080I_5994 OxtelVideoStandard = 2
	OXTEL_VIDEO_STANDARD_1080I_50   OxtelVideoStandard = 3
	OXTEL_VIDEO_STANDARD_720P_5994  OxtelVideoStandard = 4
	OXTEL_VIDEO_STANDARD_720P_50    OxtelVideoStandard = 5
	OXTEL_VIDEO_STANDARD_1080P_5994 OxtelVideoStandard = 6
	OXTEL_VIDEO_STANDARD_1080P_50   OxtelVideoStandard = 7
	OXTEL_VIDEO_STANDARD_2160P_5994 OxtelVideoStandard = 8
	OXTEL_VIDEO_STANDARD_2160P_50   OxtelVideoStandard = 9

	OXTEL_FIELD_RATE_50   OxtelFieldRate = 2
	OXTEL_FIELD_RATE_5994 OxtelFieldRate = 3
	OXTEL_FIELD_RATE_60   OxtelFieldRate = 4

	OXTEL_COLOR_SPACE_SYSTEM_MANAGER = 0x0
	OXTEL_COLOR_SPACE_BT601          = 0x1
	OXTEL_COLOR_SPACE_BT709          = 0x2
	TOXTEL_COLOR_SPACE_BT2020        = 0x3
	OXTEL_COLOR_SPACE_BT2100_HDR_PQ  = 0x4
	OXTEL_COLOR_SPACE_BT2100_HDR_HLG = 0x5
	OXTEL_COLOR_SPACE_VPID           = 0x6

	OXTEL_KANTAR_OUTPUT_PRIMARY   = 0x0
	OXTEL_KANTAR_OUTPUT_SECONDARY = 0x1
	OXTEL_KANTAR_OUTPUT_BOTH      = 0x3

	OXTEL_EXT_IO_TYPE_SDI           = 0x0
	OXTEL_EXT_IO_TYPE_2022_6        = 0x1
	OXTEL_EXT_IO_TYPE_2110          = 0x3
	OXTEL_EXT_IO_TYPE_2022_6_2022_7 = 0x4

	OXTEL_EXT_IO_DIRECTION_IN  = 0x0
	OXTEL_EXT_IO_DIRECTION_OUT = 0x1

	OXTEL_EXT_IO_INPUT_EXT_IN_1 = 0x1
	OXTEL_EXT_IO_INPUT_EXT_IN_2 = 0x7
	OXTEL_EXT_IO_INPUT_EXT_IN_3 = 0x8
	OXTEL_EXT_IO_INPUT_EXT_IN_4 = 0x9
	OXTEL_EXT_IO_INPUT_EXT_IN_5 = 0xA
	OXTEL_EXT_IO_INPUT_EXT_IN_6 = 0xB
	OXTEL_EXT_IO_INPUT_AES_IN   = 0xE

	OXTEL_EXT_IO_OUTPUT_PRIMARY         = 0x0
	OXTEL_EXT_IO_OUTPUT_SECONDARY       = 0x1
	OXTEL_EXT_IO_OUTPUT_CLEAN_PRIMARY   = 0x2
	OXTEL_EXT_IO_OUTPUT_CLEAN_SECONDARY = 0x3
)
