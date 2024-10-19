package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/devShahriar/xm/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBInstance() *CompanyDB {

	conf := config.GetAppConfig()
	spew.Dump(conf)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DbConfig.Host,
		conf.DbConfig.User,
		conf.DbConfig.Password,
		conf.DbConfig.DbName,
		conf.DbConfig.Port,
	)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(conf.DbConfig.SlowQueryThreshold) * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		log.Printf("failed to connect to the database %v", err)
	}

	log.Println("DB connection created successfully")
	return &CompanyDB{Db: db}
}
