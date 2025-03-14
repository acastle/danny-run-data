package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/acastle/danny-run-data/competitor"
	"github.com/acastle/danny-run-data/xact"
)

type Event struct {
	Name    string
	Results []RunResult
}

func (e Event) OutputPath() string {
	name := strings.ReplaceAll(e.Name, " ", "_")
	name = strings.ReplaceAll(name, "'", "")
	return filepath.Join(".", fmt.Sprintf("%s.json", name))
}

type RunResult struct {
	Gender   string
	AgeGroup string
	Time     string
}

func main() {
	events := []string{
		"24B67A0B-382D-4034-8571-8053FAF17E25",
		"AFA47BC7-B33A-4B8E-AB97-53A11B137B41",
		"EBE13675-D7D0-4E7B-9859-4CF8C003A20D",
	}
	client := competitor.NewClient("https://api.competitor.com", os.Args[1])

	for _, evt := range events {
		err := ProcessEvent(client, evt)
		if err != nil {
			panic(err)
		}
	}

	xclient := xact.NewClient("https://results.xacte.com")
	result, err := GetResultsXact(xclient, "2531")
	if err != nil {
		panic(err)
	}

	WriteOutput(result)
}

func GetResultsXact(client *xact.Client, event string) (Event, error) {
	totalResults := math.MaxInt
	val := Event{
		Results: []RunResult{},
	}
	for i := 0; i < totalResults; i = i + 1000 {
		result, err := client.GetRunData(event, i)
		if err != nil {
			return val, err
		}

		if totalResults == math.MaxInt {
			totalResults = result.ITotalRecords
		}

		for _, runner := range result.AaData {
			if val.Name == "" {
				val.Name = runner.Event
			}

			duration := time.Millisecond * time.Duration(runner.Clocktime)
			val.Results = append(val.Results, RunResult{
				Gender:   runner.Sex,
				AgeGroup: strconv.Itoa(runner.Age),
				Time:     duration.String(),
			})
		}
	}

	return val, nil
}

func ProcessEvent(client *competitor.Client, event string) error {
	result, err := GetResults(client, event)
	if err != nil {
		return err
	}

	return WriteOutput(result)
}

func GetResults(client *competitor.Client, event string) (Event, error) {
	totalResults := math.MaxInt
	val := Event{
		Results: []RunResult{},
	}
	for i := 0; i < totalResults; i = i + 1000 {
		result, err := client.GetRunData(event, 1000, i)
		if err != nil {
			return val, err
		}

		if val.Name == "" {
			val.Name = result.Data[0].SubeventName
		}
		if totalResults == math.MaxInt {
			totalResults = result.Total
		}

		for _, runner := range result.Data {
			val.Results = append(val.Results, RunResult{
				Gender:   runner.Contact.Gender,
				AgeGroup: runner.AgeGroup,
				Time:     runner.RunTimeConverted,
			})
		}
	}

	return val, nil
}

func WriteOutput(event Event) error {
	b, err := json.Marshal(event.Results)
	if err != nil {
		return err
	}

	log.Println("writing output to: ", event.OutputPath())
	err = os.WriteFile(event.OutputPath(), b, 0600)
	if err != nil {
		return err
	}
	return nil
}
