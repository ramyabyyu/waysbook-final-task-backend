package dateformat

import "time"

func ConvertStrToDate(dateString string) (time.Time) {

	layout := "2006-01-02"
	date, _ := time.Parse(layout, dateString)

	return date
}