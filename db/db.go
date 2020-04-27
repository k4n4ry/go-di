package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/knry0329/go-di/config"
)

type dbUtil struct {
	db *gorm.DB
}

var dbInstance = &dbUtil{}

// func NewDbUtil() DbUtil {
// 	return &DbUtil{}
// }

func GormConnect(cfg config.DbConfig) error {
	db, err := gorm.Open(cfg.Dbms, cfg.Dsn)
	if err != nil {
		return err
	}
	dbInstance.db = db
	return nil
}

// もどりにgormがあると、repositoryもgormであることを意識しなければいけなくなってしまう。
// gorm.DBがなにか汎用的なifを実装していればいいのだが。。。。
func GetDBInstance() *gorm.DB {
	return dbInstance.db
}
