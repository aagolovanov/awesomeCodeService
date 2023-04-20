package util

import (
	"flag"
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

	return Config{
		Port:   *port,
		DBAddr: *dbaddr + ":" + strconv.Itoa(*dbport),
		DBPass: *dbpass,
		TTL:    *ttl,
	}
}
