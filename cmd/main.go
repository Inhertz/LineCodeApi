package main

import (
	"LineCodeApi/internal/adapters/db"
	"LineCodeApi/internal/adapters/rpc"
	"LineCodeApi/internal/adapters/web"
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/domain"
	"LineCodeApi/internal/core/models"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// getEnv returns the value of the environment variable named by key,
// or fallback when the variable is unset or empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbAdapter, err := db.NewAdapter[models.Manchester](fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIME_ZONE")))

	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	domainLogic := domain.New()

	appApi := application.NewApplication(dbAdapter, domainLogic)

	webAdapter := web.NewAdapter(appApi, getEnv("WEB_PORT", "8080"))
	grpcAdapter := rpc.NewAdapter(appApi, getEnv("GRPC_PORT", "9000"))

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go webAdapter.RunAsync(ctx, wg)
	go grpcAdapter.RunAsync(ctx, wg)

	wg.Wait()
	log.Println("all servers stopped, exiting")
}
