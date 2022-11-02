package juniper

import (
	"log"
	"math"
	"os"
	"time"

	"gorm.io/gorm"
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

type Paginator struct {
	Limit int   `json:"limit,omitempty" form:"limit"`
	Page  int   `json:"page,omitempty" form:"page"`
	Pages []int `json:"pages"`

	defaultLimit int
}

func NewPaginator(limit int) *Paginator {
	return &Paginator{
		defaultLimit: limit,
	}
}

func (pag *Paginator) GetPage() int {
	if pag.Page == 0 {
		pag.Page = 1
	}

	return pag.Page
}

func (pag *Paginator) GetLimit() int {
	if pag.Limit == 0 {
		pag.Limit = pag.defaultLimit
	}

	return pag.Limit
}

func (pag *Paginator) GetOffset() int {
	return pag.GetLimit() * (pag.GetPage() - 1)
}

func Paginate(db *gorm.DB, model interface{}, inator *Paginator) func(*gorm.DB) *gorm.DB {
	var rowCount int64
	db.Model(model).Count(&rowCount)

	pageCount := int(math.Ceil(float64(rowCount) / float64(inator.GetLimit())))
	inator.Pages = make([]int, 0, pageCount)
	for i := 1; i <= pageCount; i++ {
		inator.Pages = append(inator.Pages, i)
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(inator.GetOffset()).Limit(inator.GetLimit())
	}
}
