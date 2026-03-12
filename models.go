package cwa

// County represents a Taiwan county/city name as used in the CWA API.
type County string

const (
	TaipeiCity       County = "臺北市"
	NewTaipeiCity    County = "新北市"
	TaoyuanCity      County = "桃園市"
	TaichungCity     County = "臺中市"
	TainanCity       County = "臺南市"
	KaohsiungCity    County = "高雄市"
	KeelungCity      County = "基隆市"
	HsinchuCity      County = "新竹市"
	HsinchuCounty    County = "新竹縣"
	MiaoliCounty     County = "苗栗縣"
	ChanghuaCounty   County = "彰化縣"
	NantouCounty     County = "南投縣"
	YunlinCounty     County = "雲林縣"
	ChiayiCity       County = "嘉義市"
	ChiayiCounty     County = "嘉義縣"
	PingtungCounty   County = "屏東縣"
	YilanCounty      County = "宜蘭縣"
	HualienCounty    County = "花蓮縣"
	TaitungCounty    County = "臺東縣"
	PenghuCounty     County = "澎湖縣"
	KinmenCounty     County = "金門縣"
	LienchiangCounty County = "連江縣"
)

// Response represents the standard wrapper for CWA API responses.
type Response[T any] struct {
	Success string `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
	} `json:"result"`
	Records T `json:"records"`
}

// Forecast36hRecords contains the 36-hour forecast data.
type Forecast36hRecords struct {
	DatasetDescription string     `json:"datasetDescription"`
	Location           []Location `json:"location"`
}

// Location represents weather data for a specific location.
type Location struct {
	LocationName   string           `json:"locationName"`
	WeatherElement []WeatherElement `json:"weatherElement"`
}

// WeatherElement represents a specific weather metric (e.g., Wx, MaxT).
type WeatherElement struct {
	ElementName string `json:"elementName"`
	Time        []Time `json:"time"`
}

// Time represents a forecast period.
type Time struct {
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Parameter Parameter `json:"parameter"`
}

// Parameter contains the actual forecast values.
type Parameter struct {
	ParameterName  string `json:"parameterName"`
	ParameterValue string `json:"parameterValue"`
	ParameterUnit  string `json:"parameterUnit"`
}

// IsSuccess returns true if the API request was successful.
func (r *Response[T]) IsSuccess() bool {
	return r != nil && r.Success == "true"
}

// FirstLocation safely returns the first location in the records if it exists.
func (r *Forecast36hRecords) FirstLocation() *Location {
	if r != nil && len(r.Location) > 0 {
		return &r.Location[0]
	}
	return nil
}

// Period represents the forecast period (3 days or 1 week).
type Period string

const (
	ThreeDays Period = "3days"
	OneWeek   Period = "1week"
)
