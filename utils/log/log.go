package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// 相当于宏，调用以下变量即可
var (
	Error   = errorLog.Println
	Warning = warningLog.Println
	Info    = infoLog.Println
	Infof   = infoLog.Printf
)

var (
	// error 级日志
	errorLog = log.New(os.Stdout, "dgb_error:", log.LstdFlags|log.Lshortfile)
	// warning 级日志
	warningLog = log.New(os.Stdout, "dgb_warning:", log.LstdFlags|log.Lshortfile)
	// info 级日志
	infoLog = log.New(os.Stdout, "dgb_info:", log.LstdFlags|log.Lshortfile)
	// 日志指针数组
	loggers = []*log.Logger{errorLog, warningLog, infoLog}
	//全局统一锁
	mulock sync.Mutex
)

// log levels
// iota 自增成四个级别
const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
	Disabled
)

// SetLevel 设置日志输出等级
func SetLevel(level int) {
	mulock.Lock()
	defer mulock.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	// 如果低于输出等级，则取消该等级的日志的io
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if WarningLevel < level {
		warningLog.SetOutput(ioutil.Discard)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}
