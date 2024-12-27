package db

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	mysql.Config
	IsDebug                  bool
	SlowLogThresholdMilliSec int
}

func GetDefaultConfig() Config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(fmt.Sprintf("Failed to time.LoadLocation: err=%s\n", err.Error()))
	}

	return Config{
		Config: mysql.Config{
			User:      "admin",
			Passwd:    "adminpass",
			Addr:      "localhost:3306",
			DBName:    "iam",
			Net:       "tcp",
			ParseTime: true,
			Collation: "utf8mb4_unicode_ci",
			Loc:       jst,
		},
		IsDebug:                  true,
		SlowLogThresholdMilliSec: 5, //nolint:mnd
	}
}
