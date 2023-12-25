package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	cfg := setFlagsAndParseConfig()

	setupDefaultLogger(&cfg.logLevel)

	if err := checkConfigValidity(cfg); err != nil {
		slog.Error("Invalid config", "err", err)
		return
	}

	tasksCh, workerShutdownChs, wg := setupWorkers(cfg.numWorkers)

	// set SIGTERM handler
	gotSigTerm := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		// wait for it
		<-signals
		// notify main thread about SIGTERM
		close(gotSigTerm)
	}()

	// create spam-ticker
	timeToSpam := time.Tick(cfg.tasksRate)

	// main loop
	for {
		select {
		case <-gotSigTerm:
			slog.Info("Shutting down workers...")
			for workerIx, shutdown := range workerShutdownChs {
				slog.Info("Sending shutdown signal to worker", "workerIx", workerIx)
				shutdown <- struct{}{}
			}
			slog.Info("Waiting for workers to gracefully shutdown...")
			wg.Wait()
			return
		case <-timeToSpam:
			slog.Debug("Sending task to pool!")
			tasksCh <- func() {
				fmt.Printf("Task #%d completed\n", rand.Intn(1e9))
			}
		}
	}
}

func setupDefaultLogger(logLevel slog.Leveler) {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(l)
}

type Config struct {
	numWorkers int
	tasksRate  time.Duration
	logLevel   slog.LevelVar
}

func setFlagsAndParseConfig() *Config {
	c := new(Config)
	flag.IntVar(&c.numWorkers, "workers", 0, "set number of workers")
	flag.DurationVar(&c.tasksRate, "tasks-rate", 100*time.Millisecond, "set rate of tasks generation (e.g. 1 task per 100ms)")
	logLevel := flag.Int64("log-lvl", 0, "set log level")
	flag.Parse()
	c.logLevel.Set(slog.Level(*logLevel))
	return c
}

func checkConfigValidity(cfg *Config) error {
	if cfg.numWorkers <= 0 {
		return fmt.Errorf("invalid number of workers, should be > 0, got: %d", cfg.numWorkers)
	}
	return nil
}

type Task = func()
type ShutdownCh = chan<- struct{}

func setupWorkers(numWorkers int) (chan<- Task, []chan<- struct{}, *sync.WaitGroup) {
	tasksCh := make(chan Task, numWorkers)
	shutdownChs := make([]chan struct{}, numWorkers)
	wg := sync.WaitGroup{}
	for workerIx := 0; workerIx < numWorkers; workerIx++ {
		shutdownChs[workerIx] = make(chan struct{}, 1)
		go worker(tasksCh, shutdownChs[workerIx], &wg)
		wg.Add(1)
	}
	return tasksCh, chansToWriters(shutdownChs...), &wg
}

func chansToWriters[T any](chans ...chan T) []chan<- T {
	writers := make([]chan<- T, len(chans))
	for i := range chans {
		writers[i] = chans[i]
	}
	return writers
}

func worker(tasks <-chan func(), shutdown <-chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case task := <-tasks:
			task()
		case <-shutdown:
			wg.Done()
			return
		}
	}
}
