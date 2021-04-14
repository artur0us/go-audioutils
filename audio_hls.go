package audioutils

import "path/filepath"

// ConvertSrcAudioFileToHLS : ...
func (_a *AudioUtils) ConvertSrcAudioFileToHLS(data SrcAudioToHLSRequest) SrcAudioToHLSResult {
	var result SrcAudioToHLSResult = SrcAudioToHLSResult{
		IsSuccess:      false,
		FailCode:       SrcAudioToHLSFailUnknown,
		FailMsg:        "unknown error",
		ResultDestPath: nil,
	}

	// Is source file location type valid?
	if data.SrcLocationType != AudioFileSrcLocationTypeLocal && data.SrcLocationType != AudioFileSrcLocationTypeURL {
		result.FailCode = SrcAudioToHLSFailInvalidSrcLocType
		return result
	}

	// Parse absolute path of destination directory
	var destDirPath string = "output"
	destDirPath, err := filepath.Abs(destDirPath)
	if err != nil {
		result.FailCode = SrcAudioToHLSFailParseAbsDestPath
		result.FailMsg = "failed to get absolute path of destination directory"
		return result
	}

	if data.SrcLocationType == AudioFileSrcLocationTypeLocal {
		if !_a.isFileExists(AudioFileSrcLocationTypeLocal, data.SrcLocation) {
			result.FailCode = SrcAudioToHLSFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		convertResult, err := _a.ffmpegSrcAudioFileToHLS(ffmpegAudioFileToHLSInput{
			SrcLocation: data.SrcLocation,
			DestDirPath: destDirPath,

			DestM3U8FileName:      nil,
			DestSegmentFilePrefix: nil,
			LogLevel:              nil,
			ThreadsCount:          nil,
			SegmentSeconds:        nil,
			Bitrate:               nil,
			Codec:                 nil,
		})
		if err != nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = err.Error()
			return result
		}
		if convertResult == nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "convert result is nil"
			return result
		}
		if convertResult.ResultDestDirPath == nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "result dest dir path is nil"
			return result
		}

		// TODO: some checks

		result.ResultDestPath = new(string)
		*result.ResultDestPath = *convertResult.ResultDestDirPath
	} else if data.SrcLocationType == AudioFileSrcLocationTypeURL {
		if !_a.isFileExists(AudioFileSrcLocationTypeURL, data.SrcLocation) {
			result.FailCode = SrcAudioToHLSFailFileNotFound
			result.FailMsg = "specified file is not found"
			return result
		}

		convertResult, err := _a.ffmpegSrcAudioFileToHLS(ffmpegAudioFileToHLSInput{
			SrcLocation: data.SrcLocation,
			DestDirPath: destDirPath,

			DestM3U8FileName:      nil,
			DestSegmentFilePrefix: nil,
			LogLevel:              nil,
			ThreadsCount:          nil,
			SegmentSeconds:        nil,
			Bitrate:               nil,
			Codec:                 nil,
		})
		if err != nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = err.Error()
			return result
		}
		if convertResult == nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "convert result is nil"
			return result
		}
		if convertResult.ResultDestDirPath == nil {
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "result dest dir path is nil"
			return result
		}

		// TODO: some checks

		result.ResultDestPath = new(string)
		*result.ResultDestPath = *convertResult.ResultDestDirPath
	} else {
		result.FailCode = SrcAudioToHLSFailInvalidSrcLocType
		result.FailMsg = "invalid source location type"
		return result
	}

	result.IsSuccess = true
	result.FailCode = 0
	result.FailMsg = ""
	return result
}
