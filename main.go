package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"dockerized-csv/models"
)

func main() {
	// Database connection details
	dsn := "host=db user=postgres password=Welcome@@1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Read the CSV file
	file, err := os.Open("db/fixlets.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Parse CSV records and insert into the database
	for _, record := range records[1:] { // Skip header row
		siteID, _ := strconv.Atoi(record[0])                // Convert SiteID to int
		fixletID, _ := strconv.Atoi(record[1])              // Convert FixletID to int
		relevantComputerCount, _ := strconv.Atoi(record[4]) // Convert RelevantComputerCount to int

		user := models.User{
			SiteID:                uint(siteID),
			FixletID:              uint(fixletID),
			Name:                  record[2],
			Criticality:           record[3],
			RelevantComputerCount: relevantComputerCount,
		}
		db.Create(&user)
	}

	fmt.Println("CSV data has been successfully inserted into the database!")
}
