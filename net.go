package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mdlayher/wifi"
)

var (
	net_full_text string
	net_interval  int = 1
)

func getIP(i net.Interface) (ip string) {
	addrs, _ := i.Addrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok {
			if ipnet.IP.To4() == nil {
				ip = ipnet.IP.String()
			} else {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}

func getSSID(i int, c *wifi.Client) (ssid string) {
	ifs, err := c.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range ifs {
		if iface.Index != i {
			continue
		}

		bss, err := c.BSS(iface)
		if err != nil {
			break
		}

		ssid = bss.SSID
		return
	}

	return
}

func isWifi(i int, c *wifi.Client) (ok bool) {
	wifs, err := c.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range wifs {
		if iface.Index == i {
			ok = true
			return
		}
	}

	return
}

func hasCarrier(n string) (ok bool) {
	b, err := os.ReadFile("/sys/class/net/" + n + "/carrier")
	if err != nil {
		return
	}

	if len(b) > 0 && b[0] == '1' {
		ok = true
	}
	return
}

func get_net(c *wifi.Client) (net_str string) {
	ifs, _ := net.Interfaces()
	for _, i := range ifs {
		if i.Flags&net.FlagLoopback != 0 {
			continue
		}
		if i.Flags&net.FlagUp == 0 {
			continue
		}
		if net_str != "" {
			net_str = fmt.Sprint(net_str, " ")
		}

		iswifi := isWifi(i.Index, c)
		ip := getIP(i)

		var label, ssid string
		var connected bool
		if iswifi {
			label = ""
			ssid = getSSID(i.Index, c)
			if ssid != "" {
				connected = true
			}
		} else {
			label = ""
			connected = hasCarrier(i.Name)
		}
		if connected && ip != "" {
			label = `<span foreground="#00dc1b">` + label + `</span>`
		}
		net_str = fmt.Sprint(net_str, label)
		if !connected {
			continue
		}

		if iswifi {
			if showip && ip != "" {
				net_str = fmt.Sprint(net_str, " [", ip, " ", ssid, "]")
			} else {
				net_str = fmt.Sprint(net_str, " [", ssid, "]")
			}
		} else if showip && ip != "" {
			net_str = fmt.Sprint(net_str, " [", ip, "]")
		}
	}
	return
}

func init_net() {
	c, err := wifi.New()
	if err != nil {
		return
	}

	net_full_text = get_net(c)

	go update_net(c)
}

func update_net(c *wifi.Client) {
	defer c.Close()
	ticker := time.NewTicker(time.Second * time.Duration(net_interval))
	defer ticker.Stop()

	for {
		<-ticker.C
		new_net := get_net(c)

		if net_full_text != new_net {
			m.Lock()
			net_full_text = new_net
			m.Unlock()
			update <- true
		}
	}
}
