package integrational

import (
	"context"
	"db_cp_6/config"
	"db_cp_6/internal/repo"
	"db_cp_6/pkg/logger"
	"db_cp_6/pkg/postgres"
	"os"
	"testing"
)

var (
	pgClient postgres.Client
	pgRepo   = repo.NewRepositories()
)

func setup() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	var err error
	pgClient, err = postgres.NewClient(context.Background(), 3, &cfg.Test)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	pgClient.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
