package scheduler

type Publisher interface {
	Publish() error
}

type urlPublisher struct {
	Conf    *Conf
	Encoder *EncodeService
}

func NewPublisher(conf *Conf, encoder *EncodeService) *Publisher {
	var s Publisher
	s = &urlPublisher{Conf: conf, Encoder: encoder}
	return &s
}

func (u *urlPublisher) Publish() error {
	var service SourceService
	service = *NewSourceService(u.Conf)

	var mq Mq
	mq = *NewMq(u.Conf, "url_item")
	err := mq.Connect()
	if err != nil {
		return err
	}
	defer mq.Close()
	for item := range service.Get() {
		if item.E != nil {
			return err
		}
		item := item.V.(*Item)
		msg, err := (*u.Encoder).Encode(item)
		if err != nil {
			return err
		}
		err = mq.Send(msg)
		if err != nil {
			return err
		}

	}
	return nil
}
