package main

import (
	"app/config"
	"app/models"
	"flag"
	"fmt"
	"reflect"
	"strings"
)

// 需要執行 seeder 的 model 必須實作裡面的 function
type Seeder interface {
	Seed()
}

func main() {
	model := flag.String("model", "", "This is model name")
	count := flag.Int("count", 10, "This is seeder count") // 如果沒有輸入 seed 數量，就是預設新增 10 個
	flag.Parse()

	if *model == "" {
		fmt.Println("需要輸入 '-model=' 的資料")
		return
	}

	name := strings.ToLower(*model)

	// TODO:未來要新增 seeder 的話，要在下面新增 Case
	switch name {
	case "user":
		user := new(models.User)
		runSeeder(user, *count)
	case "company":
		company := new(models.Company)
		runSeeder(company, *count)
	}
}

// 執行 seed
func runSeeder(model Seeder, count int) {
	fmt.Printf("新增 %d 個 %s 的資料\n", count, reflect.TypeOf(model))
	config.DbConfiguration()
	for i := 1; i <= count; i++ {
		model.Seed()
	}
	fmt.Println("完成 seed !")
}
