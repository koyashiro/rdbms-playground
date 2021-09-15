package env

import (
	"os"
)

const (
	networkKey = "PLAYGROUND_NETWORK"
)

var (
	Network string
)

func init() {
	var ok bool

	Network, ok = os.LookupEnv(networkKey)
	if !ok {
		panic(networkKey + " is not set")
	}
}
