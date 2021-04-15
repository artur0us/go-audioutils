package main

import (
	"log"
	"os"

	"github.com/artur0us/go-audioutils"
)

// go run .
func main() {
	if err := os.Chdir("."); err != nil {
		log.Printf("failed to change directory: %v", err)
		return
	}

	_audioUtils := audioutils.CreateAudioUtils(
		log.New(os.Stdout, "ERR: ", log.Ldate|log.Ltime|log.Llongfile),
		log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Llongfile),
		log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile),
	)

	// Duration
	durationResult := _audioUtils.GetAudioFileDuration(audioutils.AudioFileDurationRequest{
		SrcLocationType: audioutils.AudioFileSrcLocationTypeLocal,
		SrcLocation:     "source_file.mp3",
		// SrcLocationType: audioutils.AudioFileSrcLocationTypeURL,
		// SrcLocation:     "https://localhost/source_file.mp3",
	})
	log.Println(durationResult)
	if durationResult.Duration != nil {
		log.Println(*durationResult.Duration)
	}

	// Basic info
	basicInfoResult := _audioUtils.GetAudioFileBasicInfo(audioutils.AudioFileBasicInfoRequest{
		SrcLocationType: audioutils.AudioFileSrcLocationTypeLocal,
		SrcLocation:     "source_file.mp3",
	})
	log.Println(basicInfoResult)

	// HLS
	hlsM3U8FileName := "playlist.m3u8"
	hlsSegmentFilePrefix := "hls_seg_"
	threadsCount := 0
	hlsSegmentSeconds := 15
	hlsAudioBitrate := 320
	hlsAudioCodec := "aac"
	hlsAppendingSegmentFilePrefix := "https://10.0.0.1/"
	hlsResult := _audioUtils.ConvertSrcAudioFileToHLS(audioutils.SrcAudioToHLSRequest{
		DeleteDestDirPathIfExists:     true,
		DestDirPath:                   "output",
		HLSM3U8FileName:               &hlsM3U8FileName,
		HLSSegmentFilePrefix:          &hlsSegmentFilePrefix,
		ThreadsCount:                  &threadsCount,
		HLSSegmentSeconds:             &hlsSegmentSeconds,
		HLSAudioBitrate:               &hlsAudioBitrate,
		HLSAudioCodec:                 &hlsAudioCodec,
		HLSAppendingSegmentFilePrefix: &hlsAppendingSegmentFilePrefix,

		SrcLocationType: audioutils.AudioFileSrcLocationTypeLocal,
		SrcLocation:     "source_file.mp3",
		// SrcLocationType: audioutils.AudioFileSrcLocationTypeURL,
		// SrcLocation:     "https://localhost/source_file.mp3",
	})
	log.Println(hlsResult)
	if hlsResult.ResultDestPath != nil {
		log.Println(*hlsResult.ResultDestPath)
	}
}
