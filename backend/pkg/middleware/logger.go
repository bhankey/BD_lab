package middleware

import (
	"net/http"
	"time"

	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Logger return a middleware that logs HTTP request.
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			log := logger.GetLoggerWithField(logrus.Fields{
				"remote_addr": c.RealIP(),
				"request_id":  id,
			})

			start := time.Now()
			log.Infof("started %s %s", req.Method, req.RequestURI)

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			log.Infof("completed with %d %s in %v", res.Status, http.StatusText(res.Status), time.Since(start))

			return err
		}
	}
}
