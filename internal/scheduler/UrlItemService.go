package scheduler

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/reactivex/rxgo/v2"
)

type SourceService interface {
	Add(Item)

	Get() <-chan rxgo.Item
}

type Item struct {
	Url string
}

func NewItem(url *string) *Item {
	return &Item{Url: *url}
}

type UrlItemServiceImpl struct {
	DB *gorm.DB
}

func NewUrlItemServiceImpl(DB *gorm.DB) *UrlItemServiceImpl {
	return &UrlItemServiceImpl{DB: DB}
}

func (s *UrlItemServiceImpl) Add(source Item) {
	urlSource := NewUrlItem(source.Url)
	s.DB.Create(urlSource)
}

func (s *UrlItemServiceImpl) Get() <-chan rxgo.Item {
	var urlSource []UrlItem
	err := s.DB.Find(&urlSource).Error
	if err != nil {
		c := make(chan rxgo.Item)
		go func() {
			c <- rxgo.Item{
				V: nil,
				E: err,
			}
			close(c)
		}()
		return c
	} else {
		return rxgo.Just(urlSource)().Map(func(ctx context.Context, url interface{}) (ss interface{}, err error) {
			source := url.(UrlItem)
			return NewItem(&source.Url), nil
		}).Observe()
	}
}
