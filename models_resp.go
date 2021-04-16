package audioutils

// SrcAudioToHLSResult : ...
type SrcAudioToHLSResult struct {
	IsSuccess      bool    `json:"is_success"`
	FailCode       int     `json:"fail_code"`
	FailMsg        string  `json:"fail_msg"`
	ResultDestPath *string `json:"result_dest_path"`
}

// AudioFileDurationResult : ...
type AudioFileDurationResult struct {
	IsSuccess bool     `json:"is_success"`
	FailCode  int      `json:"fail_code"`
	FailMsg   string   `json:"fail_msg"`
	Duration  *float64 `json:"duration"`
}

// AudioFileBasicInfoResult : ...
type AudioFileBasicInfoResult struct {
	IsSuccess bool   `json:"is_success"`
	FailCode  int    `json:"fail_code"`
	FailMsg   string `json:"fail_msg"`

	Duration       *float64 `json:"duration"`
	Size           *int64   `json:"size"`
	Bitrate        *int     `json:"bitrate"`
	FormatName     *string  `json:"format_name"`
	FormatLongName *string  `json:"format_long_name"`
	StreamsCount   *int     `json:"streams_count"`
}

// AudioFileWaveformDataResult : ...
type AudioFileWaveformDataResult struct {
	IsSuccess bool   `json:"is_success"`
	FailCode  int    `json:"fail_code"`
	FailMsg   string `json:"fail_msg"`

	Points *[]int `json:"points"`
}
