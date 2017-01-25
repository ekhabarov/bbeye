//BigBrother Eye module for monitoring internals of app.
package bbeye

import (
	"expvar"
	"log"
	"net/http"
	"runtime"
)

//NumGoroutine returns the number of goroutines that currently exist.
func goroutines() interface{} {
	return runtime.NumGoroutine()
}

//Starts monitoring
func Run(addr string) {

	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	expvar.Publish("Goroutines", expvar.Func(goroutines))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln("unable to start monitoring: ", err)
	}
}
