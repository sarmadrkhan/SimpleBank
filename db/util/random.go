package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNumberSequence generates a random string of  of n digits
func RandomNumberSequence(n int) string {
	var sb strings.Builder
	k := len(numbers)

	for i := 0; i < n; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomMoney generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "PKR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomAccountID generates a random accound id
func RandomAccountID() int64 {
	numString := RandomNumberSequence(11)
	num, err := strconv.ParseInt(numString, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}
