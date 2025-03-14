package competitor

import "time"

type EventDataPayload struct {
	Total int    `json:"total"`
	Limit int    `json:"limit"`
	Skip  int    `json:"skip"`
	Data  []Data `json:"data"`
}

type Contact struct {
	FullName string `json:"FullName"`
	Gender   string `json:"Gender"`
}
type Country struct {
	Iso2 string `json:"ISO2"`
}

type Subevent struct {
	SubEvent string `json:"SubEvent"`
}
type Data struct {
	RunTimeConverted              string    `json:"RunTimeConverted"`
	SwimTimeConverted             string    `json:"SwimTimeConverted"`
	BikeTimeConverted             string    `json:"BikeTimeConverted"`
	FinishTimeConverted           string    `json:"FinishTimeConverted"`
	Transition1TimeConverted      string    `json:"Transition1TimeConverted"`
	Transition2TimeConverted      string    `json:"Transition2TimeConverted"`
	CountryISO2                   string    `json:"CountryISO2"`
	SubeventName                  string    `json:"SubeventName"`
	ResultID                      string    `json:"ResultId"`
	ContactID                     string    `json:"ContactId"`
	AgeGroup                      string    `json:"AgeGroup"`
	CountryRepresentingISONumeric int       `json:"CountryRepresentingISONumeric"`
	BibNumber                     int       `json:"BibNumber"`
	EventStatus                   string    `json:"EventStatus"`
	SwimTime                      int       `json:"SwimTime"`
	Transition1Time               int       `json:"Transition1Time"`
	BikeTime                      int       `json:"BikeTime"`
	Transition2Time               int       `json:"Transition2Time"`
	RunTime                       int       `json:"RunTime"`
	FinishTime                    int       `json:"FinishTime"`
	FinishRankGroup               int       `json:"FinishRankGroup"`
	FinishRankGender              int       `json:"FinishRankGender"`
	FinishRankOverall             int       `json:"FinishRankOverall"`
	SwimRankGroup                 int       `json:"SwimRankGroup"`
	SwimRankGender                int       `json:"SwimRankGender"`
	SwimRankOverall               int       `json:"SwimRankOverall"`
	BikeRankGroup                 int       `json:"BikeRankGroup"`
	BikeRankGender                int       `json:"BikeRankGender"`
	BikeRankOverall               int       `json:"BikeRankOverall"`
	RunRankGroup                  int       `json:"RunRankGroup"`
	RunRankGender                 int       `json:"RunRankGender"`
	RunRankOverall                int       `json:"RunRankOverall"`
	SyncDate                      time.Time `json:"SyncDate"`
	StatusCode                    int       `json:"StatusCode"`
	SubEventID                    string    `json:"SubEventId"`
	RankPoints                    any       `json:"RankPoints"`
	Contact                       Contact   `json:"Contact"`
	Country                       Country   `json:"Country"`
	Subevent                      Subevent  `json:"Subevent"`
}
