package scheduler

import "github.com/jinzhu/gorm"

type UrlSource struct {
	gorm.Model
	Url string
}
