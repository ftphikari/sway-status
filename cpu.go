package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	cpu_full_text string
	cpu_interval  int32 = 3
)

func get_cpu(cs []float64) (cpu_str string) {
	var hcpu float64
	for _, c := range cs {
		if cpu_str != "" {
			cpu_str += "/"
		}
		cpu_str = fmt.Sprint(cpu_str, int(c))
		if c > hcpu {
			hcpu = c
		}
	}
	cpu_str = fmt.Sprint(" ", cpu_str, "%")
	if hcpu >= 75 {
		cpu_str = fmt.Sprint(`<span foreground="#da0f00"></span>`, cpu_str)
	} else {
		cpu_str = fmt.Sprint(``, cpu_str)
	}
	return
}

func init_cpu() {
	cpus, err := cpu.Percent(0, true)
	if err != nil {
		cpu_full_text = err.Error()
	} else {
		cpu_full_text = get_cpu(cpus)
	}

	go update_cpu()
}

func update_cpu() {
	ticker := time.NewTicker(time.Second * time.Duration(cpu_interval))
	defer ticker.Stop()

	for {
		<-ticker.C

		var new_cpu string
		cpus, err := cpu.Percent(0, true)
		if err != nil {
			new_cpu = err.Error()
		} else {
			new_cpu = get_cpu(cpus)
		}

		if cpu_full_text != new_cpu {
			m.Lock()
			cpu_full_text = new_cpu
			m.Unlock()
			update <- true
		}
	}
}
