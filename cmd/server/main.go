package main

import (
	"log"
	"odesch.com/odesch/grpc-app/internal/db"
	"odesch.com/odesch/grpc-app/internal/rocket"
)

func Run() error {
	// responsible for initializing and
	// starting our gRPC server

	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)

	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
