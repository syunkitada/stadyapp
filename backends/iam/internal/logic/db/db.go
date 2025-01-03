package db

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
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
		tlog.Fatal(ctx, "failed to MustOpen")
	}
}

func (self *DB) MustOpenMock(ctx context.Context) sqlmock.Sqlmock {
	mock, err := self.OpenMock(ctx)
	if err != nil {
		tlog.Fatal(ctx, "failed to MustOpenMock")
	}

	return mock
}

func (self *DB) OpenMock(ctx context.Context) (sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to sqlmock.New")
	}

	self.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to gorm.Open")
	}

	return mock, nil
}

func (self *DB) Open(ctx context.Context) error {
	var err error
	self.DB, err = gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{
		Logger: &tlog.GormLogger{
			LogLevel:      logger.Info,
			SlowThreshold: time.Duration(self.conf.SlowLogThresholdMilliSec) * time.Millisecond,
		},
	})

	if err != nil {
		return tlog.WrapErr(ctx, err, "failed to gorm.Open")
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
		return tlog.WrapErr(ctx, err, "failed to self.DB.DB")
	} else {
		if err := db.Close(); err != nil {
			return tlog.WrapErr(ctx, err, "failed to db.Close")
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
		return tlog.WrapErr(ctx, err, "failed to gorm.Open")
	}

	if err := db.Exec("DROP DATABASE IF EXISTS " + dbName).Error; err != nil {
		return tlog.WrapErr(ctx, err, "failed to db.Exec")
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
		return tlog.WrapErr(ctx, err, "failed to gorm.Open")
	}

	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error; err != nil {
		return tlog.WrapErr(ctx, err, "failed to db.Exec")
	}

	return nil
}

func (self *DB) Transact(txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := self.DB.Begin()
	if err = tx.Error; err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil { //nolint:nestif
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on recover: %s", tmpErr.Error())
			}

			err = fmt.Errorf("recovered: %v", p) //nolint:err113
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
	Retrys           int
	RetryIntervalSec int
	Msg              string
}

func (e *RetryError) Error() string {
	return e.Msg
}

func AsRetryError(err error) (bool, *RetryError) {
	var retryErr *RetryError
	ok := errors.As(err, &retryErr)

	return ok, retryErr
}

func (self *DB) TransactWithRetry(ctx context.Context, txFunc func(tx *gorm.DB) (err error)) error {
	err := transact(self.DB, txFunc)
	if err != nil { //nolint:nestif
		ok, retryError := AsRetryError(err)
		if ok {
			if retryError.RetryIntervalSec == 0 {
				retryError.RetryIntervalSec = 3
			}

			time.Sleep(time.Duration(retryError.RetryIntervalSec) * time.Second)

			for i := range retryError.Retrys {
				tlog.Error(ctx, "failed to transact, but retrying", slog.Int("retry", i), slog.String("err", err.Error()))
				err = transact(self.DB, txFunc)
				ok, _ = AsRetryError(err)

				if ok {
					continue
				} else {
					return err
				}
			}
		} else {
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
		if p := recover(); p != nil { //nolint:nestif
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("failed rollback on recover: %s", tmpErr.Error())
			}

			err = fmt.Errorf("recovered: %v", p) //nolint:err113
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
