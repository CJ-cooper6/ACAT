package middleware



import (
"fmt"
"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
"math"
"os"
"time"
)

func Logger()gin.HandlerFunc{
	filePath :="log/log.log"

	linkName := "latest_log.log"  //这个时可以在根目录下直接可以看到最新的日志

	scr,err :=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0755)
	if err != nil{
		fmt.Println("err:",err)
	}
	logger:=logrus.New()
	logger.Out =scr

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),	//日志保存时间
		retalog.WithRotationTime(24*time.Hour), 	//日志分割时间

		retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)


	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds())/1000000)))
		hostName,err := os.Hostname()
		if err != nil{
			hostName="unkown"
		}
		statusCode := c.Writer.Status()
		clientIp :=c.ClientIP()		//请求ip
		userAgent := c.Request.UserAgent()  //请求所使用的浏览器
		dataSize := c.Writer.Size()
		if dataSize < 0{
			dataSize =0
		}
		method := c.Request.Method
		path := c.Request.RequestURI


		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}else if statusCode >= 400{
			entry.Warn()
		}else {
			entry.Info()
		}

	}
}