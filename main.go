package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/jexodusmercado/inventory-system-be/cmd"
)

func main() {
	execCtx, execCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer execCancel()

	go func() {
		<-execCtx.Done()
		log.Println("received graceful shutdown signal")
	}()

	// command is expected to obey the cancellation signal on execCtx and
	// block while it is running
	if err := cmd.RootCommand().ExecuteContext(execCtx); err != nil {
		log.Fatalln(err)
	}

}