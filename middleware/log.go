package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	LogFilePath := os.Getenv("LOG_FILE_PATH")
	LogFileName := os.Getenv("LOG_FILE_NAME")

	fileName := path.Join(LogFilePath, LogFileName)
	//src, err := os.OpenFile(fileName, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//
	logger := logrus.New()
	//logger.Out = src
	//logger.SetLevel(logrus.DebugLevel)
	logWriter, err := rotatelogs.New(
		fileName + ".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		)
	if err != nil {
		fmt.Println("err: ", err)
	}
	
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      false,
	})

	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency_time": latencyTime,
			"client_ip": clientIP,
			"req_method": reqMethod,
			"req_uri": reqUri,
		}).Info()
	}
}