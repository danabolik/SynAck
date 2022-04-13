package app

import (
	"SynAck/internal/delivery"
	"SynAck/internal/services/decorators"
	"SynAck/internal/services/producers"
	"SynAck/internal/services/workers"
	"fmt"
	"strconv"
)

type App struct {
	dialer    decorators.NetDialer
	decorator decorators.NetDecorator
	producer  producers.Generator
	delivery  delivery.Http
}

func Run(addr string, grt string) {
	app := new(App)
	decorator := decorators.NetDecorator{Dialer: app.dialer}

	count, _ := strconv.Atoi(grt)
	worker := workers.Worker{Decorator: decorator, Delivery: app.delivery, Producer: app.producer}
	openPs := worker.ScanPorts(addr, count)

	fmt.Println(openPs)
	fmt.Println("Збазиба!")
}
