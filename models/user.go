package models

import (
	"app/config"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id        int    `json:"id" faker:"-"`
	FirstName string `json:"first_name" faker:"first_name"`
	LastName  string `json:"last_name" faker:"last_name"`
	Email     string `json:"email" faker:"email"`
	CompanyID int    `json:"company_id" gorm:"default:null"`
	Company   Company
	CreatedAt time.Time `json:"created_at" faker:"-"`
	UpdatedAt time.Time `json:"updated_at" faker:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// do something...
	return
}

func (user *User) AfterSave(tx *gorm.DB) (err error) {
	// do something...
	return
}

// 執行 seeder.go 中的 seeder 的 function
func (user *User) Seed() {
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	config.DB.Create(&user)
}
