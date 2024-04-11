package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

type Flags struct {
	Address *string
	Port    *int
	Path    *string

	user        *string
	pass        *string
	disableAuth *bool

	tls *bool

	Version *bool
}

func (f *Flags) init() {
	f.Address = flag.String("server-addr", "", "Address.")
	f.Port = flag.Int("server-port", 9111, "Port.")
	f.Path = flag.String("server-path", "/metrics", "Path.")

	f.user = flag.String("server-user", "admin", "Username.")
	f.pass = flag.String("server-pass", "admin", "Password.")
	f.disableAuth = flag.Bool("disable-auth", false, "Disable basic authentication.")

	f.tls = flag.Bool("server-tls", true, "Enable TLS.")

	f.Version = flag.Bool("version", false, "Show version.")

	flag.Parse()

	// System Environment Variables
	flag.VisitAll(func(f *flag.Flag) {
		name := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if value, ok := os.LookupEnv(name); ok {
			err := flag.Set(f.Name, value)
			if err != nil {
				log.Panicln("failed setting flag from environment:", err)
			}
		}
	})

}
