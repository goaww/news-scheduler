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
			log.Println(item.V.(*Item).Url)
			if item.E != nil {

			} else {
				failOnError(item.E, "error when get url item")
			}
		}
	} else {
		failOnError(err, "error when connect mq")

	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
