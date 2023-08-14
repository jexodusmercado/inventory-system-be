package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	execCtx, execCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer execCancel()

	go func() {
		<-execCtx.Done()
		log.Println("received graceful shutdown signal")
	}()

}