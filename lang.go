package main

import (
	"context"

	"github.com/joshuarubin/go-sway"
)

type handler struct {
	sway.EventHandler
}

var (
	lang_full_text string
)

func (h handler) Input(ctx context.Context, e sway.InputEvent) {
	if e.Change != "xkb_layout" {
		return
	}

	l := get_layout(*e.Input.XKBActiveLayoutName)
	if lang_full_text != l {
		m.Lock()
		lang_full_text = l
		m.Unlock()
		update <- true
	}
}

func get_layout(l string) string {
	r, ok := repl[l]
	if ok {
		return " " + r
	} else {
		return " " + l
	}
}

func init_lang() {
	ctx, cancel := context.WithCancel(context.Background())

	client, err := sway.New(ctx)
	if err != nil {
		lang_full_text = err.Error()
		return
	}

	inputs, err := client.GetInputs(ctx)
	if err != nil {
		lang_full_text = err.Error()
		return
	}

	for _, i := range inputs {
		if i.XKBActiveLayoutName != nil && *i.XKBActiveLayoutName != "" {
			l := get_layout(*i.XKBActiveLayoutName)
			lang_full_text = l
		}
	}

	go update_lang(ctx, cancel)
}

func update_lang(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	err := sway.Subscribe(ctx, handler{sway.NoOpEventHandler()}, sway.EventTypeInput)
	if err != nil {
		m.Lock()
		lang_full_text = err.Error()
		m.Unlock()
		return
	}
}
