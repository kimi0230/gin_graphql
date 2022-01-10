package main

import (
	"flag"
	"fmt"
	"gin_graphql/app/models"
	db "gin_graphql/config/mysql"
)

var (
	h bool
	m string
)

var tables []interface{} = []interface{}{
	// models.User{},
	// models.Meetup{},
	models.Staff{},
	models.StaffRole{},
	models.Role{},
	models.RolePermission{},
	models.Permission{},
}

func init() {
	flag.BoolVar(&h, "h", false, "migration help")
	flag.StringVar(&m, "m", "default", "migrate command: auto/drop/create/refresh")
}

func auto() {
	db.GormDB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(tables...)
}
func drop() {
	db.GormDB.Set("gorm:table_options", "CHARSET=utf8mb4").DropTable(tables...)
}
func refresh() {
	auto()
	drop()
	auto()
}

func migrate() {
	fmt.Println("--- Start Migrate ---")
	switch m {
	case "auto":
		fmt.Println("--- AutoMigrate ---")
		auto()
	case "drop":
		fmt.Println("--- DropTable ---")
		drop()
	case "refresh":
		fmt.Println("--- FreshTable ---")
		refresh()
	case "create":
		fmt.Println("create (not yet)")
	default:
		fmt.Println("Please input auto/drop/create/refresh")
	}
	fmt.Println("--- End Migrate ---")
}
func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	migrate()
}
