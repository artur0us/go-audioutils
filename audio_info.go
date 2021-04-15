package audioutils

import (
	"encoding/json"
	"fmt"
)

// GetAudioFileBasicInfo : ...
func (_a *AudioUtils) GetAudioFileBasicInfo(data AudioFileBasicInfoRequest) AudioFileBasicInfoResult {
	var result AudioFileBasicInfoResult = AudioFileBasicInfoResult{
		IsSuccess: false,
		FailCode:  AudioFileDurationGetFailUnknown,
		FailMsg:   "unknown error",

		Duration: nil,
	}

	// Is source file location type valid?
	_a.InfoLogger.Printf("specified source location type: %v\n", data.SrcLocationType)
	if data.SrcLocationType != AudioFileSrcLocationTypeLocal && data.SrcLocationType != AudioFileSrcLocationTypeURL {
		_a.ErrLogger.Printf("unknown source location type: %v\n", data.SrcLocationType)
		result.FailCode = AudioFileDurationGetFailInvalidSrcLocType
		return result
	}

	var ffprobeAnswer *FFProbeProbeData = nil
	var err error = nil

	_a.InfoLogger.Printf("starting obtaining audio file duration: %v\n", data.SrcLocation)
	if data.SrcLocationType == AudioFileSrcLocationTypeLocal {
		if !isFileExists(AudioFileSrcLocationTypeLocal, data.SrcLocation) {
			_a.ErrLogger.Printf("specified file (local) is not found: %v\n", data.SrcLocation)
			result.FailCode = AudioFileDurationGetFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		_a.InfoLogger.Println("calling ffprobe...")
		ffprobeAnswer, err = _a.ffprobeGetAudioFileInfo(data.SrcLocation)
	} else if data.SrcLocationType == AudioFileSrcLocationTypeURL {
		if !isFileExists(AudioFileSrcLocationTypeURL, data.SrcLocation) {
			_a.ErrLogger.Printf("specified file (URL) is not found: %v\n", data.SrcLocation)
			result.FailCode = AudioFileDurationGetFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		_a.InfoLogger.Println("calling ffprobe...")
		ffprobeAnswer, err = _a.ffprobeGetAudioFileInfo(data.SrcLocation)
	} else {
		_a.ErrLogger.Printf("unknown source location type: %v\n", data.SrcLocationType)
		result.FailCode = AudioFileDurationGetFailInvalidSrcLocType
		result.FailMsg = "invalid source location type"
		return result
	}

	if err != nil {
		_a.ErrLogger.Printf("failed to get audio file info with ffprobe: %v\n", err)
		result.FailCode = AudioFileDurationGetFailUnknown
		result.FailMsg = err.Error()
		return result
	}
	if ffprobeAnswer == nil {
		_a.ErrLogger.Println("ffprobeAnswer is nil")
		result.FailCode = AudioFileDurationGetFailUnknown
		result.FailMsg = "ffprobeAnswer is nil"
		return result
	}
	if ffprobeAnswer.Error != nil {
		_a.ErrLogger.Printf("ffprobe returned error: (%v, %v)\n", ffprobeAnswer.Error.Code, ffprobeAnswer.Error.String)
		result.FailCode = AudioFileDurationGetFailFFProbeReturnedErr
		result.FailMsg = fmt.Sprintf("ffprobe returned error: (%v, %v)", ffprobeAnswer.Error.Code, ffprobeAnswer.Error.String)
		return result
	}
	if ffprobeAnswer.Format == nil {
		_a.ErrLogger.Println("ffprobeAnswer.Format is nil")
		result.FailCode = AudioFileDurationGetFailFFProbeFormatRespEmpty
		result.FailMsg = "ffprobeAnswer.Format is nil"
		return result
	}

	result.Duration = new(float64)
	*result.Duration = ffprobeAnswer.Format.DurationSeconds

	result.Size = new(int64)
	json.Unmarshal([]byte(ffprobeAnswer.Format.Size), &result.Size)

	result.Bitrate = new(int)
	json.Unmarshal([]byte(ffprobeAnswer.Format.BitRate), &result.Bitrate)

	result.FormatName = new(string)
	*result.FormatName = ffprobeAnswer.Format.FormatName

	result.FormatLongName = new(string)
	*result.FormatLongName = ffprobeAnswer.Format.FormatLongName

	result.StreamsCount = new(int)
	*result.StreamsCount = len(ffprobeAnswer.Streams)

	_a.InfoLogger.Printf("audio file basic information obtaining finished successfully: (file: %v)\n", data.SrcLocation)
	result.IsSuccess = true
	result.FailCode = 0
	result.FailMsg = ""
	return result
}
