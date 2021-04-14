package audioutils

// Source audio file location types enums
const (
	AudioFileSrcLocationTypeLocal int = 1
	AudioFileSrcLocationTypeURL   int = 2
)

// Source audio file to HLS convert fail enums
const (
	SrcAudioToHLSFailNotImpl           int = 1 // not implemented
	SrcAudioToHLSFailUnknown           int = 2 // unknown error
	SrcAudioToHLSFailInvalidSrcLocType int = 3 // invalid source location type
	SrcAudioToHLSFailParseAbsDestPath  int = 4 // invalid source location type
	SrcAudioToHLSFailFileNotFound      int = 5 // specified file is not found
)

// Audio file duration getting fail enums
const (
	AudioFileDurationGetFailNotImpl                int = 1 // not implemented
	AudioFileDurationGetFailUnknown                int = 2 // unknown error
	AudioFileDurationGetFailInvalidSrcLocType      int = 3 // invalid source location type
	AudioFileDurationGetFailFileNotFound           int = 4 // specified file is not found
	AudioFileDurationGetFailFFProbeReturnedErr     int = 5 // ffprobe returned error
	AudioFileDurationGetFailFFProbeFormatRespEmpty int = 6 // ffprobe format section is empty
)
