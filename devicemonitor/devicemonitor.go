package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	value := "{\"fields\": [\"kismet.device.base.key\", \"kismet.device.base.macaddr\", \"kismet.device.base.last_time\", \"kismet.device.base.signal/kismet.common.signal.last_signal_dbm\", \"dot11.device/dot11.device.last_beaconed_ssid\", \"dot11.device/dot11.device.last_probed_ssid\"]}"

	resp, err := http.PostForm("http://192.168.0.24:2501/devices/last-time/-10/devices.ekjson", url.Values{"json": {value}})
	if err != nil {
		log.Fatal("Error connecting to server.")
	}

	if resp.StatusCode != 200 {
		log.Fatal("Error response from server:", resp.StatusCode)
	}

	resp.Body.Close()
}
