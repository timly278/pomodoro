package util

import (
	"math/rand"
	"strings"
)

const (
	alphabet = "asdfghjklzxcvbnmqwertyuiop"
)

// RandomInt generate a random integer within [min,max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a random string with length n
func RandomString(n int) string {
	k := len(alphabet)
	var str strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		str.WriteByte(c)
	}

	return str.String()
}

func RandomColor() string {
	var list = []string{"Red", "Yellow", "Orange", "Green"}
	return list[RandomInt(0, int64(len(list) - 1))]
}

func RandomAlarmSound() string {
	var list = []string{"Kitchen", "Bell", "Bird", "Digital"}
	return list[RandomInt(0, int64(len(list) - 1))]
}