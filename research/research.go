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
	srvc      *service.MemberService
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

	client, err = postgres.NewClient(context.Background(), 3, &cfg.Admin)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repos := repo.NewRepositories()
	srvc = service.NewMemberService(repos.MemberRepo)

	step := 0

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 10; i <= 1000; i += step {
		fmt.Println("Members count: ", i)

		resultTimeInd, errorCountInd, err := researchGetExpeditionMembersWithIndices(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research with indices end ok!")
		}

		resultTime, errorCount, err := researchGetExpeditionMembersWithoutIndices(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research without indices end ok!")
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

		if i < 100 {
			i += 10
		} else if i == 100 {
			step = 100
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

func researchGetExpeditionMembersWithIndices(count int) (int, int, error) {
	err := dropIndices(context.Background(), clientInd)
	if err != nil {
		return 0, 0, err
	}
	err = setupData(count, clientInd)
	if err != nil {
		return 0, 0, err
	}
	err = createIndices(context.Background(), clientInd)
	if err != nil {
		return 0, 0, err
	}

	var result int64
	var errorCount int64
	var successCount int64

	// for i := 0; i < N; i++ {
	for successCount != N {
		expeditionId := rand.Intn(100) + 1

		duration, err := srvc.GetExpeditionMembersTime(context.Background(), clientInd, expeditionId)

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

func researchGetExpeditionMembersWithoutIndices(count int) (int, int, error) {
	err := dropIndices(context.Background(), clientInd)
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
		expeditionId := rand.Intn(100) + 1

		duration, err := srvc.GetExpeditionMembersTime(context.Background(), clientInd, expeditionId)

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
		TRUNCATE members, locations, expeditions, expeditions_members
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func dropIndices(ctx context.Context, client postgres.Client) error {
	q := `
		DROP INDEX IF EXISTS idx_expeditions_members_member_id;
		DROP INDEX IF EXISTS idx_expeditions_members_expedition_id;
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func createIndices(ctx context.Context, client postgres.Client) error {
	q := `
		CREATE INDEX idx_expeditions_members_member_id ON expeditions_members(member_id);
		CREATE INDEX idx_expeditions_members_expedition_id ON expeditions_members(expedition_id);
	`

	if _, err := client.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}
