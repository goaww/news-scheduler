package scheduler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Handle() {
	//load config
	var config Conf
	config.getConf()

	db, err := connectDB(config)
	if err == nil {
		defer db.Close()
		log.Println("connect database successful")

		err = db.AutoMigrate(UrlSource{}).Error
		if err != nil {
			log.Fatal("failed to migrate table todo")
		}

	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}
func connectDB(config Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
