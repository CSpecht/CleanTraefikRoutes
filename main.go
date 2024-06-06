package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	filename := getEnv("FILENAME", "services.yml")
	interval := getEnv("INTERVAL", "* * * * *")
	period := getEnv("PERIOD", "MONTHLY")
	ticks := getEnv("TICKS", "[3,9]")
	db_host := getEnv("DB_HOST", "localhost")
	db_port := getEnv("DB_PORT", "5432")
	db_username := getEnv("DB_USERNAME", "postgres")
	db_name := getEnv("DB_NAME", "postgres")
	db_pw := getEnv("DB_PW", "postgres")
	db_tablename := getEnv("DB_TABLE", "principles")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		db_host,
		db_port,
		db_username,
		db_name,
		db_pw)

	t, err := parseIntSlice(ticks)
	if err != nil {
		panic(err)
	}

	c := cron.New()
	c.AddFunc(interval, func() { execFunctions(filename, db_tablename, dsn, t, period) })
	c.Start()
	select {}
	// time.Sleep(10 * time.Minute)
	// c.Stop()
}

func execFunctions(filename, tablename, dsn string, ticks []int, period string) {
	now := time.Now()
	if strings.ToLower(period) == "daily" {
		if slices.Contains(ticks, now.Day()) {
			fmt.Println(time.Now())
			truncateTable(tablename, dsn)
			truncateYaml(filename)
		}
	}

	if strings.ToLower(period) == "monthly" {
		if slices.Contains(ticks, int(now.Month())) {
			fmt.Println(time.Now())
			truncateTable(tablename, dsn)
			truncateYaml(filename)
		}
	}
}

func parseIntSlice(str string) ([]int, error) {
	str = strings.TrimSpace(str) // Remove leading/trailing spaces
	if !strings.HasPrefix(str, "[") || !strings.HasSuffix(str, "]") {
		return nil, fmt.Errorf("invalid format: missing square brackets")
	}

	str = strings.TrimPrefix(str, "[") // Remove opening bracket
	str = strings.TrimSuffix(str, "]") // Remove closing bracket

	// Split the string by comma
	parts := strings.Split(str, ",")
	var intSlice []int
	for _, part := range parts {
		part = strings.TrimSpace(part) // Remove spaces around each number
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}

func getEnv(name string, def string) string {
	if len(os.Getenv(name)) > 0 {
		return os.Getenv(name)
	} else {
		return def
	}
}
