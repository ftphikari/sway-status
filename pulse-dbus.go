package main

import (
	"fmt"
	"math"
	"os"

	"github.com/godbus/dbus/v5"
)

const (
	// Sink represents devices used for audio output, e.g. headphones.
	Sink = iota

	// Source represents devices used for audio input, e.g. microphones.
	Source
)

type controlDevice struct {
	conn     *dbus.Conn
	core     dbus.BusObject
	dev      dbus.BusObject
	dev_type int
	max, vol int64
	muted    bool
}

func listen(core dbus.BusObject, signal string, objects ...dbus.ObjectPath) error {
	return core.Call(
		"org.PulseAudio.Core1.ListenForSignal",
		0,
		"org.PulseAudio.Core1."+signal,
		objects,
	).Err
}

func openFallbackDevice(cd *controlDevice) (err error) {
	var path dbus.Variant
	dt := ([]string{"Sink", "Source"})[cd.dev_type]
	path, err = cd.core.GetProperty("org.PulseAudio.Core1.Fallback" + dt)
	if err != nil {
		return
	}
	device_path := path.Value().(dbus.ObjectPath)
	cd.dev = cd.conn.Object("org.PulseAudio.Core1."+dt, device_path)

	err = listen(cd.core, "Device.VolumeUpdated", device_path)
	if err != nil {
		return
	}
	err = listen(cd.core, "Device.MuteUpdated", device_path)
	if err != nil {
		return
	}
	err = listen(cd.core, "Fallback"+dt+"Updated")
	if err != nil {
		return
	}

	cd.max, cd.vol, cd.muted, err = get_volume_manual(cd.dev, cd.dev_type)
	if err != nil {
		return
	}
	return
}

func openPulseAudio() (*dbus.Conn, error) {
	// Pulse defaults to creating its socket in a well-known place under
	// XDG_RUNTIME_DIR. For Pulse instances created by systemd, this is the
	// only reliable way to contact Pulse via D-Bus, since Pulse is created
	// on a per-user basis, but the session bus is created once for every
	// session, and a user can have multiple sessions.
	xdgDir := os.Getenv("XDG_RUNTIME_DIR")
	if xdgDir != "" {
		addr := fmt.Sprintf("unix:path=%s/pulse/dbus-socket", xdgDir)
		conn, err := dbus.Dial(addr)
		if err != nil {
			return nil, err
		}
		err = conn.Auth(nil)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}

	// Couldn't find the PulseAudio bus on the fast path, so look for it
	// by querying the session bus.
	bus, err := dbus.SessionBusPrivate()
	if err != nil {
		return nil, err
	}
	defer bus.Close()
	err = bus.Auth(nil)
	if err != nil {
		return nil, err
	}

	locator := bus.Object("org.PulseAudio1", "/org/pulseaudio/server_lookup1")
	path, err := locator.GetProperty("org.PulseAudio.ServerLookup1.Address")
	if err != nil {
		return nil, err
	}

	conn, err := dbus.Dial(path.Value().(string))
	if err != nil {
		return nil, err
	}
	err = conn.Auth(nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Take the volume as the average across all channels.
func calcVol(channels []uint32, max int64) int64 {
	var totalVol int64
	for _, ch := range channels {
		totalVol += int64(ch)
	}
	vol := (totalVol / int64(len(channels)))
	return int64(math.Ceil(float64(vol*100) / float64(max)))
}

func get_volume_manual(device dbus.BusObject, dev_type int) (max, vol int64, muted bool, err error) {
	max = 0x10000
	/*	if dev_type == Sink {
		var maxVol dbus.Variant
		maxVol, err = device.GetProperty("org.PulseAudio.Core1.Device.BaseVolume")
		if err != nil {
			return
		}
		max = int64(maxVol.Value().(uint32))
	}*/

	currentVol, err := device.GetProperty("org.PulseAudio.Core1.Device.Volume")
	if err != nil {
		return
	}

	vol = calcVol(currentVol.Value().([]uint32), max)

	mute, err := device.GetProperty("org.PulseAudio.Core1.Device.Mute")
	if err != nil {
		return
	}
	muted = mute.Value().(bool)

	return
}

func openControlDevice(dev_type int) (cd controlDevice, err error) {
	cd.dev_type = dev_type
	cd.conn, err = openPulseAudio()
	if err != nil {
		return
	}

	cd.core = cd.conn.Object("org.PulseAudio.Core1", "/org/pulseaudio/core1")
	err = openFallbackDevice(&cd)
	if err != nil {
		return
	}
	return
}
