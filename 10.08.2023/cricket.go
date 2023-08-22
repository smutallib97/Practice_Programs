package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type PlayersData struct {
	PlayerName    string
	Period        string
	MatchesPlayed int
	TotalScore    int
	AvgScore      int
}

func FilterFunction(records []PlayersData, operation func(PlayersData) bool) []PlayersData {
	var result []PlayersData
	for _, value := range records {
		if operation(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	file, err := os.Open("CricketData.csv")
	if err != nil {
		fmt.Println("Error while opening csv file:", err)
		return
	}
	defer file.Close() // Close the file when done
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading the csv file:", err)
		return
	}

	var playersData []PlayersData
	// Convert CSV records to PlayersData structs
	for _, record := range records {
		matchesPlayed, _ := strconv.Atoi(record[2])
		totalScore, _ := strconv.Atoi(record[3])
		avgScore, _ := strconv.Atoi(record[4])
		player := PlayersData{
			PlayerName:    record[0],
			Period:        record[1],
			MatchesPlayed: matchesPlayed,
			TotalScore:    totalScore,
			AvgScore:      avgScore,
		}
		playersData = append(playersData, player)
	}

	threshold := 100
	aboveThreshold := func(p PlayersData) bool {
		return p.MatchesPlayed >= threshold
	}

	aboveThresholdMatches := FilterFunction(playersData, aboveThreshold)

	fmt.Println("Players who played greater than 100 matches:")
	for _, match := range aboveThresholdMatches {
		fmt.Printf("%s:%d\n", match.PlayerName, match.MatchesPlayed)
	}
}
