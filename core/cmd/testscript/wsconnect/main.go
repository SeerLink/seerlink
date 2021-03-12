package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/SeerLink/seerlink/core/services"
	"github.com/SeerLink/seerlink/core/services/eth"
	"github.com/SeerLink/seerlink/core/store/models"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cb := func(log models.Log) {}
	c, err := eth.NewClient("ws://localhost:8546")
	panicErr(err)
	err = c.Dial(context.Background())
	panicErr(err)
	sub, err := services.NewManagedSubscription(c, ethereum.FilterQuery{}, cb)
	panicErr(err)
	fmt.Println(sub)
	time.Sleep(30 * time.Second)
	// While this is connected run:
	// docker stop <id of node container>
	// docker start <id of node container>
	// and ensure you see reconnection logs.
}
