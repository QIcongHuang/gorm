package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db := NewDB()
	if err := BatchCreate(db); err != nil {
		fmt.Printf("db action fail: %v\n", err)
	}
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("gorm init fail: %v\n", err))
	}
	return db
}

func NewDBWithCustomConfig() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DefaultDatetimePrecision: nil,
		//DisableWithReturning:          false,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("gorm init fail: %v\n", err))
	}
	return db
}
