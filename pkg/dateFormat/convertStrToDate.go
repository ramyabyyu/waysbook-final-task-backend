package dateformat

import "time"

func ConvertStrToDate(dateString string) (time.Time, error) {

	layout := "2006-01-02"
	date, err := time.Parse(layout, dateString)
	
	if err != nil {
		return time.Now(), err
	}

	return date, err
}