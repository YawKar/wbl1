package main

import (
	"fmt"
	"log/slog"
)

type Order interface {
	DeliveryAddress() string
	BidCost() uint32
}

type BuyNowPayLaterOrder struct {
	address string
	cost    uint32
}

func (o *BuyNowPayLaterOrder) DeliveryAddress() string {
	return o.address
}

func (o *BuyNowPayLaterOrder) BidCost() uint32 {
	return o.cost
}

// Adapter interface
type PartnerDeliveryServiceAdapter interface {
	PlaceOrder(order Order) bool
}

// Samokat adapter
type SamokatAdapter struct {
	s *Samokat // adaptee
}

func (s *SamokatAdapter) PlaceOrder(order Order) bool {
	if err := s.s.MakeNewDrop(order.BidCost(), order.DeliveryAddress()); err != nil {
		slog.Error("Samokat refused", "err", err)
		return false
	}
	return true
}

// Yandex adapter
type YandexAdapter struct {
	y *YandexDelivery
}

func (y *YandexAdapter) PlaceOrder(order Order) bool {
	y.y.RegisterNewBid(order.DeliveryAddress(), order.BidCost())
	return true
}

// Now it is possible to use them under `PartnerDeliveryServiceAdapter` interface
//
// Example output:
// 2023/12/29 11:03:00 ERROR Samokat refused err="lol, i ain't gonn' do tha', man"
// Yandex placed bid: address=Молочный переулок, Каштановый дом cost=588
// Order was placed: true
func main() {
	partners := []PartnerDeliveryServiceAdapter{
		&SamokatAdapter{&Samokat{}},
		&YandexAdapter{&YandexDelivery{}},
	}
	bnplOrder := BuyNowPayLaterOrder{
		address: "Молочный переулок, Каштановый дом",
		cost:    588,
	}
	placed := false
	for _, partner := range partners {
		if partner.PlaceOrder(&bnplOrder) {
			placed = true
			break
		}
	}
	fmt.Printf("Order was placed: %t\n", placed)
}
