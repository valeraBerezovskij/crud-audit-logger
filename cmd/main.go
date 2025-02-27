package main

import (
	"context"
	"crud-audit/internal/config"
	"fmt"
	"log"
	"time"

	"crud-audit/internal/repository"
	"crud-audit/internal/server"
	"crud-audit/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Config initialization
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// Context initialization
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mongo connection
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	// Database initialization
	db := dbClient.Database(cfg.DB.Database)

	//Repository, service, server initialization
	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	// Server start
	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
