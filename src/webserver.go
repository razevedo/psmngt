package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"utils"

	"github.com/gorilla/mux"
)

var (
	listenAddr        = flag.String("http_addr", "localhost", "HTTP listen address. It overrides HTTP Port specified in the config file.")
	listenPort        = flag.String("http_port", "8080", "HTTP listen port. It overrides HTTP Port specified in the config file.")
	configClientsFile = flag.String("clients_config", "./clients.conf", "Specifies the config file to configure authorized applications.")
)

var runningProcesses map[int]Command

func main() {
	start := time.Now()

	flag.Parse()
	addr := utils.ConcateStrings(*listenAddr, ":", *listenPort)
	if err := loadAuthorizedClients(*configClientsFile); err != nil {
		log.Fatal("Impossible to load Clients Configuration File at - ", *configClientsFile)
	} else {
		log.Printf("Clients Configuration File loaded")
	}

	runningProcesses = make(map[int]Command)
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = authorize(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	log.Printf("Server Started - %s\t\t%s", addr, time.Since(start))

	log.Fatal(http.ListenAndServe(addr, router))
}

func authorize(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if !isAuthorized(r.Header.Get("access_token")) {
			log.Printf("Error: Not authorized")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
		} else {
			inner.ServeHTTP(w, r)
		}

		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			r.Method,
			r.Header.Get("access_token"),
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func prettyPrintPSList() {
	for key, value := range runningProcesses {
		log.Printf("Key: %d - Value: %s", key, value)
	}
}
