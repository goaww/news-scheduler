package scheduler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB interface {
	Connect() (*gorm.DB, error)
}
type dBImpl struct {
	Config *Conf
}

func NewDB(config *Conf) *DB {
	var impl DB
	impl = &dBImpl{Config: config}
	return &impl
}

func (d *dBImpl) Connect() (*gorm.DB, error) {
	db, err := connectDB(d.Config)
	if err == nil {
		err = db.AutoMigrate(UrlItem{}).Error
		if err != nil {
			return nil, err
		} else {
			return db, nil
		}
	} else {
		return nil, err
	}
}
func connectDB(config *Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
