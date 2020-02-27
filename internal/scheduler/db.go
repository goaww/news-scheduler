package scheduler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type DB struct {
	Config Conf
	Con    *gorm.DB
}

func NewDB(config Conf) *DB {
	return &DB{Config: config}
}

func (d *DB) Connect() {
	db, err := connectDB(d.Config)
	if err == nil {
		log.Println("connect database successful")

		err = db.AutoMigrate(UrlItem{}).Error
		if err != nil {
			log.Fatal("failed to migrate table todo")
		}
		d.Con = db
	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}
func (d *DB) Close() {
	_ = d.Con.Close()
}
func connectDB(config Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
