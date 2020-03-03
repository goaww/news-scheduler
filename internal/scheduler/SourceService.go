package scheduler

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

type SourceService interface {
	Add(Item) error

	Get() <-chan rxgo.Item
}

type sourceServiceImpl struct {
	Config *Conf
}

func NewSourceService(config *Conf) *SourceService {
	var s SourceService
	s = &sourceServiceImpl{Config: config}
	return &s
}

func (s *sourceServiceImpl) Add(source Item) error {
	db := *NewDB(s.Config)
	conn, err := db.Connect()
	if err != nil {
		return err
	} else {
		defer conn.Close()

		urlSource := NewUrlItem(source.Url)
		conn.Create(urlSource)
		return nil
	}
}

func (s *sourceServiceImpl) Get() <-chan rxgo.Item {
	db := *NewDB(s.Config)
	conn, err := db.Connect()
	if err != nil {
		return makeErrorItem(err)
	} else {
		defer conn.Close()

		var urlSource []UrlItem
		err := conn.Find(&urlSource).Error
		if err != nil {
			return makeErrorItem(err)
		} else {
			return rxgo.Just(urlSource)().Map(func(ctx context.Context, url interface{}) (ss interface{}, err error) {
				source := url.(UrlItem)
				return NewItem(&source.Url), nil
			}).Observe()
		}
	}

}

func makeErrorItem(err error) <-chan rxgo.Item {
	c := make(chan rxgo.Item)
	go func() {
		c <- rxgo.Item{
			V: nil,
			E: err,
		}
		close(c)
	}()
	return c
}
