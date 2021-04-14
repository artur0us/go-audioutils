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
