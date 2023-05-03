package main

import (
	"fmt"
	"gorm.io/gorm"
)

func TscBatchCreate(db *gorm.DB) error {
	tx := func(tdb *gorm.DB) error {
		students := []*Student{
			{
				Name: "Jack",
				Age:  38,
			},
			{
				Name: "Mack",
				Age:  34,
			},
		}
		result := db.Debug().Create(students)
		fmt.Printf("effect row: %d", result.RowsAffected)
		return result.Error
	}
	return db.Transaction(tx)
}
