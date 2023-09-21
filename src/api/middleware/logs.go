package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/gin-gonic/gin"
)


func CustomLogger(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start:= time.Now()
		path := ctx.FullPath()
		raw := ctx.Request.URL.RawQuery
		reqBdy, _ := io.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBdy))

		ctx.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.BodySize = ctx.Writer.Size()

		if raw != "" {
			path += "?" + raw
		}
		param.Path = path

		key := map[logger.ExtraKey]interface{}{}

		key[logger.ClientIp] = param.ClientIP
		key[logger.Latency] = param.Latency
		key[logger.Method] = param.Method
		key[logger.BodySize] = param.BodySize
		key[logger.Path] = param.Path
		key[logger.RequestBody] = string(reqBdy)

		log.Info(logger.RequestResponse, logger.Api, "", key)
	}
}