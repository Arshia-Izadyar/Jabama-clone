package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp() string {
	rand.Seed(time.Now().Unix())
	min := int(math.Pow10(6 - 1))
	max := int(math.Pow10(6) - 1)
	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
