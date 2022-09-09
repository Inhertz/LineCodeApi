package main

import (
	"LineCodeApi/internal/adapters/db"
	"LineCodeApi/internal/adapters/rpc"
	"LineCodeApi/internal/adapters/web"
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/domain"
	"fmt"
	"log"
	"os"
	"sync"
)

var dbConnStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
	os.Getenv("DB_SERVER"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_SSL_MODE"),
	os.Getenv("DB_TIME_ZONE"))

func main() {

	dbAdapter, err := db.NewAdapter(dbConnStr)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	domainLogic := domain.New()

	appApi := application.NewApplication(dbAdapter, domainLogic)

	webAdapter := web.NewAdapter(appApi)
	grpcAdapter := rpc.NewAdapter(appApi)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go webAdapter.RunAsync(wg)
	go grpcAdapter.RunAsync(wg)

	wg.Wait()
}
