package audioutils

// GetAudioFileWaveformData : ...
func (_a *AudioUtils) GetAudioFileWaveformData(data AudioFileWaveformDataRequest) AudioFileWaveformDataResult {
	var result AudioFileWaveformDataResult = AudioFileWaveformDataResult{
		IsSuccess: false,
		FailCode:  AudioFileWaveformDataGetFailUnknown,
		FailMsg:   "unknown error",

		Points: nil,
	}

	// Input data validation
	_a.InfoLogger.Printf("validating data: %v\n", data)
	validationCode, err := data.validate()
	if err != nil {
		_a.ErrLogger.Printf("data validation failed: %v\n", err)
		result.FailCode = validationCode
		result.FailMsg = err.Error()
		return result
	}

	var audiowaveformAnswer *audiowaveformJSON = nil

	if data.SrcLocationType == AudioFileSrcLocationTypeLocal {
		if !isFileExists(AudioFileSrcLocationTypeLocal, data.SrcLocation) {
			_a.ErrLogger.Printf("specified file (local) is not found: %v\n", data.SrcLocation)
			result.FailCode = AudioFileWaveformDataGetFailFileNotFound
			result.FailMsg = "specified file (local) is not found"
			return result
		}

		audiowaveformAnswer, err = _a.audiowaveformGetWaveformData(audiowaveformWaveformDataInput{
			SrcLocation: data.SrcLocation,

			PointsPerSecond: data.PointsPerSecond,
			DataBitsCount:   data.DataBitsCount,
		})
	} else if data.SrcLocationType == AudioFileSrcLocationTypeURL {
		if !isFileExists(AudioFileSrcLocationTypeURL, data.SrcLocation) {
			_a.ErrLogger.Printf("specified file (URL) is not found: %v\n", data.SrcLocation)
			result.FailCode = AudioFileWaveformDataGetFailFileNotFound
			result.FailMsg = "specified file (URL) is not found"
			return result
		}

		audiowaveformAnswer, err = _a.audiowaveformGetWaveformData(audiowaveformWaveformDataInput{
			SrcLocation: data.SrcLocation,

			PointsPerSecond: data.PointsPerSecond,
			DataBitsCount:   data.DataBitsCount,
		})
	} else {
		_a.ErrLogger.Printf("unknown source location type: %v\n", data.SrcLocationType)
		result.FailCode = AudioFileWaveformDataGetFailInvalidSrcLocType
		result.FailMsg = "invalid source location type"
		return result
	}

	if err != nil {
		_a.ErrLogger.Printf("failed to get audio file info with audiowaveform: %v\n", err)
		result.FailCode = AudioFileWaveformDataGetFailUnknown
		result.FailMsg = err.Error()
		return result
	}
	if audiowaveformAnswer == nil {
		_a.ErrLogger.Println("audiowaveformAnswer is nil")
		result.FailCode = AudioFileWaveformDataGetFailUnknown
		result.FailMsg = "audiowaveformAnswer is nil"
		return result
	}

	result.Points = &[]int{}
	*result.Points = audiowaveformAnswer.Data

	_a.InfoLogger.Printf("audio file waveform data obtaining finished successfully: (file: %v)\n", data.SrcLocation)
	result.IsSuccess = true
	result.FailCode = 0
	result.FailMsg = ""
	return result
}
