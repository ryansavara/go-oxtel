package oxtel

import (
	"testing"
)

var oxtel = NewOxtel("10.43.192.11", 9100)

const templateName = "1080i60-EAS-Warning.swf"
const layer = OXTEL_LAYER_0

func TestConnection(t *testing.T) {
	err := oxtel.Connect()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannelMasks(t *testing.T) {
	mask := MakeEightChannelMask()
	str := buildChannelMask(mask)
	if str != "00FF" {
		t.Fatalf("Channel mask does not match: %v", str)
	}
}

// func TestEnquireMediaTallies(t *testing.T) {
// 	val, err := oxtel.EnquireMediaTallies()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Images {
// 		t.Fatal("Media Tallies should be off")
// 	}
// }

// func TestMediaTallies(t *testing.T) {
// 	s := MediaTallies{
// 		Images: true,
// 	}
// 	err := oxtel.EnableMediaTallies(s)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquirePlayStateTally(t *testing.T) {
// 	val, err := oxtel.EnquirePlayStateTally()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val {
// 		t.Fatal("Play State Tally should be off")
// 	}
// }

// func TestEnablePlayStateTally(t *testing.T) {
// 	err := oxtel.EnablePlayStateTally(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquireFileInfo(t *testing.T) {
// 	val, err := oxtel.EnquireFileInfo(templateName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !val.Exists || val.Filename != templateName {
// 		t.Fatalf("Unexpected file info response: %v", val)
// 	}
// }

// func TestEnquireExtendedFileInfo(t *testing.T) {
// 	val, err := oxtel.EnquireExtendedFileInformation(templateName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !val.Exists || val.Filename != templateName {
// 		t.Fatalf("Unexpected file info response: %v", val)
// 	}
// }

// func TestValidateTemplate(t *testing.T) {
// 	val, err := oxtel.ValidateTemplate(templateName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Filename != templateName || !val.FileExists {
// 		t.Fatal("validate template failed")
// 	}
// }

// func TestPreloadImage(t *testing.T) {
// 	err := oxtel.PreloadImage(layer, templateName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquirePreloadImage(t *testing.T) {
// 	fmt.Println("Sleeping to ensure Oxtel processes the preload command")
// 	time.Sleep(3 * time.Second)

// 	val, err := oxtel.EnquirePreloadImage(layer)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Layer != 0 || val.Filename != templateName {
// 		t.Fatalf("Incorrect template preloaded: %v", val)
// 	}
// }

// func TestLoadImage(t *testing.T) {
// 	err := oxtel.LoadImage(0, "1080i60-EAS-Warning.swf")

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquireLoadImage(t *testing.T) {
// 	fmt.Println("Sleeping to ensure Oxtel processes the load command")
// 	time.Sleep(3 * time.Second)

// 	val, err := oxtel.EnquireLoadImage(layer)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Layer != 0 || val.Filename != templateName {
// 		t.Fatalf("Incorrect template loaded: %v", val)
// 	}
// }

// func TestSetImagePosition(t *testing.T) {
// 	err := oxtel.SetImagePosition(layer, 10, 10)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquireImagePosition(t *testing.T) {
// 	fmt.Println("Sleeping to ensure Oxtel processes the set position command")
// 	time.Sleep(3 * time.Second)

// 	val, err := oxtel.EnquireImagePosition(layer)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Layer != 0 || val.XOffset != 10 || val.YOffset != 10 {
// 		t.Fatalf("Invalid position image: %v", val)
// 	}
// }

// func TestEraseStore(t *testing.T) {
// 	err := oxtel.EraseStore(layer)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestSelectMixerInput(t *testing.T) {
// 	arc := OXTEL_ARC_DEFAULT
// 	err := oxtel.SelectMixerInput(0, OXTEL_VIDEO_SOURCE_EXT_IN_3, &arc)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// }

// func TestEnquireMixerInput(t *testing.T) {
// 	time.Sleep(2 * time.Second)
// 	val, err := oxtel.EnquireMixerInput(0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Input != 0 || val.Source != OXTEL_VIDEO_SOURCE_EXT_IN_3 {
// 		t.Fatalf("Invalid mixer input %v", val)
// 	}
// }

// func TestEnquireMixMode(t *testing.T) {
// 	val, err := oxtel.EnquireMixMode()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.WipeSoftness != 0 {
// 		t.Fatalf("Invalid mix mode response %v", val)
// 	}
// }

// func TestSetColorGenerator(t *testing.T) {
// 	err := oxtel.SetColorGeneratorColor(0, 255, 0, 0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquireColorGeneratorColor(t *testing.T) {
// 	time.Sleep(2 * time.Second)
// 	val, err := oxtel.EnquireColorGeneratorColor(0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Unit != 0 || val.Red != 255 || val.Green != 0 || val.Blue != 0 {
// 		t.Fatalf("Invalid color generate response %v", val)
// 	}
// }

// func TestEnquireVideoTallies(t *testing.T) {
// 	val, err := oxtel.EnquireVideoTallies()

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val {
// 		t.Fatal("video tally should be false")
// 	}
// }

// func TestEnquireAudioProfile(t *testing.T) {
// 	val, err := oxtel.EnquireAudioProfile(0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Profile != 0 || val.Source != 0 {
// 		t.Fatalf("invalid audio profile %v", val)
// 	}
// }

// func TestAudioFollowingVideo(t *testing.T) {
// 	err := oxtel.SetAudioABFollowVideoAB(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestGetLoudness(t *testing.T) {
// 	val, err := oxtel.GetAudioLoudness(0, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Channel1 {
// 		t.Fatal("audio loudness program failed")
// 	}
// }

// func TestGetLoudnessLicenseStatus(t *testing.T) {
// 	val, err := oxtel.GetLoudnessLicenseStatus()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_ = val
// }

// func TestEnquireAudioFollowingVideo(t *testing.T) {
// 	time.Sleep(2 * time.Second)
// 	val, err := oxtel.EnquireAudioABFollowVideoAB()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !val.Enabled {
// 		t.Fatal("audio should be following video")
// 	}
// }

// func TestEnquireAudioGain(t *testing.T) {
// 	mask := MakeEightChannelMask()
// 	val, err := oxtel.EnquireAudioGain(OXTEL_AUDIO_OUTPUT_EXT_IN_3, mask)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Source != OXTEL_AUDIO_SOURCE_EXT_IN_3 || val.Gain != 0 {
// 		t.Fatal("audio gain failed")
// 	}
// }

// func TestEnquireLatency(t *testing.T) {
// 	val, err := oxtel.EnquireLatency(OXTEL_LATENCY_SOURCE_OXTEL)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Source != OXTEL_LATENCY_SOURCE_OXTEL {
// 		t.Fatal("latency response failed")
// 	}
// }

// func TestEnquireNumberOfGraphicLayers(t *testing.T) {
// 	val, err := oxtel.EnquireNumberOfGraphicLayers()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val != 8 {
// 		t.Fatal("number of graphic layers failed")
// 	}
// }

// func TestEnquireSystemStatus(t *testing.T) {
// 	val, err := oxtel.EnquireSystemStatus()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.VideoStandard != OXTEL_VIDEO_STANDARD_1080I_5994 {
// 		t.Fatal("system status failed")
// 	}

// }

// func TestFadeKeyer(t *testing.T) {
// 	err := oxtel.FadeKeyer(OXTEL_LAYER_0, OXTEL_DIR_UP, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEnquireVideoLayerStatus(t *testing.T) {
// 	time.Sleep(2 * time.Second)
// 	val, err := oxtel.EnquireVideoLayerStatus(OXTEL_LAYER_0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	oxtel.FadeKeyer(OXTEL_LAYER_0, OXTEL_DIR_DOWN, nil)

// 	if val.LayerFaderAngle != 512 {
// 		t.Fatal("video layer status failed")
// 	}
// }

// func TestEnquireCommandAvailability(t *testing.T) {
// 	val, err := oxtel.EnquireCommandAvailability('X', '3')
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.CommandByte1 != 'X' || val.CommandByte2 != '3' || !val.Supported {
// 		t.Fatal("command availability failed")
// 	}
// }

// func TestEnquireSlaveLayerStatus(t *testing.T) {
// 	val, err := oxtel.EnquireSlaveLayerStatus()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val.Layer7State {
// 		t.Fatal("layer status failed")
// 	}
// }

// func TestEnquireFullVersionNumber(t *testing.T) {
// 	val, err := oxtel.EnquireFullVersionNumber()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// if val.Major == 0 || len(val.AsString) == 0 {
// 		t.Fatal("full version failed")
// 	}
// }

// func TestEnquireProductName(t *testing.T) {
// 	val, err := oxtel.EnquireProductName()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(val) == 0 {
// 		t.Fatal("product name failed")
// 	}
// }

// func TestEnquireMediaPortName(t *testing.T) {
// 	val, err := oxtel.EnquireMediaPortName()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if val[:3] != "Tap" {
// 		t.Fatal("media port name failed")
// 	}
// }

func TestTemperature(t *testing.T) {
	val, err := oxtel.EnquireTemperature()
	if err != nil {
		t.Fatal(err)
	}

	if val != 0.0 {
		t.Fatalf("temperature error")
	}
}

// Last Function to call
func TestDisconnect(t *testing.T) {
	err := oxtel.Disconnect()

	if err != nil {
		t.Fatal(err)
	}
}
