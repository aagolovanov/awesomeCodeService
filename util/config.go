package util

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port   int
	DBAddr string
	DBPass string
	TTL    int // TTL in seconds
}

func LoadConfig() Config {
	port := flag.Int("port", 8080, "port for server")
	dbaddr := flag.String("dbaddr", "localhost", "KeyDB address")
	dbport := flag.Int("dbport", 6379, "KeyDB port")
	dbpass := flag.String("dbpass", "", "KeyDB password")
	ttl := flag.Int("TTL", 30, "TimeToLive for codes")

	flag.Parse()

	// костыль для оверрайда флагов енвами, переделать на viper или подобные
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			*port = p
		}
	}
	if envDBAddr := os.Getenv("DB_ADDR"); envDBAddr != "" {
		*dbaddr = envDBAddr
	}
	if envDBPass := os.Getenv("DB_PASS"); envDBPass != "" {
		*dbpass = envDBPass
	}
	if envDBPort := os.Getenv("DB_PORT"); envDBPort != "" {
		if p, err := strconv.Atoi(envDBPort); err == nil {
			*dbport = p
		}
	}
	if envTTL := os.Getenv("TTL"); envTTL != "" {
		if t, err := strconv.Atoi(envTTL); err == nil {
			*ttl = t
		}
	}

	return Config{
		Port:   *port,
		DBAddr: *dbaddr + ":" + strconv.Itoa(*dbport),
		DBPass: *dbpass,
		TTL:    *ttl,
	}
}
