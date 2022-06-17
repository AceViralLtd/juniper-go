package juniper

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const LogTimeFormat = "2006-01-02 15:04:05"

func DefaultLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var (
				err       error
				errString string
				req       = ctx.Request()
				res       = ctx.Response()
				start     = time.Now()
			)
			if err = next(ctx); err != nil {
				ctx.Error(err)
				errString = err.Error()
			}
			stop := time.Now()

			var (
				ipAddress   = ctx.RealIP()
				countryCode = "A1"
			)

			if cfIp := req.Header.Get("cf-connecting-ip"); cfIp != "" {
				ipAddress = cfIp
			} else if originalIp := req.Header.Get("x-original-forwarded-for"); originalIp != "" {
				ipList := strings.Split(originalIp, ",")
				ipAddress = ipList[0]
			}

			if cfCountry := req.Header.Get("cf-ipcountry"); cfCountry != "" {
				countryCode = cfCountry
			}

			_, err = ctx.Logger().Output().Write([]byte(fmt.Sprintf(
				"%v | %d | %s | %s | %s | %s %#v\n%s",
				stop.Format(LogTimeFormat),
				res.Status,
				stop.Sub(start).String(),
				ipAddress,
				countryCode,
				req.Method,
				req.URL.Path,
				errString,
			)))

			return err
		}
	}
}

func AccountIdLogger(accountKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var (
				err       error
				errString string
				req       = ctx.Request()
				res       = ctx.Response()
				start     = time.Now()
			)
			if err = next(ctx); err != nil {
				ctx.Error(err)
				errString = err.Error()
			}
			stop := time.Now()

			var (
				ipAddress   = ctx.RealIP()
				accountId   = "-guest-"
				countryCode = "A1"
			)

			if cfIp := req.Header.Get("cf-connecting-ip"); cfIp != "" {
				ipAddress = cfIp
			} else if originalIp := req.Header.Get("x-original-forwarded-for"); originalIp != "" {
				ipList := strings.Split(originalIp, ",")
				ipAddress = ipList[0]
			}

			if cfCountry := req.Header.Get("cf-ipcountry"); cfCountry != "" {
				countryCode = cfCountry
			}

			if id, ok := ctx.Get(accountKey).(string); ok {
				accountId = id
			}

			_, err = ctx.Logger().Output().Write([]byte(fmt.Sprintf(
				"%v | %d | %s | %s | %s | %s | %s %#v\n%s",
				stop.Format(LogTimeFormat),
				res.Status,
				stop.Sub(start).String(),
				ipAddress,
				countryCode,
				accountId,
				req.Method,
				req.URL.Path,
				errString,
			)))

			return err
		}
	}
}
