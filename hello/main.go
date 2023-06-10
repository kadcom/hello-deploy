package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
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
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// Setup Handler
	r.Get("/", helloWorldHandler)

	// Setup Server
	listenPort := lookupEnvInt("HELLO_PORT", defaultPort)
	listenAddr := fmt.Sprintf(":%d", listenPort)

	log.Info().Uint16("port", listenPort).Msg("Staring Server...")

	err := http.ListenAndServe(listenAddr, r)
	if err != nil {
		log.Fatal().Err(err).Msg("Server is failed to start")
	}

	log.Info().Msg("Server Terminated.")
}
