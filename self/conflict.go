package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateDoNothing(db *gorm.DB) error {
	student := &Student{
		ID:   1,
		Name: "Jack",
		Age:  38,
	}
	result := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Debug().Create(student)
	fmt.Printf("create do nothing while conflict: affect row: %d", result.RowsAffected)
	return result.Error
}

func CreateUseSQL(db *gorm.DB) error {
	student := &Student{
		ID:   1,
		Name: "Jack",
		Age:  38,
	}
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"age": gorm.Expr("`age` + 100")}),
	}).Debug().Create(student)
	fmt.Printf("create do nothing while conflict: affect row: %d", result.RowsAffected)
	return result.Error
}

func CreateUpdateColumns(db *gorm.DB) error {
	student := &Student{
		ID:   1,
		Name: "Jackxxx",
		Age:  11138,
	}
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Debug().Create(student)
	fmt.Printf("create do nothing while conflict: affect row: %d", result.RowsAffected)
	return result.Error
}

func CreateUpdateAll(db *gorm.DB) error {
	student := &Student{
		ID:   1,
		Name: "Jack#####",
		Age:  99839,
	}
	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Debug().Create(student)
	fmt.Printf("create do nothing while conflict: affect row: %d", result.RowsAffected)
	return result.Error
}
