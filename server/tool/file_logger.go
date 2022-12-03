package tool

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

// Logger 日志实例
var Logger *FileLogger

type LogTemp struct {
	Level   LogLevel `json:"level"`
	Message string   `json:"message"`
}

//
// init
//  @Description: 初始化日志实例
//
func init() {
	Logger = NewFileLogger("warning", "./", "log.log", 100*1024*1024)
}

// LogLevel 日志等级
type LogLevel uint16

//  日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//
// paraLogLevel
//  @Description: 解析日志
//  @param s 日志级别字符串
//  @return LogLevel
//  @return error
//
func paraLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("无效日志级别")
	}
}

//
// getLogStr
//  @Description: 获取日志的字符串格式
//  @param level 日志级别
//  @return string
//
func getLogStr(level LogLevel) string {
	switch level {
	case DEBUG:
		return "debug"
	case TRACE:
		return "trace"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "unknown"
	}
}

//
// FileLogger
//  @Description: 日志结构体
//
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//
// initFile
//  @Description: 初始化要打开和写入的日志文件
//  @receiver fl FileLogger
//  @return error
//
func (fl *FileLogger) initFile() error {
	join := path.Join(fl.filePath, fl.fileName)
	fileObj, err := os.OpenFile(join, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	errFileObj, err := os.OpenFile(join+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fl.fileObj = fileObj
	fl.errFileObj = errFileObj
	return nil
}

//
// NewFileLogger
//  @Description: 生成FileLogger
//  @param LeverStr 日志等级字符串
//  @param filePath 文件路径
//  @param fileName 文件名称
//  @param size 文件最大容量
//  @return *FileLogger
//
func NewFileLogger(LeverStr, filePath, fileName string, size int64) *FileLogger {
	level, err := paraLogLevel(LeverStr)
	if err != nil {
		panic(err)
	}

	f := &FileLogger{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: size,
	}
	err = f.initFile()
	if err != nil {
		panic(err)
	}

	return f
}

//
// enable
//  @Description: 判断级别大小
//  @receiver fl FileLogger
//  @param logLevel 级别
//  @return bool
//
func (fl *FileLogger) enable(logLevel LogLevel) bool {
	return fl.Level >= logLevel
}

//
// Log
//  @Description: 打印日志
//  @receiver fl FileLogger
//  @param level 等级
//  @param msg
//
func (fl *FileLogger) Log(level LogLevel, msg string) {
	now := time.Now()
	if fl.enable(level) {
		//  等级在error之下写在.log文件
		_, err := fmt.Fprintf(fl.fileObj, "[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogStr(level), msg)
		if err != nil {
			return
		}
	} else {
		//  等级在error及其之上写在.log.err文件
		_, err := fmt.Fprintf(fl.errFileObj, "[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogStr(level), msg)
		if err != nil {
			return
		}
	}
}

//
// Debug
//  @Description: 调试
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Debug(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(DEBUG, msg)
}

//
// Trace
//  @Description: 跟踪
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Trace(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(TRACE, msg)
}

//
// Info
//  @Description: 消息
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Info(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(INFO, msg)
}

//
// Warning
//  @Description: 警告
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Warning(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(WARNING, msg)
}

//
// Error
//  @Description: 错误
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Error(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(ERROR, msg)
}

//
// Fatal
//  @Description: 致命
//  @receiver fl FileLogger
//  @param msg 信息
//  @param s 其他补充信息
//
func (fl *FileLogger) Fatal(msg string, s ...interface{}) {
	msg = fmt.Sprint(msg, s)
	fl.Log(FATAL, msg)
}

//
// Close
//  @Description: 关闭
//  @receiver fl FileLogger
//
func (fl *FileLogger) Close() {
	err := fl.fileObj.Close()
	if err != nil {
		return
	}
	err = fl.errFileObj.Close()
	if err != nil {
		return
	}
}
