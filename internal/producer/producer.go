package producer

type Producer interface {
	Publish(message map[string]interface{}) error
}

type ProducerClient struct {
	Producer
}

func (p *ProducerClient) Publish(message map[string]interface{}) error {
	return p.Producer.Publish(message)
}