package main

import "fmt"

type YandexDelivery struct{}

func (y *YandexDelivery) RegisterNewBid(orderAddress string, orderCost uint32) {
	fmt.Printf("Yandex placed bid: address=%s cost=%d\n", orderAddress, orderCost)
}
