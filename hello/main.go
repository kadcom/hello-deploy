package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const defaultPort = 3000

func lookupEnvInt(key string, defaultValue uint16) uint16 {
	valStr, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	val, err := strconv.Atoi(valStr)

	if err != nil || val < 0 || val > 65535 {
		return defaultValue
	}

	return uint16(val)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!\n"))
	})

	listenAddr := fmt.Sprintf(":%d", lookupEnvInt("HELLO_PORT", defaultPort))

	http.ListenAndServe(listenAddr, r)
}
