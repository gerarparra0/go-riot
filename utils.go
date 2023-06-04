package goriot

import (
	"os"
	"strings"
)

func intRef(i int) *int {
	return &i
}

func loadKey() string {
	bts, err := os.ReadFile(".key")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bts))
}
