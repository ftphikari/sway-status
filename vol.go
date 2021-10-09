package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

var (
	speaker_full_text string
	mic_full_text     string
)

func get_vol(vol int64, muted bool, dt int) (vol_str string) {
	vol_str = fmt.Sprint(" ", vol, "%")
	if dt == Source {
		if muted {
			vol_str = fmt.Sprint("", vol_str)
		} else {
			vol_str = fmt.Sprint("", vol_str)
		}
	} else {
		if muted {
			vol_str = fmt.Sprint("", vol_str)
		} else if vol <= 10 {
			vol_str = fmt.Sprint("", vol_str)
		} else if vol <= 50 {
			vol_str = fmt.Sprint("", vol_str)
		} else if vol <= 100 {
			vol_str = fmt.Sprint("", vol_str)
		} else {
			vol_str = fmt.Sprint(`<span foreground="#da0f00">`, "", "</span>", vol_str)
		}
	}
	return
}

func init_speaker() {
	sink, err := openControlDevice(Sink)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		speaker_full_text = "X"
		return
	}
	speaker_full_text = get_vol(sink.vol, sink.muted, sink.dev_type)
	go update_vol(&sink, &speaker_full_text)
}

func init_mic() {
	source, err := openControlDevice(Source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		mic_full_text = "X"
		return
	}
	mic_full_text = get_vol(source.vol, source.muted, source.dev_type)
	go update_vol(&source, &mic_full_text)
}

func update_vol(cd *controlDevice, vol_str *string) {
	defer cd.conn.Close()
	dt := ([]string{"Sink", "Source"})[cd.dev_type]
	signals := make(chan *dbus.Signal)
	cd.conn.Signal(signals)

	// Listen for signals from D-Bus, and update appropriately.
	for signal := range signals {
		// If the fallback device changed, open the new one.
		switch signal.Name {
		case "org.PulseAudio.Core1.Fallback" + dt + "Updated":
			err := openFallbackDevice(cd)
			if err != nil {
				m.Lock()
				*vol_str = err.Error()
				m.Unlock()
				update <- true
				continue
			}
		case "org.PulseAudio.Core1.Device.VolumeUpdated":
			channels := signal.Body[0].([]uint32)
			cd.vol = calcVol(channels, cd.max)
		case "org.PulseAudio.Core1.Device.MuteUpdated":
			cd.muted = signal.Body[0].(bool)
		}
		new_vol := get_vol(cd.vol, cd.muted, cd.dev_type)
		if *vol_str != new_vol {
			m.Lock()
			*vol_str = new_vol
			m.Unlock()
			update <- true
		}
	}
}
