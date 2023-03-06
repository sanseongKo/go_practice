package Context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"os"
	"time"
)

type Echo struct {
	Echo *echo.Echo
}

type Gin struct {
}

type EchoContext struct {
	Context echo.Context
}

type GinContext struct {
	Context *gin.Context
}

type Router interface {
	RequestBody(obj any) error
	Response(obj any) error
	RequestQueryParam(key string) string
}

func (c EchoContext) Response(obj any) error {
	err := c.Context.JSON(http.StatusOK, obj)

	return err
}

func (c EchoContext) RequestBody(obj any) error {
	err := c.Context.Bind(&obj)

	if err != nil {
		_ = c.Context.String(http.StatusBadRequest, "bad Request")
	}

	return err
}

func (c EchoContext) RequestQueryParam(key string) string {
	return c.Context.QueryParam(key)
}

func (c *GinContext) Response(obj any) error {
	c.Context.JSON(http.StatusOK, obj)

	return nil
}

func (c *GinContext) RequestBody(obj any) error {
	err := c.Context.Bind(obj)

	if err != nil {
		c.Context.String(http.StatusBadRequest, "bad Request")
		panic(err)
	}

	return err
}

func (c *GinContext) RequestQueryParam(key string) string {
	return c.Context.Query(key)
}

func CustomEchoLogger() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout)
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("method", v.Method).
				Err(v.Error).
				Msg("request")

			return nil
		},
	})
}

func CustomGinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 사용자 정의 형식
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func WriteLogFile() {
	gin.DisableConsoleColor()

	// 파일에 로그를 작성합니다.
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
