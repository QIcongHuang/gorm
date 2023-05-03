package main

import (
	"fmt"
	"gorm.io/gorm"
)

type UUIDStudent struct {
	UUID uint `gorm:"primaryKey"`
	Name string
	Age  uint
}

// GetAllUUIDStudents 查询所有用户
func GetAllUUIDStudents(db *gorm.DB) error {
	var students []*UUIDStudent
	err := db.Debug().Find(&students).Error
	for i, s := range students {
		fmt.Printf("student index: %d, info: %v\n", i, s)
	}
	return err
}

func BatchCreateUUIDStudent(db *gorm.DB) error {
	students := []*UUIDStudent{
		{
			UUID: 7,
			Name: "Leo",
			Age:  38,
		},
		{
			UUID: 8,
			Name: "Marry",
			Age:  34,
		},
	}
	result := db.Debug().Create(students)
	fmt.Printf("effect row: %d", result.RowsAffected)
	return result.Error
}

func UpdatesUUIDStudent(db *gorm.DB) error {
	result := db.Debug().Model(&UUIDStudent{
		UUID: 1,
	}).Updates(UUIDStudent{
		Name: "Jack1",
		Age:  38,
	})
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}
