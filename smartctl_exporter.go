package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	devices    []*Device
	results    = sync.Map{}
	collectors = map[string]*prometheus.GaugeVec{}

	flags = Flags{}

	AppName   = "smartctl_exporter"
	Version   = ""
	BuildDate = ""
)

func WithMetrics(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wg := sync.WaitGroup{}
		wg.Add(len(devices))
		for _, d := range devices {
			go func(d *Device) {
				results.Store(d.Name, GetAll(d))
				wg.Done()
			}(d)
		}
		wg.Wait()

		for _, d := range devices {
			if r, ok := results.Load(d.Name); ok {

				r := r.(*Result)
				passed := 0
				if r.Passed {
					passed = 1
				}
				collectors["device_status"].
					WithLabelValues(d.Name, r.ModelName, r.SerialNumber, r.FirmwareVersion).
					Set(float64(passed))

				for k, v := range r.Attributes {
					c, ok := collectors[k]
					if !ok {
						c = promauto.NewGaugeVec(
							prometheus.GaugeOpts{
								Namespace: "smartctl",
								Name:      k,
							},
							[]string{"device"},
						)
						collectors[k] = c
					}
					c.WithLabelValues(d.Name).Set(v)
				}
			}
		}

		handler.ServeHTTP(w, r)
	}
}

func main() {
	flags.init()

	if *flags.Version {
		fmt.Printf(
			"%s\n"+
				"Version: \t%s\n"+
				"Build date: \t%s\n",
			AppName,
			Version,
			BuildDate)
		return
	}

	serverTLSConf, err := certsetup()
	if err != nil {
		panic(err)
	}

	devices = GetDevices()

	deviceStatusCollector := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "smartctl",
			Name:      "device_status",
			Help:      "Device Status",
		},
		[]string{
			"device",
			"model_name",
			"serial_number",
			"firmware_version",
		},
	)
	collectors["device_status"] = deviceStatusCollector

	promHandler := promhttp.Handler()

	hf := WithMetrics(promHandler)

	if !*flags.disableAuth {
		hf = BasicAuth(hf)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(*flags.Path, hf)

	addr := fmt.Sprintf("%s:%d", *flags.Address, *flags.Port)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	if *flags.tls {
		server.TLSConfig = serverTLSConf
	}

	log.Printf("Listen: %s\n", addr)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
