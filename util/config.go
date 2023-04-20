package util

type Config struct {
	Port   int
	DBAddr string
	DBPass string
	TTL    int // TTL in seconds
}

func LoadConfig() (Config, error) {
	panic("TODO")
}
