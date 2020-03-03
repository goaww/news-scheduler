package scheduler

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
)

type Scheduler interface {
	Run(func())
}
type schedulerImpl struct {
	cronExpr string
}

func (s *schedulerImpl) Run(f func()) {
	c := cron.New(cron.WithSeconds())
	addFunc, err := c.AddFunc(s.cronExpr, f)
	if err != nil {
		failOnError(err, "cannot create scheduler job")
	}
	log.Println("add scheduler function successful", addFunc)
	c.Start()
	defer c.Stop()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func NewScheduler(cronExpr string) *Scheduler {
	var s Scheduler
	s = &schedulerImpl{cronExpr: cronExpr}
	return &s
}

func Handle() {
	conf := NewConf()

	var s Scheduler
	s = *NewScheduler(conf.SDL.Cron)
	var p Publisher

	p = *NewPublisher(conf, NewJsonEncoder())
	s.Run(func() {
		err := p.Publish()
		failOnError(err, "cannot publish message")
	})

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
