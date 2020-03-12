package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sethvargo/go-password/password"
)

func main() {
	length := func() int {
		var l int64 = 16
		if len(os.Args) > 1 {
			l, _ = strconv.ParseInt(os.Args[1], 10, 64)
		}
		return int(l)
	}()

	res, err := password.Generate(length, length/2, length/2, false, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(res)
}
