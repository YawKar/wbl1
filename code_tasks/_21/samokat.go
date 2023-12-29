package main

import (
	"errors"
	"fmt"
)

type Samokat struct{}

func (s *Samokat) MakeNewDrop(cost uint32, address string) error {
	if cost < 1000 {
		return errors.New("lol, i ain't gonn' do tha', man")
	}
	fmt.Printf("Samokat made new drop: cost is %d, address is %s\n", cost, address)
	return nil
}
