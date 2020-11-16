package log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var performanceLogger *PerformanceLogger

var logEnabled = false

type PerformanceLogger struct {
	logger *log.Logger
	enable bool
	file   *os.File
	index  int
	start  time.Time
}

func Enable(enable bool) {
	logEnabled = enable
	if logEnabled && performanceLogger == nil {
		performanceLogger = new(PerformanceLogger).init()
	}
}

func GetInstance() *PerformanceLogger {
	return performanceLogger
}

func (p *PerformanceLogger) init() *PerformanceLogger {
	if logEnabled {
		var err error
		fileName := time.Now().Format("20060102_150405.txt")
		p.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open file: ", fileName)
		}
		p.logger = log.New(p.file, "", 0)
		p.index = 1
	}

	return p
}

func (p *PerformanceLogger) Start() {
	if logEnabled {
		p.start = time.Now()
	}
}

func (p *PerformanceLogger) StopAndLog(method string, url string) {
	if logEnabled {
		duration := time.Since(p.start).Milliseconds()

		// Strip parameters
		i := strings.IndexByte(url, '?')
		var path string
		if i > 0 {
			path = url[0:i]
		} else {
			path = url
		}

		// Log the header before first record
		if p.index == 1 {
			p.logger.Println("Index, Duration(ms), URL")
		}
		p.logger.Println(fmt.Sprintf("%d, %d, %s %s", p.index, duration, method, path))

		p.index++
	}
}
