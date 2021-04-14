package audioutils

import (
	"errors"
	"strings"
)

// ffmpegAudioFileToHLSInput : ...
type ffmpegAudioFileToHLSInput struct {
	// Required
	SrcLocation string `json:"src_location"`
	DestDirPath string `json:"dest_dir_path"`

	// Optional
	DestM3U8FileName      *string `json:"dest_m3u8_file_name"`
	DestSegmentFilePrefix *string `json:"dest_segment_file_prefix"`
	LogLevel              *string `json:"log_level"`
	ThreadsCount          *int    `json:"threads_count"`
	SegmentSeconds        *int    `json:"segment_seconds"`
	Bitrate               *int    `json:"bitrate"`
	Codec                 *string `json:"codec"`
}

// validate : ...
func (_f *ffmpegAudioFileToHLSInput) validate() error {
	if _f == nil {
		return errors.New("input is nil")
	}

	// [ Required fields ]
	// SrcLocation
	if strings.ReplaceAll(_f.SrcLocation, " ", "") == "" {
		return errors.New("source location is empty")
	}
	// DestDirPath
	if strings.ReplaceAll(_f.DestDirPath, " ", "") == "" {
		return errors.New("destination directory path is empty")
	}

	// [ Optional fields ]
	// DestM3U8FileName
	if _f.DestM3U8FileName != nil {
		if strings.ReplaceAll(*_f.DestM3U8FileName, " ", "") == "" {
			return errors.New("destination M3U8 file name is empty")
		}
	}
	// DestSegmentFilePrefix
	if _f.DestSegmentFilePrefix != nil {
		if strings.ReplaceAll(*_f.DestSegmentFilePrefix, " ", "") == "" {
			return errors.New("destination segment file prefix is empty")
		}
	}
	// LogLevel
	if _f.LogLevel != nil {
		if strings.ReplaceAll(*_f.LogLevel, " ", "") == "" {
			return errors.New("log level is empty")
		}
		if *_f.LogLevel != "fatal" && *_f.LogLevel != "error" && *_f.LogLevel != "verbose" {
			return errors.New("log level is unknown")
		}
	}
	// ThreadsCount
	if _f.ThreadsCount != nil {
		if *_f.ThreadsCount > 128 {
			return errors.New("threads count is invalid")
		}
	}
	// SegmentSeconds
	if _f.SegmentSeconds != nil {
		if *_f.SegmentSeconds < 1 {
			return errors.New("segment seconds cannot be zero or negative")
		}
	}
	// Bitrate
	if _f.Bitrate != nil {
		var isBitrateAvail bool = false
		for _, oneAllowedBitrate := range ffmpegAllowedAudioBitrates {
			if oneAllowedBitrate == *_f.Bitrate {
				isBitrateAvail = true
				break
			}
		}
		if !isBitrateAvail {
			return errors.New("specified bitrate is unavailable")
		}
	}
	// Codec
	if _f.Codec != nil {
		var isCodecAvail bool = false
		for _, oneAllowedCodec := range ffmpegAllowedAudioCodecs {
			if oneAllowedCodec == *_f.Codec {
				isCodecAvail = true
				break
			}
		}
		if !isCodecAvail {
			return errors.New("specified codec is unavailable")
		}
	}

	return nil
}

// ffmpegAudioFileToHLSResult : ...
type ffmpegAudioFileToHLSResult struct {
	FFMpegRespStdout  string  `json:"ffmpeg_resp_stdout"`
	FFMpegRespStderr  string  `json:"ffmpeg_resp_stderr"`
	ResultDestDirPath *string `json:"result_dest_dir_path"`
}
