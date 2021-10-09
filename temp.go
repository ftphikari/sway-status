package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

var (
	temp_full_text string
	temp_interval  int32 = 3
)

func get_temp(ts []host.TemperatureStat) (temp_str string) {
	var htemp float64
	for _, t := range ts {
		if strings.HasPrefix(t.SensorKey, "coretemp_core") {
			if temp_str != "" {
				temp_str += "/"
			}
			temp_str = fmt.Sprint(temp_str, t.Temperature)
			if t.Temperature > htemp {
				htemp = t.Temperature
			}
		}
	}
	temp_str = fmt.Sprint(" ", temp_str, "°C")
	if htemp >= 75 {
		temp_str = fmt.Sprint(`<span foreground="#da0f00"></span>`, temp_str)
	} else {
		temp_str = fmt.Sprint(``, temp_str)
	}
	return
}

func init_temp() {
	temps, _ := host.SensorsTemperatures()
	temp_full_text = get_temp(temps)

	go update_temp()
}

func update_temp() {
	ticker := time.NewTicker(time.Second * time.Duration(temp_interval))
	defer ticker.Stop()

	for {
		<-ticker.C

		temps, _ := host.SensorsTemperatures()
		new_temp := get_temp(temps)

		if temp_full_text != new_temp {
			m.Lock()
			temp_full_text = new_temp
			m.Unlock()
			update <- true
		}
	}
}
