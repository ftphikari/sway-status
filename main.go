package main

import (
	"errors"
	"fmt"
	"sync"
	"strings"
	"flag"
)

type replacements map[string]string

var (
	m       sync.Mutex
	update  chan bool
	format  string
	repl    = make(replacements)
)

func (r replacements) String() string {
	var s string
	for k, v := range repl {
		s = fmt.Sprint(s, k, ":", v, ";")
	}
	return s
}

func (r replacements) Set(value string) error {
	v := strings.Split(value, ":")
	if len(v) < 2 {
		return errors.New("cannot parse: " + value)
	}
	r[v[0]] = v[1]
	return nil
}

// print modules
func print_bar() {
	m.Lock()
	replacer := strings.NewReplacer("cpu", cpu_full_text, "temp", temp_full_text, "bat", bat_full_text, "speaker", speaker_full_text, "mic", mic_full_text, "date", date_full_text, "time", time_full_text, "lang", lang_full_text, "net", net_full_text)
	status := replacer.Replace(format)
	fmt.Println(status)
	m.Unlock()
}

func main() {
	flag.StringVar(&format, "f", "cpu temp | net | bat | speaker mic | date | time | lang ", "format")
	flag.Var(repl, "r", "language replacements (e.g. -r 'English (US):US' -r 'Ukrainian:UA')")
	flag.Parse()

	// setup modules
	update = make(chan bool)

	if strings.Contains(format, "lang") {
		init_lang()
	}
	if strings.Contains(format, "speaker") {
		init_speaker()
	}
	if strings.Contains(format, "mic") {
		init_mic()
	}
	if strings.Contains(format, "bat") {
		init_bat()
	}
	if strings.Contains(format, "cpu") {
		init_cpu()
	}
	if strings.Contains(format, "temp") {
		init_temp()
	}
	if strings.Contains(format, "net") {
		init_net()
	}
	if strings.Contains(format, "date") || strings.Contains(format, "time") {
		init_date()
	}

	print_bar()

	for {
		ok := <-update
		if !ok {
			return
		}
		print_bar()
	}
}
