package scheduler

import "github.com/jinzhu/gorm"

type UrlItem struct {
	gorm.Model
	Url string
}

func NewUrlItem(url string) *UrlItem {
	return &UrlItem{Url: url}
}
