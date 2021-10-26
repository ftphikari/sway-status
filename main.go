package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/goodsign/monday"
)

type replacements map[string]string

type Module struct {
	f    func()
	text *string
}

var (
	format string
	showip bool
	repl   = make(replacements)

	m       sync.Mutex
	update  = make(chan bool)
	modules = map[string]Module{
		"cpu":     {init_cpu, &cpu_full_text},
		"temp":    {init_temp, &temp_full_text},
		"net":     {init_net, &net_full_text},
		"bat":     {init_bat, &bat_full_text},
		"speaker": {init_speaker, &speaker_full_text},
		"mic":     {init_mic, &mic_full_text},
		"lang":    {init_lang, &lang_full_text},
		"date":    {init_date, &date_full_text},
		"time":    {init_date, &time_full_text},
	}
	locale monday.Locale
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

func print_bar() {
	m.Lock()
	var mrepl []string
	for k, v := range modules {
		mrepl = append(mrepl, k, *v.text)
	}
	fmt.Println(strings.NewReplacer(mrepl[:]...).Replace(format))
	m.Unlock()
}

func main() {
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		locale = monday.Locale(strings.Split(envlang, ".")[0])
	} else {
		locale = monday.LocaleEnUS
	}

	flag.StringVar(&format, "f", "cpu temp | net | bat | speaker mic | lang | date | time ", "format")
	flag.BoolVar(&showip, "i", false, "show ip of network interfaces")
	flag.Var(repl, "r", "language replacements (e.g. -r 'English (US):US' -r 'Ukrainian:UA')")
	flag.Parse()

	for k, v := range modules {
		if strings.Contains(format, k) {
			v.f()
		}
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
