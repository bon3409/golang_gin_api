package models

import (
	"app/config"
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type Company struct {
	Id    int    `json:"id" gorm:"primaryKey" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Phone string `json:"phone" faker:"phone_number"`
}

// 執行 seeder.go 中的 seeder 的 function
func (company *Company) Seed() {
	err := faker.FakeData(&company)
	if err != nil {
		fmt.Println(err)
	}
	config.DB.Create(&company)
}
