package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbUtil struct {
	db *gorm.DB
}

var dbInstance = &dbUtil{}

// func NewDbUtil() DbUtil {
// 	return &DbUtil{}
// }

func GormConnect(dbms string, dsn string) error {
	db, err := gorm.Open(dbms, dsn)
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
