package audioutils

import (
	"os"
	"path/filepath"
)

// ConvertSrcAudioFileToHLS : ...
func (_a *AudioUtils) ConvertSrcAudioFileToHLS(data SrcAudioToHLSRequest) SrcAudioToHLSResult {
	var result SrcAudioToHLSResult = SrcAudioToHLSResult{
		IsSuccess:      false,
		FailCode:       SrcAudioToHLSFailUnknown,
		FailMsg:        "unknown error",
		ResultDestPath: nil,
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

	// Parse absolute path of destination directory
	_a.InfoLogger.Printf("parsing absolute path of destination directory: %v\n", data.DestDirPath)
	destDirPath, err := filepath.Abs(data.DestDirPath)
	if err != nil {
		_a.ErrLogger.Printf("destination directory absolute path parse failed: %v\n", err)
		result.FailCode = SrcAudioToHLSFailParseAbsDestPath
		result.FailMsg = "failed to get absolute path of destination directory"
		return result
	}
	_a.InfoLogger.Println("checking if destination directory is already exists...")
	if _, err := os.Stat(destDirPath); !os.IsNotExist(err) {
		// destDirPath exists
		_a.InfoLogger.Println("destination directory is exists")
		if data.DeleteDestDirPathIfExists {
			_a.InfoLogger.Println("*DeleteDestDirPathIfExists* flag is true, deleting existing destination directory...")
			os.RemoveAll(destDirPath)
			_a.InfoLogger.Println("creating new destination directory...")
			err = os.MkdirAll(destDirPath, os.ModePerm)
			if err != nil {
				_a.ErrLogger.Printf("failed to create new destination directory: %v\n", err)
				result.FailCode = SrcAudioToHLSFailUnknown
				result.FailMsg = err.Error()
				return result
			}
		}
	} else {
		// destDirPath does not exists
		_a.InfoLogger.Println("destination directory is not found, creating new...")
		err = os.MkdirAll(destDirPath, os.ModePerm)
		if err != nil {
			_a.ErrLogger.Printf("failed to create destination directory: %v\n", err)
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = err.Error()
			return result
		}
	}

	_a.InfoLogger.Printf("starting processing HLS from source audio file: %v\n", data.SrcLocation)
	_a.InfoLogger.Printf("specified source location type: %v\n", data.SrcLocationType)
	if data.SrcLocationType == AudioFileSrcLocationTypeLocal {
		_a.InfoLogger.Println("calling ffmpeg...")
		convertResult, err := _a.ffmpegSrcAudioFileToHLS(ffmpegAudioFileToHLSInput{
			SrcLocation: data.SrcLocation,
			DestDirPath: destDirPath,

			DestM3U8FileName:           data.HLSM3U8FileName,
			DestSegmentFilePrefix:      data.HLSSegmentFilePrefix,
			LogLevel:                   nil,
			ThreadsCount:               data.ThreadsCount,
			SegmentSeconds:             data.HLSSegmentSeconds,
			Bitrate:                    data.HLSAudioBitrate,
			Codec:                      data.HLSAudioCodec,
			AppendingSegmentFilePrefix: data.HLSAppendingSegmentFilePrefix,
		})
		if err != nil {
			_a.ErrLogger.Printf("failed to process HLS with ffmpeg: %v\n", err)
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = err.Error()
			return result
		}
		if convertResult == nil {
			_a.ErrLogger.Println("convertResult is nil")
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "convertResult is nil"
			return result
		}
		if convertResult.ResultDestDirPath == nil {
			_a.ErrLogger.Println("convertResult.ResultDestDirPath is nil")
			result.FailCode = SrcAudioToHLSFailUnknown
			result.FailMsg = "convertResult.ResultDestDirPath is nil"
			return result
		}

		// TODO: some checks

		result.ResultDestPath = new(string)
		*result.ResultDestPath = *convertResult.ResultDestDirPath
	} else if data.SrcLocationType == AudioFileSrcLocationTypeURL {
		convertResult, err := _a.ffmpegSrcAudioFileToHLS(ffmpegAudioFileToHLSInput{
			SrcLocation: data.SrcLocation,
			DestDirPath: destDirPath,

			DestM3U8FileName:           data.HLSM3U8FileName,
			DestSegmentFilePrefix:      data.HLSSegmentFilePrefix,
			LogLevel:                   nil,
			ThreadsCount:               data.ThreadsCount,
			SegmentSeconds:             data.HLSSegmentSeconds,
			Bitrate:                    data.HLSAudioBitrate,
			Codec:                      data.HLSAudioCodec,
			AppendingSegmentFilePrefix: data.HLSAppendingSegmentFilePrefix,
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
		_a.ErrLogger.Printf("unknown source location type: %v\n", data.SrcLocationType)
		result.FailCode = SrcAudioToHLSFailInvalidSrcLocType
		result.FailMsg = "invalid source location type"
		return result
	}

	if result.ResultDestPath == nil {
		_a.ErrLogger.Println("result.ResultDestPath is nil")
		result.FailCode = SrcAudioToHLSFailUnknown
		result.FailMsg = "result.ResultDestPath is nil"
		return result
	}

	_a.InfoLogger.Printf("HLS processing finished successfully: (file: %v) (result dest path: %v)\n", data.SrcLocation, *result.ResultDestPath)
	result.IsSuccess = true
	result.FailCode = 0
	result.FailMsg = ""
	return result
}
