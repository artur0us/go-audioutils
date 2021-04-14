package audioutils

import "fmt"

// GetAudioFileDuration : ...
func (_a *AudioUtils) GetAudioFileDuration(data AudioFileDurationRequest) AudioFileDurationResult {
	var result AudioFileDurationResult = AudioFileDurationResult{
		IsSuccess: false,
		FailCode:  AudioFileDurationGetFailUnknown,
		FailMsg:   "unknown error",
		Duration:  nil,
	}

	// Is source file location type valid?
	if data.SrcLocationType != AudioFileSrcLocationTypeLocal && data.SrcLocationType != AudioFileSrcLocationTypeURL {
		result.FailCode = AudioFileDurationGetFailInvalidSrcLocType
		return result
	}

	if data.SrcLocationType == AudioFileSrcLocationTypeLocal {
		if !_a.isFileExists(AudioFileSrcLocationTypeLocal, data.SrcLocation) {
			result.FailCode = AudioFileDurationGetFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		ffprobeAnswer, err := _a.ffprobeGetAudioFileInfo(data.SrcLocation)
		if err != nil {
			result.FailCode = AudioFileDurationGetFailUnknown
			result.FailMsg = err.Error()
			return result
		}
		if ffprobeAnswer == nil {
			result.FailCode = AudioFileDurationGetFailUnknown
			result.FailMsg = "ffprobe answer is nil"
			return result
		}

		if ffprobeAnswer.Error != nil {
			result.FailCode = AudioFileDurationGetFailFFProbeReturnedErr
			result.FailMsg = fmt.Sprintf("ffprobe returned error: (%v, %v)", ffprobeAnswer.Error.Code, ffprobeAnswer.Error.String)
			return result
		}
		if ffprobeAnswer.Format == nil {
			result.FailCode = AudioFileDurationGetFailFFProbeFormatRespEmpty
			result.FailMsg = "ffprobeAnswer.Format is nil"
			return result
		}

		result.Duration = new(float64)
		*result.Duration = ffprobeAnswer.Format.DurationSeconds
	} else if data.SrcLocationType == AudioFileSrcLocationTypeURL {
		if !_a.isFileExists(AudioFileSrcLocationTypeURL, data.SrcLocation) {
			result.FailCode = AudioFileDurationGetFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		ffprobeAnswer, err := _a.ffprobeGetAudioFileInfo(data.SrcLocation)
		if err != nil {
			result.FailCode = AudioFileDurationGetFailUnknown
			result.FailMsg = err.Error()
			return result
		}
		if ffprobeAnswer == nil {
			result.FailCode = AudioFileDurationGetFailUnknown
			result.FailMsg = "ffprobe answer is nil"
			return result
		}

		if ffprobeAnswer.Error != nil {
			result.FailCode = AudioFileDurationGetFailFFProbeReturnedErr
			result.FailMsg = fmt.Sprintf("ffprobe returned error: (%v, %v)", ffprobeAnswer.Error.Code, ffprobeAnswer.Error.String)
			return result
		}
		if ffprobeAnswer.Format == nil {
			result.FailCode = AudioFileDurationGetFailFFProbeFormatRespEmpty
			result.FailMsg = "ffprobeAnswer.Format is nil"
			return result
		}

		result.Duration = new(float64)
		*result.Duration = ffprobeAnswer.Format.DurationSeconds
	} else {
		result.FailCode = AudioFileDurationGetFailInvalidSrcLocType
		result.FailMsg = "invalid source location type"
		return result
	}

	result.IsSuccess = true
	result.FailCode = 0
	result.FailMsg = ""
	return result
}
