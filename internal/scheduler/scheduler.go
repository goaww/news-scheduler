package scheduler

import (
	"log"
)

type Scheduler interface {
	Execute(string)
}
type SchImpl struct {
	Conf    *Conf
	Encoder *EncodeService
}

func NewSchImpl(conf *Conf, encoder *EncodeService) *SchImpl {
	return &SchImpl{Conf: conf, Encoder: encoder}
}

func (s *SchImpl) Execute(string) {
	err := s.handle()
	if err != nil {
		failOnError(err, "cannot execute scheduler")
	} else {
		log.Println("completed")
	}
}

func (s *SchImpl) handle() error {
	service := NewUrlItemServiceImpl(s.Conf)

	mq := NewMq(s.Conf, "url_item")
	err := mq.Connect()
	if err == nil {
		defer mq.Close()
		for item := range service.Get() {
			if item.E == nil {
				item := item.V.(*Item)
				msg, err := (*s.Encoder).Encode(item)
				if err != nil {
					return err
				} else {
					err := mq.Send(msg)
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

func Handle() {
	impl := NewSchImpl(NewConf(), NewJsonEncoder())
	impl.Execute("")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
