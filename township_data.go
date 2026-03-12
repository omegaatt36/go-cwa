package cwa

// townshipDatasetIDs maps county and period to the CWA dataset ID (F-D0047-XXX).
var townshipDatasetIDs = map[County]map[Period]string{
	YilanCounty:      {ThreeDays: "F-D0047-001", OneWeek: "F-D0047-003"},
	TaoyuanCity:      {ThreeDays: "F-D0047-005", OneWeek: "F-D0047-007"},
	HsinchuCounty:    {ThreeDays: "F-D0047-009", OneWeek: "F-D0047-011"},
	MiaoliCounty:     {ThreeDays: "F-D0047-013", OneWeek: "F-D0047-015"},
	ChanghuaCounty:   {ThreeDays: "F-D0047-017", OneWeek: "F-D0047-019"},
	NantouCounty:     {ThreeDays: "F-D0047-021", OneWeek: "F-D0047-023"},
	YunlinCounty:     {ThreeDays: "F-D0047-025", OneWeek: "F-D0047-027"},
	ChiayiCounty:     {ThreeDays: "F-D0047-029", OneWeek: "F-D0047-031"},
	PingtungCounty:   {ThreeDays: "F-D0047-033", OneWeek: "F-D0047-035"},
	TaitungCounty:    {ThreeDays: "F-D0047-037", OneWeek: "F-D0047-039"},
	HualienCounty:    {ThreeDays: "F-D0047-041", OneWeek: "F-D0047-043"},
	PenghuCounty:     {ThreeDays: "F-D0047-045", OneWeek: "F-D0047-047"},
	KeelungCity:      {ThreeDays: "F-D0047-049", OneWeek: "F-D0047-051"},
	HsinchuCity:      {ThreeDays: "F-D0047-053", OneWeek: "F-D0047-055"},
	ChiayiCity:       {ThreeDays: "F-D0047-057", OneWeek: "F-D0047-059"},
	TaipeiCity:       {ThreeDays: "F-D0047-061", OneWeek: "F-D0047-063"},
	KaohsiungCity:    {ThreeDays: "F-D0047-065", OneWeek: "F-D0047-067"},
	NewTaipeiCity:    {ThreeDays: "F-D0047-069", OneWeek: "F-D0047-071"},
	TaichungCity:     {ThreeDays: "F-D0047-073", OneWeek: "F-D0047-075"},
	TainanCity:       {ThreeDays: "F-D0047-077", OneWeek: "F-D0047-079"},
	LienchiangCounty: {ThreeDays: "F-D0047-081", OneWeek: "F-D0047-083"},
	KinmenCounty:     {ThreeDays: "F-D0047-085", OneWeek: "F-D0047-087"},
}
