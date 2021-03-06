package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

// FileLogger 文件日志结构体
type FileLogger struct {
	level       LogLevel
	filePath    string //日志文件保存的路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// initFile 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file failed, err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// checkSize 判断文件是否需要被切割
func (f *FileLogger) checkSize(file *os.File) bool {
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	//如果当前文件大小大于等于日志文件的最大值，就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

// splitFile 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割文件
	// 2 备份一下 rename xx.log -> xx.log.bak202107130821
	// 2.1 获取当前时间的字符串，如20060102130405
	nowStr := time.Now().Format("20060102130405")
	// 2.2 拿到当前的日志文件路径
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1 关闭当前文件
	err = file.Close()
	if err != nil {
		fmt.Printf("close file failed, err is %v\n", err)
		return nil, err
	}
	// 2.3拼接一个日志文件备份的名字
	err = os.Rename(logName, newLogName)
	if err != nil {
		fmt.Printf("rename file failed, err is %v\n", err)
		return nil, err
	}
	// 3 打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log failed, err:%v\n", err)
		return nil, err
	}
	// 4 将打开的新的日志文件对象赋值给f.fileObj
	return fileObj, nil
}

// log 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			// 日志文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"),
			getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				// 错误日志文件
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"),
				getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

// enable 判断是否需要记录该日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.level <= logLevel
}

// Debug Debug级别日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info Info级别日志
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Waring Waring级别日志
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error Error级别日志
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal Fatal级别日志
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
