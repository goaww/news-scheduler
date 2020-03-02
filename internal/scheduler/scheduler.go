package scheduler

import "log"

func Handle() {
	conf := NewConf()
	db := NewDB(*conf)
	db.Connect()
	defer db.Close()
	service := NewUrlItemServiceImpl(db.Con)

	mq := NewMqImpl(conf, "url_source")
	err := mq.Connect()
	if err == nil {
		defer mq.Close()
		for item := range service.Get() {
			if item.E == nil {
				url := item.V.(*Item).Url
				log.Println("Send:", url)
				err := mq.Send(url)
				if err != nil {
					failOnError(err, "error when send message")
				}
			} else {
				failOnError(item.E, "error when get url item")
			}
		}
	} else {
		failOnError(err, "error when connect mq")

	}
	log.Println("Completed")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
