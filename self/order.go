package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Order struct {
	OrderId uint `gorm:"primaryKey;<-:false"`
	UserId  uint `gorm:"primaryKey;size:23;"`
	//OrderId uint `gorm:"primaryKey;autoIncrement:true"`
	//UserId  uint `gorm:"primaryKey;autoIncrement:false"`
	Name string `gorm:"not null;size:200"`
}

// GetAllOrders 查询所有用户
func GetAllOrders(db *gorm.DB) error {
	var students []*Order
	err := db.Debug().Find(&students).Error
	for i, s := range students {
		fmt.Printf("student index: %d, info: %v\n", i, s)
	}
	return err
}

func BatchCreateOrder(db *gorm.DB) error {
	students := []*Order{
		//{
		//	OrderId: 2,
		//	UserId:  2,
		//	Name:    "22",
		//},
		//{
		//	OrderId: 2,
		//	UserId:  3,
		//	Name:    "23",
		//},
		{
			UserId: 247329567,
			Name:   "夏天短袖",
		},
		{
			UserId: 247329567,
			Name:   "跑鞋",
		},
	}
	result := db.Debug().Create(students)
	fmt.Printf("effect row: %d", result.RowsAffected)
	return result.Error
}
