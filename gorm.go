package juniper

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

// GormLoggerProduction will generate a configured logging instance for use with the gorm ORM
func GormLoggerProduction() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}

// GormLoggerDebug will print all queries
func GormLoggerDebug(logPath *string) logger.Interface {
	var (
		err error
		fh  = os.Stdout
	)

	if logPath != nil {
		fh, err = os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fh = os.Stdout
		}
	}

	return logger.New(
		log.New(fh, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Microsecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
}
