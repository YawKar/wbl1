package _09

import (
	"fmt"
	"slices"
	"testing"
)

const numbersLen = 100

var numbers []int

func init() {
	numbers = make([]int, 0, numbersLen)
	for i := 0; i < numbersLen; i++ {
		numbers = append(numbers, i)
	}
}

func TestDoubleNumbersUnbuffered(t *testing.T) {
	producer := make(chan int)
	consumer := make(chan int)
	double := func(v int) int { return v * 2 }
	go MapConnect(producer, double, consumer)

	// run producer
	go func() {
		defer close(producer)
		for _, v := range numbers {
			producer <- v
		}
	}()

	// consume doubled numbers from consumer
	consumed := make([]int, 0, numbersLen)
	for v := range consumer {
		consumed = append(consumed, v)
	}

	if !slices.EqualFunc(numbers, consumed, func(n int, c int) bool { return double(n) == c }) {
		t.Errorf("numbers were not doubled after 1-step map connection:\n%v\n != \n%v", numbers, consumed)
	}
}

func TestStringifyAndParseUnbuffered(t *testing.T) {
	producer := make(chan int)
	stringify := func(v int) string { return fmt.Sprintf("%d", v) }
	stringed := make(chan string)
	go MapConnect(producer, stringify, stringed)
	parse := func(v string) (out int) {
		if _, err := fmt.Sscan(v, &out); err != nil {
			out = 0
		}
		return
	}
	parsed := make(chan int)
	go MapConnect(stringed, parse, parsed)

	// run producer
	go func() {
		defer close(producer)
		for _, v := range numbers {
			producer <- v
		}
	}()

	// consume final numbers
	consumed := make([]int, 0, numbersLen)
	for v := range parsed {
		consumed = append(consumed, v)
	}

	if !slices.Equal(numbers, consumed) {
		t.Errorf("numbers are not equal after 2-step map connection:\n%v\n != \n%v", numbers, consumed)
	}
}

func TestConnectNumbersUnbuffered(t *testing.T) {
	producer := make(chan int)
	consumer := make(chan int)
	go Connect(producer, consumer) // bind 'em

	// run producer
	go func() {
		defer close(producer)
		for _, v := range numbers {
			producer <- v
		}
	}()

	// consume numbers from consumer
	consumed := make([]int, 0, numbersLen)
	for v := range consumer {
		consumed = append(consumed, v)
	}

	if !slices.Equal(numbers, consumed) {
		t.Errorf("numbers are not equal after 1-step connection:\n%v\n != \n%v", numbers, consumed)
	}
}
