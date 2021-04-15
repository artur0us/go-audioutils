package audioutils

import (
	"errors"
	"strings"
)

// SrcAudioToHLSRequest : ...
type SrcAudioToHLSRequest struct {
	SrcLocationType           int    `json:"src_location_type"`
	SrcLocation               string `json:"src_location"`
	DestDirPath               string `json:"dest_dir_path"`
	DeleteDestDirPathIfExists bool   `json:"delete_dest_dir_path_if_exists"`

	HLSM3U8FileName               *string `json:"hls_m3u8_file_name"`
	HLSSegmentFilePrefix          *string `json:"hls_segment_file_prefix"`
	ThreadsCount                  *int    `json:"threads_count"`
	HLSSegmentSeconds             *int    `json:"hls_segment_seconds"`
	HLSAudioBitrate               *int    `json:"hls_audio_bitrate"`
	HLSAudioCodec                 *string `json:"hls_audio_codec"`
	HLSAppendingSegmentFilePrefix *string `json:"hls_appending_segment_file_prefix"`
}

// validate : ...
func (_s *SrcAudioToHLSRequest) validate() (int, error) {
	if _s == nil {
		return SrcAudioToHLSFailInvalidInputData, errors.New("input is nil")
	}

	// SrcLocationType
	if _s.SrcLocationType != AudioFileSrcLocationTypeLocal && _s.SrcLocationType != AudioFileSrcLocationTypeURL {
		return SrcAudioToHLSFailInvalidSrcLocType, errors.New("unknown source location type")
	}

	// SrcLocation
	if strings.ReplaceAll(_s.SrcLocation, " ", "") == "" {
		return SrcAudioToHLSFailInvalidInputData, errors.New("source location is empty")
	}

	// SrcLocation : files existence
	if _s.SrcLocationType == AudioFileSrcLocationTypeLocal {
		if !isFileExists(AudioFileSrcLocationTypeLocal, _s.SrcLocation) {
			return SrcAudioToHLSFailFileNotFound, errors.New("specified file is not found")
		}
	} else if _s.SrcLocationType == AudioFileSrcLocationTypeURL {
		if !isFileExists(AudioFileSrcLocationTypeURL, _s.SrcLocation) {
			return SrcAudioToHLSFailFileNotFound, errors.New("specified file is not found")
		}
	}

	// DestDirPath
	if strings.ReplaceAll(_s.DestDirPath, " ", "") == "" {
		return SrcAudioToHLSFailInvalidInputData, errors.New("destination directory path is empty")
	}

	return 0, nil
}

// AudioFileDurationRequest : ...
type AudioFileDurationRequest struct {
	SrcLocationType int    `json:"src_location_type"`
	SrcLocation     string `json:"src_location"`
}

// AudioFileBasicInfoRequest : ...
type AudioFileBasicInfoRequest struct {
	SrcLocationType int    `json:"src_location_type"`
	SrcLocation     string `json:"src_location"`
}
