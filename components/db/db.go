package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Driver db init driver config
type Driver struct {
	User       string `json:"user"        mapstructure:"user"`
	Password   string `json:"password"    mapstructure:"password"`
	Host       string `json:"host"        mapstructure:"host"`
	Port       string `json:"port"        mapstructure:"port"`
	Name       string `json:"name"        mapstructure:"name"`        // database name
	DisableLog bool   `json:"disable_log" mapstructure:"disable_log"` // disable log
}

// Register init db connect
func Register(driver Driver) {
	db = &gorm.DB{}

	var err error

	db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", driver.User, driver.Password, driver.Host, driver.Port, driver.Name))

	if err != nil {
		panic(err)
	}

	if !driver.DisableLog {
		db.LogMode(true)
	}
}

// Get return db
func Get() *gorm.DB {
	return db
}
