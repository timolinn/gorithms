package algos

import (
	"math/rand"
	"time"
)

type Side int

const (
	Head Side = iota
	Tail
)

func (h Side) String() string {
	if h == Head {
		return "heads"
	}
	return "tails"
}

func NewCoin() [2]Side {
	return [2]Side{Head, Tail}
}

func Toss(coin [2]Side) Side {
	rand.Seed(time.Now().UnixNano())
	return coin[rand.Intn(2)]
}
