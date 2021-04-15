package audioutils

import (
	"io/ioutil"
	"log"
)

// AudioUtils : ...
type AudioUtils struct {
	ErrLogger  *log.Logger
	WarnLogger *log.Logger
	InfoLogger *log.Logger
}

// CreateAudioUtils : ...
func CreateAudioUtils(
	errLogger *log.Logger,
	warnLogger *log.Logger,
	infoLogger *log.Logger,
) *AudioUtils {
	var _audioUtils *AudioUtils = &AudioUtils{
		ErrLogger:  log.New(ioutil.Discard, "ERR: ", log.Ldate|log.Ltime|log.Llongfile),
		WarnLogger: log.New(ioutil.Discard, "WARN: ", log.Ldate|log.Ltime|log.Llongfile),
		InfoLogger: log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Llongfile),
	}

	if errLogger != nil {
		_audioUtils.ErrLogger = errLogger
	}
	if warnLogger != nil {
		_audioUtils.WarnLogger = warnLogger
	}
	if infoLogger != nil {
		_audioUtils.InfoLogger = infoLogger
	}

	return _audioUtils
}
