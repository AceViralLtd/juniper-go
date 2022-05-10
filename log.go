package juniper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogFormatter for gin to use with aditional data
func LogFormatter(param gin.LogFormatterParams) string {
	var (
		ipAddress   = param.ClientIP
		countryCode = "A1"
	)

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	if cfIp := param.Request.Header.Get("cf-connecting-ip"); cfIp != "" {
		ipAddress = cfIp
	} else if originalIp := param.Request.Header.Get("x-original-forwarded-for"); originalIp != "" {
		ipList := strings.Split(originalIp, ",")
		ipAddress = ipList[0]
	}

	if cfCountry := param.Request.Header.Get("cf-ipcountry"); cfCountry != "" {
		countryCode = cfCountry
	}

	return fmt.Sprintf(
		"%v | %3d | %13v | %15s | %2s | %-7s %#v\n%s",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.StatusCode,
		param.Latency,
		ipAddress,
		countryCode,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

// LogFormatterWithAccountId takes the standard juniper log formatter an adds a user id to it
func LogFormatterWithAccountId(accountKey string) func(gin.LogFormatterParams) string {
	return func(param gin.LogFormatterParams) string {
		var (
			accountId   = "-guest-"
			ipAddress   = param.ClientIP
			countryCode = "A1"
		)

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		if id, ok := param.Keys[accountKey]; ok {
			accountId, _ = id.(string)
		}

		if cfIp := param.Request.Header.Get("cf-connecting-ip"); cfIp != "" {
			ipAddress = cfIp
		} else if originalIp := param.Request.Header.Get("x-original-forwarded-for"); originalIp != "" {
			ipList := strings.Split(originalIp, ",")
			ipAddress = ipList[0]
		}

		if cfCountry := param.Request.Header.Get("cf-ipcountry"); cfCountry != "" {
			countryCode = cfCountry
		}

		return fmt.Sprintf(
			"%v | %3d | %13v | %15s | %s | %2s | %-7s %#v\n%s",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.StatusCode,
			param.Latency,
			ipAddress,
			countryCode,
			accountId,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}
}
