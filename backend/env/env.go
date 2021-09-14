package env

import (
	"os"
)

const (
	portKey    = "PLAYGROUND_BACKEND_PORT"
	networkKey = "PLAYGROUND_NETWORK"
)

var (
	Port    string
	Network string
)

func init() {
	var ok bool

	Port, ok = os.LookupEnv(portKey)
	if !ok {
		panic(portKey + " is not set")
	}

	Network, ok = os.LookupEnv(networkKey)
	if !ok {
		panic(networkKey + " is not set")
	}
}
