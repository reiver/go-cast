package cast

import (
	"math/rand"
	"time"
)

var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
)

const (
	maxUint uint = ^uint(0)
	maxInt   int = int(maxUint >> 1)
	minInt   int = -maxInt - 1
)
