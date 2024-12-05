package oxtel

import (
	"fmt"
	"strconv"
	"strings"
)

// EnquireFileInfo queries information about the existence of the template on the file system. The folder location searched
// is defined by the Effects configuration section in SystemManager.
//
// Response is a FileInfoResponse.
func (o *Oxtel) EnquireFileInfo(fileName string) (FileInfoResponse, error) {
	val, err := o.sendCommandExpectResponse("R3", fileName)
	if err != nil {
		return FileInfoResponse{}, err
	}

	exists, err := strconv.ParseBool(string(val[0]))
	if err != nil {
		return FileInfoResponse{}, err
	}

	return FileInfoResponse{
		Exists:   exists,
		Filename: val[1:],
	}, nil
}

// EnquireFileInfo_AsString returns the command string used to query information about the existence of the template on
// the file system. The folder location searched is defined by the Effects configuration section in SystemManager.
//
// For use with scheduled commands.
func EnquireFileInfo_AsString(fileName string) string {
	return fmt.Sprintf("R3%s", fileName)
}

// QueryFirstFile queries the name of the first file within the specified folder name alias.
//
// In order to enumerate the files on the system, this command should be issued first followed by calls to QuerySubsequentFile.
// If folderName is nil, the template folder specified in the Effects configuration section of SystemManager is used.
//
// The order of filenames returned is "older file first".
// Response is a FileQueryResponse.
func (o *Oxtel) QueryFirstFile(folderName *string) (FileQueryResponse, error) {
	if folderName == nil {
		defaultVal := "$VIDEO"
		folderName = &defaultVal
	}
	val, err := o.sendCommandExpectResponse("R4", *folderName)
	if err != nil {
		return FileQueryResponse{}, err
	}

	eod, err := strconv.ParseBool(string(val[0]))
	if err != nil {
		return FileQueryResponse{}, err
	}

	return FileQueryResponse{
		EndOfDir: eod,
		Filename: val[2:],
	}, nil
}

// QueryFirstFile_AsString returns the command string used to query the name of the first file within the specified folder name alias.
//
// For use with scheduled commands.
func QueryFirstFile_AsString(folderName *string) string {
	if folderName == nil {
		defaultVal := "$VIDEO"
		folderName = &defaultVal
	}
	return fmt.Sprintf("R4%s", *folderName)
}

// QuerySubsequentFile is used in conjunction with QueryFirstFile to enumerate the names of all files within the media folder.
// It should be used following a single QueryFirstFile command using the same folder name alias.
//
// If folderName is nil, the template folder specified in the Effects configuration section of SystemManager is used.
//
// The order of filenames returned is "older file first".
// Once all files have been enumerated, subsequent changes to files can be tracked using Media Tallies (YB).
//
// Response is a FileQueryResponse.
func (o *Oxtel) QuerySubsequentFile(folderName *string) (FileQueryResponse, error) {
	if folderName == nil {
		defaultVal := "$VIDEO"
		folderName = &defaultVal
	}
	val, err := o.sendCommandExpectResponse("R5", *folderName)
	if err != nil {
		return FileQueryResponse{}, err
	}

	eod, err := strconv.ParseBool(string(val[0]))
	if err != nil {
		return FileQueryResponse{}, err
	}

	return FileQueryResponse{
		EndOfDir: eod,
		Filename: val[2:],
	}, nil

}

// QuerySubsequentFile_AsString returns the command string used to enumerate the names of all files within the media folder.
//
// For use with scheduled commands.
func QuerySubsequentFile_AsString(folderName *string) string {
	if folderName == nil {
		defaultVal := "$VIDEO"
		folderName = &defaultVal
	}
	return fmt.Sprintf("R5%s", *folderName)
}

// EnquireExtendedFileInformation queries for information about the specified template.
//
// If no extension is specified, the command is treated as if the request was for the file with one of the supported extensions.
//
// Response is a ExtendedFileInfoResponse. Only "File exists" and "Filename" are valid.
func (o *Oxtel) EnquireExtendedFileInformation(fileName string) (ExtendedFileInfoResponse, error) {
	val, err := o.sendCommandExpectResponse("R6", fileName)

	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	boolVal, err := strconv.Atoi(string(val[0]))
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	exists, err := intToBool(boolVal)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	xpos, err := strconv.ParseInt(val[1:4], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}
	ypos, err := strconv.ParseInt(val[4:7], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	width, err := strconv.ParseInt(val[7:10], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	height, err := strconv.ParseInt(val[10:13], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	clip, err := strconv.ParseInt(val[13:16], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	gain, err := strconv.ParseInt(val[16:19], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	transparency, err := strconv.ParseInt(val[19:22], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	imageType, err := strconv.ParseInt(val[22:24], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	frames, err := strconv.ParseInt(val[24:28], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	animationMode, err := strconv.ParseInt(string(val[28]), 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	loadTime, err := strconv.ParseInt(val[29:31], 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	associatedAudio, err := strconv.ParseInt(string(val[31]), 16, 64)
	if err != nil {
		return ExtendedFileInfoResponse{}, err
	}

	templateName := val[32:]

	return ExtendedFileInfoResponse{
		Exists:          exists,
		XPosition:       xpos,
		YPosition:       ypos,
		Width:           width,
		Height:          height,
		Clip:            clip,
		Gain:            gain,
		Transparency:    transparency,
		ImageType:       imageType,
		Frames:          frames,
		AnimationMode:   animationMode,
		LoadTime:        loadTime,
		AssociatedAudio: associatedAudio,
		Filename:        templateName,
	}, nil
}

// EnquireExtendedFileInformation_AsString returns the command string used to query for information about the specified template.
//
// For use with scheduled commands.
func EnquireExtendedFileInformation_AsString(fileName string) string {
	return fmt.Sprintf("R6%s", fileName)
}

// ValidateTemplate validates the presence/absence of the specified template.
// Assets referenced by the template are not checked.
//
// Response is a ValidateTemplateResponse. MissingAssets is not supported.
func (o *Oxtel) ValidateTemplate(fileName string) (ValidateTemplateResponse, error) {
	val, err := o.sendCommandExpectResponse("RA", fileName)
	if err != nil {
		return ValidateTemplateResponse{}, err
	}

	data := strings.Split(val, "|")
	exists, err := strconv.ParseBool(string(data[1][0]))
	if err != nil {
		return ValidateTemplateResponse{}, err
	}

	missingAssets, err := strconv.ParseInt(string(data[1][1:5]), 16, 16)
	if err != nil {
		return ValidateTemplateResponse{}, err
	}

	return ValidateTemplateResponse{
		Filename:      data[0],
		FileExists:    exists,
		MissingAssets: int16(missingAssets),
	}, err

}

// ValidateTemplate_AsString returns the command string used to validates the presence/absence of the specified template.
//
// For use with scheduled commands.
func ValidateTemplate_AsString(fileName string) string {
	return fmt.Sprintf("RA%s", fileName)
}

// EnableMediaTallies enables or disables media tallies for the connection on which the command was received.
//
// Media tallies are used to track media management as files are added, deleted, or modified on the file system.
func (o *Oxtel) EnableMediaTallies(data MediaTallies) error {

	mask := buildMediaTallies(data)

	return o.sendCommand(fmt.Sprintf("YB%x", mask))
}

// EnableMediaTallies_AsString returns the command string used to enable or disable media tallies for the connection on which the
// command was received.
//
// For use with scheduled commands.
func EnableMediaTallies_AsString(data MediaTallies) string {
	mask := buildMediaTallies(data)

	return fmt.Sprintf("YB%x", mask)
}

// EnquireMediaTallies queries the enable/disable state of the media tallies for the connection.
//
// Response is MediaTallies. Only Images is supported.
func (o *Oxtel) EnquireMediaTallies() (MediaTallies, error) {
	val, err := o.sendCommandExpectResponse("YB", "")
	if err != nil {
		return MediaTallies{}, err
	}

	u1, err := strconv.ParseBool(string(val[0]))
	if err != nil {
		return MediaTallies{}, err
	}

	u2, err := strconv.ParseBool(string(val[1]))
	if err != nil {
		return MediaTallies{}, err
	}

	u3, err := strconv.ParseBool(string(val[2]))
	if err != nil {
		return MediaTallies{}, err
	}

	u4, err := strconv.ParseBool(string(val[3]))
	if err != nil {
		return MediaTallies{}, err
	}

	u5, err := strconv.ParseBool(string(val[4]))
	if err != nil {
		return MediaTallies{}, err
	}

	images, err := strconv.ParseBool(string(val[5]))
	if err != nil {
		return MediaTallies{}, err
	}
	return MediaTallies{
		Unused1: u1,
		Unused2: u2,
		Unused3: u3,
		Unused4: u4,
		Unused5: u5,
		Images:  images,
	}, err
}

// EnquireMediaTallies_AsString returns the command string used to query the enable/disable state of the media tallies for
// the connection.
//
// For use with scheduled commands.
func EnquireMediaTallies_AsString() string {
	return "YB"
}
