package middleware

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/libs/constants"
	"github.com/kaolnwza/proj-blueprint/libs/utils"
	"github.com/kaolnwza/proj-blueprint/pkg/logger"
)

type responseWriter struct {
	gin.ResponseWriter
	ResponseBody *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.ResponseBody.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestLogger(logConf config.LogConfig) func(c *gin.Context) {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.URL.String(), "/health") && !logConf.IncludeHealth {
			c.Next()
			return
		}

		startTime := time.Now()
		log := logger.NewRequestLogger(c.Request)
		log.LogRequest()
		setCtx(c, log.ID())

		writer := &responseWriter{
			ResponseWriter: c.Writer,
			ResponseBody:   bytes.NewBufferString(""),
		}
		writer.Header().Add("version", config.VERSION)
		c.Writer = writer

		c.Next()

		logger.GetById(log.ID()).Infof("Completed with status code %v %v; latency = %v ms",
			c.Writer.Status(),
			http.StatusText(c.Writer.Status()),
			time.Since(startTime).Milliseconds(),
		)
		log.LogResponse(writer.ResponseBody.String(), writer.Status())
	}
}

// setup requestId on context for logging
func setCtx(c *gin.Context, requestId string) {
	ctx := utils.SetContext[string](c.Request.Context(), constants.CtxRequestId, requestId)
	c.Request = c.Request.WithContext(ctx)
}
