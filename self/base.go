package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	dsn = "root:root@tcp(localhost:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
)

type Student struct {
	ID        uint
	Name      string
	Age       uint
	DeletedAt gorm.DeletedAt
}

func Rows(db *gorm.DB) error {
	var err error
	rows, err := db.Model(&Student{}).Debug().Select("*").Rows()
	if err != nil {
		fmt.Printf("rows fail: %v\n", err)
	}
	defer rows.Close()
	var results []Student
	for rows.Next() {
		var s Student
		if err = db.ScanRows(rows, &s); err != nil {
			fmt.Printf("rows should get no error, but got %v", err)
		}
		results = append(results, s)
	}
	fmt.Printf("rows result: %v\n", results)
	return err
}

func Row(db *gorm.DB) error {
	var err error
	var student Student
	row := db.Model(&Student{}).Debug().Select("id, name, age, deleted_at").Row()
	err = row.Scan(&student.ID, &student.Name, &student.Age, &student.DeletedAt)
	if err != nil {
		fmt.Printf("row should get no error, but got %v", err)
	}
	fmt.Printf("row result: %v\n", student)
	return err
}

func First(db *gorm.DB) error {
	s := &Student{}
	result := db.Debug().Where("id=?", 28).First(s)
	fmt.Printf("student: %v\n", s)
	return result.Error
}

func Take(db *gorm.DB) error {
	s := &Student{
		ID: 9999,
	}
	result := db.Debug().First(s)
	fmt.Printf("student: %v\n", s)
	return result.Error
}

func Find(db *gorm.DB) error {
	s := []Student{
		{
			ID: 1,
		},
		{
			ID: 28,
		},
	}
	result := db.Debug().Find(&s)
	fmt.Printf("student: %v\n", s)
	return result.Error
}

func FindInBatches(db *gorm.DB) error {
	//s := []Student{
	//	{
	//		ID: 1,
	//	},
	//	{
	//		ID: 28,
	//	},
	//	{
	//		ID: 42,
	//	},
	//}
	//result := db.Debug().FindInBatches(&s, 2)
	//fmt.Printf("student: %v\n", s)
	//return result.Error
	return nil
}

// GetAllStudents 查询所有用户
func GetAllStudents(db *gorm.DB) error {
	var student []*Student
	err := db.Debug().Find(&student).Error
	for i, user := range student {
		fmt.Printf("student index: %d, info: %v\n", i, user)
	}
	return err
}

func BatchCreate(db *gorm.DB) error {
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
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func CreateInBatches(db *gorm.DB) error {
	students := []*Student{
		{
			Name: "Jack",
			Age:  38,
		},
		{
			Name: "Mack",
			Age:  34,
		},
		{
			Name: "Marry",
			Age:  18,
		},
	}
	result := db.Debug().CreateInBatches(students, 2)
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func Save(db *gorm.DB) error {
	students := []*Student{
		{
			ID:   39,
			Name: "Jack",
			Age:  38,
		},
		{
			ID:   28,
			Name: "Mack",
			Age:  34,
		},
	}
	result := db.Debug().Save(students)
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func Update(db *gorm.DB) error {
	result := db.Debug().Model(&Student{}).Where("id=1").Update("age", 18)
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func Updates(db *gorm.DB) error {
	result := db.Debug().Model(&Student{
		ID: 28,
	}).Updates(Student{
		Name: "Jack1",
		Age:  38,
	})
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func delete(db *gorm.DB) error {
	student := &Student{
		ID: 1,
	}
	result := db.Debug().Delete(student)
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func deleteAll(db *gorm.DB) error {
	result := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Debug().Delete(&Student{})
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func deleteWithReturning(db *gorm.DB) error {
	student := &Student{
		ID: 1,
	}
	result := db.Debug().Clauses(clause.Returning{}).Delete(student)
	fmt.Printf("effect row: %d, student: %v", result.RowsAffected, student)
	return result.Error
}

func deleteByPrimaryKey(db *gorm.DB) error {
	result := db.Debug().Delete(&Student{}, 1)
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}

func deleteBySQL(db *gorm.DB) error {
	result := db.Debug().Where("`age`=18").Delete(&Student{})
	fmt.Printf("effect row: %d\n", result.RowsAffected)
	return result.Error
}
