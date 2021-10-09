package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type battery struct {
	Path       string
	OnBattery  bool
	Percentage int64
	Capacity   int64
}

var (
	bat_full_text string

	bat = battery{"UNKNOWN", false, -1, -1}
)

func get_bat() (bat_str string) {
	if bat.OnBattery {
		if bat.Percentage <= 10 {
			bat_str = fmt.Sprint(bat_str, "")
		} else if bat.Percentage <= 25 {
			bat_str = fmt.Sprint(bat_str, "")
		} else if bat.Percentage <= 50 {
			bat_str = fmt.Sprint(bat_str, "")
		} else if bat.Percentage <= 75 {
			bat_str = fmt.Sprint(bat_str, "")
		} else {
			bat_str = fmt.Sprint(bat_str, "")
		}
	} else {
		bat_str = fmt.Sprint(bat_str, "")
	}
	if !bat.OnBattery {
		bat_str = fmt.Sprint(`<span foreground="#00dc1b">`, bat_str, `</span>`)
	} else if bat.Percentage <= 10 {
		bat_str = fmt.Sprint(`<span foreground="#da0f00">`, bat_str, `</span>`)
	} else if bat.Percentage <= 30 {
		bat_str = fmt.Sprint(`<span foreground="#da7300">`, bat_str, `</span>`)
	} else {
		bat_str = fmt.Sprint(bat_str)
	}
	bat_str = fmt.Sprint(bat_str, " ", bat.Percentage, "% ")
	if bat.Capacity <= 10 {
		bat_str = fmt.Sprint(bat_str, `<span foreground="#da0f00"></span>`)
	} else if bat.Capacity <= 30 {
		bat_str = fmt.Sprint(bat_str, `<span foreground="#da7300"></span>`)
	} else {
		bat_str = fmt.Sprint(bat_str, "")
	}
	bat_str = fmt.Sprint(bat_str, " ", bat.Capacity, "%")
	return
}

func init_bat() {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		bat_full_text = err.Error()
		return
	}

	call := conn.Object("org.freedesktop.UPower", dbus.ObjectPath("/org/freedesktop/UPower")).Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UPower")
	if call.Err != nil {
		bat_full_text = call.Err.Error()
		return
	}
	opts := call.Body[0].(map[string]dbus.Variant)
	if v, ok := opts["OnBattery"]; ok {
		bat.OnBattery = v.Value().(bool)
	}

	call = conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower").Call("org.freedesktop.UPower.EnumerateDevices", 0)
	if call.Err != nil {
		bat_full_text = call.Err.Error()
		return
	}
	devices := call.Body[0].([]dbus.ObjectPath)
	for _, dev := range devices {
		call := conn.Object("org.freedesktop.UPower", dev).Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UPower.Device")
		if call.Err != nil {
			bat_full_text = call.Err.Error()
			return
		}
		opts := call.Body[0].(map[string]dbus.Variant)
		if t, ok := opts["Type"]; ok {
			/*
				0: Unknown
				1: Line Power
				2: Battery
				3: Ups
				4: Monitor
				5: Mouse
				6: Keyboard
				7: Pda
				8: Phone
			*/
			if t.Value().(uint32) == 2 {
				bat.Path = string(dev)
				if v, ok := opts["Percentage"]; ok {
					bat.Percentage = int64(v.Value().(float64))
				}
				if v, ok := opts["Capacity"]; ok {
					bat.Capacity = int64(v.Value().(float64))
				}
			}
		}
	}
	bat_full_text = get_bat()

	// Wait for any "PropertiesChanged" signal sent by UPower. Normally
	// you'd wait for this on a particular object, and then listen to
	// DeviceAdded/DeviceRemoved to remove/add new things to listen for.
	err = conn.AddMatchSignal(
		dbus.WithMatchInterface("org.freedesktop.DBus.Properties"),
		dbus.WithMatchSender("org.freedesktop.UPower"),
		dbus.WithMatchMember("PropertiesChanged"))
	if err != nil {
		bat_full_text = err.Error()
		return
	}

	go update_bat(conn)
}

func update_bat(conn *dbus.Conn) {
	signals := make(chan *dbus.Signal)
	conn.Signal(signals)
	for signal := range signals {
		//		sender := signal.Body[0].(string)
		body := signal.Body[1].(map[string]dbus.Variant)
		//		fmt.Println("Sender: ", sender)
		//		fmt.Println("Path: ", signal.Path)
		//		fmt.Println("Name: ", signal.Name)
		//		fmt.Println("Body: ", body)
		old_bat := bat

		switch string(signal.Path) {
		case bat.Path:
			if v, ok := body["Percentage"]; ok {
				bat.Percentage = int64(v.Value().(float64))
			}
			if v, ok := body["Capacity"]; ok {
				bat.Capacity = int64(v.Value().(float64))
			}
		case "/org/freedesktop/UPower":
			if v, ok := body["OnBattery"]; ok {
				bat.OnBattery = v.Value().(bool)
			}
		}

		if old_bat != bat {
			m.Lock()
			bat_full_text = get_bat()
			m.Unlock()
			update <- true
		}
	}
}
