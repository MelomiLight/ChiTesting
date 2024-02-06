package main

import (
	"chicha/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	app := application.New(application.LoadConfig()) //First of all initialize application

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx) //start application
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
	
}
