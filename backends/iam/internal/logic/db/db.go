package db

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type DB struct {
	conf *Config
	DB   *gorm.DB
}

func New(conf *Config) db.IDB {
	return &DB{
		conf: conf,
	}
}

func (self *DB) MustOpen(ctx context.Context) {
	if err := self.Open(ctx); err != nil {
		fmt.Println("failed to MustOpen")
		os.Exit(1)
	}
}

func (self *DB) MustOpenMock(ctx context.Context) sqlmock.Sqlmock {
	mock, err := self.OpenMock(ctx)
	if err != nil {
		fmt.Println("failed to MustOpenMock")
		os.Exit(1)
	}

	return mock
}

func (self *DB) OpenMock(ctx context.Context) (sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to sqlmock.New")
	}

	self.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to gorm.Open")
	}

	return mock, nil
}

func (self *DB) Open(ctx context.Context) error {
	var err error
	self.DB, err = gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{
		Logger: &tlog.Logger{
			LogLevel:      logger.Info,
			SlowThreshold: time.Duration(self.conf.SlowLogThresholdMilliSec) * time.Millisecond,
		},
	})

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to gorm.Open")
	}

	if self.conf.IsDebug {
		self.DB.Logger.LogMode(logger.Info)
		self.DB = self.DB.Debug()
	}

	return nil
}

func (self *DB) MustClose(ctx context.Context) {
	if err := self.Open(ctx); err != nil {
		log.Fatalf("failed Close")
	}
}

func (self *DB) Close(ctx context.Context) error {
	if db, err := self.DB.DB(); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.DB.DB")
	} else {
		if err := db.Close(); err != nil {
			return tlog.WrapError(ctx, err, "failed to db.Close")
		}
	}

	return nil
}

func (self *DB) MustCreateDatabase(ctx context.Context) {
	if err := self.CreateDatabase(ctx); err != nil {
		log.Fatalf("failed to CreateDatabase")
	}
}

func (self *DB) MustDropDatabase(ctx context.Context) {
	if err := self.DropDatabase(ctx); err != nil {
		log.Fatalf("failed DropDatabase")
	}
}

func (self *DB) DropDatabase(ctx context.Context) error {
	dbName := self.conf.DBName
	self.conf.DBName = ""

	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return tlog.WrapError(ctx, err, "failed to gorm.Open")
	}

	if err := db.Exec("DROP DATABASE IF EXISTS " + dbName).Error; err != nil {
		return tlog.WrapError(ctx, err, "failed to db.Exec")
	}

	return nil
}

func (self *DB) MustRecreateDatabase(ctx context.Context) {
	self.MustDropDatabase(ctx)
	self.MustCreateDatabase(ctx)
}

func (self *DB) CreateDatabase(ctx context.Context) error {
	dbName := self.conf.DBName
	self.conf.DBName = ""

	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return tlog.WrapError(ctx, err, "failed to gorm.Open")
	}

	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error; err != nil {
		return tlog.WrapError(ctx, err, "failed to db.Exec")
	}

	return nil
}

func (self *DB) Transact(txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := self.DB.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on recover: %s", tmpErr.Error())
			}

			err = fmt.Errorf("recovered: %v", p)
		} else if err != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on err: %s", tmpErr.Error())
			} else {
				log.Printf("rollbacked because of err: %s", err.Error())
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				log.Printf("Failed commit: %s", err.Error())

				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					log.Printf("failed rollback on commit: %s", tmpErr.Error())
				}
			}
		}
	}()

	err = txFunc(tx)

	return err
}

type RetryError struct {
	TTL int
	Msg string
}

func (e *RetryError) Error() string {
	return e.Msg
}

func (self *DB) TransactWithRetry(txFunc func(tx *gorm.DB) (err error)) error {
	err := transact(self.DB, txFunc)
	if err != nil {
		switch err.(type) {
		case *RetryError:
			ttl := err.(*RetryError).TTL
			n := rand.Intn(ttl)
			time.Sleep(time.Duration(n) * time.Second)

			for i := range ttl {
				fmt.Printf("retry count=%d, %s\n", i, err.Error())
				err = transact(self.DB, txFunc)

				switch err.(type) {
				case *RetryError:
					continue
				default:
					return err
				}
			}
		default:
			return err
		}
	}

	return nil
}

//nolint:nolintlint,nonamedreturns
func transact(db *gorm.DB, txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on recover: %s", tmpErr.Error())
			}

			err = fmt.Errorf("recovered: %v", p)
		} else if err != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on err: %s", tmpErr.Error())
			} else {
				log.Printf("rollbacked because of err: %s", err.Error())
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				log.Printf("failed commit: %s", err.Error())

				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					log.Printf("failed rollback on commit: %s", tmpErr.Error())
				}
			}
		}
	}()

	err = txFunc(tx)

	return err
}
