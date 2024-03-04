package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

var (
	flagTick = flag.Duration("tick", 10*time.Second, "how often to print time message")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("ERROR: no command provided") //nolint:forbidigo
		os.Exit(1)
	}
	program := args[0]
	if len(args) > 1 {
		args = args[1:]
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancelFunc()

	cmd := exec.CommandContext(ctx, program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		fmt.Println(err) //nolint:forbidigo
		os.Exit(1)
	}

	start := time.Now()
	tick := time.NewTicker(*flagTick)
	defer tick.Stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case tt := <-tick.C:
				fmt.Printf("  command has taken %v\n", tt.Sub(start).Truncate(time.Second)) //nolint:forbidigo
			}
		}
	}()
	defer func() {
		fmt.Printf("  command took %v\n", time.Since(start).Truncate(time.Second)) //nolint:forbidigo
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err) //nolint:forbidigo
		os.Exit(1)
	}
}
