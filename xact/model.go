package xact

type Result struct {
	SEcho                int      `json:"sEcho"`
	Timestamp            int64    `json:"timestamp"`
	ITotalRecords        int      `json:"iTotalRecords"`
	ITotalDisplayRecords int      `json:"iTotalDisplayRecords"`
	AaData               []AaData `json:"aaData"`
}

type Distance struct {
	ID          int    `json:"id"`
	Distance    int    `json:"distance"`
	Label       string `json:"label"`
	IsMetric    bool   `json:"is_metric"`
	LegDistance int    `json:"leg_distance"`
}

type Split struct {
	ID           int      `json:"id"`
	Elapsed      int      `json:"elapsed"`
	DeltaNet     int      `json:"delta_net"`
	DeltaLeg     int      `json:"delta_leg"`
	DeltaTrans   any      `json:"delta_trans"`
	TimeMs       int      `json:"time_ms"`
	Locktime     int      `json:"locktime"`
	Pace         any      `json:"pace"`
	Distance     Distance `json:"distance"`
	Filterstatus int      `json:"filterstatus"`
}

type AaData struct {
	ExternalID string           `json:"externalId"`
	EventID    int              `json:"eventId"`
	SubeventID int              `json:"subeventId"`
	CategoryID int              `json:"categoryId"`
	WaveID     int              `json:"waveId"`
	AgegroupID any              `json:"agegroupId"`
	Tagcode    string           `json:"tagcode"`
	Bib        string           `json:"bib"`
	Firstname  string           `json:"firstname"`
	Lastname   string           `json:"lastname"`
	City       string           `json:"city"`
	State      string           `json:"state"`
	Country    string           `json:"country"`
	Age        int              `json:"age"`
	Sex        string           `json:"sex"`
	Event      string           `json:"event"`
	Subevent   string           `json:"subevent"`
	EventDate  string           `json:"event_date"`
	Chiptime   int              `json:"chiptime"`
	Clocktime  int              `json:"clocktime"`
	Pace       any              `json:"pace"`
	Diff       any              `json:"diff"`
	DistanceID int              `json:"distanceId"`
	Overall    int              `json:"overall"`
	Oversex    int              `json:"oversex"`
	Overdiv    int              `json:"overdiv"`
	Dq         bool             `json:"dq"`
	Data       string           `json:"data"`
	Data1      string           `json:"data1"`
	Data2      string           `json:"data2"`
	Data3      string           `json:"data3"`
	Data4      string           `json:"data4"`
	Data5      string           `json:"data5"`
	Data6      string           `json:"data6"`
	Data7      string           `json:"data7"`
	ThisPerson any              `json:"thisPerson"`
	Splits     map[string]Split `json:"splits,omitempty"`
	Listings   any              `json:"listings"`
	Start      any              `json:"start"`
}
