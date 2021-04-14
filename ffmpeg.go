package audioutils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// ffmpeg binary path
const (
	FFMpegPath string = "ffmpeg"
)

// ffmpegSrcAudioFileToHLS : ...
func (_a *AudioUtils) ffmpegSrcAudioFileToHLS(input ffmpegAudioFileToHLSInput) (*ffmpegAudioFileToHLSResult, error) {
	// Local:
	// ffmpeg -y -i source_file.mp3 -hide_banner -y -threads 0 -v quiet -progress - -loglevel verbose -c:a aac -b:a 128k -muxdelay 0 -f segment -sc_threshold 0 -segment_time 45 -segment_list ./output/playlist.m3u8 -segment_format mpegts ./output/result_file_%d.m4a
	// URL:
	// ffmpeg -y -i https://localhost/source_file.mp3 -hide_banner -y -threads 0 -v quiet -progress - -loglevel verbose -c:a aac -b:a 128k -muxdelay 0 -f segment -sc_threshold 0 -segment_time 45 -segment_list ./output/playlist.m3u8 -segment_format mpegts ./output/result_file_%d.m4a

	if err := input.validate(); err != nil {
		return nil, err
	}

	var srcFilePath string = input.SrcLocation // `"` + input.SrcLocation + `"`
	var destDirPath string = input.DestDirPath

	var destM3U8FileName string = "playlist.m3u8"
	if input.DestM3U8FileName != nil {
		destM3U8FileName = *input.DestM3U8FileName
	}
	var destM3U8FileDirPath string = fmt.Sprintf("%v/%v", destDirPath, destM3U8FileName) // fmt.Sprintf(`"%v/playlist.m3u8"`, destDirPath)

	var destSegmentFilePrefix string = "result_file_"
	if input.DestSegmentFilePrefix != nil {
		destSegmentFilePrefix = *input.DestSegmentFilePrefix
	}
	var destSegmentsFilesDirPath string = destDirPath + "/" + destSegmentFilePrefix + "%d.m4a" // `"` + destDirPath + `/result_file_%d.m4a"`

	var ffmpegThreadsCount int = 0
	if input.ThreadsCount != nil {
		ffmpegThreadsCount = *input.ThreadsCount
	}

	var ffmpegLogLevel string = "error"
	if input.LogLevel != nil {
		ffmpegLogLevel = *input.LogLevel
	}

	var ffmpegOutputHLSAudioBitrate string = "128k"
	if input.Bitrate != nil {
		ffmpegOutputHLSAudioBitrate = fmt.Sprintf("%vk", *input.Bitrate)
	}

	var ffmpegOutputHLSAudioCodec string = "aac"
	if input.Codec != nil {
		ffmpegOutputHLSAudioCodec = *input.Codec
	}

	var ffmpegHLSOneSegmentSeconds int = 15
	if input.SegmentSeconds != nil {
		ffmpegHLSOneSegmentSeconds = *input.SegmentSeconds
	}

	ctx := context.Background()
	args := []string{
		"-y", "-i", srcFilePath,
		"-hide_banner", "-y",
		"-threads", strconv.Itoa(ffmpegThreadsCount),
		"-v", "quiet",
		"-progress", "-",
		"-loglevel", ffmpegLogLevel,
		"-c:a", ffmpegOutputHLSAudioCodec,
		"-b:a", ffmpegOutputHLSAudioBitrate,
		"-muxdelay", "0",
		"-f", "segment",
		"-sc_threshold", "0",
		"-segment_time", strconv.Itoa(ffmpegHLSOneSegmentSeconds),
		"-segment_list", destM3U8FileDirPath,
		"-segment_format", "mpegts", destSegmentsFilesDirPath,
	}

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer

	cmd := exec.CommandContext(ctx, FFMpegPath, args...)
	cmd.SysProcAttr = nil
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	// cmd.Stdin = reader

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}
	cmd.Dir = cwd

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg process: %w", err)
	}

	var res *ffmpegAudioFileToHLSResult = &ffmpegAudioFileToHLSResult{}
	res.FFMpegRespStdout = stdoutBuf.String()
	res.FFMpegRespStderr = stderrBuf.String()
	res.ResultDestDirPath = &destDirPath

	return res, nil
}
