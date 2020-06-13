package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level FOR LOG
type Level int

// LOG CFG
var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// sET LOG LEVEL.
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug LOG SET
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info log set
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn log set
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error log set
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal log set
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
