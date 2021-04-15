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

	_a.InfoLogger.Printf("validating input: %v\n", input)
	if err := input.validate(); err != nil {
		_a.ErrLogger.Printf("input validation failed: %v\n", err)
		return nil, err
	}

	var srcFilePath string = input.SrcLocation // `"` + input.SrcLocation + `"`
	var destDirPath string = input.DestDirPath

	var destM3U8FileName string = "playlist.m3u8"
	if input.DestM3U8FileName != nil {
		destM3U8FileName = *input.DestM3U8FileName
	}
	var destM3U8FileDirPath string = fmt.Sprintf("%v/%v", destDirPath, destM3U8FileName) // fmt.Sprintf(`"%v/playlist.m3u8"`, destDirPath)

	var destSegmentFileExt string = "m4a"
	var destSegmentFilePrefix string = "result_file_"
	if input.DestSegmentFilePrefix != nil {
		destSegmentFilePrefix = *input.DestSegmentFilePrefix
	}
	var destSegmentsFilesDirPath string = destDirPath + "/" + destSegmentFilePrefix + "%d." + destSegmentFileExt // `"` + destDirPath + `/result_file_%d.m4a"`

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

	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> srcFilePath: %v\n", srcFilePath)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destDirPath: %v\n", destDirPath)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destM3U8FileName: %v\n", destM3U8FileName)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destM3U8FileDirPath: %v\n", destM3U8FileDirPath)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destSegmentFileExt: %v\n", destSegmentFileExt)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destSegmentFilePrefix: %v\n", destSegmentFilePrefix)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> destSegmentsFilesDirPath: %v\n", destSegmentsFilesDirPath)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> ffmpegThreadsCount: %v\n", ffmpegThreadsCount)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> ffmpegLogLevel: %v\n", ffmpegLogLevel)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> ffmpegOutputHLSAudioBitrate: %v\n", ffmpegOutputHLSAudioBitrate)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> ffmpegOutputHLSAudioCodec: %v\n", ffmpegOutputHLSAudioCodec)
	_a.InfoLogger.Printf("ffmpegSrcAudioFileToHLS internal variables -> ffmpegHLSOneSegmentSeconds: %v\n", ffmpegHLSOneSegmentSeconds)

	_a.InfoLogger.Println("building new ffmpeg process config...")
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

	_a.InfoLogger.Println("running ffmpeg process...")
	if err := cmd.Run(); err != nil {
		_a.ErrLogger.Printf("failed to start ffmpeg process: %v\n", err)
		_a.ErrLogger.Printf("ffmpeg stdout:\n%v\n", stdoutBuf.String())
		_a.ErrLogger.Printf("ffmpeg stderr:\n%v\n", stderrBuf.String())
		return nil, fmt.Errorf("failed to start ffmpeg process: %w", err)
	}

	_a.InfoLogger.Println("checking appending segment file prefix...")
	if input.AppendingSegmentFilePrefix != nil {
		_a.InfoLogger.Printf("appending segment file prefix is set: %v\n", *input.AppendingSegmentFilePrefix)

		selectedLinesPatterns := []string{destSegmentFilePrefix, destSegmentFileExt}

		_a.InfoLogger.Printf("calling *addStrBeforeFileNameInM3U8* function with params: %v; %v; %v\n", destM3U8FileDirPath, selectedLinesPatterns, *input.AppendingSegmentFilePrefix)
		err = addStrBeforeFileNameInM3U8(
			destM3U8FileDirPath,
			selectedLinesPatterns,
			*input.AppendingSegmentFilePrefix,
		)
		if err != nil {
			_a.ErrLogger.Printf("failed to add string before segment files names in M3U8 playlist file: %v\n", err)
			return nil, fmt.Errorf("failed to add string before segment files names in M3U8 playlist file: %w", err)
		}
	}

	_a.InfoLogger.Printf("audio file is successfully converted to HLS: (destDirPath: %v)\n", destDirPath)
	var res *ffmpegAudioFileToHLSResult = &ffmpegAudioFileToHLSResult{}
	res.FFMpegRespStdout = stdoutBuf.String()
	res.FFMpegRespStderr = stderrBuf.String()
	res.ResultDestDirPath = &destDirPath

	return res, nil
}
