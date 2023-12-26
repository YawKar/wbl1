package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

func main() {
	cfg := setFlagsAndParseConfig()

	setupDefaultLogger(&cfg.logLevel)

	// this is the main channel,
	// I made the task a bit harder and more entertainment
	// So there are 2 players who *sequentially* pass a ping-pong ball to each other
	gameTable := make(chan Ball)

	wg := sync.WaitGroup{}

	go player("Jules Winnfield", &wg, gameTable, cfg.pingDelay, time.After(cfg.gametime))
	wg.Add(1)
	go player("Vincent Vega", &wg, gameTable, cfg.pingDelay, time.After(cfg.gametime))
	wg.Add(1)

	// drop the ball (e.g. start the game)
	gameTable <- Ball{}

	// wait until both players stop gaming
	wg.Wait()
}

func setupDefaultLogger(logLevel slog.Leveler) {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(l)
}

type Config struct {
	gametime  time.Duration
	pingDelay time.Duration
	logLevel  slog.LevelVar
}

func setFlagsAndParseConfig() *Config {
	c := new(Config)
	flag.DurationVar(&c.gametime, "game-time", 10*time.Second, "set gaming time")
	flag.DurationVar(&c.pingDelay, "ping-delay", 800*time.Millisecond, "set delay between ping-pong passes")
	logLevel := flag.Int64("log-lvl", 0, "set log level")
	flag.Parse()
	c.logLevel.Set(slog.Level(*logLevel))
	return c
}

type Ball struct {
	count int
}

func player(name string, wg *sync.WaitGroup, gameTable chan Ball, delay time.Duration, haltGame <-chan time.Time) {
	for {
		select {
		case ball := <-gameTable:
			time.Sleep(delay)
			ball.count++
			fmt.Printf("%s hits the ball! Ball hit %d times!\n", name, ball.count)
			select {
			case gameTable <- ball: // try to kick the ball if the other player is still on
			default:
				// other player have halted the game
			}
		case <-haltGame:
			fmt.Printf("%s stops playing! Halt!\n", name)
			wg.Done()
			return
		}
	}
}
