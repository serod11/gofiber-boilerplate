/*
 * @author serod
 */

package db

import (
	"github.com/serod11/gofiber-boilerplate/pkg/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Init(config config.Config) *gorm.DB {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Millisecond * 100, // Slow SQL threshold
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,                   // Disable color
			},
		),
		SkipDefaultTransaction: true,
	}
	p := config.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormConfig); err != nil {
		return nil
	} else {
		RegisterTables(db)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
	//&model.Book{},
	)
	if err != nil {
		return
	}
}
