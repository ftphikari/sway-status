package main

import (
	"time"

	"github.com/goodsign/monday"
)

var (
	date_full_text string
	time_full_text string
	date_interval  int32  = 1
	date_format    string = " Mon 2006-01-02"
	time_format    string = " 15:04"
	ticker_enabled bool
)

func init_date() {
	// a single instance updates both date and time
	// so do not allow running 2 instances at the same time
	if ticker_enabled {
		return
	}

	ticker_enabled = true
	t := time.Now()
	date_full_text = monday.Format(t, date_format, locale)
	time_full_text = monday.Format(t, time_format, locale)

	go update_date()
}

func update_date() {
	ticker := time.NewTicker(time.Second * time.Duration(date_interval))
	defer ticker.Stop()

	for {
		t := <-ticker.C
		// for testing time difference
		//fmt.Println("Current time: ", t)
		new_date := monday.Format(t, date_format, locale)
		new_time := monday.Format(t, time_format, locale)
		upd := false

		if date_full_text != new_date {
			m.Lock()
			date_full_text = new_date
			m.Unlock()
			upd = true
		}
		if time_full_text != new_time {
			m.Lock()
			time_full_text = new_time
			m.Unlock()
			upd = true
		}

		if upd {
			update <- true
		}
	}
}
