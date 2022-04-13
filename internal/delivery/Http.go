package delivery

type Delivery interface {
	GetTcpNetwork() string
}

type Http struct {
}

func (http Http) GetTcpNetwork() string {
	return "tcp"
}
