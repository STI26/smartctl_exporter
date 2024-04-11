package main

import (
	"log"
	"net/http"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok {
			log.Println(r.RemoteAddr, "| error parsing basic auth")
			w.Header().Set("WWW-Authenticate", `Basic realm=""`)
			w.WriteHeader(401)
			return
		}

		if user != *flags.user {
			log.Println(r.RemoteAddr, "| username provided is incorrect:", user)
			w.Header().Set("WWW-Authenticate", `Basic realm=""`)
			w.WriteHeader(401)
			return
		}

		if pass != *flags.pass {
			log.Println(r.RemoteAddr, "| password provided is incorrect:", user)
			w.Header().Set("WWW-Authenticate", `Basic realm=""`)
			w.WriteHeader(401)
			return
		}

		handler(w, r)
	}
}
