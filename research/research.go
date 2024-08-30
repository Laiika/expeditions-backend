package main

import (
	"context"
	"db_cp_6/config"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"db_cp_6/pkg/postgres"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	client    postgres.Client
	clientInd postgres.Client
	srvc      *service.ExpeditionService
)

const N = 100

func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	var err error
	clientInd, err = postgres.NewClient(context.Background(), 3, &cfg.Test)
	if err != nil {
		log.Fatal(err)
	}
	defer clientInd.Close()

	client, err = postgres.NewClient(context.Background(), 3, &cfg.Test)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repos := repo.NewRepositories()
	srvc = service.NewExpeditionService(repos.ExpeditionRepo)

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	step := 0
	for i := 100; i <= 10000; i += step {
		fmt.Println("Expeditions count: ", i)

		resultTimeInd, errorCountInd, err := researchGetWithIndex(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research with index end ok!")
		}

		resultTime, errorCount, err := researchGetWithoutIndex(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research without index end ok!")
		}

		_, err = file.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(resultTimeInd) + " " + strconv.Itoa(errorCountInd) + " ")
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.WriteString(strconv.Itoa(resultTime) + " " + strconv.Itoa(errorCount) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}

		if i == 100 {
			step = 400
		} else if i == 500 {
			step = 500
		} else if i == 1000 {
			step = 1000
		}
	}
}

func setupData(count int, cl postgres.Client) error {
	if err := truncateTables(context.Background(), cl); err != nil {
		return err
	}

	path := fmt.Sprintf("./research/scripts/%s.sql", strconv.Itoa(count))
	text, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if _, err = cl.Exec(context.Background(), string(text)); err != nil {
		return err
	}
	return nil
}

func researchGetWithIndex(count int) (int, int, error) {
	err := dropIndex(context.Background(), clientInd)
	if err != nil {
		return 0, 0, err
	}
	err = setupData(count, clientInd)
	if err != nil {
		return 0, 0, err
	}
	err = createIndex(context.Background(), clientInd)
	if err != nil {
		return 0, 0, err
	}

	var result int64
	var errorCount int64
	var successCount int64

	// for i := 0; i < N; i++ {
	for successCount != N {
		locationId := rand.Intn(100) + 1

		duration, err := srvc.GetLocationExpeditionsTime(context.Background(), clientInd, locationId)

		if err != nil {
			errorCount += 1
		} else {
			successCount += 1
			result += duration.Nanoseconds()
		}
	}

	fmt.Println("итоговое время: ", result/N)
	fmt.Println("итого ошибок:", errorCount)
	return int(result), int(errorCount), err
}

func researchGetWithoutIndex(count int) (int, int, error) {
	err := dropIndex(context.Background(), client)
	if err != nil {
		return 0, 0, err
	}
	err = setupData(count, client)
	if err != nil {
		return 0, 0, err
	}

	var result int64
	var errorCount int64
	var successCount int64

	for successCount != N {
		// for i := 0; i < N; i++ {
		locationId := rand.Intn(100) + 1

		duration, err := srvc.GetLocationExpeditionsTime(context.Background(), clientInd, locationId)

		if err != nil {
			errorCount += 1
		} else {
			successCount++
			result += duration.Nanoseconds()
		}
	}

	fmt.Println("итоговое время - ", result/N)
	fmt.Println("ошибок - ", errorCount)
	return int(result), int(errorCount), err
}

func truncateTables(ctx context.Context, client postgres.Client) error {
	q := `
		TRUNCATE locations, expeditions RESTART IDENTITY
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func dropIndex(ctx context.Context, client postgres.Client) error {
	q := `
		DROP INDEX IF EXISTS idx_expeditions_location_id;
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func createIndex(ctx context.Context, client postgres.Client) error {
	q := `
		CREATE INDEX idx_expeditions_location_id ON expeditions(location_id);
	`

	if _, err := client.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}
