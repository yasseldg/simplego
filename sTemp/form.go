package sTemp

import (
	"net/http"
	"time"
)

type Form struct {
	TsFrom int64
	TsTo   int64
}

func ReadDate(field, layout string, r *http.Request, msg *FlashMessages) int64 {
	if layout == "" {
		layout = "2006-01-02"
	}
	date := r.FormValue(field)
	date_time, err := time.Parse(layout, date)
	if err != nil {
		msg.Error("Supplied ( %s ) value ( %s ) is not in a date format: %s", field, date, err)
		date_time = time.Unix(0, 0)
	}
	return date_time.Unix()
}

func DatesForm(form *Form, layout string, msg *FlashMessages, r *http.Request) {
	form.TsFrom = ReadDate("ts_from", layout, r, msg)
	form.TsTo = ReadDate("ts_to", layout, r, msg)
}
