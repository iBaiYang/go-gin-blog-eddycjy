package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/iBaiYang/go-gin-blog-eddycjy/global"
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)

		/*日志追踪*/
		// 把 *gin.Context 传入日志方法中
		//s := "access log: method: %s, status_code: %d, " + "begin_time: %d, end_time: %d"
		//global.Logger.WithFields(fields).Infof(c,
		//	s,
		//	c.Request.Method,
		//	bodyWriter.Status(),
		//	beginTime,
		//	endTime,
		//)
	}
}
