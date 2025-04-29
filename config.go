package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Main             string
	BackendAddresses []string
}

const BACKENDADDRES = "BackendAddres"

func NewConfig() *Config {
	n := os.Getenv("NumberOfServers")

	if n == "" {
		log.Println("not successful reading env n ")
		return &Config{BackendAddresses: nil}
	}
	m := os.Getenv(BACKENDADDRES)
	if m == "" {
		log.Println("not successful reading env main addres ")

	}
	adrs := make([]string, 0)
	N, err := strconv.Atoi(n)
	if err != nil {
		log.Println("not successful converting in config")
	}
	for i := 1; i < N; i++ {
		s := BACKENDADDRES + strconv.Itoa(i)
		tmp := os.Getenv(s)

		if tmp == "" {
			log.Println("not successful reading env")
			continue
		}
		adrs = append(adrs, tmp)
	}
	return &Config{Main: m,
		BackendAddresses: adrs,
	}
}
