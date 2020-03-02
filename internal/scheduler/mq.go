package scheduler

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Mq interface {
	Connect() error
	Close()
	Send(string) error
}

type MqImpl struct {
	Config *Conf
	name   string

	Conn *amqp.Connection
	Ch   *amqp.Channel
	Q    *amqp.Queue
}

func NewMq(config *Conf, name string) *MqImpl {
	return &MqImpl{Config: config, name: name}
}

func (m *MqImpl) Connect() error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", m.Config.MSG.User, m.Config.MSG.Password, m.Config.MSG.Host, m.Config.MSG.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	} else {
		m.Conn = conn
		ch, err := m.Conn.Channel()
		if err != nil {
			return err
		} else {
			m.Ch = ch

			q, err := m.Ch.QueueDeclare(
				m.name, // name
				false,  // durable
				false,  // delete when unused
				false,  // exclusive
				false,  // no-wait
				nil,    // arguments
			)
			if err == nil {
				m.Q = &q
				return nil
			} else {
				return err
			}

		}
	}
}

func (m *MqImpl) Close() {
	_ = m.Ch.Close()
	_ = m.Conn.Close()
}

func (m *MqImpl) Send(msg string) error {
	return m.Ch.Publish(
		"",       // exchange
		m.Q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}
